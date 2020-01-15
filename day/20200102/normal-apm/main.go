package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	fmt.Println("normal-apm-go")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	var i int = 0
	for {

		go func(i *int) {
			fmt.Println("index :", *i)
			fmt.Println("i address", i)
		}(&i)

		i++
		if i%1000 == 0 {
			time.Sleep(time.Second * 10)
		}
		if i > 10000 {
			break
		}
	}

}
