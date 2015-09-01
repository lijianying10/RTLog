package main

import (
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
	"github.com/ActiveState/tail"
)

// connection manager
var connected []websocket.Conn

// Echo response for connecter
func Echo(ws *websocket.Conn) {
	connected = append(connected, *ws)
}

// WsRun server 如果是集成的情况加上新的路由就可以了。
func main() {
	// setting router
	http.Handle("/c", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// LogSender 日志发送
func LogSender() {
	var err error
	// log file watcher , ATTENTION: seekinfo
	t, _ := tail.TailFile("./test.log",
		tail.Config{Follow: true, Location: &tail.SeekInfo{0, 2}})
	for line := range t.Lines {
		//先不管，直接推出去，因为在调试中丢log无所谓，no need cache
		for index := 0; index < len(connected); index++ {
			if err = websocket.Message.Send(&connected[index], line.Text); err != nil {
				fmt.Println("Can't send")
				// delete error client
				connected = append([]websocket.Conn{}, connected[:index-1]...)
				connected = append(connected, connected[index+1:]...)
				continue
			}
		}
	}
}
