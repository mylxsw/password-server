package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/sethvargo/go-password/password"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	length := flag.Int("len", 0, "password length, only used if not set -sn")
	digitParam := flag.Int("digit", 0, "number of digits in the password, only used if not set -sn")
	symbolParam := flag.Int("symbol", 0, "number of symbols in the password, only used if not set -sn")
	serialNumberStyle := flag.Bool("sn", false, "output in serial number style")
	numberToGen := flag.Int("n", 1, "number of passwords to generate")
	syncClipboard := flag.Bool("clipboard", false, "copy password to clipboard")

	flag.Parse()

	if *numberToGen < 1 {
		*numberToGen = 1
	}

	if *serialNumberStyle {
		generator, err := password.NewGenerator(&password.GeneratorInput{Symbols: "-=.@#$:/+"})
		if err != nil {
			fmt.Println("Error creating password generator:", err)
			return
		}

		for i := 0; i < *numberToGen; i++ {
			pass := fmt.Sprintf("%s-%s-%s", generator.MustGenerate(4, 1, 0, false, true), generator.MustGenerate(4, 1, 0, true, true), generator.MustGenerate(3, 1, 0, false, true))
			fmt.Println(pass)

			if *syncClipboard && i == 0 {
				_ = clipboard.WriteAll(pass)
			}
		}

	} else {

		digitNumber := func() int {
			if *digitParam <= 0 {
				return rand.Intn(3) + 1
			}

			return *digitParam
		}

		symbolNumber := func() int {
			if *symbolParam <= 0 {
				return rand.Intn(2) + 1
			}

			return *symbolParam
		}

		lengthNumber := func() int {
			if *length < 6 {
				return 8 + rand.Intn(5)
			}

			return *length
		}

		for i := 0; i < *numberToGen; i++ {
			pass, err := password.Generate(lengthNumber(), digitNumber(), symbolNumber(), false, true)
			if err != nil {
				fmt.Println("Error generating password:", err)
				return
			}
			fmt.Println(pass)

			if *syncClipboard && i == 0 {
				_ = clipboard.WriteAll(pass)
			}
		}
	}
}
