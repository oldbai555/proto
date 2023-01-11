package gen

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
)

var nullStruct struct{}

type ParseGoCode struct {
	IsNew bool
	Path  string

	SvrDef  map[string]struct{}
	TypeDef map[string]struct{}

	FuncSvr map[string]struct{}
}

func NewParseGoCode() *ParseGoCode {
	return &ParseGoCode{
		SvrDef:  make(map[string]struct{}),
		TypeDef: make(map[string]struct{}),

		FuncSvr: make(map[string]struct{}),
	}
}

// 解析go文件
func parserGoCode(dir string, goFileName string) (*ParseGoCode, error) {
	// step 1: 拿到go 文件的 ast 语法抽象树
	g := NewParseGoCode()
	fSet := token.NewFileSet()

	// step 2: 创建文件目录
	utils.CreateDir(dir)

	// step 3: 拼文件路径
	g.Path = path.Join(dir, goFileName)

	// step 4: 判断文件是否存在
	if !fileExists(g.Path) {
		file, err := os.Create(g.Path)
		if err != nil {
			log.Errorf("err is %v", err)
			return nil, err
		}
		defer file.Close()
		// step 4.1:标记是新建的文件
		g.IsNew = true
		// step 4.2:写包名
		content := "package impl\n"
		err = writeFile(file, content)
		if err != nil {
			log.Errorf("err is %v", err)
			return nil, err
		}
	}
	// step 5: 解析go 文件,得到抽象结构体
	f, err := parser.ParseFile(fSet, g.Path, nil, 0)
	if err != nil {
		log.Errorf("err is %v", err)
		return nil, err
	}

	// step 6: 遍历go 文件结构
	for _, decl := range f.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			g.FuncSvr[t.Name.Name] = nullStruct
		case *ast.GenDecl:
			for _, spec := range t.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println("Import", spec.Path.Value)
				case *ast.TypeSpec:
					fmt.Println("Type", spec.Name.String())
					g.TypeDef[spec.Name.String()] = nullStruct
				case *ast.ValueSpec:
					for _, id := range spec.Names {
						g.SvrDef[id.Name] = nullStruct
					}
				default:
				}
			}
		}
	}
	return g, nil
}

// 文件是否存在
func fileExists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// 写内容进入文件
func writeFile(file *os.File, content string) error {
	if _, err := file.Write([]byte(content)); err != nil {
		return err
	}
	return nil
}
