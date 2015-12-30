package main

import (
	"github.com/codegangsta/cli"
	"github.com/deepglint/backbone-cmd/controller"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	//"strconv"
	"time"
)

var (
	Version = "0.1.3"
)

///
///
///
///
///
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

///
///The Git Watch Helper
///Using /pull to pull the code from repo
///Using /release?tag=... to add the tag and release the code
///
func gitWatch(ctx *cli.Context) {
	go func() {
		for {
			o := execCommand("git", []string{"pull"})
			//i, _ := strconv.ParseInt(ctx.String("span"))
			log.Println(string(o))
			time.Sleep(time.Second * time.Duration(10))
			continue
		}
	}()
	fs := http.FileServer(http.Dir("../build/"))
	//http.Handle("/", fs)
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("/pull", handleUpdate)
	mux.HandleFunc("/build", handleTag)
	mux.HandleFunc("/release", handleRelease)
	mux.HandleFunc("/ls", handleLs)
	http.ListenAndServe("0.0.0.0:"+ctx.String("port"), mux)
}

func handleLs(w http.ResponseWriter, r *http.Request) {
	o := execCommand("ls", []string{"../build/"})
	w.Write(o)
}

func handleRelease(w http.ResponseWriter, r *http.Request) {
	var t = r.URL.Query().Get("tag")
	println("Getting Tag :" + t)
	attach := "../build/backbone_" + t + ".tar.gz"
	// o := execCommand("git", []string{"tag", t})
	// o = execCommand("git", []string{"push", "--tag"})
	//o := execCommand("bash", []string{"build.sh", t})
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")
	sub := "Backbone " + t + " Release"
	c, _ := ioutil.ReadFile("release-note.html")
	//sendMail("Backbone "+t+" Release", string(c), "../build/backbone_"+t+".tar.gz", user, pass)
	sendMail2(sub, string(c), attach, user, pass)
	w.Write([]byte("done"))
}

func handleTag(w http.ResponseWriter, r *http.Request) {
	var t = r.URL.Query().Get("tag")
	println("Getting Tag :" + t)
	// o := execCommand("git", []string{"tag", t})
	// o = execCommand("git", []string{"push", "--tag"})
	o := execCommand("bash", []string{"build.sh", t})
	w.Write(o)
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
func sendMail2(sub string, content string, attach string, user string, pass string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "backbone@deepglint.com")
	m.SetHeader("To", "yanhuang@deepglint.com")
	//m.SetHeader("To", "weiranyuan@deepglint.com", "zhenyuchen@deepglint.com", "yunhou@deepglint.com", "huiyanliu@deepglint.com", "yanzhang@deepglint.com")
	//m.SetAddressHeader("Cc", "libra@deepglint.com")
	m.SetHeader("Subject", sub)
	m.SetBody("text/html", content)
	m.Attach(attach)

	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, user, pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
func sendMail(sub string, content string, attach string, user string, pass string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "backbone@deepglint.com")
	//m.SetHeader("To", "yanhuang@deepglint.com")
	//m.SetHeader("To", "weiranyuan@deepglint.com", "zhenyuchen@deepglint.com", "yunhou@deepglint.com", "huiyanliu@deepglint.com", "yanzhang@deepglint.com")
	//m.SetAddressHeader("Cc", "libra@deepglint")
	m.SetHeader("Subject", sub)
	m.SetBody("text/html", content)
	m.Attach(attach)
	log.Println(sub, content, attach, user, pass)
	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, user, pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
