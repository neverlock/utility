package httpjs

import (
	"fmt"
	"testing"
)

type Foo struct {
	Args    struct{}
	Headers Bar
	Origin  string
	Url     string
}

type Bar struct {
	ACCEPT          string `json:"Accept"`
	ACCEPTENCODING  string `json:"Accept-Encoding"`
	ACCEPTLANGUAGE  string `json:"Accept-Language"`
	CACHECONTRAL    string `json:"Cache-Control"`
	CONNECTION      string `json:"Connection"`
	HOST            string `json:"Host"`
	IFMODIFIEDSINCE string `json:"If-Modified-Since"`
	USERAGENT       string `json:"User-Agent"`
	XPROXYID        string `json:"X-Proxy-Id"`
}

func Test_GetJS(t *testing.T) {
	fmt.Println("Testing GetJS")
	foo1 := new(Foo)
	ret := GetJS("http://httpbin.org/get", foo1)
	if ret != nil {
		t.Errorf("Error occured: %v\n", ret)
	} else {
		fmt.Println(foo1.Headers.ACCEPT)
		fmt.Println(foo1.Origin)
	}
}
