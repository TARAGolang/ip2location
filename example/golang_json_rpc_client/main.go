package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Response struct {
	Country   string
	Province  string
	City      string
	ISP       string
	Latitude  float64
	Longitude float64
	TimeZone  string
}
type Agrs struct {
	IpString string
}

func main() {
	client, err := jsonrpc.Dial("tcp", "121.40.238.123:3344")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	var res Response
	err = client.Call("Ip2addr.Address", Agrs{"219.140.227.235"}, &res)
	if err != nil {
		log.Fatal("Ip2addr error:", err)
	}
	fmt.Println(res)

}
