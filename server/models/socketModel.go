package models

import (
	"GolangServer/server/drivers"
	"fmt"
	"net"
	"time"
)

type msg struct {
	FirstNum  int    `json:"firstNum"`
	SecondNum int    `json:"secondNum"`
	Msg       string `json:"msg"`
}

func connect() net.Conn {
	conn, err := net.Dial("tcp", drivers.Viper.GetString("FlaskServer.IP")+":"+drivers.Viper.GetString("FlaskServer.SocketPort"))
	if err != nil {
		panic(err)
	}
	return conn
}

func SendMsg(firstNum int, secondNum int, _msg string) {
	var _conn = connect()
	NewMsg := new(msg)
	NewMsg.FirstNum = firstNum
	NewMsg.SecondNum = secondNum
	NewMsg.Msg = _msg
	fmt.Println(firstNum, "_", secondNum, "msg:", NewMsg.Msg)
	_conn.Write(StruckToJson(NewMsg))

	res := make([]byte, 64)
	_conn.Read(res)

	_conn.Close()
	time.Sleep(100 * time.Millisecond)
}
