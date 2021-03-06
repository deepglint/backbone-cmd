package controller

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func CreateVulcand(name string, component string, vulcandaddr string, localaddr string, p string) error {
	log.Println(os.Getenv("GOPATH"))
	var gopath string
	if gopath = os.Getenv("GOPATH"); gopath == "" {
		return errors.New("The env variable GOPATH is not existed")
	}

	var proj = path.Join(gopath, "src", "github.com/deepglint/backbone-cmd")
	//log.Println(proj)
	if _, err := os.Stat(proj); err != nil {

		return errors.New("The Project github.com/deepglint/backbone-cmd is not exist")
	} else {

		if err := execVulcandScript(name, component, vulcandaddr, localaddr, path.Join(proj, "template", "vulcand", "gen_to_vulcand.sh"), "./gen_to_vulcand.sh"); err != nil {

			return err
		}

	}
	return nil
}

func execVulcandScript(name string, component string, vulcandaddr string, localaddr string, source string, dest string) error {
	cmd := exec.Command("cp", "-rf", source, dest)
	if err := cmd.Run(); err != nil {
		log.Println("error for copy files")
		log.Println(err)
		return err
	}
	if err := replaceScript(name, component, vulcandaddr, localaddr, dest); err != nil {
		return err
	}
	return nil
}

func replaceScript(name string, component string, vulcand string, localaddr string, p string) error {
	input, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	output := strings.Replace(string(input), "{{component}}", component, -1)
	output = strings.Replace(string(output), "{{username}}", name, -1)
	output = strings.Replace(string(output), "{{vulcandaddr}}", vulcand, -1)
	output = strings.Replace(string(output), "{{localaddr}}", localaddr, -1)

	err = ioutil.WriteFile(p, []byte(output), 0777)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func CreateComponent(name string, p string) error {
	log.Println(os.Getenv("GOPATH"))
	var gopath string
	if gopath = os.Getenv("GOPATH"); gopath == "" {
		return errors.New("The env variable GOPATH is not existed")
	}
	if _, err := os.Stat(path.Join(p, name)); err == nil {
		return errors.New("The Folder has already existed")
	} else {
		var cmd = exec.Command("mkdir", path.Join(p, name))
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	var proj = path.Join(gopath, "src", "github.com/deepglint/backbone-cmd")
	log.Println(proj)
	if _, err := os.Stat(proj); err != nil {

		return errors.New("The Project github.com/deepglint/backbone-cmd is not exist")
	} else {
		if err := execFile(name, path.Join(proj, "template", "component", "component.js"), path.Join(p, name, name+".js")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component", "component.vue"), path.Join(p, name, name+".vue")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component", "component.go"), path.Join(p, name, name+".go")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component", "component.html"), path.Join(p, name, name+".html")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component", "component_test.vue"), path.Join(p, name, name+"_test.vue")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component", "webpack.config.js"), path.Join(p, name, "webpack.config.js")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component", "package.json"), path.Join(p, name, "package.json")); err != nil {

			return err
		}

	}
	return nil
}
func CreateComponent2(name string, p string) error {
	log.Println(os.Getenv("GOPATH"))
	var gopath string
	if gopath = os.Getenv("GOPATH"); gopath == "" {
		return errors.New("The env variable GOPATH is not existed")
	}
	if _, err := os.Stat(path.Join(p, name)); err == nil {
		return errors.New("The Folder has already existed")
	} else {
		var cmd = exec.Command("mkdir", path.Join(p, name))
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	var proj = path.Join(gopath, "src", "github.com/deepglint/backbone-cmd")
	log.Println(proj)
	if _, err := os.Stat(proj); err != nil {

		return errors.New("The Project github.com/deepglint/backbone-cmd is not exist")
	} else {
		if err := execFile(name, path.Join(proj, "template", "component2", "component.js"), path.Join(p, name, name+".js")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2", "component.vue"), path.Join(p, name, name+".vue")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2", "component.go"), path.Join(p, name, name+".go")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2", "component.html"), path.Join(p, name, name+".html")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2", "component_test.vue"), path.Join(p, name, name+"_test.vue")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2", "webpack.config.js"), path.Join(p, name, "webpack.config.js")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2", "package.json"), path.Join(p, name, "package.json")); err != nil {

			return err
		}

	}
	return nil
}
func CreateComponent2x(name string, p string) error {
	log.Println(os.Getenv("GOPATH"))
	var gopath string
	if gopath = os.Getenv("GOPATH"); gopath == "" {
		return errors.New("The env variable GOPATH is not existed")
	}
	if _, err := os.Stat(path.Join(p, name)); err == nil {
		return errors.New("The Folder has already existed")
	} else {
		var cmd = exec.Command("mkdir", path.Join(p, name))
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	var proj = path.Join(gopath, "src", "github.com/deepglint/backbone-cmd")
	log.Println(proj)
	if _, err := os.Stat(proj); err != nil {

		return errors.New("The Project github.com/deepglint/backbone-cmd is not exist")
	} else {
		if err := execFile(name, path.Join(proj, "template", "component2x", "component.js"), path.Join(p, name, name+".js")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2x", "component.vue"), path.Join(p, name, name+".vue")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2x", "component.go"), path.Join(p, name, name+".go")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2x", "component.html"), path.Join(p, name, name+".html")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2x", "component_test.vue"), path.Join(p, name, name+"_test.vue")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2x", "webpack.config.js"), path.Join(p, name, "webpack.config.js")); err != nil {

			return err
		}
		if err := execFile(name, path.Join(proj, "template", "component2x", "package.json"), path.Join(p, name, "package.json")); err != nil {

			return err
		}

	}
	return nil
}

func execFile(name string, source string, dest string) error {
	cmd := exec.Command("cp", "-rf", source, dest)
	if err := cmd.Run(); err != nil {
		log.Println("error for copy files")
		log.Println(err)
		return err
	}
	if err := replaceName(name, dest); err != nil {
		return err
	}
	return nil
}

func replaceName(name string, p string) error {
	input, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	output := strings.Replace(string(input), "{{component}}", name, -1)

	err = ioutil.WriteFile(p, []byte(output), 0777)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
