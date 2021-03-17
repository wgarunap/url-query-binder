package querybinder_test

import (
	querybinder "github.com/wgarunap/url-query-binder"
	"net/url"
	"testing"
)

func TestBind(t *testing.T) {
	type Obj struct {
		Query       string   `bind:"query"`
		StringParam string   `bind:"string_param"`
		SliceParam  []string `bind:"slice_param"`
		IntParam    int      `bind:"int_param"`
	}

	var obj Obj

	u, _ := url.Parse("/get?query=something&string_param=testing&slice_param=param1&slice_param=param2&int_param=12")

	err := querybinder.Bind(&obj, u, "bind")
	if err != nil {
		t.Error(err)
	}
	if obj.Query != "something" ||
		obj.StringParam != "testing" ||
		obj.SliceParam[0] != "param1" ||
		obj.SliceParam[1] != "param2" ||
		obj.IntParam != 12 {
		t.Fail()
	}
}
