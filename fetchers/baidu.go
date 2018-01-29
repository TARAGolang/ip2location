package fetchers

import (
	"encoding/json"
	"errors"
	"net"
)

//还要集成Redis
var baiduAk = "0cXtwomkf844vYYtjyec37hozGlfg1am"

type Location struct {
	Country   string  `json:"country"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	ISP       string  `json:"isp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone  string  `json:"time_zone"`
}

type jsonBaidu struct {
	Address string `json:"address"`
	Content struct {
		AddressDetail struct {
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
			CityCode     int    `json:"city_code"`
		} `json:"address_detail"`
		Address string `json:"address"`
		Point   struct {
			Y string `json:"y"`
			X string `json:"x"`
		} `json:"point"`
	} `json:"content"`
	Status int `json:"status"`
}

//http://ip.taobao.com/instructions.php
func FetcherBaidu(ip net.IP) (loc *Location, err error) {
	//http://api.map.baidu.com/location/ip?ak=0cXtwomkf844vYYtjyec37hozGlfg1am&ip=119.96.5.68

	body, err := getRequestFormat("http://api.map.baidu.com/location/ip?ak="+baiduAk+"&ip=%s", ip)
	if err != nil {
		return nil, err
	}
	var ipJson jsonBaidu
	json.Unmarshal(body, &ipJson)
	if ipJson.Status != 0 {
		return nil, errors.New("百度ip地址解析失败!")
	}
	loc = &Location{
		Province: ipJson.Content.AddressDetail.Province,
		City:     ipJson.Content.AddressDetail.City,
	}

	return
}
