package main

import (
	"log"
	"net/url"
	"os"

	querybinder "github.com/wgarunap/url-query-binder"
)

type Obj struct {
	Query       string   `bind:"query,required"`
	StringParam string   `bind:"string_param"`
	SliceParam  []string `bind:"slice_param"`
	IntParam    int      `bind:"int_param"`
}

func main() {
	var obj Obj
	u, _ := url.Parse("/get?query=something&string_param=testing&slice_param=param1&slice_param=param2&int_param=12")

	qb := querybinder.NewQueryBinder()
	err := qb.Bind(&obj, u)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(obj.Query)
	log.Println(obj.StringParam)
	log.Println(obj.SliceParam)
	log.Println(obj.IntParam)
}
