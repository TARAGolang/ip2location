package fetchers

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

var db *geoip2.Reader
var GeoIpMmdbPath = "C:/GOPATH/src/github.com/mojocn/ip2location/geoIpData/GeoLite2-City.mmdb"

func CreateGeoIpDB() {
	var err error
	db, err = geoip2.Open(GeoIpMmdbPath)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
}

func FetcherGeoip(ip net.IP) (loc *Location, err error) {

	// If you are using strings that may be invalid, check that ip is not nil
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	loc = &Location{
		City:      record.City.Names["zh-CN"],
		Province:  record.Subdivisions[0].Names["zh-CN"],
		Country:   record.Country.Names["zh-CN"],
		Latitude:  record.Location.Latitude,
		Longitude: record.Location.Longitude,
		TimeZone:  record.Location.TimeZone,
	}

	return loc, nil
}
