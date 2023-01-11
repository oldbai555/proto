func {{.RpcName}}(c *gin.Context) {
	var req {{.Server}}.{{.RpcReq}}
	handler := NewHandler(c)

	err := handler.BindAndValidateReq(&req)
	if err != nil {
		log.Errorf("err is : %v", err)
		handler.RespError(err)
		return
	}

	if err = req.ValidateAll(); err != nil {
		log.Errorf("err is : %v", err)
		handler.RespError(err)
		return
	}

	rsp, err := {{.Server}}Server.{{.RpcName}}(c, &req)
	if err != nil {
		log.Errorf("err is : %v", err)
		handler.RespError(err)
		return
	}
	handler.Success(rsp)
}