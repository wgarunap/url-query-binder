package querybinder

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func Bind(obj interface{}, url *url.URL, tagKey string) error {
	elem := reflect.ValueOf(obj).Elem()
	if !elem.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	rt := elem.Type()

	params := url.Query()

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tag, ok := f.Tag.Lookup(tagKey)

		if ok {
			val, ok := params[tag]
			if ok {
				var v interface{}

				switch f.Type.Kind() {
				case reflect.String:
					v = val[0]
				case reflect.Slice:
					v = val
				case reflect.Int:
					i, err := strconv.Atoi(val[0])
					if err != nil {
						return err
					}
					v = i
				}

				reflect.ValueOf(obj).Elem().Field(i).Set(reflect.ValueOf(v))
			}

		}
	}
	return nil
}
