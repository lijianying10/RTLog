package main

import (
	"container/list"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"code.google.com/p/go.net/websocket"
	"github.com/ActiveState/tail"
)

// connection manager
var connected *list.List

// Echo response for connecter
func Echo(ws *websocket.Conn) {
	times := 0
	for {
		if times == 0 {
			connected.PushBack(*ws)
		}
		Push(`{"Timestamp":` + strconv.FormatInt(time.Now().Unix(), 10) + `,"Codeline":"null","Level":4,"Info":"new clent connected","Detail":null}`)
		ws.Read(nil)
		fmt.Println("code here")
		times++
	}
}

// WsRun server 如果是集成的情况加上新的路由就可以了。
func main() {
	// initial connection
	connected = list.New()
	// setting router
	http.Handle("/c", websocket.Handler(Echo))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	go func() {
		if err := http.ListenAndServe(":1234", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()
	LogSender()
}

// LogSender 日志发送
func LogSender() {
	// log file watcher , ATTENTION: seekinfo
	t, _ := tail.TailFile("./test.log",
		tail.Config{Follow: true, Location: &tail.SeekInfo{0, 2}})
	for line := range t.Lines {
		//先不管，直接推出去，因为在调试中丢log无所谓，no need cache
		Push(line.Text)
	}
}

// Push push data to client
func Push(text string) {
	for e := connected.Front(); e != nil; e = e.Next() {
		conn := e.Value.(websocket.Conn)
		if err := websocket.Message.Send(&conn, text); err != nil {
			fmt.Println("Can't send", err.Error(), connected.Len())
			// delete error client
			connected.Remove(e)
			continue
		}
	}
}
