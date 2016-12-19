package controller

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// 是否存在目录
func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("not reached")
}

// 复制某个文件夹的文件到另一个文件夹
func CopyDir(srcDir, destDir string) {
	if isDirExists(srcDir) {
		//存在附件文件夹
		tmpSrc := strings.TrimSpace(srcDir)
		files, _ := ioutil.ReadDir(tmpSrc)
		var name string
		for _, f := range files {
			name = f.Name()
			// 如果是目录
			if f.IsDir() {
				dir := path.Join(destDir, name)
				// 如果目标目录不存在
				if !isDirExists(dir) {
					// 创建目录
					err := os.Mkdir(dir, os.ModePerm)
					if err != nil {
						log.Println(err)
					}
				}
				CopyDir(path.Join(srcDir, name), path.Join(destDir, name))
			} else {
				CopyFile(path.Join(srcDir, name), path.Join(destDir, name))
			}
		}
	}
}

// 复制文件到指定地方
func CopyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()
	return io.Copy(desFile, srcFile)
}

// 校验路径是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取backbone-cmd的路径
func getBackboneCmdPath() (string, error) {
	var gopath string
	if gopath = os.Getenv("GOPATH"); gopath == "" {
		return "", errors.New("The env variable GOPATH is not existed")
	}

	var cmdPath = path.Join(gopath, "src", "github.com/deepglint/backbone-cmd")

	return cmdPath, nil
}

// 脚手架初始化
func InitScaffold() error {
	cmdPath, err := getBackboneCmdPath()
	log.Println(cmdPath, err)
	if err != nil {
		return err
	}
	log.Println(cmdPath)
	// 获取pwd
	pwd, _ := os.Getwd()

	if err = ScaffoldTarget(pwd, cmdPath); err != nil {
		return err
	}
	return nil
}

/**
 * 通过name创建工程
 */
func InitScaffoldByName(name string) error {
	cmdPath, err := getBackboneCmdPath()
	if err != nil {
		return err
	}
	log.Println(cmdPath)
	// 获取pwd
	pwd, _ := os.Getwd()
	// 目录地址
	sdir := path.Join(pwd, name)

	log.Println(sdir, "sdir")
	// 检测文件夹是否存在
	dirExist, _ := pathExists(sdir)
	if dirExist {
		log.Println("文件夹已存在")
		return errors.New("文件夹已存在")
	}
	// 创建目录
	err = os.Mkdir(sdir, os.ModePerm)
	if err != nil {
		log.Println(err, "err")
		return err
	}
	log.Println(dirExist, "dirExist")

	if err = ScaffoldTarget(sdir, cmdPath); err != nil {
		return err
	}

	return nil
}

/**
 * ReplaceName，替换目标文件的关键词
 * file string 文件地址
 * replaceName string 被替换的关键字
 * name string 目标文字
 */
func ReplaceName(file string, replaceName string, name string) error {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	output := strings.Replace(string(input), replaceName, name, -1)

	err = ioutil.WriteFile(file, []byte(output), 0777)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

/**
 * ScaffoldTarget
 * 拷贝内容到指定目录
 * sdir string 目标目录
 * cmdPath string 命令行包目录位置
 */
func ScaffoldTarget(sdir string, cmdPath string) error {

	projectDir := path.Join(cmdPath, "template", "demo_project", "./")
	log.Println(projectDir)
	// 复制文件夹
	CopyDir(projectDir, sdir)

	_, name := path.Split(sdir)

	log.Println(name)
	// 替换名字
	err := ReplaceName(path.Join(sdir, "conf", "app.conf"), "{{project_name}}", name)
	if err != nil {
		return err
	}
	return nil
}

/**
 * ScaffoldTarget
 * 拷贝内容到指定目录
 * sdir string 目标目录
 * cmdPath string 命令行包目录位置
 */
func ScaffoldBackbone(sdir string, cmdPath string) error {

	projectDir := path.Join(cmdPath, "template", "backbone_project", "./")
	log.Println(projectDir)
	// 复制文件夹
	CopyDir(projectDir, sdir)

	_, name := path.Split(sdir)

	log.Println(name)
	// 替换名字
	err := ReplaceName(path.Join(sdir, "conf", "app.conf"), "{{project_name}}", name)
	if err != nil {
		return err
	}
	return nil
}

func InitBackbone() error {
	cmdPath, err := getBackboneCmdPath()
	log.Println(cmdPath, err)
	if err != nil {
		return err
	}
	log.Println(cmdPath)
	// 获取pwd
	pwd, _ := os.Getwd()

	if err = ScaffoldBackbone(pwd, cmdPath); err != nil {
		return err
	}
	return nil
}

/**
 * 通过name创建工程
 */
func InitBackboneByName(name string) error {
	cmdPath, err := getBackboneCmdPath()
	if err != nil {
		return err
	}
	log.Println(cmdPath)
	// 获取pwd
	pwd, _ := os.Getwd()
	// 目录地址
	sdir := path.Join(pwd, name)

	log.Println(sdir, "sdir")
	// 检测文件夹是否存在
	dirExist, _ := pathExists(sdir)
	if dirExist {
		log.Println("文件夹已存在")
		return errors.New("文件夹已存在")
	}
	// 创建目录
	err = os.Mkdir(sdir, os.ModePerm)
	if err != nil {
		log.Println(err, "err")
		return err
	}
	log.Println(dirExist, "dirExist")

	if err = ScaffoldBackbone(sdir, cmdPath); err != nil {
		return err
	}

	return nil
}
