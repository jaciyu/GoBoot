package lib
/* mongodb封装类，依赖mgo: http://labix.org/mgo, 如需要使用，请保证gopath路径中有mgo, 然后将此注释打开， 否则编译无法通过
import (
    "time"
    "fmt"
    "labix.org/v2/mgo/bson"
    "labix.org/v2/mgo"
)

var MongoSrv = &Mongo{}

type Mongo struct {
	server string
	timeout int
	session *mgo.Session
}
//init函数，整个程序生命周期只需要执行一次
func (this *Mongo) Init(server string, timeout int, mode int)  {
	//服务器
    this.server = server
    //超时时长
    this.timeout = timeout
    session, err := mgo.DialWithTimeout(this.server, time.Duration(this.timeout)*time.Millisecond)
    if err != nil {
       panic(err)
    }
    smode := mgo.Monotonic
    if mode == 0 {
    	smode = mgo.Eventual
    }else if mode == 2 {
    	smode = mgo.Strong
    }
    session.SetMode(smode, true)
    this.session = session
    fmt.Println("Mongo Init: ", this.server)
}

func (this *Mongo) FindOne(id int, db string, col string) bson.M {
    session := this.session.Copy()
    defer session.Close()
    c := session.DB(db).C(col)
    result := bson.M{}
    err := c.FindId(id).One(result)
    //fmt.Println(result);
    if err != nil && err != mgo.ErrNotFound{
        panic(ErrorInfo{ERR_MONGODB, err})
    } 
    return result
}
//其他Insert,Update,Upsert方法请参考http://godoc.org/gopkg.in/mgo.v2#Collection自行封装
*/
