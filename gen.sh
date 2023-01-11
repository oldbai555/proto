#!/bin/sh

#使用说明，用来提示输入参数
usage() {
  echo "Usage: sh 执行脚本.sh [gf]"
  exit 1
}

get_model() {
  go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest

  go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

  go get -u github.com/envoyproxy/protoc-gen-validate@latest

  go get -u github.com/favadi/protoc-go-inject-tag@latest

  go install google.golang.org/protobuf/cmd/protoc-gen-go

  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

  go install github.com/envoyproxy/protoc-gen-validate

  go install github.com/favadi/protoc-go-inject-tag

  go mod tidy
}

gen_all() {
  for file in $(git ls-files '*.proto'); do
    protoc --go_out=../ --go-grpc_out=../ --gorm_out=../ ./"$file"
  done
}

gen_one() {
  protoc --go-grpc_out=../ --validate_out="lang=go:../" --go_out=../ ./"$1"

}

gen_faster() {
  protoc --go-grpc_out=../ --gogofaster_out=../ ./"$1"
}

#根据输入参数，选择执行对应方法，不输入则执行使用说明
case "$1" in
"gf")
  gen_faster "$2"
  ;;
"go")
  gen_one "$2"
  ;;
"gm")
  get_model
  ;;
*)
  usage
  ;;

esac
