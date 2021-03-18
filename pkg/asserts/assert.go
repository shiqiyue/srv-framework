package asserts

import "errors"

// 断言，不为nil
func NotNil(i interface{}, err error) {
	if i == nil {
		if err != nil {
			panic(err)
		} else {
			panic(errors.New("value is  nil"))
		}
	}
}

// 断言，为nil
func Nil(i interface{}, err error) {
	if i != nil {
		if err != nil {
			panic(err)
		} else {
			panic(errors.New("value is not nil"))
		}

	}
}
