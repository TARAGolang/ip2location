package ip2location

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"github.com/oschwald/maxminddb-golang"
	"log"
	"net"
	"testing"
)

func TestGeoIp(t *testing.T) {

}

func TestGeoIp2(t *testing.T) {
	db, err := geoip2.Open("C:/GOPATH/src/github.com/mojocn/ip2location/geoIpData/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("119.96.211.173")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}

// This example demonstrates how to decode to an interface{}
func TestGeo2ipGolang(t *testing.T) {
	db, err := maxminddb.Open("C:/GOPATH/src/github.com/mojocn/ip2location/geoIpData/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("119.96.211.173")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(record)
}
