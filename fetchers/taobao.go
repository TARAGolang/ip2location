package fetchers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

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

func FetcherTaobao(ip net.IP) (loc *Location, err error) {

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

func getRequestFormat(urlFormate string, ip net.IP) ([]byte, error) {
	url := fmt.Sprintf(urlFormate, ip.String())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
