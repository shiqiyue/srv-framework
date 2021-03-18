package errors

import (
	"fmt"
	"github.com/shinedone/srv-framework/pkg/jsons"
	"testing"
)

func TestNew(t *testing.T) {
	e := New("test")
	bs, _ := jsons.Marshal(e)
	str := string(bs)
	fmt.Println(str)
}
