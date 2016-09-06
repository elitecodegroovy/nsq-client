package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"math"
	"errors"
)
//Go’s mechanism for grouping and naming related sets of methods: interfaces.

type user struct {
	name string

	age int                     //Omitted fields will be zero-valued.
}

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim()float64 {
	return 2 * math.Pi * c.radius
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
/*
Go automatically handles conversion between values and pointers for method calls.
You may want to use a pointer receiver type to avoid copying on method calls
or to allow the method to mutate the receiving struct.
 */
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func doInterface(){
	r := rect{width: 20, height: 40}
	c := circle{radius : 15}

	measure(r)
	measure(c)
}


func ferror(n int)(int, error){
	if n== 99 {
		return -1, errors.New("can't reach the point 99")
	}
	return n + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 99 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func doError(){

	for _, i := range []int{7, 99} {
		if c, e := ferror(i); e != nil {
			fmt.Println("f1 failed:", e)
		}else {
			fmt.Println("value i ", c )
		}
	}

	for _, i := range []int{99, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(99)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}


const GOOD_FAQ string= "How do you know that？"

func plus(x, y int) int{
	return x + y
}

func plusPlus(x, y, z int) int {
	return x + y + z
}

func vals()(int, int ){
	return 100, 200
}

func increaseInt()func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}
//Go supports recursive functions. Here’s a classic factorial example.
func multipleN(n int) int{
	if n== 1 {
		return 1
	}
	return n * multipleN(n -1)
}

func zeroVal(val int){
	val = 1000
}

func zeroOpr(opr *int){
	*opr = 1000
}

func doPoints(){
	i := 0
	zeroVal(i)
	fmt.Println("current zereVal value i= ", i)
	zeroOpr(&i)
	fmt.Println("current zeroptr value i = ", i)

	//struct definition
	fmt.Println(user{name: "John.Lau", age: 29})
	fmt.Println(user{name: "John"})
	s := user{name: "Sean", age: 50}

	u := &s
	fmt.Println("u dereference, name ", u.name)

	//Methods
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	pr := &r
	fmt.Println("pr area:", pr.area())
	fmt.Println("pr perim:", pr.perim())

	//interface
	doInterface()
	doError()
}

func doCloure(){
	nextInt := increaseInt()
	fmt.Println("nextInt", nextInt())
	fmt.Println("nextInt", nextInt())
	fmt.Println("nextInt", nextInt())

	//recursive func
	fmt.Println("recursive func 10*9*8...1 ", multipleN( 10))
	doPoints()
}

// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.\
func sum(nums ...int)int {
	total := 0
	fmt.Println("input", nums)
	for _, v := range nums {
		total += v
	}
	return total
}

func doArrange(){
	num := []int{1 ,3, 5, 7}
	total := 0
	for _, v := range num {
		total += v;
	}
	fmt.Println("sume: ", total)

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
	s := plus(20, 40)
	z := plusPlus(20, 40, 80)
	fmt.Println("plus(20, 40)= ", s, "plusPlus(20, 40 ,80)=", z)
	a, b := vals()
	fmt.Println("multiple return case:",a, b )
	fmt.Println("sum(10, 100, 1000)", sum(10, 100, 1000))
	nums := []int{1, 10, 100, 1000, 10000}
	//If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	fmt.Println("sum(nums): ", sum(nums...))
	doCloure()
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

// A goroutine is a lightweight thread of execution.

func pIndex(form string){
	for i:=0 ; i < 3; i++ {
		fmt.Println(form, ":",  i)
	}
}

func doGo(){
	go func(msg string){
		fmt.Println("msg:", msg)
	}("golang")

	go pIndex("let's go")
	pIndex("show me ")

	//var input string
	//fmt.Scanln(&input)
	//fmt.Println("done")
}
//Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and
// receive those values into another goroutine.
func doChan(){
	message := make(chan string)

	go func(){message <- "ping!"}()

	//This property allowed us to wait at the end of our program for the "ping"
	// message without having to use any other synchronization.
	msg := <- message
	fmt.Println("received msg: ", msg)

	/*
	By default channels are unbuffered, meaning that they will only accept
	 sends (chan <-) if there is a corresponding receive (<- chan) ready
	 to receive the sent value. Buffered channels accept a limited number
	 of values without a corresponding receiver for those values.
	 */
	message2 := make(chan string, 2)
	message2 <- "ping2"
	message2 <- "pong2"
	fmt.Println("received one ", <- message2)
	fmt.Println("received two ", <- message2)

	doChanSync()
}
//This is the function we’ll run in a goroutine. The done channel will be used to
// notify another goroutine that this function’s work is done.
func doChanSync(){
	done := make(chan bool, 1)
	go doAsWorker(done)

	//sync until it has been done.
	<- done
}

func doAsWorker(done chan bool){
	fmt.Println("start to work ...")
	time.Sleep(time.Second)
	fmt.Println("done!")
	done <- true
}

func doConcurrent(){
	doGo()
	doChan()
}

func main() {
	doSimple()
	doConcurrent()
}




