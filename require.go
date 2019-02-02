package require

import (
	"errors"
	"flag"
	"github.com/gofunct/require/api"
	"github.com/gofunct/require/decider"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
	"os"
)

type Initializer func(e *Enforcer)

type Enforcer struct {
	Name 				string
	Paths 				[]string
	Ext 				string
	EnvPrefix 			string
	Requirements 		[]Value
	dcdr 				*decider.Decider
	v 					*viper.Viper
}

func NewEnforcer(inits ...Initializer) *Enforcer {
	e := &Enforcer{}
	for _, i := range inits {
		i(e)
	}
	if e.Name == "" {
		e.Name = "require"
	}
	if len(e.Paths) == 0 {
		e.Paths = []string{".", os.Getenv("HOME"), "..", os.Getenv("REQUIRE_PATH")}
	}
	if e.v == nil {
		e.v = viper.New()
		_ = e.v.MergeInConfig()
	}
	if e.dcdr == nil {
		e.dcdr = decider.NewDecider()
	}
	if e.Ext == "" {
		e.Ext = "yaml"
	}
	return e
}

func(v Value) Type() {
	switch v.(type) {

	}
}

func NewHelmEnforcer(reqs ...Value) *Enforcer {
	e := &Enforcer{
		Name:      "values",
		Paths:     []string{"./helm", "helm", "deploy", "./deploy", os.Getenv("HOME")+"/helm", "../helm", os.Getenv("REQUIRE_HELM_PATH")},
		Ext:       "yaml",
		EnvPrefix: "helm",
		dcdr:     decider.NewDecider(),
		v:         viper.New(),
		Requirements: reqs,
	}
	return e
}

func (e *Enforcer)  Init() error {
	if len(e.Requirements)== 0 {
		return errors.New("no requirements were found")
	}
	for _, key := range e.Requirements {
		if !e.v.IsSet(key) || e.v.Get(key) == "" || e.v.Get(key) == nil {
			if val, exists := os.LookupEnv(key); val != "" && exists == true {
				e.v.Set(key, val)
			} else {
				typ, err := e.input.Select("Please provide a type for the following key: "+key, []string{"string", "int", "bool", "[]slice"}, &input.Options{
					Default: "string",
					Loop:         true,
					Required:     true,
					ValidateFunc: e.Ensure(),
				})
				ans, err := e.query.Ask("Please provide a value for the following key: "+key, &input.Options{
					Loop:         true,
					Required:     true,
					ValidateFunc: e.Ensure(),
				})
				if err != nil {
					panic(err)
				}
				e.v.Set(key, ans)
				_ = os.Setenv(key, ans)
			}
		}
	}

		}
	}
	return nil
}

func (e *Enforcer) Sub(key string) *Enforcer {
	return &Enforcer{
		Name:      key,
		Paths:     e.Paths,
		EnvPrefix: e.EnvPrefix,
		dcdr:     decider.NewDecider(),
		v:      	e.v.Sub(key),
	}
}

func (e *Enforcer) RequireBool(key string) {
	panic("implement me")
}

func (e *Enforcer) RequireInt(key string) {
	panic("implement me")
}

func (e *Enforcer) RequireStringSlice(key string) {
	panic("implement me")
}

func (e *Enforcer) RequireStringMapString(key string) {
	panic("implement me")
}

func (e *Enforcer) GetRequiredString(key string) string {
	panic("implement me")
}

func (e *Enforcer) GetRequiredBool(key string) string {
	panic("implement me")
}

func (e *Enforcer) GetRequiredStringSlice(key string) string {
	panic("implement me")
}

func (e *Enforcer) GetRequiredStringMapString(key string) string {
	panic("implement me")
}

func (e *Enforcer) RequireAll(i ...interface{}) {
	panic("implement me")
}

func (e *Enforcer) BindCobra(cmd *cobra.Command) {
	panic("implement me")
}

func (e *Enforcer) BindFlagSet(f *flag.FlagSet) {
	panic("implement me")
}

func (e *Enforcer) Debug() {
	panic("implement me")
}

func (e *Enforcer) UpdateConfigs() error {
	panic("implement me")
}

func (e *Enforcer) RequireString(key string) {
	if !e.v.IsSet(key) || e.v.GetString(key) == "" {
		if val, exists := os.LookupEnv(key); val != "" && exists == true {
			e.v.Set(key, val)
		} else {
			ans, err := e.query.Ask("Please provide a value for the following key: "+key, &input.Options{
				Loop:         true,
				Required:     true,
				ValidateFunc: e.Ensure(),
			})
			if err != nil {
				panic(err)
			}
			e.v.Set(key, ans)
			_ = os.Setenv(key, ans)
		}
	}
}

func (e *Enforcer) RequireDef(key, def string) {
	e.v.SetDefault(key, def)
	_ = os.Setenv(key, def)
}

func (e *Enforcer) GetString(key string) string {
	if !e.v.IsSet(key) || e.v.Get(key) == nil {
		if k, exists := os.LookupEnv(key); k == "" || exists == false {
			ans, err := e.query.Ask("Please provide a value for the following key: "+key, &input.Options{
				Loop:         true,
				Required:     true,
				ValidateFunc: e.Ensure(),
			})
			if err != nil {
				panic(err)
			}
			e.v.Set(key, ans)
			_ = os.Setenv(key, ans)
			return ans
		}
	}
	if k, exists := os.LookupEnv(key); k == "" || exists == false {
		ans, err := e.query.Ask("Please provide a value for the following key: "+key, &input.Options{
			Loop:         true,
			Required:     true,
			ValidateFunc: e.Ensure(),
		})
		if err != nil {
			panic(err)
		}
		e.v.Set(key, ans)
		_ = os.Setenv(key, ans)
		return ans
	}
	return e.v.GetString(key)
}

func (e *Enforcer) RequireKeys() {
	for _, key := range e.v.AllKeys() {
		if !e.v.IsSet(key) || e.v.Get(key) == "" {
			if val, exists := os.LookupEnv(key); val != "" && exists == true {
				e.v.Set(key, val)
			} else {
				ans, err := e.query.Ask("Please provide a value for the following key: "+key, &input.Options{
					Loop:         true,
					Required:     true,
					ValidateFunc: e.Ensure(),
				})
				if err != nil {
					panic(err)
				}
				e.v.Set(key, ans)
				_ = os.Setenv(key, ans)
			}
		}
	}
}

