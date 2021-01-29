package main

import "fmt"

const pai = 3.1415

// not working so far
// const {
// 	firstOne = 1
// 	secondOne = "second"
// }

func main() {
	// hello world!
	fmt.Println("Hello World!")

	// primitive data types
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var pointer1 *string = new(string)
	*pointer1 = "derefereing"
	fmt.Println(*pointer1)

	lastName := "Arthur"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Tricia"
	fmt.Println(ptr, *ptr)

	const pi = 3.1415
	fmt.Println(pi)
	// pi = 1.2 => invalid

	const d int = 3
	fmt.Println(d + 3)
	fmt.Println(float32(d) + 1.2)

	fmt.Println(pai)

	// collections
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1]) // don't go beyond the boundary

	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)

	// slice := arr2[:]
	// arr2[1] = 42
	// slice[2] = 27
	// fmt.Println(slice)

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 4, 42, 27)
	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)

	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Yue"
	u.LastName = "Liu"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{ID: 1,
		FirstName: "Yue",
		LastName:  "Liu", // be careful where to put the closing curly braces
	}
	fmt.Println(u2)

}
