package test

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

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

// parseParam 这里需要自定义string转化
//
// 参数是string / reflect.Type，返回值是reflect.Value
// string(1), int => int(1)
// string([1,2,3]), slice(int) => []int{1,2,3}
func parseParam(t *testing.T, param string, typ reflect.Type) reflect.Value {
	var r reflect.Value
	var as = assert.New(t)
	param = strings.TrimSpace(param)

	switch typ.Kind() {
	case reflect.Int:
		i, err := strconv.Atoi(param)
		as.Nil(err)
		r = reflect.ValueOf(i)
	case reflect.String:
		r = reflect.ValueOf(param)
	case reflect.Bool:
		b, err := strconv.ParseBool(param)
		as.Nil(err)
		r = reflect.ValueOf(b)
	case reflect.Slice:
		as.True(len(param) > 1)
		as.True(strings.HasPrefix(param, "["))
		as.True(strings.HasSuffix(param, "]"))
		param = strings.TrimPrefix(param, "[")
		param = strings.TrimSuffix(param, "]")
		s2 := strings.Split(param, ",")

		r = reflect.MakeSlice(reflect.SliceOf(typ.Elem()), 0, 0)
		for _, v := range s2 {
			r = reflect.Append(r, parseParam(t, v, typ.Elem()))
		}
	default:
		panic(fmt.Sprintf("not support %s", typ.Kind()))
	}

	return r
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

		ithParamIn := parseParam(t, input[i], ithCallInType)

		in = append(in, ithParamIn)
	}

	out := fv.Call(in)

	// out 有三个，call返回，ft.Out(i)的，output的
	for i := 0; i < ft.NumOut(); i++ {
		ithCallRealOut := out[i] // 比input多的
		ithCallOutType := ft.Out(i)
		ithCallOut := parseParam(t, output[i], ithCallOutType)

		as.Equal(ithCallOut.Kind(), ithCallRealOut.Kind())
		as.Equal(ithCallOut.Kind(), ithCallRealOut.Convert(ithCallOutType).Kind())
		as.Equal(ithCallOut.Interface(), ithCallRealOut.Convert(ithCallOutType).Interface())
	}
}
