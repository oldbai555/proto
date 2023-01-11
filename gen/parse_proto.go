package gen

import (
	"errors"
	"fmt"
	"github.com/emicklei/proto"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"io"
	"os"
	"strings"
)

const (
	ProtoFileNameSuffix       = ".proto"
	ProtoFileNamePrefixWithLb = "lb"
	MessagePrefixModel        = "Model"
)

// TrimProtoFileNameSuffix 去除proto文件的后缀 .proto
func TrimProtoFileNameSuffix(protoFileName string) string {
	return strings.TrimSuffix(protoFileName, ProtoFileNameSuffix)
}

// TrimProtoFileNamePrefixWithLb 去除proto文件 前缀 lb
func TrimProtoFileNamePrefixWithLb(protoFileName string) string {
	return strings.TrimPrefix(protoFileName, ProtoFileNamePrefixWithLb)
}

// TrimPrefixMessageNameWithModel 去除 Message 前缀 Model
func TrimPrefixMessageNameWithModel(msgName string) string {
	return strings.TrimPrefix(msgName, MessagePrefixModel)
}

// 生成表名
func genTableName(protoFileName, msgName string) string {
	return fmt.Sprintf("%s_%s", TrimProtoFileNameSuffix(protoFileName), utils.Camel2UnderScore(TrimPrefixMessageNameWithModel(msgName)))
}

type RpcNode struct {
	Server  string
	RpcName string
	RpcReq  string
	RpcRsp  string

	Rpc *proto.RPC
}

type ProtoCtx struct {
	Content string

	RpcMap  map[string]*RpcNode
	RpcList []*RpcNode
}

func parseProtoCode(protoFileName string) (*ProtoCtx, error) {
	var protoCtx = ProtoCtx{
		RpcMap: make(map[string]*RpcNode),
	}
	// step 1: 检查文件是否存在
	if !strings.HasSuffix(ProtoFileNameSuffix, protoFileName) {
		protoFileName = protoFileName + ProtoFileNameSuffix
	}

	// step 2: 打开文件
	reader, err := os.Open(protoFileName)
	if err != nil {
		log.Errorf("err is %v", err)
		return nil, err
	}

	// step 3: 解析go 文件
	definition, err := proto.NewParser(reader).Parse()
	proto.Walk(definition,
		// step 3.1: 只解析 rpc
		proto.WithRPC(func(rpc *proto.RPC) {
			r := &RpcNode{
				Server:  TrimProtoFileNameSuffix(protoFileName),
				RpcName: rpc.Name,
				RpcReq:  rpc.RequestType,
				RpcRsp:  rpc.ReturnsType,

				Rpc: rpc,
			}
			protoCtx.RpcMap[rpc.Name] = r
			protoCtx.RpcList = append(protoCtx.RpcList, r)
		}),
	)

	// step 4: 重置读文件的指针
	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		log.Errorf("err is %v", err)
		return nil, err
	}
	buf, err := io.ReadAll(reader)
	if err != nil {
		log.Errorf("read proto file error %v", err)
		return nil, err
	}
	protoCtx.Content = string(buf)

	return &protoCtx, nil
}

// InsertRpcBlock 插入方法
func (p *ProtoCtx) InsertRpcBlock(block string) error {
	if len(p.RpcList) > 0 {
		last := p.RpcList[len(p.RpcList)-1]
		b := last.Rpc.Position.Offset
		pos := -1
		for b < len(p.Content) {
			e := b
			for e < len(p.Content) && p.Content[e] != '\n' {
				e++
			}
			line := p.Content[b:e]
			if strings.TrimSpace(line) == "};" || strings.TrimSpace(line) == "}" {
				pos = e
				break
			}
			b = e + 1
		}
		if pos < 0 {
			return errors.New("invalid proto format, not found rpc block ending }")
		}
		left := p.Content[:pos]
		right := p.Content[pos:]
		p.Content = fmt.Sprintf("%s\n\n%s%s", left, block, right)
	} else {
		pos := strings.Index(p.Content, "service ")
		if pos < 0 {
			return errors.New("not found `service ` block")
		}
		b := pos
		for b < len(p.Content) {
			e := b
			for e < len(p.Content) && p.Content[e] != '\n' {
				e++
			}
			line := p.Content[b:e]
			if strings.TrimSpace(line) == "}" {
				// pos = e
				break
			}
			b = e + 1
		}
		left := p.Content[:b]
		right := p.Content[b:]
		block = fmt.Sprintf("\t%s\n", strings.TrimSpace(block))
		p.Content = fmt.Sprintf("%s%s%s", left, block, right)
	}
	return nil
}
