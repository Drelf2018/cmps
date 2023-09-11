package cmps_test

import (
	"sync"
	"testing"

	"github.com/Drelf2018/cmps"
)

func TestSliceSafe(t *testing.T) {
	s := cmps.SafeSlice[*Student]{I: make([]*Student, 0)}

	wg := sync.WaitGroup{}
	wg.Add(1)
	task := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			task.Add(1)
			defer task.Done()
			wg.Wait()
			stu := &Student{ID: int64(1000 - i)}
			s.Insert(stu)
			s.Search(stu)
			s.Delete(stu)
		}(i)
	}

	go func() {
		for {
			l := len(s.I)
			if l != 0 {
				print(l, " | ")
			}
		}
	}()

	wg.Done()
	task.Wait()
	println("len(s):", len(s.I))
	if len(s.I) != 0 {
		t.Fail()
	}
}

func TestSlice(t *testing.T) {
	s := make([]*Student, 0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	task := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			task.Add(1)
			defer task.Done()
			wg.Wait()
			stu := &Student{ID: int64(1000 - i)}
			s = cmps.Insert(s, stu)
			cmps.Search(s, stu)
			s = cmps.Delete(s, stu)
		}(i)
	}

	go func() {
		for {
			print(len(s), ", ")
		}
	}()

	wg.Done()
	task.Wait()
	println("len(s):", len(s))
	if len(s) != 0 {
		t.Fail()
	}
}

// go test -run ^TestSlice\w*
