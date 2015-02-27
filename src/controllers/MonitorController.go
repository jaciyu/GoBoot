package controllers

import (
	"GoBoot/src/lib"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

//监控专用Controller
type MonitorController struct {
	BaseController
}

//pprof监控
func (this *MonitorController) PprofAction() {
	http.DefaultServeMux.ServeHTTP(this.rw, this.r)
}

//用于lvs监控
func (this *MonitorController) StatusAction() {
	this.OutputString("ok\n")
}

//在这里实现自己的业务监控逻辑, 这里的代码只是一个示例
func (this *MonitorController) GetHealthAction() {
	ret, msg := this.getDemoMsg()
	REV := "FAILED"
	if ret {
		REV = "OK"
	}

	data := map[string]interface{}{
		"REV":        REV,
		"MAC":        lib.GetLocalIp(),
		"TOTAL_TIME": strconv.FormatInt(this.cost(), 10),
		"DATA":       msg,
	}
	format := this.GetString("format", "json")
	if format == "json" {
		this.Output(data)
	} else { //这里为了兼容只支持php序列化格式的监控系统
		datastr := lib.SerializePhp(data)
		this.OutputString(datastr)
	}
}

func (this *MonitorController) getDemoMsg() (ret bool, msg string) {
	defer func() {
		if err := recover(); err != nil {
			ret = false
			msg = fmt.Sprintf("%+v", err)
		}
	}()
	demoController := &DemoController{}
	demoController.Init(this.rw, this.r)
	demoController.OutputDirect = false
	id := "156565966"
	this.r.Form.Set("id", id)
	data := demoController.GetMsgAction()
	msg = data["data"].(string)
	if msg == "hello "+id {
		ret = true
	} else {
		ret = false
	}
	return
}
