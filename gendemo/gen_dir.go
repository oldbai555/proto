package gendemo

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"github.com/oldbai555/proto/gendemo/temp"
	"github.com/oldbai555/proto/gendemo/temp/gtpl"
	"os"
	"path"
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
