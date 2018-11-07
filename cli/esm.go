package cli

import (
	"github.com/tucnak/climax"
	"fmt"
	"github.com/nabakirov/MethodsOfOptimization/core"
	"strconv"
)

var EsmCliCommand climax.Command = climax.Command{
	Name: "esm-min",
	Brief: "even search method min",
	Usage: `[-f=] your function goes here
			[-s=] step
			[-x=] x0
			[-k=] iteration limit`,
	Help: "esm-min",

	Flags: []climax.Flag{
		{
			Name: "function",
			Short: "f",
			Variable: true,
		},
		{
			Name: "step",
			Short: "s",
			Variable: true,
		},
		{
			Name: "x",
			Short: "x",
			Variable: true,
		},
		{
			Name: "k",
			Short: "k",
			Variable: true,
		},
	},

	Handle: func(ctx climax.Context) int {
		var f func(float64) float64
		var s, x float64
		var k int
		var err error
		if fstring, ok := ctx.Get("function"); ok {
			f = core.Parser(fstring)
		}
		if step, ok := ctx.Get("step"); ok {
			s, err = strconv.ParseFloat(step, 64)
			if err != nil {
				panic(err)
			}
		}
		if x0, ok := ctx.Get("x"); ok {
			x, err = strconv.ParseFloat(x0, 64)
			if err != nil {
				panic(err)
			}
		}
		if k_max, ok := ctx.Get("k"); ok {
			k, err = strconv.Atoi(k_max)
			if err != nil {
				panic(err)
			}
		}
		result := core.EvenSearchMethodMax(s, x, 0.1, k, f)

		fmt.Println("result ", result)


		return 0
	},
}