package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"git.sr.ht/~mendelmaleh/go-prompt"
	"git.sr.ht/~mendelmaleh/patterns"
	"github.com/kr/pretty"
)

func main() {
	// flags
	d := flag.Bool("d", false, "debug mode")
	i := flag.Bool("i", false, "interactive mode")
	l := flag.Int("l", 128, "upper limit for quantifiers like * and +")
	x := flag.Int("x", 3, "number of strings to generate")
	flag.Parse()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if *i {
		prompt.New(
			func(s string) {
				if s == "" {
					return
				}

				g, err := patterns.NewGenerator(s, r)
				if err != nil {
					log.Println(err)
					return
				}

				g.Max = *l
				fmt.Println(g.Generate())
			},
			func(_ prompt.Document) (s []prompt.Suggest) { return },
		).Run()
	} else {
		pattern := flag.Arg(0)
		if pattern == "" {
			log.Fatal("no pattern")
		}

		g, err := patterns.NewGenerator(pattern, r)
		if err != nil {
			log.Fatal(err)
		}

		if *d {
			pretty.Println(g.Tokens)
		}

		g.Max = *l
		for i := 0; i < *x; i++ {
			fmt.Println(g.Generate())
		}
	}
}
