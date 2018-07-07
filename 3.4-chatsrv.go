package main

import (
	"bufio"
	"net"
    "fmt"
    "strings"
)

messages := make([]string, 0, 100)
var msgWriter chan string
var count chan int

func messagesHandler() {
    for {
        msg := <- msgWriter
        append(messages, msg)
        count <- len(mensages)
    }
}

func handle(conn net.Conn) {
    defer conn.Close()
	r := bufio.NewReader(conn)
    msg, _ := r.ReadString('\n')
    last := strconv.Atoi(msg)
    msg, _ := r.ReadString('\n')
    msgWriter <- msg
    newlast := <- count
    fmt.Fprintf(conn, "%d", newlast)
    for i := last + 1; i < newlast; i++ {
        // io.Copy(conn, conn)
        fmt.Fprintf(conn, "%s", message[i])
    }
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
    for {
        conn, _ := ln.Accept()
        go handle(conn)
    }
}
