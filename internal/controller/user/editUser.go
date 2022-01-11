package user

//EditUser 编辑用户
//func EditUser(ctx *gin.Context) {
//	var param *model.User
//	err := ctx.ShouldBindJSON(&param)
//	if err != nil {
//		global.GBLog.Error("查询用户列表参数保定失败")
//		vs.SendParamParseError(ctx)
//		return
//	}
//	userList, count, err := userService.NewAppFormService().GetUserList(param.UserName, param.PageIndex, param.PageSize)
//	if err != nil {
//		global.GBLog.Error("查询用户列表失败：", zap.Error(err))
//		vs.SendBad(ctx, err)
//		return
//	}
//	res := vs.NewResData(param.PageIndex, param.PageSize, count, userList)
//	vs.SendOkData(ctx, res)
//}
//
