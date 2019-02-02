package option

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
	"gopkg.in/dixonwille/wmenu.v4"
	"time"
)

type API interface {
	viper.FlagValue
}

type Option struct {
	Pointer interface{}
	ID int
	Key string
	Env string
	Usage string
	Required bool
	Default string
	menu *wmenu.Opt
	clif cli.Flag
}

func (o *Option) HasChanged() bool {
	panic("implement me")
}

func (o *Option) Name() string {
	return o.Key
}

func (o *Option) ValueString() string {
	return o.clif.String()
}

func (o *Option) ValueType() string {
	return fmt.Sprintf("%T", o.Pointer)
}

func NewOption(dest interface{}, id int, name, env,  def, usage string, required bool) *Option {
	viper.SetDefault(name, def)
	viper.BindEnv(name, def)
	o := &Option{
		ID:       id,
		Key:     name,
		Env:      env,
		Usage:    usage,
		Required: required,
		Default: def,
		menu:     &wmenu.Opt{
			ID:    id,
			Text:  usage,
			Value: dest,
		},
	}
	switch dest.(type) {
	case (*bool):
		o.clif = &cli.BoolFlag{
			Name:        name,
			Usage:       usage,
			EnvVar:      env,
			Destination: dest.(*bool),
		}
	case (*string):
		o.clif = &cli.StringFlag{
			Name:        name,
			Usage:       usage,
			EnvVar:      env,
			Value:       def,
			Destination: dest.(*string),
		}
	case (time.Duration):
		o.clif = &cli.DurationFlag{
			Name:   name,
			Usage:  usage,
			EnvVar: env,
			Value: dest.(time.Duration),
		}
	case (*int):
		o.clif = &cli.IntFlag{
			Name:   name,
			Usage:  usage,
			EnvVar: env,
			Value: dest.(int),
		}
	}
	return o
}
