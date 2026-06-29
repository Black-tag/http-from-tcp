package main

import (
	"fmt"
	"net"
	"os"
)


type TCPAddr struct {
    Ip net.IP
    Port int
}


func main() {
    fmt.Println("Starting TCP server on 8080")
    l, err := net.Listen("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Print("couldnt make listener")
        os.Exit(1)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
			continue

        }
        fmt.Println("Connected to tcp server running on 8080")
        fmt.Fprint(conn)
        
    
    }
     

}