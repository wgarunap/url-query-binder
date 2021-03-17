package main

import (
	"fmt"
	querybinder "github.com/wgarunap/url-query-binder"
	"net/url"
	"os"
)

type Obj struct {
	Query       string   `bind:"query"`
	StringParam string   `bind:"string_param"`
	SliceParam  []string `bind:"slice_param"`
	IntParam    int      `bind:"int_param"`
}

func main() {
	var obj Obj
	u, _ := url.Parse("/get?query=something&string_param=testing&slice_param=param1&slice_param=param2&int_param=12")

	err := querybinder.Bind(&obj, u, "bind")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(obj.Query)
	fmt.Println(obj.StringParam)
	fmt.Println(obj.SliceParam)
	fmt.Println(obj.IntParam)
}
