package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/instrumenta/kubeval/kubeval"
	"github.com/k0kubun/pp"
)

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cnf := kubeval.NewDefaultConfig()
	cnf.Strict = true
	results, err := kubeval.Validate(buf, cnf)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(results)
}
