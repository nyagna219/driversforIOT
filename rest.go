package main

import (
	"fmt"
	"log"
	"net/http"
	"net"
	"strconv"
	"strings"
	"bufio"
	"os"
)

type query struct {
	Glowled int
	Status bool
}

var sensor int
var led_found int
var photo_found int
var clientaddr string
var ledaddr string



func parseBool(s string, dest *bool) error {
	// assume error = false
	*dest, _ = strconv.ParseBool(s)
	return nil
}


/* A Simple function to verify error */
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}


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

func setup_bcast () {
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
		if photo_found == 1 &&  led_found == 1  {
			fmt.Printf("WE DONE\n")
			ServerConn.Close()
			return
		}
	}

}

func communicate2 ( led int , status bool , addr string) string {
	cmd := 0
	ans := "NULL"
	
	conn, _ := net.Dial("udp", addr)

	if  status == true {
		cmd =1
	} else {
		cmd =0
	}	
	msg :=  fmt.Sprintf("%d:%d",led,cmd)
	fmt.Fprintf(conn, msg)
	fmt.Printf("data sent %s to %s",msg,addr)
	/*
	_, err = bufio.NewReader(conn).Read(p)
	ack := fmt.Sprintf("%s",p)
	if err == nil {
		if strings.Contains(ack,msg) {
			fmt.Printf("ACK received %s\n", p)
			s := strings.Split(ack, ":")
			conn.Close()
			fmt.Printf("Values are %s:%s:%s\n",s[0],s[1],s[2])
			return s[2]
		} else {
			fmt.Printf("Corrupt Ack\n");
		}
	} else {
		fmt.Printf("Some error %v\n", err)
	}*/

	conn.Close()
	return ans

}


func communicate ( led int , status bool , addr string) string {
	p :=  make([]byte, 2048)
	cmd := 0
	ans := "NULL"
	
	conn, err := net.Dial("udp", addr)

	if  status == true {
		cmd =1
	} else {
		cmd =0
	}	
	msg :=  fmt.Sprintf("%d:%d",led,cmd)
	fmt.Fprintf(conn, msg)
	fmt.Printf("data sent %s to %s",msg,addr)
	_, err = bufio.NewReader(conn).Read(p)
	ack := fmt.Sprintf("%s",p)
	if err == nil {
		if strings.Contains(ack,msg) {
			fmt.Printf("ACK received %s\n", p)
			s := strings.Split(ack, ":")
			conn.Close()
			fmt.Printf("Values are %s:%s:%s\n",s[0],s[1],s[2])
			return s[2]
		} else {
			fmt.Printf("Corrupt Ack\n");
		}
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
	return ans

}

func parseInt(s string, dest *int) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*dest = n
	return nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}
	q := &query{}
	ans := "2000"
	if err := parseBool(req.Form.Get("Status"), &q.Status); err != nil {
		log.Printf("Error parsing dryrun: %s", err)
		return
	}
	if err := parseInt(req.Form.Get("Glowled"), &q.Glowled); err != nil {
		log.Printf("Error parsing limit: %s", err)
		return
	}

	if  q.Glowled ==1 || q.Glowled == 9 {
		ans = communicate(q.Glowled,q.Status,clientaddr)
	} else if q.Glowled == 2 {
		communicate2(q.Glowled,q.Status,ledaddr)
	}

	fmt.Print(ans)
	fmt.Fprintf(w,ans)
}

func main() {
	setup_bcast()
	fmt.Printf("LED ADDR %s\n",ledaddr)
	fmt.Printf("Sensor LED %s\n",clientaddr)
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(handler)))
}
