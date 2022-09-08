package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"git.sr.ht/~mendelmaleh/patterns"
	"github.com/kr/pretty"
)

func main() {
	// flags
	x := flag.Int("x", 3, "number of strings to generate")
	d := flag.Bool("d", false, "debug mode")
	flag.Parse()

	// args
	pattern := flag.Arg(0)
	if pattern == "" {
		log.Fatal("no pattern")
	}

	// generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	g, err := patterns.NewGenerator(pattern, r)
	if err != nil {
		log.Fatal(err)
	}

	if *d {
		pretty.Println(g.Tokens)
	}

	// generate
	for i := 0; i < *x; i++ {
		fmt.Println(g.Generate())
	}
}
