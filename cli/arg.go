package cli

import (
	"github.com/alexflint/go-arg"
	"fmt"
)

var args struct {
	A string `arg:"required"`
	B int
}


func main() {
	arg.MustParse(&args)
	fmt.Print(args.A, args.B)
}