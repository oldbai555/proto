package gen

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"github.com/oldbai555/proto/gen/temp"
	"github.com/oldbai555/proto/gen/temp/gtpl"
	"os"
	"strings"
)

// 生成 server 代码
func genServerFuncCode(protoFileName string, code *ParseGoCode, protoCtx *ProtoCtx) error {
	protoFileName = TrimProtoFileNameSuffix(protoFileName)
	var content strings.Builder

	// step 1: 新文件 需要初始化一下
	if code.IsNew {
		template, err := temp.GenCodeByTemplate(gtpl.GenRpcClientFuncHeaderCode, &temp.ServerFuncHeaderCode{
			Server: protoFileName,
		})
		if err != nil {
			log.Errorf("err is %v", err)
			return err
		}
		content.WriteString(template + "\n")

		template, err = temp.GenCodeByTemplate(gtpl.GenRpcServerStructCode, &temp.ServerStructCode{
			Server:     protoFileName,
			TypePrefix: utils.UpperFirst(protoFileName),
			Variable:   protoFileName,
		})
		if err != nil {
			log.Errorf("err is %v", err)
			return err
		}
		content.WriteString(template + "\n")
	}

	// step 2: 非新文件 判断是否有这个方法
	if !code.IsNew {
		_, ok := code.TypeDef[fmt.Sprintf("%sServer", utils.UpperFirst(protoFileName))]
		if !ok {
			template, err := temp.GenCodeByTemplate(gtpl.GenRpcServerStructCode, &temp.ServerStructCode{
				Server:     protoFileName,
				TypePrefix: utils.UpperFirst(protoFileName),
				Variable:   protoFileName,
			})
			if err != nil {
				log.Errorf("err is %v", err)
				return err
			}
			content.WriteString(template + "\n")
		}
	}

	// step 3: 遍历查看有哪个方法还没有
	for _, model := range protoCtx.RpcMap {
		if _, ok := code.FuncSvr[model.RpcName]; !ok {
			template, err := temp.GenCodeByTemplate(gtpl.GenRpcServerFuncCode, &temp.ServerFuncCode{
				RpcName: model.RpcName,
				RpcReq:  model.RpcReq,
				RpcRsp:  model.RpcRsp,
				Server:  model.Server,
				NewSev:  utils.UpperFirst(model.Server),
			})
			if err != nil {
				log.Errorf("err is %v", err)
				return err
			}
			content.WriteString(template + "\n")
		}
	}
	log.Infof("content is %s", content.String())

	// step 4: 打开文件
	f, err := os.OpenFile(code.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("can not generate file %s,Error :%v\n", code.Path, err)
		return err
	}
	defer f.Close()

	// step 5: 追加文件内容
	err = writeFile(f, content.String())
	if err != nil {
		log.Errorf("err is %v", err)
		return err
	}
	return nil
}
