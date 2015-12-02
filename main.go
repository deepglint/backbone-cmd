package main

import (
	"github.com/codegangsta/cli"
	"github.com/deepglint/backbone-cmd/controller"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	//"strconv"
	"time"
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
		}, {
			Name:  "git",
			Usage: "For github things",
			Subcommands: []cli.Command{
				{
					Name:   "sync",
					Usage:  "make a client to update",
					Action: gitSync,
				}, {
					Name:   "watch",
					Usage:  "always update the github repo",
					Action: gitWatch,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "port,p",
							Value: "8030",
							Usage: "The update port to listen",
						}, cli.StringFlag{
							Name:  "span,s",
							Value: "3",
							Usage: "Sleep time for every update",
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func gitSync(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Println("Error for arg number,please tell the address for remote backbone")
	}
	r, err := http.Get(ctx.Args()[0] + "/pull")
	if err != nil {
		println(err.Error())
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(b))
}

func gitWatch(ctx *cli.Context) {
	go func() {
		o := execCommand("git", []string{"pull"})
		//i, _ := strconv.ParseInt(ctx.String("span"))
		log.Println(string(o))
		time.Sleep(time.Second * time.Duration(60))
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/pull", handleUpdate)
	http.ListenAndServe("0.0.0.0:"+ctx.String("port"), mux)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	o := execCommand("git", []string{"pull"})
	w.Write(o)
}

func execCommand(name string, args []string) []byte {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		return []byte(err.Error())
	}
	return out
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
