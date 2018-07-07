package main

import (
	"bufio"
	"fmt"
	"net"
    "os"
    "strings"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	r := bufio.NewReader(conn)
    cons := bufio.NewReader(os.Stdin)
    var msg string
    for {
        fmt.Print("Mensaje: ")
        msg, _ = cons.ReadString('\n')
        fmt.Fprintf(conn, "%s", msg)
        msg, _ = r.ReadString('\n')
        fmt.Printf("Respuesta: %s", msg)
        if strings.Contains(msg, "exit") {
            break
        }
    }
}