package test

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/Chyroc/algorithms-go/lib"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	Func   interface{}
	Input  string
	Output string
}

func Runs(t *testing.T, f interface{}, cs []*Case) {
	for _, c := range cs {
		Run(t, &Case{
			Func:   f,
			Input:  c.Input,
			Output: c.Output,
		})
	}
}

// 定义字符串能够通过转成定的类型: string -> typ
// 特定的类型也能转成字符串:      typ -> string
type Serialization interface {
	Unmarshal(data string) (interface{}, error)
	Marshal() (string, error)
}

var nilValue = reflect.New(reflect.TypeOf((*interface{})(nil)).Elem()).Elem()

// ParseParam 这里需要自定义string转化
//
// 参数是string / reflect.Type，返回值是reflect.Value
// string("") => zero value
// string(1), int => int(1)
// string([1,2,3]), slice(int) => []int{1,2,3}
// string(2 -> 4 -> 3)/Serialization => Serialization()
func ParseParam(t *testing.T, param string, typ reflect.Type) (reflect.Value, error) {
	var r reflect.Value
	param = strings.TrimSpace(param)

	if param == "" {
		return reflect.Zero(typ), nil
	}

	switch typ.Kind() {
	case reflect.Ptr:
		typecType := reflect.TypeOf((*Serialization)(nil)).Elem()
		if typ.Implements(typecType) {
			method, ok := typ.MethodByName("Unmarshal")
			if !ok {
				return nilValue, fmt.Errorf("not such method")
			}

			f := method.Func

			zhi := reflect.New(typ).Elem()
			out := f.Call([]reflect.Value{
				zhi,
				reflect.ValueOf(param),
			})

			var returnInterface = out[0]
			var errValue = out[1]

			if !errValue.IsNil() {
				return nilValue, fmt.Errorf("%s", errValue.Interface())
			} else if returnInterface.IsNil() {
				return reflect.New(typ).Elem(), nil
			}

			r = returnInterface.Elem().Convert(typ)
		} else {
			return ParseParam(t, param, typ.Elem())
		}
	case reflect.Int:
		i, err := strconv.Atoi(param)
		if err != nil {
			return nilValue, err
		}
		r = reflect.ValueOf(i)
	case reflect.String:
		r = reflect.ValueOf(param)
	case reflect.Bool:
		b, err := strconv.ParseBool(param)
		if err != nil {
			return nilValue, err
		}
		r = reflect.ValueOf(b)
	case reflect.Slice:
		if len(param) <= 1 {
			return nilValue, fmt.Errorf("invalid params(%s) length(%d) when mathched %v", param, len(param), reflect.Slice)
		}

		r = reflect.MakeSlice(reflect.SliceOf(typ.Elem()), 0, 0)
		if param == `` || param == `[]` {
			return r, nil
		}

		s2, err := lib.SplitWithToken(param, '[', ']')
		if err != nil {
			return nilValue, err
		}

		for _, v := range s2 {
			paramValue, err := ParseParam(t, v, typ.Elem())
			if err != nil {
				return nilValue, err
			}
			r = reflect.Append(r, paramValue)
		}
	case reflect.Float64:
		f, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return nilValue, err
		}
		return reflect.ValueOf(f), nil
	default:
		return nilValue, fmt.Errorf("not support %s", typ.Kind())
	}

	return r, nil
}

func Run(t *testing.T, c *Case) {
	as := assert.New(t)
	input := strings.Split(strings.TrimSpace(c.Input), `\n`)
	output := strings.Split(strings.TrimSpace(c.Output), `\n`)

	ft := reflect.TypeOf(c.Func)
	fv := reflect.ValueOf(c.Func)
	as.Equal(reflect.Func, ft.Kind())
	as.Equal(reflect.Func, fv.Kind())
	as.Len(input, ft.NumIn())
	as.Len(output, ft.NumOut())

	var in []reflect.Value
	for i := 0; i < ft.NumIn(); i++ {
		ithCallInType := ft.In(i)

		ithParamIn, err := ParseParam(t, input[i], ithCallInType)
		as.Nil(err)

		in = append(in, ithParamIn)
	}

	out := fv.Call(in)

	// out 有三个，call返回，ft.Out(i)的，output的
	for i := 0; i < ft.NumOut(); i++ {
		ithCallRealOut := out[i] // 比input多的
		ithCallOutType := ft.Out(i)
		ithCallOut, err := ParseParam(t, output[i], ithCallOutType)
		as.Nil(err)

		as.Equal(ithCallOut.Kind(), ithCallRealOut.Kind())
		as.Equal(ithCallOut.Kind(), ithCallRealOut.Convert(ithCallOutType).Kind())
		if ithCallOut.Kind() == reflect.Slice && ithCallOut.Len() == 0 {
			as.Equal(0, ithCallRealOut.Convert(ithCallOutType).Len())
			continue
		}
		as.Equal(ithCallOut.Interface(), ithCallRealOut.Convert(ithCallOutType).Interface())
	}
}
