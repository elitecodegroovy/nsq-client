package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const GOOD_FAQ string= "How do you know that？"

func doArrange(){
	num := []int{1 ,3, 5, 7}
	sum := 0
	for _, v := range num {
		sum += v;
	}
	fmt.Println("sume: ", sum)

	//index value
	for i, v := range num {
		if v == 7 {
			fmt.Println("value 7 , index ", i)
		}
	}

	//iterate the map kv
	jobs := map[string]string{"C++": "8k", "Java": "10"}
	for k, v := range jobs {
		fmt.Println("k:", k, ",v:", v)
	}

	//range on strings iterates over Unicode code points. The first value is
	// the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
func doMaps(){
	m := make(map[string]int)
	m["ok"] = 1
	m["failed"] = 0
	fmt.Println("map[string]int", m)

	m["success"] = 2
	fmt.Println("len:", len(m))

	v := m["ok"]
	fmt.Println("key:ok, value:", v)

	//delete
	delete(m, "ok")
	fmt.Println("Delete kv [\"ok\"]", m)

	//judge it exits or not.
	_, pros := m["success"]
	fmt.Println("result search for a key 'success':", pros)

	//init values
	x := map[string]int{"return": 1, "rssult":0}
	fmt.Println("map init:", x)

	doArrange()
}
func doSlice(){
	s:= make([]string, 3)
	fmt.Println("string slice", s)

	s[0] = "0"
	s[2] = "2"
	s[1] = "1"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Print("len s", len(s))
	s = append(s, "c")
	fmt.Println("s appended,", s)

	s = append(s, "d", "e")
	fmt.Println("s appended many string", s)

	//copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy c :", c)

	//cut arrange
	x := s[2:5]
	fmt.Println("arrange :", x)

	x = s[2:]
	fmt.Println("s[2:]: ", x)

	x = s[:2]
	fmt.Println("s[2:]:", x)
	doMaps()
}

func doArray(){
	var a [5]int
	fmt.Println("emp",a )

	//assign value
	a[4] = 100
	fmt.Println("set ", a, ", value: a[4]:", a[4] )

	b := [5]int{12, 13, 14, 15,17}
	fmt.Println("b  value:", b)
	doSlice()
}

func doSimple(){
	who := "World!"
	if len(os.Args) > 1 {
		/* os.Args[0] is "hello" or "hello.exe" */
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println(who)
	fmt.Println(fmt.Sprintf("print statement: %s", GOOD_FAQ))
	const n = 50000000
	const d = 3e20 / n
	fmt.Println(d)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
	doArray()
}
func main() {
	doSimple()
}




