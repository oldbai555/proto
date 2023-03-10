package gen

import (
	"fmt"
	"github.com/emicklei/proto"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"github.com/oldbai555/proto/gen/temp"
	"github.com/oldbai555/proto/gen/temp/gtpl"
	"github.com/oldbai555/proto/parse"
	"os"
	"path"
	"strings"
)

func ModelTableName(protoFileName string, pathDir string) {
	// step 0: 判断文件是否存在
	if !strings.HasSuffix(parse.ProtoFileNameSuffix, protoFileName) {
		protoFileName = protoFileName + parse.ProtoFileNameSuffix
	}
	// step 1: 打开proto 文件
	file, err := os.Open(protoFileName)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	// step 2: 解析proto 文件
	definition, err := proto.NewParser(file).Parse()
	var tableNameMap = make(map[string]*temp.Table)
	// step 3: 遍历解析后的proto 文件
	proto.Walk(definition,
		// step 3.1: 处理 message
		proto.WithMessage(func(message *proto.Message) {
			// step 3.1.1: 如果 message 包含前缀 Model , 表示为表结构, 特殊处理
			if strings.HasPrefix(message.Name, parse.MessagePrefixModel) {
				tableNameMap[message.Name] = &temp.Table{
					TableName:   parse.GenTableName(protoFileName, message.Name),
					MessageName: message.Name,
				}
			}
		}),
	)

	// step 4: 构造文件内容
	var lines strings.Builder
	lines.WriteString("// Code generated by gen_model_table.go, DO NOT EDIT.\n ")
	lines.WriteString("\n")
	lines.WriteString(fmt.Sprintf("package %s", parse.TrimProtoFileNameSuffix(protoFileName)))
	lines.WriteString("\n")
	for _, table := range tableNameMap {
		template, err := temp.GenCodeByTemplate(gtpl.GenModelTable, table)
		if err != nil {
			log.Errorf("err is %v", err)
			return
		}
		lines.WriteString(template)
		lines.WriteString("\n")
	}

	// step 5: 构造文件路径
	var dir = path.Join(pathDir, strings.TrimSuffix(protoFileName, ".proto"))

	utils.CreateDir(dir)
	var fileName = fmt.Sprintf("%s_table.go", strings.TrimSuffix(protoFileName, ".proto"))

	// step 6: 创建并写入文件
	f, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(lines.String())
}
