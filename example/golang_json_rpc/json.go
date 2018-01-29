package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
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

type Ip2addr struct {
	db *geoip2.Reader
}
type Agrs struct {
	IpString string
}

func (t *Ip2addr) Address(agr *Agrs, res *Response) error {
	netIp := net.ParseIP(agr.IpString)
	record, err := t.db.City(netIp)
	res.City = record.City.Names["zh-CN"]
	res.Province = record.Subdivisions[0].Names["zh-CN"]
	res.Country = record.Country.Names["zh-CN"]
	res.Latitude = record.Location.Latitude
	res.Longitude = record.Location.Longitude
	res.TimeZone = record.Location.TimeZone
	return err
}

func main() {
	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	ip2addr := &Ip2addr{db}

	rpc.Register(ip2addr)
	address := ":3344"
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	log.Println("json rpc is listening", tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
