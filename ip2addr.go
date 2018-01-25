package ip2location

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type CacheDriver interface {
	Set(ip string, lo *Location)
	Get(ip string) *Location
}

//还要集成Redis
var baiduLbsAk string

type Location struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	ISP      string `json:"isp"`
}

type locationFetcher func(string) (*Location, error)

type IpParser struct {
	Handlers        []locationFetcher
	CacheExpireTime time.Duration
	CackeKey        string
	Store           CacheDriver
}

func (parser *IpParser) FetchIpAddress(ip string) (*Location, error) {

	if len(parser.Handlers) < 1 {
		return nil, errors.New("you must add ip handler function!")
	}

	if net.ParseIP(ip) == nil {
		return nil, errors.New("ip string is invalid!")
	}

	if lo := parser.Store.Get(ip); lo != nil {
		return lo, nil
	}

	for _, f := range parser.Handlers {
		lo, err := f(ip)
		if err == nil {
			parser.Store.Set(ip, lo)
			return lo, nil
		}
	}

	return nil, errors.New("all handlers failed, returns no address!")
}

type jsonTaobao struct {
	Code int `json:"code"`
	Data struct {
		Country   string `json:"country"`
		CountryID string `json:"country_id"`
		Area      string `json:"area"`
		AreaID    string `json:"area_id"`
		Region    string `json:"region"`
		RegionID  string `json:"region_id"`
		City      string `json:"city"`
		CityID    string `json:"city_id"`
		County    string `json:"county"`
		CountyID  string `json:"county_id"`
		Isp       string `json:"isp"`
		IspID     string `json:"isp_id"`
		IP        string `json:"ip"`
	} `json:"data"`
}

func fetcherTaobao(ip string) (loc *Location, err error) {

	body, err := getRequestFormat("http://ip.taobao.com/service/getIpInfo.php?ip=%s", ip)
	if err != nil {
		return nil, err
	}
	var ipJson jsonTaobao
	json.Unmarshal(body, &ipJson)
	if ipJson.Code != 0 {
		return nil, errors.New("淘宝ip地址解析失败!")
	}
	loc = &Location{
		Country:  ipJson.Data.Country,
		Province: ipJson.Data.Region,
		City:     ipJson.Data.City,
		ISP:      ipJson.Data.Isp,
	}

	return
}

func getRequestFormat(urlFormate, ip string) ([]byte, error) {
	url := fmt.Sprintf(urlFormate, ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
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
func fetcherBaidu(ip string) (loc *Location, err error) {
	//http://api.map.baidu.com/location/ip?ak=0cXtwomkf844vYYtjyec37hozGlfg1am&ip=119.96.5.68
	baiduAk := "0cXtwomkf844vYYtjyec37hozGlfg1am"
	if baiduLbsAk != "" {
		baiduAk = baiduLbsAk
	}
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

func fetcherSina(ip string) (loc *Location, err error) {
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
