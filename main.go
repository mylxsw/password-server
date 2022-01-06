package main

import (
	"fmt"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/listener"
	"github.com/mylxsw/glacier/starter/application"
	"github.com/mylxsw/glacier/web"
	"github.com/sethvargo/go-password/password"
	"math/rand"
	"time"
)

var Version, GitRev string

func main() {
	rand.Seed(time.Now().Unix())

	app := application.Create(fmt.Sprintf("%s (%s)", Version, GitRev))
	app.AddStringFlag("listen", ":8080", "http listen address")

	app.Singleton(func() (*password.Generator, error) {
		return password.NewGenerator(&password.GeneratorInput{Symbols: "-=.@#$:/+"})
	})

	app.Provider(web.Provider(
		listener.FlagContext("listen"),
		web.SetRouteHandlerOption(func(cc infra.Resolver, router web.Router, mw web.RequestMiddleware) {
			router.Get("/-", func(ctx web.Context, gen *password.Generator) web.Response {
				pass := fmt.Sprintf("%s-%s-%s", gen.MustGenerate(4, 1, 0, false, true), gen.MustGenerate(4, 1, 0, true, true), gen.MustGenerate(3, 1, 0, false, true))
				return ctx.JSON(web.M{"password": pass})
			})
			router.Get("/", func(ctx web.Context, gen *password.Generator) web.Response {
				digitParam := ctx.IntInput("digit", 0)
				if digitParam < 0 {
					digitParam = rand.Intn(3)
					if digitParam <= 0 {
						digitParam = 1
					}
				}

				symbolParam := ctx.IntInput("symbol", 0)
				if symbolParam < 0 {
					symbolParam = rand.Intn(2)
					if symbolParam <= 0 {
						symbolParam = 1
					}
				}

				length := ctx.IntInput("length", 0)
				if length < 6 {
					length = 8 + rand.Intn(6)
				}

				return ctx.JSON(web.M{
					"password": gen.MustGenerate(length, digitParam, symbolParam, false, true),
				})
			})
		}),
	))

	application.MustRun(app)
}
