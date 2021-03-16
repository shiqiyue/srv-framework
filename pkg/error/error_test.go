package error

import (
	"fmt"
	"github.com/shinedone/srv-framework/pkg/json"
	"testing"
)

func TestNew(t *testing.T) {
	e := New("test")
	bs, _ := json.Marshal(e)
	str := string(bs)
	fmt.Println(str)
}
