package game_server

import (
	"UnityGame_AB_System/dao"
	"fmt"
	"net"
	"os"
	"time"
)

//game_server.StartRunServer
func StartRunServer() {
	service := ":20013"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

var count = 0

func handleClient(conn net.Conn) {
	fmt.Println("Client connect success :", conn.RemoteAddr().String())
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack
	defer conn.Close()                                    // close connection before exit
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("rev client len:", read_len)
		if read_len == 0 {
			fmt.Println("connection already closed by client")
			break // connection already closed by client
		} else {
			// daytime := time.Now().String()
			count = count + 1
			//
			rev_untiy_msg := string(request[:read_len])
			//
			sendClient := fmt.Sprintf("go server:%d from unity=%s", count, rev_untiy_msg)
			conn.Write([]byte(sendClient))
			//
			var logType = "log"
			var logTxt =rev_untiy_msg
			dao.AutoCreateTodayAndYestodayRptEmptyRow(logType,logTxt)
			//
			fmt.Println("msg2=", sendClient)
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/*
else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
			fmt.Println("msg1=", daytime)
*/
