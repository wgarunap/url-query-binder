package querybinder_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	querybinder "github.com/wgarunap/url-query-binder"
)

func TestBinder_Bind(t *testing.T) {
	type Obj struct {
		Query       string   `bind:"query,required"`
		StringParam string   `bind:"string_param"`
		SliceParam  []string `bind:"slice_param"`
		IntParam    int      `bind:"int_param"`
	}

	tests := map[string]struct {
		obj         interface{}
		url         string
		expect      interface{}
		expectedErr error
	}{
		"url has multiple parameters": {
			obj: &Obj{},
			url: "/get?query=something&string_param=testing&slice_param=param1&slice_param=param2&int_param=12",
			expect: &Obj{
				Query:       "something",
				StringParam: "testing",
				SliceParam:  []string{"param1", "param2"},
				IntParam:    12,
			},
			expectedErr: nil,
		},
		"url has multiple parameters with comma separation, (should not treated as separate params)": {
			obj: &Obj{},
			url: "/get?query=something&string_param=testing&slice_param=param1,param2&int_param=12",
			expect: &Obj{
				Query:       "something",
				StringParam: "testing",
				SliceParam:  []string{"param1,param2"},
				IntParam:    12,
			},
			expectedErr: nil,
		},
		"url has only few params": {
			obj: &Obj{},
			url: "/get?query=something&string_param=testing",
			expect: &Obj{
				Query:       "something",
				StringParam: "testing",
			},
			expectedErr: nil,
		},
		"url is missing required query parameter": {
			obj:         &Obj{},
			url:         "/get?string_param=testing&slice_param=param1&slice_param=param2&int_param=12",
			expect:      &Obj{},
			expectedErr: querybinder.ErrMissingQueryParam,
		},
		"invalid object value passed": {
			obj:         "",
			url:         "/get?query=something&string_param=testing&slice_param=param1&slice_param=param2&int_param=12",
			expect:      "",
			expectedErr: querybinder.ErrInvalidObjectType,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			qb := querybinder.NewQueryBinder()
			u, _ := url.Parse(test.url)
			err := qb.Bind(test.obj, u)
			require.ErrorIs(t, err, test.expectedErr)
			require.Equal(t, test.expect, test.obj)
		})
	}
}
