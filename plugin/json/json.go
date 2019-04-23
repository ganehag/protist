package main

import (
	"encoding/json"
	"errors"
	"github.com/ganehag/protist/filter"
	"github.com/ganehag/protist/plugin"
	"io/ioutil"
	"os"
	"sort"
)

type JsonDefinition struct {
	Id       string   `json:"id"`
	Function string   `json:"function"`
	Args     string   `json:"args"`
	Order    int32    `json:"order"`
	Source   []string `json:"source"`
}

type jsonPlugin struct {
	author  string
	name    string
	version string
	file    string
}

func (p *jsonPlugin) Author() string {
	return p.author
}

func (p *jsonPlugin) Name() string {
	return p.name
}

func (p *jsonPlugin) Version() string {
	return p.version
}

func (p *jsonPlugin) Get(id string) ([]plugin.FilterDefinition, error) {
	var jsonItems []JsonDefinition
	var filterItems []plugin.FilterDefinition

	if jsonFile, err := os.Open(id); err != nil {
		return nil, err
	} else {
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		if err := json.Unmarshal(byteValue, &jsonItems); err != nil {
			return nil, err
		}
	}

	sort.Slice(jsonItems, func(i, j int) bool {
		return jsonItems[i].Order < jsonItems[j].Order
	})

	for _, item := range jsonItems {
		filterItems = append(filterItems, plugin.FilterDefinition{
			Id:       item.Id,
			Function: item.Function,
			Args:     item.Args,
			Order:    item.Order,
			Source:   item.Source,
		})
	}

	return filterItems, nil
}

func (p *jsonPlugin) Chains() ([]string, error) {
	chain := make([]string, 1)
	chain[0] = p.file
	return chain, nil
}

func (p *jsonPlugin) Factory(instance int64, t string, cfg string) (filter.Filter, error) {
	return nil, errors.New("invalid type")
}

func Init(args interface{}) plugin.I {
	var argv map[string]interface{}
	var file string

	if a, ok := args.(map[string]interface{}); ok {
		argv = a
	}

	if v, ok := argv["file"].(string); ok {
		file = v
	}

	return &jsonPlugin{
		author:  "Mikael Ganehag Brorsson",
		name:    "json-loader",
		version: "0.0.1",
		file:    file,
	}
}
