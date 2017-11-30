package test-service

import(
	"github.com/1backend/go-client"
)

var Token string

type Test struct {
	Name	string
	Foods	[]string
}


func GetHi(howManyTimes int64, allCool bool) error {
	var ret 
	return ret, goclient.New(Token).Call("dobika-test", "test-service", "GET", "/hi", map[string]interface{}{ "howManyTimes": howManyTimes, "allCool": allCool }, &ret)
}

func GetImportedHi() error {
	var ret 
	return ret, goclient.New(Token).Call("dobika-test", "test-service", "GET", "/imported-hi", map[string]interface{}{  }, &ret)
}

func GetSqlExample() (&#34;String&#34;, error) {
	var ret &#34;String&#34;
	return ret, goclient.New(Token).Call("dobika-test", "test-service", "GET", "/sql-example", map[string]interface{}{  }, &ret)
}

