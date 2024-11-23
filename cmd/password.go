package main

import (
	"flag"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	length := flag.Int("len", 0, "password length")
	digitParam := flag.Int("digit", 0, "number of digits in the password")
	symbolParam := flag.Int("symbol", 0, "number of symbols in the password")
	serialNumberStyle := flag.Bool("sn-style", false, "output in serial number style")

	flag.Parse()

	if *digitParam < 0 {
		*digitParam = rand.Intn(3) + 1
	}

	if *symbolParam < 0 {
		*symbolParam = rand.Intn(3) + 1
	}

	if *length < 6 {
		*length = 8 + rand.Intn(5)
	}

	if *serialNumberStyle {
		generator, err := password.NewGenerator(&password.GeneratorInput{Symbols: "-=.@#$:/+"})
		if err != nil {
			fmt.Println("Error creating password generator:", err)
			return
		}
		pass := fmt.Sprintf("%s-%s-%s", generator.MustGenerate(4, 1, 0, false, true), generator.MustGenerate(4, 1, 0, true, true), generator.MustGenerate(3, 1, 0, false, true))
		fmt.Println(pass)
	} else {
		pass, err := password.Generate(*length, *digitParam, *symbolParam, false, true)
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}
		fmt.Println(pass)
	}
}
