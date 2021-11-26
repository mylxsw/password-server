package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/sethvargo/go-password/password"
)

func main() {
	var listenAddr string

	flag.StringVar(&listenAddr, "listen", ":18921", "http listen address")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	gen, err := password.NewGenerator(&password.GeneratorInput{Symbols: "-=.@#$:/+"})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/-", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "text/html")
		rw.Write([]byte(fmt.Sprintf("%s-%s-%s", gen.MustGenerate(4, 1, 0, false, true), gen.MustGenerate(4, 1, 0, true, true), gen.MustGenerate(3, 1, 0, false, true))))
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		digitParam, err := strconv.Atoi(r.FormValue("digit"))
		if digitParam < 0 || err != nil {
			digitParam = rand.Intn(3)
			if digitParam <= 0 {
				digitParam = 1
			}
		}

		symbolParam, err := strconv.Atoi(r.FormValue("symbol"))
		if symbolParam < 0 || err != nil {
			symbolParam = rand.Intn(2)
			if symbolParam <= 0 {
				symbolParam = 1
			}
		}

		length, err := strconv.Atoi(r.FormValue("length"))
		if length < 6 || err != nil {
			length = 8 + rand.Intn(6)
		}

		rw.Header().Add("Content-Type", "text/html")
		rw.Write([]byte(gen.MustGenerate(length, digitParam, symbolParam, false, true)))
	})

	log.Fatal(http.ListenAndServe(listenAddr, nil))

}
