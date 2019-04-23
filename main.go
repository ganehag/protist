package main

import (
	"errors"
	"fmt"
	"github.com/ganehag/protist/filter"
	"github.com/ganehag/protist/plugin"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	goplugin "plugin"
	"sync"
	"time"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/protist")
	viper.AddConfigPath("$HOME/.protist")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault("plugins", "")

	viper.SetDefault("update", "5m")
	viper.SetDefault("verbose", false)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("ERROR [config.yaml] %s\n", err)
	}
}

func loadPlugin(sofile string, cfg interface{}) (plugin.I, error) {
	so, err := goplugin.Open(sofile)
	if err != nil {
		return nil, err
	}

	symbol, err := so.Lookup("Init")
	if err != nil {
		return nil, err
	}

	p := symbol.(func(interface{}) plugin.I)(cfg)
	if p == nil {
		return nil, errors.New("unable to initialize plugin")
	}

	return p, nil
}

func main() {
	var chains []string
	var wg sync.WaitGroup
	var plugIn plugin.I
	var err error

	intr := make(chan os.Signal, 1)
	signal.Notify(intr, os.Interrupt, os.Kill)
	termchan := make(chan int, 1)

	/*
	 * Plug-ins
	 */
	if enabledPlugin := viper.Get("enabled_plugin"); enabledPlugin != nil {
		var so string
		var pluginConfig map[string]interface{}

		if s, ok := viper.Get(fmt.Sprintf("plugins.%s.sofile", enabledPlugin)).(string); ok {
			so = s
		} else {
			return // ERROR
		}

		if s, ok := viper.Get(fmt.Sprintf("plugins.%s", enabledPlugin)).(map[string]interface{}); ok {
			pluginConfig = s
		} else {
			return // ERROR
		}

		plugIn, err = loadPlugin(so, pluginConfig)
		if err != nil {
			log.Println("ERROR", err)
			return
		}
	}

	/*
	 * Load chains from plugin
	 */
	if chains, err = plugIn.Chains(); err != nil {
		log.Println("ERROR", err)
		return
	}

	for instance, id := range chains {
		wg.Add(1)

		go func() {
			defer wg.Done()

			var ticker *time.Ticker
			if dur, ok := viper.Get("update").(string); ok {
				if updateDuration, err := time.ParseDuration(dur); err != nil {
					log.Println("ERROR", err)
					return

				} else {
					ticker = time.NewTicker(updateDuration)
				}
			}

		loop:
			for {
				var filterItems []plugin.FilterDefinition
				var err error

				filterLut := make(map[string]filter.Filter)

				if filterItems, err = plugIn.Get(id); err != nil {
					log.Println("ERROR", err)
					return
				}

				// Build filter, create LUT
				for _, item := range filterItems {
					var f filter.Filter
					var err error

					if f, err = plugIn.Factory(int64(instance), item.Function, item.Args); err == nil {
						filterLut[item.Id] = f
						continue // continue loop as we have now handled this function
					}

					if f, err = filter.Factory(int64(instance), item.Function, item.Args); err != nil {
						log.Printf("ERROR %+v\n", err)
					} else {
						filterLut[item.Id] = f
					}
				}

				// Map sources to filters
				for _, item := range filterItems {
					if targetFilter, ok := filterLut[item.Id]; ok && item.Source != nil {
						for _, source := range item.Source {
							if filter, ok := filterLut[source]; ok {
								targetFilter.AddSource(filter)
							}
						}
					}
				}

				// Execute the entire chain
				for _, item := range filterItems {
					if filter, ok := filterLut[item.Id]; ok {
						if err := filter.Eval(); err != nil {
							log.Println("ERROR", err)
						} else {
							rv := filter.GetRetval()
							if rv != nil {
								log.Printf("DEBUG %s %+v\n", item.Function, rv)
							}
						}
					}
				}

				/*
				 * Imperfect solution. Will not execute every Nth duration. But Nth duration + exectime.
				 */
				select {
				case <-ticker.C:
					log.Println("INFO", "updating")
					continue
				case <-termchan: // triggered when the stop channel is closed
					break loop // exit
				}
			}
		}()
	}

	// Blocking main loop
	select {
	case s := <-intr:
		log.Println("INFO", "got signal:", s)
		close(termchan)
		break
	}

	wg.Wait()

	log.Println("INFO", "done")
}
