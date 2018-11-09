package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/nabakirov/MethodsOfOptimization/core"
)

var esmArgs struct {
	Step float64 `arg:"required,-s,help:step size of iterations"`
	X0 float64 `arg:"required,-x,help:x0 value"`
	Tollerance float64 `arg:"-t,help:tollerance"`
	K_Max int `arg:"required,-k,help:max number of iterations"`
	Func string `arg:"required,-f,help:function"`


}


func ESM() {
	arg.MustParse(&esmArgs)
	f := core.Parse(esmArgs.Func)
	result := core.EvenSearchMethodMin(esmArgs.Step, esmArgs.X0, esmArgs.Tollerance, esmArgs.K_Max, f)

	fmt.Println(result)

}

func main() {
	ESM()
}