*** Example

inspire from https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang

```golang
package main
import (
	"github.com/neverlock/utility/httpjs"
	"fmt"
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
func main(){
	foo1 := new(Foo) // or &Foo{}  
	err := httpjs.GetJS("http://httpbin.org/get", foo1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", foo1.Headers)
	fmt.Println(foo1.Headers.ACCEPT) 
	println(foo1.Headers.XPROXYID)

	//alternately:
	/*
	foo2 := Foo{}
	httpjs.GetJS("http://example.com",&foo2)
	*/
}
```
