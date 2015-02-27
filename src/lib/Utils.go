package lib

import (
    "net"
    "fmt"
)
//获取本机ip
func GetLocalIp()string{
	addrs, _ := net.InterfaceAddrs()
	var ip string = ""
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			ip = ipnet.IP.String()
			if ip != "127.0.0.1" {
			}
		}
	}
	return ip
}

//简单序列化成php的格式, 和已有的php系统交互的时候可能会用到
func SerializePhp(data map[string]interface{}) string{
	ret := fmt.Sprintf("a:%d:{", len(data))
	for key, value := range data {
		ret = ret + fmt.Sprintf("s:%d:\"%s\";", len(key), key)
		if valuemap, ok := value.(map[string]interface{}); ok{
			ret = ret + SerializePhp(valuemap)
		}else{
			valuestr := value.(string)
			ret = ret + fmt.Sprintf("s:%d:\"%s\";", len(valuestr), valuestr)
		}
	}
	ret = ret + "}"
	return ret
}