package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"


	//"path"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	var directory  string
	args:=os.Args
	for i,arg :=range args {
		if arg =="--directory" && i+1<len(args){
			directory=args[i+1]
		}

	}
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
		//Concurrent connections
		go handleConnection(conn,directory)
	}
}

	
	func handleConnection(conn net.Conn,directory string){
		defer conn.Close()

	
	//Read Request Line 
	reader :=bufio.NewReader(conn)
	reqline,err :=reader.ReadString('\n')
	if err !=nil{
		fmt.Println("Error Reading Request ",err)
        return
	}
	parts :=strings.Split(reqline," ")
	if len(parts)<2 {
		return
	}
	path :=parts[1]
	var response string 
	// Handle endpoints
	if path == "/" {
	response = "HTTP/1.1 200 OK\r\n\r\n"
}else if strings.HasPrefix(path,"/echo/") {
message:=strings.TrimPrefix(path,"/echo/")
contentlen :=len(message)
response = fmt.Sprintf(
				"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
				contentlen, message)

}else if strings.HasPrefix(path,"/files/"){
	filename:=strings.TrimPrefix(path,"/files/")
	filepath:=filepath.Join(directory ,filename) 

	// Check if file exists
	data,err :=os.ReadFile(filepath)
	if err !=nil {
		response = "HTTP/1.1 404 Not Found\r\n\r\n"
	}else{
		response = fmt.Sprintf(
				"HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s",
				len(data), string(data))
	}	

}else {
	response = "HTTP/1.1 404 Not Found\r\n\r\n"
}


conn.Write([]byte(response))
	

	}
	 

