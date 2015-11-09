package main

import (
	"github.com/codegangsta/cli"
	"github.com/deepglint/backbone-cmd/controller"
	"log"
	"os"
)

var (
	Version = "0.1.2"
)

func main() {
	app := cli.NewApp()
	app.Name = "backbone-cli"
	app.Usage = "The Backbone Framework Helper"
	app.Author = "YanHuang"
	app.Email = "yanhuang@deepglint.com"
	app.Commands = []cli.Command{
		{
			Name:  "component",
			Usage: "Components actions",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "Create a Vue Component",
					Action: componentCreate,
				},
			},
		}, {
			Name:  "vulcand",
			Usage: "Generate Vulcand Script",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "yourname componentname vulcandaddr localhostaddr, like :backbone vulcand create huangyan helloworld http://192.168.5.46:8182 http://192.168.1.24:8004",
					Action: vulcandCreate,
				},
			},
		},
	}
	app.Run(os.Args)
}

func componentCreate(ctx *cli.Context) {
	arg := ctx.Args()
	if len(arg) != 1 {
		log.Println("Please just tell the name of component")
		return
	}
	controller.CreateComponent(arg[0], "./")
}
func vulcandCreate(ctx *cli.Context) {
	arg := ctx.Args()
	if len(arg) != 4 {
		log.Println("Please tell the name of you and the name of component")
		return
	}
	controller.CreateVulcand(arg[0], arg[1], arg[2], arg[3], "./")

}
