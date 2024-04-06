package models

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"sync"

// 	"github.com/gorilla/websocket"
// 	"gopkg.in/fatih/set.v0"
// 	"gorm.io/gorm"
// )

// // 消息
// type Message struct {
// 	gorm.Model
// 	FromId   string //发送者
// 	TargetId string // 接收者
// 	Type     string //消息类型  群聊  私聊  广播
// 	Media    int    //消息类型  文字 图片  音频
// 	Content  string //消息内容
// 	Pic      string
// 	Url      string
// 	Desc     string
// 	Amount   int //其他数字统计
// }

// func (table *Message) TableName() string {
// 	return "message"
// }

// type Node struct {
// 	Conn      *websocket.Conn
// 	DataQueue chan []byte
// 	GroupSets set.Interface
// }

// // 映射关系
// var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// // 读写锁
// var rwLocker sync.RWMutex

// func Chat(writer http.ResponseWriter, request *http.Request) {
// 	//1. 检验 token
// 	query := request.URL.Query()
// 	Id := query.Get("userId")
// 	userId, _ := strconv.ParseInt(Id, 10, 64)
// 	token := query.Get("token")
// 	targetId := query.Get("targetId")
// 	context := query.Get("context")
// 	isvalida := true // checkToken() Todo
// 	msgType := query.Get("type")
// 	conn, err := (&websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			return isvalidate
// 		},
// 	}).Upgrade(writer, request, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//2. 获取conn
// 	node := &Node{
// 		Conn:      conn,
// 		DataQueue: make(chan []byte, 50),
// 		GroupSets: set.New(set.ThreadSafe),
// 	}
// 	//3. 用户关系
// 	//4. userid跟node绑定
// 	rwLocker.Lock()
// 	clientMap[userId] = node
// 	rwLocker.Unlock()

// 	//5. 完成发送逻辑
// 	go sendProc(node)
// 	//6. 完成接收逻辑
// 	go recvProc(node)
// }
