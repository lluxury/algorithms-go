package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func (r *Interval) Unmarshal(data string) (interface{}, error) {
	invalidErr := fmt.Errorf("invalid data(%s)", data)

	if !strings.HasPrefix(data, "[") || !strings.HasSuffix(data, "]") {
		return nil, invalidErr
	}

	data = strings.Replace(data[1:len(data)-1], " ", "", -1)

	datas := strings.Split(data, ",")
	if len(datas) != 2 {
		return nil, invalidErr
	}

	var s = new(Interval)
	var err error
	s.Start, err = strconv.Atoi(datas[0])
	if err != nil {
		return nil, err
	}
	s.End, err = strconv.Atoi(datas[1])
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Interval) Marshal() (string, error) {
	if r == nil {
		return "[]", nil
	}

	return fmt.Sprintf("[%d,%d]", r.Start, r.End), nil
}
