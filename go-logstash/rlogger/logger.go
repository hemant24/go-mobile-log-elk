package rlogger


import (
	"net"
	"fmt"
	"time"
	"errors"
)

type Logstash struct {
	Hostname string
	Port int
	Connection *net.TCPConn
	Timeout int
}


func New(hostname string, port int, timeout int) *Logstash {
	l := Logstash{}
	l.Hostname = hostname
	l.Port = port
	l.Connection = nil
	l.Timeout = timeout
	return &l
}

func (l *Logstash) Dump() {
	fmt.Println("Hostname:   ", l.Hostname)
	fmt.Println("Port:       ", l.Port)
	fmt.Println("Connection: ", l.Connection)
	fmt.Println("Timeout:    ", l.Timeout)
}

func (l *Logstash) SetTimeouts() {
	deadline := time.Now().Add(time.Duration(l.Timeout) * time.Millisecond)
	l.Connection.SetDeadline(deadline)
	l.Connection.SetWriteDeadline(deadline)
	l.Connection.SetReadDeadline(deadline)
}

func (l *Logstash) Connect() (bool) {
	var connection *net.TCPConn
	service := fmt.Sprintf("%s:%d", l.Hostname, l.Port)
	addr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		fmt.Println(err)
		return  false
	}
	connection, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if connection != nil {
		l.Connection = connection
		l.Connection.SetLinger(0) // default -1
		l.Connection.SetNoDelay(true)
		l.Connection.SetKeepAlive(true)
		l.Connection.SetKeepAlivePeriod(time.Duration(5) * time.Second)
		l.SetTimeouts()
	}
	return true
}

func (l *Logstash) Writeln(app, level, tag, message string) (bool) {
	var err = errors.New("TCP Connection is nil.")
	message = fmt.Sprintf("{\"app\" : \"%s\", \"level\" : \"%s\", \"tag\" : \"%s\", \"msg\" : \"%s\"}\n", app, level, tag, message)
	fmt.Print(message)
	if l.Connection != nil {
		_, err = l.Connection.Write([]byte(message))
		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				l.Connection.Close()
				l.Connection = nil
				if err != nil {
					fmt.Println(err)
					return false
				}
			} else {
				l.Connection.Close()
				l.Connection = nil
				fmt.Println(err)
				return false
			}
		} else {
			// Successful write! Let's extend the timeout.
			l.SetTimeouts()
			return true
		}
	}
	fmt.Println(err)
	return false
}