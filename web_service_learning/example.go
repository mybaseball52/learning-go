package main

import (
	"errors"
	"fmt"
)

const (
	first   = 1
	second2 = "second"
	third   = iota
	fourth  = iota
)

func main2() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32
	f = 1.01
	fmt.Println(f)

	firstName := "Kevin"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//pointer
	var firstN *string = new(string)

	*firstN = "Kevin"
	//fmt.Println(firstN) //0xc000010210
	fmt.Println(*firstN)

	secondName := "HUang"

	ptr := &secondName
	fmt.Println(ptr, *ptr)

	secondName = "Huang"
	fmt.Println(ptr, *ptr)

	const pi float32 = 3.14
	fmt.Println(pi)
	fmt.Println(pi + 1.2)

	//Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr[:]
	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

	slice = append(slice, 1, 2, 3)
	fmt.Println(arr, slice)

	s2 := slice[1:]
	s3 := slice[:2]
	fmt.Println(s2, s3)

	//Array map
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	//struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	fmt.Println(u)

	u.ID = 1
	u.FirstName = "Kevin"
	u.LastName = "Huang"
	fmt.Println(u)

	u2 := user{
		ID:        2,
		FirstName: "Kevin",
		LastName:  "Huang",
	}
	fmt.Println(u2)

	fmt.Println(test(1000))
	_, err := test(300) //_ can be not used but still compiled

	fmt.Println(err)

	panicFunc()
}

func test(port int) (int, error) {
	fmt.Println("start the service")

	return port, errors.New("Something is wrong")
}

func panicFunc() {
	panic("Something went wrong")
}

func loops() {
	for i := 0; i < 5; i++ {
		println(i)
		if i == 3 {
			break
		}
	}

	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	//or
	for i, v := range slice {
		println(i, v)
	}

	wellKnowPorts := map[string]int{"Http": 80, "Https": 443}

	for _, v := range wellKnowPorts {
		println(v)
	}
}
