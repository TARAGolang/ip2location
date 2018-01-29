package ip2location

import (
	"errors"
	"github.com/mojocn/ip2location/fetchers"

	"net"
)

type CacheDriver interface {
	SetLocation(ip string, lo *fetchers.Location)
	GetLocation(ip string) *fetchers.Location
}

type locationFetcher func(ip net.IP) (*fetchers.Location, error)

type IpParser struct {
	Handlers   []locationFetcher
	Store      CacheDriver
	IsUseCache bool
}

func NewIpParser() {
	instance := IpParser{}
	instance.Handlers = []locationFetcher{fetchers.FetcherGeoip, fetchers.FetcherBaidu, fetchers.FetcherTaobao, fetchers.FetcherSina}
}

func (parser *IpParser) FetchIpAddress(ip string) (*fetchers.Location, error) {

	if len(parser.Handlers) < 1 {
		return nil, errors.New("you must add ip handler function!")
	}

	if net.ParseIP(ip) == nil {
		return nil, errors.New("ip string is invalid!")
	}

	if lo := parser.Store.GetLocation(ip); lo != nil {
		return lo, nil
	}

	for _, f := range parser.Handlers {
		lo, err := f(ip)
		if err != nil {
			return nil, err
		}
		if parser.IsUseCache {
			parser.Store.SetLocation(ip, lo)
		}
		return lo, nil
	}

	return nil, errors.New("all handlers failed, returns no address!")
}
