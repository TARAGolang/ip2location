package ip2location

import (
	"log"
	"testing"
)

func TestIpParser_FetchIpAddress(t *testing.T) {

	parser := IpParser{
		Handlers: []locationFetcher{
			fetcherTaobao,
			fetcherSina,
			fetcherBaidu,
		},
		//Store: newIpCacheDriver(),
	}

	parser.FetchIpAddress("119.96.211.173")

	loc, err := fetcherBaidu("119.96.211.173")
	if err != nil {
		t.Error(err)
	}
	log.Print(loc)

	loc, err = fetcherSina("119.96.211.173")
	if err != nil {
		t.Error(err)
	}
	log.Print(loc)

	loc, err = fetcherTaobao("119.96.211.173")
	if err != nil {
		t.Error(err)
	}
	log.Print(loc)

}
