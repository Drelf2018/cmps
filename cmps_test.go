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

func test1(t *testing.T) {
	fs := []cmps.Field{
		{Cmps: 1.14},
		{Cmps: 5.14},
		{Cmps: 4},
		{Cmps: 2},
		{Cmps: 3},
		{Cmps: 0},
	}
	fmt.Printf("fs: %v\n", fs)
	cmps.Slice(fs)
	fmt.Printf("fs: %v\n", fs)
}

func test2(t *testing.T) {
	fmt.Printf("cmps.Compare(true, true): %v\n", cmps.Compare(true, true))
	fmt.Printf("cmps.Compare(true, false): %v\n", cmps.Compare(true, false))
	fmt.Printf("cmps.Compare(false, true): %v\n", cmps.Compare(false, true))
	fmt.Printf("cmps.Compare(false, false): %v\n", cmps.Compare(false, false))
}

func TestMain(t *testing.T) {
	s1 := Student{
		Name: "张三2",
		ID:   1,
		Male: true,
		Scores: struct {
			Chinese float64
			Math    float64
			English float64
		}{
			Chinese: 100,
			Math:    40,
			English: 79,
		},
	}
	s2 := Student{
		Name: "张三2",
		ID:   1,
		Male: true,
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
