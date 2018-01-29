package fetchers

import (
	"encoding/json"
	"errors"
	"net"
)

//http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json
type jsonSina struct {
	Ret      int    `json:"ret"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Isp      string `json:"isp"`
	Type     string `json:"type"`
	Desc     string `json:"desc"`
}

func FetcherSina(ip net.IP) (loc *Location, err error) {
	//http://api.map.baidu.com/location/ip?ak=0cXtwomkf844vYYtjyec37hozGlfg1am&ip=119.96.5.68
	body, err := getRequestFormat("http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json&ip=%s", ip)
	if err != nil {
		return nil, err
	}
	var ipJson jsonSina
	json.Unmarshal(body, &ipJson)
	if ipJson.Ret != 1 {
		return nil, errors.New("新浪ip地址解析失败!")
	}
	loc = &Location{
		Country:  ipJson.Country,
		Province: ipJson.Province,
		City:     ipJson.City,
	}

	return
}
