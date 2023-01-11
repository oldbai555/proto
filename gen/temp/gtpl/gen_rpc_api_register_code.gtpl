func register{{.VariablePrefix}}Api(h *gin.Engine) {
	// 可以利用反射来映射函数进去
	group := h.Group("{{.Path}}").Use(RegisterJwt())
}