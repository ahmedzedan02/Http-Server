package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	//"path"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Server is running on port 4221...")
	listener,err :=net.Listen("tcp","0.0.0.0:4221")
	if err !=nil{
		os.Exit(1)
	}

	defer listener.Close()
	for   {
		
		conn,err :=listener.Accept()
		if err != nil{
			fmt.Println("Error accepting connection:", err)
			continue
		}
	
	
	//Read Request Line 
	reader :=bufio.NewReader(conn)
	reqline,err :=reader.ReadString('\n')
	if err !=nil{
		fmt.Println("Error Reading Request ",err)
		conn.Close()
        continue
	}
	parts :=strings.Split(reqline," ")
	if len(parts)<2 {
		conn.Close()
		continue
		
	}
	path :=parts[1]
	var response string 
	if path == "/" {
	response = "HTTP/1.1 200 OK\r\n\r\n"
} else {
	response = "HTTP/1.1 404 Not Found\r\n\r\n"
}


conn.Write([]byte(response))
conn.Close()
	
	}
	 }

