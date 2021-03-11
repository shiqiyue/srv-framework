package assert

import "errors"

func NotNil(i interface{}, err error) {
	if i == nil {
		if err != nil {
			panic(err)
		} else {
			panic(errors.New("value is  nil"))
		}
	}
}

func Nil(i interface{}, err error) {
	if i != nil {
		if err != nil {
			panic(err)
		} else {
			panic(errors.New("value is not nil"))
		}

	}
}
