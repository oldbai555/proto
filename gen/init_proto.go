package gen

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/proto/gen/temp"
	"github.com/oldbai555/proto/gen/temp/gtpl"
	"os"
	"strings"
)

func InitProto(protoName string) {
	code, err := temp.GenCodeByTemplate(gtpl.InitProto, &temp.InitProto{
		ServerName: protoName,
	})
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}

	code = strings.Replace(code, "    ", "\t", -1)
	err = os.WriteFile(fmt.Sprintf("%s.proto", strings.TrimSuffix(protoName, ".proto")), []byte(code), 0666)
	if err != nil {
		log.Errorf("write file err %s", err)
		return
	}

	log.Infof("gen success")
	return
}
