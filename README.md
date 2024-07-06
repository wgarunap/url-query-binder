<div align="center">
    <img src="logo.webp" alt="Logo" width="250"/>
</div>

# Go URL Query Binder
Simple, yet Powerful library. This library binds the URL request query parameters to Go structs based on the given tag key.

### How to use
Add dependency to your repository
```shell
go get -u github.com/wgarunap/url-query-binder
```
Import dependency to your code
```shell
import querybinder "github.com/wgarunap/url-query-binder"
```

### Example
```go
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
	_ = qb.Bind(&obj, u)

	log.Println(obj.Query)
	log.Println(obj.StringParam)
	log.Println(obj.SliceParam)
	log.Println(obj.IntParam)
}
```

##### Output
```shell
something
testing
[param1 param2]
12
```

### Benchmark
Benchmark was done with 4 query parameters and results as below. 

```bash
goos: darwin
goarch: amd64
pkg: github.com/wgarunap/url-query-binder
cpu: Intel(R) Core(TM) i7-5557U CPU @ 3.10GHz
BenchmarkBinderBind-4             581227              2125 ns/op
PASS
ok      github.com/wgarunap/url-query-binder    2.320s

```

## Contributions
Contributions are welcome :) 
