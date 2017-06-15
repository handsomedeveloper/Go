package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	beautifulSeparationLine("int Type")
	var forCycleData = []int{1, 2, 3}
	x, y := intType(5, 3)
	fmt.Println(x, y)
	forCycle(forCycleData)

	beautifulSeparationLine("string type")
	stringType()

	beautifulSeparationLine("float point type")
	floatPointType()

	beautifulSeparationLine("slice type")
	sliceType()

	beautifulSeparationLine("map type")
	mapType()

	beautifulSeparationLine("struct type")
	structType()

	beautifulSeparationLine("switch type")
	switchType()

	beautifulSeparationLine("chan type")
	chanType()

	beautifulSeparationLine("go-routine")
	goRoutine()
}

//Explain Go language string type
func stringType() {
	var world string

	I, am, name := "I", "am", "SheJie"
	world = "world"
	hello := "Hello"
	rawString := "rawString"
	toSlice := []byte(rawString)
	toString := string(toSlice)

	numbers, errors := fmt.Printf("%d + %d = %d\n", 1, 1, 2)
	if errors == nil {
		fmt.Print("This a error.")
	} else {
		fmt.Print("succese.")
	}
	fmt.Println(numbers, errors)

	fmt.Println(hello, world, toSlice, toString)
	fmt.Println(I, am, name)
}

//Explain Go language floating-point type
func floatPointType() {
	fmt.Println("float32: 7")
	fmt.Println("float64: 15")
}

func forCycle(arg []int) {
	fmt.Println("Tish is number", arg)
	for _, n := range arg {
		fmt.Println("Tish is number", n)
	}
}

func intType(x, y int) (addition int, reduce int) {
	addition = x + y
	reduce = x - y

	var two int = 10  //1010
	var three int = 6 //110
	var result int = two ^ three
	var s = [10]byte{10, 4, 6, 8}
	fmt.Println(s[0])
	s[0] <<= 2
	fmt.Println(s[0])
	fmt.Println(s)
	fmt.Println(result)
	fmt.Println(5 ^ 2)
	fmt.Println(5 & 2)
	fmt.Println(5 & 2)
	fmt.Println(5 | 2)
	return
}

func switchType() {
	var testInterface interface{}
	testInterface = "I love you."
	switch ty := testInterface.(type) {
	case int:
		fmt.Println(ty)
	case string:
		fmt.Println(ty)
	default:
		fmt.Println(ty)
	}

	switch 1 {
	case 0, 1, 2, 3:
		fmt.Println("0~3")
		fallthrough
	case 4, 5, 6, 7:
		fmt.Println("4~7")
	}

	switch {
	case 3 < 1:
		fmt.Println(111)
	case 1 < 3:
		fmt.Println(222)
	case 3 == 4:
		fmt.Println(111)
	}
}

func sliceType() {
	var array [10]int
	var size [5]int64
	fmt.Println(unsafe.Sizeof(size))
	fmt.Println(array)
	slice := array[2:4]
	slice2 := array[2:]
	fmt.Println(unsafe.Sizeof(slice))

	fmt.Println(slice)
	array[3] = 99
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println("slice2:", slice2)
	slice = append(slice, 11, 22, 33, 44, 55)
	fmt.Println(slice)
	fmt.Println(slice2)

	add := append(slice, 11, 22, 33, 44, 55)
	fmt.Println(add)
	fmt.Println(len(add))
	fmt.Println(cap(add))
	fmt.Println(unsafe.Sizeof(add))
}

func mapType() {
	map2 := make(map[string]int)
	map2["one"] = 1
	map2["two"] = 2
	fmt.Println(map2)
	mp := map[string]int{"aa": 11, "bb": 22}
	fmt.Println(mp)
}

type person struct {
	name   string
	age    uint8
	gender string
}

func structType() {
	var p1 person
	p1.name = "one"
	p1.age = 1
	p1.gender = "boy"
	p2 := person{"two", 2, "girl"}
	p3 := person{name: "three", age: 3, gender: "boy"}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}

func chanType() {
	ch := make(chan int, 10)
	go func() { ch <- 3 + 4 }()
	va := <-ch
	fmt.Println("chan", va)

}

func testGo(ch chan int) {
	fmt.Println("This is test.")
	ch <- 1
}

func say(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		//runtime.Gosched()
		fmt.Println(s)
	}
	ch <- 2
}

func goRoutine() {
	ch := make(chan int)
	go say("world", ch)
	go say("Hi", ch)
	go testGo(ch)
	//say("hello", ch)
	gt := runtime.NumGoroutine()
	fmt.Println(gt)
	ch1, ch2, ch3 := <-ch, <-ch, <-ch
	fmt.Println(ch1, ch2, ch3)
}

func beautifulSeparationLine(comment string) {
	fmt.Println("-------------------------------", comment, "------------------------------------")
}
