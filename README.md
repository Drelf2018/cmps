# cmps

通过 tag 比较相同结构体不同实例大小

详见 cmps_test.go

```go
package cmps_test

import (
	"fmt"
	"testing"

	"github.com/Drelf2018/cmps"
)

type Student struct {
	Name   string `cmps:"114"`
	ID     int64  `cmps:"810"`
	Male   bool   `cmps:"514"`
	Scores struct {
		Chinese float64
		Math    float64
		English float64
	} `cmps:"1919;Math,Chinese,English"`
}

func TestMain(t *testing.T) {
	s1 := Student{
		Name: "张三2",
		ID:   1,
		Male: false,
		Scores: struct {
			Chinese float64
			Math    float64
			English float64
		}{
			Chinese: 100,
			Math:    40,
			English: 81,
		},
	}
	s2 := Student{
		Name: "张三2",
		ID:   1,
		Male: false,
		Scores: struct {
			Chinese float64
			Math    float64
			English float64
		}{
			Chinese: 100.00,
			Math:    40,
			English: 80,
		},
	}
	match := cmps.Compare(s1, s2)
	fmt.Printf("match: %v\n", match)
}
```