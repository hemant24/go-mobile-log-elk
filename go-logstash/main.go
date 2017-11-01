package main

import (
	"fmt"
	"github.com/hemant24/go-logstash/rlogger"
)



func main()  {

	//msg := `{"first_name" : "Douglas","last_name" : "Fir","age" : 35,"about": "I like to build cabinets", "create_date" : "2015-01-03","interests": [ "forestry" ]}`
	msg := "hello world"
	l := rlogger.New("0.tcp.ngrok.io", 13632, 5)
	err := l.Connect()
	if err == false {
		fmt.Println("error while connecting")
	}
	err2 := l.Writeln("go", "INFO", "test", msg)
	if err2 == false {
		fmt.Println("error while writing to logstash")
	}
}
