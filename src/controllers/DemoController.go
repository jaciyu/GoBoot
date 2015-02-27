package controllers

import (
	"GoBoot/src/lib"
	"GoBoot/src/logic"
)

//示例Controller
type DemoController struct {
	BaseController
}

func (this *DemoController) GetMsgAction() map[string]interface{} {
	id := this.GetInt("id", 0)
	if id <= 0 {
		panic(lib.ERR_INPUT) //有错误直接抛出异常，HttpApiServer和BaseController会统一处理异常
	}
	demoLogic := logic.NewDemoLogic()
	data := demoLogic.GetMsg(int(id))
	ret := map[string]interface{}{
		"errno":  0,
		"errmsg": "",
		"data":   data,
	}
	this.Output(ret)
	return ret //通过函数返回时为了兼容从MonitorController发起的调用
}
