package util

import "net"
import "fmt"
import "bufio"

//RequestString requests a String from CDS Server
func RequestString(toRequest string) string{
 
	// connect to this socket
  	conn, _ := net.Dial("tcp", "88.214.57.8:1905")
	// send to socket
    fmt.Fprintf(conn, toRequest)
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    return message
	

}