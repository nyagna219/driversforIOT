package main

import (
	"fmt"
	"net"
	"os"
	//"bufio"
	//"strings"
)

var clientaddr string
var ledaddr string

/* A Simple function to verify error */
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

var sensor int

var led_found int
var photo_found int

func responseBcast (client string)  {
	conn, err := net.Dial("udp", client)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return 
	}
	msg := "9"
	fmt.Fprintf(conn, msg)
	fmt.Printf("Sent the ACK to the Sensor\n")

	fmt.Printf("WE HAVE A NEW SENSOR\n")
	sensor++
	conn.Close()
	return
}

func main() {
	/* Lets prepare a address at any address at port 9*/   

	led_found=0
	photo_found=0
	ServerAddr,err := net.ResolveUDPAddr("udp",":9")
	CheckError(err)

	sensor =0
	buf := make([]byte, 1024)
	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	for {
		_,addr,_ := ServerConn.ReadFromUDP(buf)
		defer ServerConn.Close()
		if buf[0] == '3' && photo_found == 0 {
			clientaddr = fmt.Sprintf("%s",addr)
			responseBcast(clientaddr)
			photo_found = 1
		} else if buf[0] == '1' && led_found == 0{
			ledaddr = fmt.Sprintf("%s",addr)
			responseBcast(ledaddr)
			led_found = 1
		}
		if photo_found == 1 && led_found == 1  {
			fmt.Printf("WE DONE\n")
			return
		}
	}

}
