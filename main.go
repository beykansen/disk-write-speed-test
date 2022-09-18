package main

import (
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
	"github.com/beykansen/disk-write-speed-test/pkg"
)

func main() {
	args := new(pkg.ProgramArguments)
	arg.MustParse(args)

	speedResult, err := pkg.Run(args)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(speedResult)
}
