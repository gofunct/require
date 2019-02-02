package decider

import (
	"fmt"
	"github.com/gofunct/gofs"
	"github.com/tcnksm/go-input"
	"gopkg.in/dixonwille/wmenu.v4"
	"strconv"
	"strings"
)

type Decider struct {
	ask  *input.UI
	menu *wmenu.Menu
}

func NewDecider(q string) *Decider {
	return &Decider{
		ask: input.DefaultUI(),
		menu: wmenu.NewMenu(q),
	}
}

func (d *Decider) AskString(q string, def string, required bool) string {
	ans, err := d.ask.Ask(q, &input.Options{
		Default:      def,
		Loop:         required,
		Required:     required,
		ValidateFunc: Ensure(required),
	})
	if err != nil {
		panic(err)
	}
	return ans
}

func (d *Decider) AskInt(q string, def string, required bool) int {
	ans, err := d.ask.Ask(q, &input.Options{
		Default:      def,
		Loop:         required,
		Required:     required,
		ValidateFunc: Ensure(required),
	})
	if err != nil {
		panic(err)
	}
	intans, err := strconv.Atoi(ans)
	if err != nil {
		panic(err)
	}
	return intans
}

func (d *Decider) AskStringSlice(q string, def string, required bool) []string {
	ans, err := d.ask.Ask(q, &input.Options{
		Default:      def,
		Loop:         required,
		Required:     required,
		ValidateFunc: EnsureSlice(required),
	})
	if err != nil {
		panic(err)
	}
	sliceans, err := gofs.ReadAsCSV(ans)
	if err != nil {
		panic(err)
	}

	return sliceans
}

func (d *Decider) AskStringMapString(q string, def string, required bool) map[string]string {
	ans, err := d.ask.Ask(q, &input.Options{
		Default:      def,
		Loop:         required,
		Required:     required,
		ValidateFunc: EnsureSlice(required),
	})
	if err != nil {
		panic(err)
	}
	newMap, err := gofs.ReadAsMap(ans)
	if err != nil {
		panic(err)
	}
	return newMap
}


func (d *Decider) AskYn(q string, def int) bool {
	q = "y/n | "+q
	var ans bool
	var errrr error
	d.menu = wmenu.NewMenu(q)
	actFunc := func(opts []wmenu.Opt) error {
		if strings.Contains(opts[0].Text, "y") || strings.Contains(opts[0].Text, "Y") {
			ans = true
		} else {
			ans = false
		}
		fmt.Printf("registered response of %s.\n", opts[0].Value.(string))
		return nil
	}
	d.menu.Action(actFunc)
	d.menu.IsYesNo(def)
	d.menu.LoopOnInvalid()
	if errrr != nil {
		panic(errrr)
	}
	return ans
}


func (d *Decider) AskTF(q string, def string) bool {
	var ans bool
	var errrr error
	q = "t/f | "+q
	d.menu = wmenu.NewMenu(q)
	d.menu.Option("true", 0, false, func(opt wmenu.Opt) error {

		return nil
	})
	actFunc := func(opts []wmenu.Opt) error {
		if len(opts[0].Text) > 1 || len(def) > 1{
			fmt.Printf("value must be either t or f")
		}

		if strings.Contains(opts[0].Text, "t") || strings.Contains(opts[0].Text, "T") {
			ans = true
		} else {
			ans = false
		}
		fmt.Printf("registered response of %s.\n", opts[0].Value.(string))
		return nil
	}
	d.menu.Action(actFunc)
	d.menu.LoopOnInvalid()
	if errrr != nil {
		panic(errrr)
	}
	return ans
}