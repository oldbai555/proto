package gen

import (
	"fmt"
	"github.com/emicklei/proto"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"github.com/oldbai555/proto/parse"
	"os"
	"path"
	"strings"
)

func ProtoError(protoFileName string, pathDir string) {
	// step 1: 判断文件是否存在
	if !strings.HasSuffix(parse.ProtoFileNameSuffix, protoFileName) {
		protoFileName = protoFileName + parse.ProtoFileNameSuffix
	}

	// step 2: 打开文件
	file, err := os.Open(protoFileName)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}

	// step 3: 解析go 文件
	definition, err := proto.NewParser(file).Parse()
	var fields []*proto.EnumField
	proto.Walk(definition,
		// step 3.1: 只解析 Enum
		proto.WithEnum(func(enum *proto.Enum) {
			if enum.Name != "ErrCode" {
				return
			}
			for _, ele := range enum.Elements {
				switch ele.(type) {
				case *proto.EnumField:
					fields = append(fields, ele.(*proto.EnumField))
				}
			}
		}),
	)

	// step 5: 构造文件内容
	var lines, camelLines []string
	lines = append(lines, "// Code generated by gen_errorcode.go, DO NOT EDIT.\n ")
	lines = append(lines, fmt.Sprintf("package %s", strings.TrimSuffix(protoFileName, ".proto")))
	lines = append(lines, "import (")
	lines = append(lines, "\"github.com/oldbai555/lbtool/pkg/lberr\"")
	lines = append(lines, ")")
	lines = append(lines, "var (")

	for _, v := range fields {
		var comment string
		if v.InlineComment != nil {
			comment = strings.Join(v.InlineComment.Lines, " ")
		}
		comment = strings.TrimSpace(comment)
		if comment == "" {
			comment = v.Name
		}
		camelLines = append(camelLines, fmt.Sprintf("\t%s = lberr.NewErr(int32(ErrCode_%s), \"%s\")", v.Name, v.Name, comment))
	}
	lines = append(lines, camelLines...)
	lines = append(lines, ")")

	// step 6: 构造文件路径
	var dir = path.Join(pathDir, strings.TrimSuffix(protoFileName, ".proto"))
	utils.CreateDir(dir)
	var fileName = fmt.Sprintf("%s_error_autogen.go", strings.TrimSuffix(protoFileName, ".proto"))

	// step 7: 创建并写入文件
	f, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(strings.Join(lines, "\n"))
}