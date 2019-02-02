package option

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/dixonwille/wmenu.v4"
	"os"
	"strings"
)

type API interface {
	cli.Flag
	viper.FlagValue
}

type OptionSet struct {
	Key string
	Default string
	*wmenu.Opt
}

func NewOptionSet(key string, def string, opt *wmenu.Opt) *OptionSet {
	return &OptionSet{Key: key, Default: def, Opt: opt}
}

func  (o *OptionSet) String() string {
	return fmt.Sprint(o)
}

func (o *OptionSet) Apply(set *flag.FlagSet) {
	set.Var(o, o.Key, o.Text)
}

func (o *OptionSet) GetName() string {
	return o.Key
}

func (o *OptionSet) HasChanged() bool {
	if o.Default != o.Opt.Value {
		return true
	}
	return false
}

func (o *OptionSet) Name() string {
	return o.Key
}

func (o *OptionSet) ValueString() string {
	return o.Value.(string)
}

func (o *OptionSet) ValueType() string {
	return fmt.Sprintf("%T", o.Value)
}

func (o *OptionSet) Set(s string) error {
	viper.Set(o.Key, s)
	os.Setenv(strings.ToUpper(o.Key), s)
	if !viper.IsSet(o.Key) {
		return errors.New("failed to set value")
	}
	return nil
}


