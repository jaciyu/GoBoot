package logic

import (
	"strconv"
)

type DemoLogic struct {
	msg string
}

func NewDemoLogic() *DemoLogic{
	ret := &DemoLogic{
		msg : "hello ",
		}
	return ret	
}

func (this *DemoLogic) GetMsg(id int) string {
	//这里也可以调用models层获取数据
    return this.msg + strconv.Itoa(id)
}
