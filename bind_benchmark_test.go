package querybinder

import (
	"net/url"
	"testing"
)

func BenchmarkBinderBind(b *testing.B) {
	type Obj struct {
		Query       string   `bind:"query"`
		StringParam string   `bind:"string_param"`
		SliceParam  []string `bind:"slice_param"`
		IntParam    int      `bind:"int_param"`
	}

	var obj Obj
	u, _ := url.Parse("/get?query=something&string_param=testing&slice_param=param1&slice_param=param2&int_param=12")
	qb := NewQueryBinder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = qb.Bind(&obj, u)
	}
}
