func (a *{{.NewSev}}Server) {{.RpcName}}(ctx context.Context, req *{{.Server}}.{{.RpcReq}}) (*{{.Server}}.{{.RpcRsp}}, error) {
	var rsp {{.Server}}.{{.RpcRsp}}
	return &rsp, nil
}