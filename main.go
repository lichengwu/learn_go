package main

import "fmt"

const CT_STR = "good"

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

type Vex  struct {
	x int
	y string
}

func main() {
	//fmt.Print(split(10))
	//var i int = 4
	//var f float64 = float64(i)
	//var d uint = uint(f)
	//fmt.Printf("%T,%T,%T\n", i, f, d)
	//fmt.Println(CT_STR)

	//max := 10
	//for i := 0; i < max; i++ {
	//	fmt.Println(i)
	//}
	//
	//if v := max; v < 100 {
	//	fmt.Print(v)
	//}

	//x := Vex{1, "ok"}
	//x.x = 12
	//
	//fmt.Println(x)
	//fmt.Print(x.y)

	a := make([]int, 0, 10)
	fmt.Println(cap(a))
	fmt.Println(len(a))
	a[9] = 10
	fmt.Println(a[1:8])

}