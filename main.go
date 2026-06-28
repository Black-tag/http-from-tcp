package main

import (
	"bytes"
	"fmt"
	"net"
)


type TCPAddr struct {
    Ip net.IP
    Port int
}


func main() {
    fmt.Print("hello")

    for {
        ln, err := net.Listen("tcp", "127.0.0.1:8080")
        if err != nil {
            fmt.Print("couldnt make listener")
            return 
        }
        // fmt.Printf(ln.Accept())
        // ln.Close()
        tcpAddr := TCPAddr{
            Ip: bytes{"127.0.0.1"},
        }


        
        conn, err := net.DialTCP("tcp", nil, nil)
        if err != nil {
            fmt.Print("could get connaection")
            
        }
        fmt.Print(conn)
        ln.Accept()
    
    }
    
    
    

}