package plugin

import(
	"github.com/ganehag/protist/filter"
)

type FilterDefinition struct {
	Id       string
	Function string
	Args     string
	Order    int32
	Source   []string
}

type I interface {
	Get(string) ([]FilterDefinition, error)
	Factory(int64, string, string) (filter.Filter, error)
	Chains() ([]string, error)

	// Plugin specificatioon
	Author() string
	Version() string
	Name() string
}
