package gendemo

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"github.com/oldbai555/proto/gendemo/temp"
	"github.com/oldbai555/proto/gendemo/temp/gtpl"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

const projectDir = "../github.com/oldbai555/bgg/server"

func GenDir(serverName string) {
	// 创建目录 github.com/oldbai555/bgg/server/serverName
	var proDir = fmt.Sprintf("%s/%s", projectDir, serverName)
	utils.CreateDir(proDir)
	// 创建子目录
	utils.CreateDir(path.Join(proDir, "cmd"))
	utils.CreateDir(path.Join(proDir, "impl"))
	// 判断文件是否存在
	if !utils.FileExists(path.Join(proDir, "cmd", "main.go")) {
		template, err := temp.GenCodeByTemplate(gtpl.GenTempMain, &temp.TempMain{
			ServerName: serverName,
		})
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}

		// step 7: 创建并写入文件
		f, err := os.Create(path.Join(proDir, "cmd", "main.go"))
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}
		defer f.Close()
		_, err = f.WriteString(template)
	}

	if !utils.FileExists(path.Join(proDir, "cmd", "application.yaml")) {
		template, err := temp.GenCodeByTemplate(gtpl.GenTempYaml, &temp.TempYaml{
			ServerName: serverName,
		})
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}

		// step 7: 创建并写入文件
		f, err := os.Create(path.Join(proDir, "cmd", "application.yaml"))
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}
		defer f.Close()
		_, err = f.WriteString(template)
	}

	if !utils.FileExists(path.Join(proDir, "impl", "sys_cmd.go")) {
		template, err := temp.GenCodeByTemplate(gtpl.GenImplSysCmd, &temp.TempImplSysCmd{
			ServerName:      serverName,
			UpperServerName: utils.UpperFirst(serverName),
		})
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}

		// step 7: 创建并写入文件
		f, err := os.Create(path.Join(proDir, "impl", "sys_cmd.go"))
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}
		defer f.Close()
		_, err = f.WriteString(template)
	}

	if !utils.FileExists(path.Join(proDir, "impl", "db.go")) {
		template, err := temp.GenCodeByTemplate(gtpl.GEnImplDb, &temp.TempYaml{
			ServerName: serverName})
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}

		// step 7: 创建并写入文件
		f, err := os.Create(path.Join(proDir, "impl", "db.go"))
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}
		defer f.Close()
		_, err = f.WriteString(template)
	}

	if !utils.FileExists(path.Join(proDir, "impl", "util_orm.go")) {
		template, err := temp.GenCodeByTemplate(gtpl.GEnImplUtilOrm, &temp.TempYaml{
			ServerName: serverName})
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}

		// step 7: 创建并写入文件
		f, err := os.Create(path.Join(proDir, "impl", "util_orm.go"))
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}
		defer f.Close()
		_, err = f.WriteString(template)
	}

	if !utils.FileExists(path.Join(proDir, "impl", "const.go")) {
		template, err := temp.GenCodeByTemplate(gtpl.GEnImplConst, &temp.TempYaml{
			ServerName: serverName})
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}

		// step 7: 创建并写入文件
		f, err := os.Create(path.Join(proDir, "impl", "const.go"))
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}
		defer f.Close()
		_, err = f.WriteString(template)
	}
}

func Exec(rootDir, name string, args ...string) {
	_ = os.Setenv("GOSUMDB", "off")
	_ = os.Setenv("GOPROXY", "https://goproxy.cn,https://admin:pinkuai1228@goproxy.aquanliang.com,direct")

	cmd := exec.Command(name, args...)
	cmd.Dir = rootDir
	cmd.Env = os.Environ()
	log.Infof(strings.Join(cmd.Args, " "))
	log.Infof("run at:%s", cmd.Dir)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		}
	}
}

func GoFmt(path string) {
	var w sync.WaitGroup
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d == nil {
			return nil
		}

		if !d.IsDir() {
			return nil
		}

		w.Add(1)
		go func(path string) {
			defer w.Done()
			cmd := exec.Command("gofmt", "-w", ".")
			cmd.Dir = path
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			log.Infof("run %s", strings.Join(cmd.Args, " "))
			log.Infof("run on %s", path)

			err = cmd.Run()
			if err != nil {
				log.Errorf("err:%v", err)
				return
			}
		}(path)

		return nil
	})
	if err != nil {
		log.Errorf("err:%v", err)
	}

	w.Wait()
}
