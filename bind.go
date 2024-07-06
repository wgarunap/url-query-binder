package querybinder

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Binder interface {
	SetTag(tag string)
	Bind(obj interface{}, url *url.URL) error
}

const (
	defaultTagKey    = "bind"
	requiredTagValue = "required"
)

var ErrInvalidObjectType = errors.New("invalid object type")
var ErrMissingQueryParam = errors.New("missing required query param")
var ErrInvalidQueryValue = errors.New("invalid query value")

type binder struct {
	tagKey string
}

// Bind unmarshal the url into given object struct based on the provided tags.
func (b *binder) Bind(obj interface{}, url *url.URL) error {
	val := reflect.ValueOf(obj)
	if !isStructPointer(val) {
		return ErrInvalidObjectType
	}

	rt := val.Elem().Type()

	params := url.Query()

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)

		tagStr, ok := f.Tag.Lookup(b.tagKey)
		if !ok {
			continue
		}

		tag, required := isRequiredField(tagStr)

		val, ok := params[tag]
		if !ok {
			if required {
				return fmt.Errorf("%w, '%s' is required", ErrMissingQueryParam, tag)
			}
			continue
		}

		var v interface{}

		switch f.Type.Kind() {
		case reflect.String:
			v = val[0]
		case reflect.Slice:
			v = val
		case reflect.Int:
			i, err := strconv.Atoi(val[0])
			if err != nil {
				return fmt.Errorf("%w, err:%s", ErrInvalidQueryValue, err.Error())
			}
			v = i
		}

		reflect.ValueOf(obj).Elem().Field(i).Set(reflect.ValueOf(v))
	}
	return nil
}

func isStructPointer(v reflect.Value) bool {
	// Check if the value is a pointer
	if v.Kind() == reflect.Ptr {
		// Check if the element the pointer points to is a struct
		return v.Elem().Kind() == reflect.Struct
	}
	return false
}

func isRequiredField(tagStr string) (tag string, required bool) {
	tagSplit := strings.Split(tagStr, ",")
	switch len(tagSplit) {
	case 1:
		return tagSplit[0], false
	case 2:
		if tagSplit[1] == requiredTagValue {
			return tagSplit[0], true
		}
	}
	return tagStr, false
}

// SetTag sets a new struct parameter read tag.
func (b *binder) SetTag(tag string) {
	b.tagKey = tag
}

// NewQueryBinder returns a new QueryBinder object
func NewQueryBinder() Binder {
	return &binder{
		tagKey: defaultTagKey,
	}
}
