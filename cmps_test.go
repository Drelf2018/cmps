package cmps_test

import (
	"fmt"
	"testing"

	"github.com/Drelf2018/cmps"
)

type Scores struct {
	Chinese float64 `cmps:"1;groups:chinese"`
	Math    float64 `cmps:"2;groups:math"`
	English float64
}

type Student struct {
	Name   string
	ID     int64 `cmps:"810"`
	Male   bool  `cmps:"514"`
	Scores `cmps:"1919;fields:Math,Chinese,English"`
}

func testCmpsOrdered() bool {
	// int
	if cmps.Compare(1, 2) != -1 {
		return false
	}
	if cmps.Compare(1, 1) != 0 {
		return false
	}
	if cmps.Compare(2, 1) != 1 {
		return false
	}
	// float
	if cmps.Compare(1.5, 2.5) != -1 {
		return false
	}
	if cmps.Compare(1.5, 1.5) != 0 {
		return false
	}
	if cmps.Compare(2.5, 1.5) != 1 {
		return false
	}
	// bool
	if cmps.Compare(false, true) != -1 {
		return false
	}
	if cmps.Compare(false, false) != 0 {
		return false
	}
	if cmps.Compare(true, true) != 0 {
		return false
	}
	if cmps.Compare(true, false) != 1 {
		return false
	}
	// string
	if cmps.Compare("1", "2") != -1 {
		return false
	}
	if cmps.Compare("1", "1") != 0 {
		return false
	}
	if cmps.Compare("2", "1") != 1 {
		return false
	}
	return true
}

func testStruct() bool {
	s1 := Student{
		Name: "张三1",
		ID:   2,
		Male: false,
		Scores: Scores{
			Chinese: 90,
			Math:    60,
			English: 80,
		},
	}
	s2 := Student{
		Name: "张三2",
		ID:   2,
		Male: false,
		Scores: Scores{
			Chinese: 90,
			Math:    60,
			English: 80,
		},
	}
	// Order: Male ID Math Chinese English
	if cmps.Compare(s1, s2) != 0 {
		return false
	}
	s2.Male = true
	if cmps.Compare(s1, s2) != -1 {
		return false
	}
	s1.Male = true
	if cmps.Compare(s1, s2) != 0 {
		return false
	}
	s2.Male = false
	if cmps.Compare(s1, s2) != 1 {
		return false
	}
	s1.Male = false
	s1.ID = 1
	if cmps.Compare(s1, s2) != -1 {
		return false
	}
	s1.ID = 3
	if cmps.Compare(s1, s2) != 1 {
		return false
	}
	s1.ID = 2
	s1.Math = 50
	if cmps.Compare(s1, s2) != -1 {
		return false
	}
	s1.Math = 70
	if cmps.Compare(s1, s2) != 1 {
		return false
	}
	s1.Math = 60
	s1.Chinese = 80
	if cmps.Compare(s1, s2) != -1 {
		return false
	}
	s1.Chinese = 100
	if cmps.Compare(s1, s2) != 1 {
		return false
	}
	s1.Chinese = 90
	s1.English = 70
	if cmps.Compare(s1, s2) != -1 {
		return false
	}
	s1.English = 90
	return cmps.Compare(s1, s2) == 1
}

type Product struct {
	Name  string `cmps:"1"`
	Price string `cmps:"2;order:desc"`
}

func TestInsert(t *testing.T) {
	var ps []*Product
	ps = cmps.Insert(ps, &Product{"2", "1.99"})
	ps = cmps.Insert(ps, &Product{"2", "2.99"})
	ps = cmps.Insert(ps, &Product{"1", "3.99"})
	ps = cmps.Insert(ps, &Product{"1", "0.99"})
	ps = cmps.Insert(ps, &Product{"3", "4.99"})

	i, ok := cmps.SearchFunc(ps, func(p *Product) {
		p.Name = "2"
		p.Price = "2.99"
	})
	fmt.Printf("i: %v ok: %v\n", i, ok)
	for _, p := range ps {
		fmt.Printf("p: %v\n", p)
	}
	fmt.Println()
	ps = cmps.Delete(ps, &Product{"3", "4.99"})
	for _, p := range ps {
		fmt.Printf("p: %v\n", p)
	}
}

func TestMain(t *testing.T) {
	i := 1
	for ; i > 0; i-- {
		if !testCmpsOrdered() {
			break
		}
		if !testStruct() {
			break
		}
	}
	if i != 0 {
		t.Fail()
	}
}
