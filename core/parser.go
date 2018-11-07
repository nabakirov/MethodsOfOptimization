package core

import "github.com/Knetic/govaluate"


func Parser(fstring string) func(float64) float64 {
	expression, err := govaluate.NewEvaluableExpression(fstring)
	if err != nil {
		panic(err)
	}
	params := make(map[string]interface{}, 1)
	return func(x float64) float64 {
		
		params["x"] = x
		result, err := expression.Evaluate(params)
		if err != nil {
			panic(err)
		}
		resultf := result.(float64)
		return resultf
	}
}