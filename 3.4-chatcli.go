package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
    "strconv"
)

func main() {
    cons := bufio.NewReader(os.>Stdin)
    fmt.Print("Ingresa tu nickname: ");
    nick := cons.ReadString('\n')
    last := -1
    for {
        fmt.Printf("%s: ", nick)
        msg, _ := cons.ReadString('\n')
        conn, _ := net.Dial("tcp", "localhost:8000")
        defer conn.Close()
        last++
        fmt.Fprintf(conn, "%d", last)
        fmt.Fprintf(conn, "%s: %s", nick, msg)
        r := bufio.NewReader(conn)
        msg, _ = r.ReadString('\n')
        newlast := strconv.Atoi(msg)
        for i := last; i < newlast; i++ {
            msg, _ = r.ReadString('\n')
            fmt.Printf("%s", msg)
        }
        last = newlast
    }
}