// 指定proto版本
syntax = "proto3";

// 指定golang包名
option go_package = "github.com/oldbai555/bgg/{{ .ServerName }}";

// 指定默认包名
package {{ .ServerName }};

// 定义 {{ .ServerName }} 服务
service {{ .ServerName }} {
}

enum ErrCode {
	Success = 0;
}