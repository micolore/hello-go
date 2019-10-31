package def

import (
	"flag"
	"fmt"
	"strings"
)
var n = flag.Bool("n",false,"omit trailing new line")

var sep = flag.String("s"," " , "separtor")


func Echo1(){

	flag.Parse()

	fmt.Println(strings.Join(flag.Args(),*sep))

	if !*n {
		fmt.Println()
	}


}
