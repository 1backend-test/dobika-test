package test-service

import(
	"github.com/1backend/go-client"
)

var Token string

type Test struct {
	Name	string
	Foods	[]string
}


func GetHi() error {
	var ret 
	return ret, goclient.New(Token).Call("dobika-test", "test-service", "GET", "/hi", map[string]interface{}{  }, &ret)
}

func GetImportedHi() error {
	var ret 
	return ret, goclient.New(Token).Call("dobika-test", "test-service", "GET", "/imported-hi", map[string]interface{}{  }, &ret)
}

func GetSqlExample() error {
	var ret 
	return ret, goclient.New(Token).Call("dobika-test", "test-service", "GET", "/sql-example", map[string]interface{}{  }, &ret)
}

