package main
import "fmt"
import "slices"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    changeLastElement(myTutors, "Bobby")
    fmt.Println(myTutors)
    theirHours := [4]int{20,10,30,15}
    theirHoursSlice := theirHours[:]
    changeFirst(theirHoursSlice, 25)
    fmt.Println(theirHoursSlice)
    printFirstLastArray(theirHours)


    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)


    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    // twoD := make([][]int, 3)
    // for i := 0; i < 3; i++ {
    //     innerLen := i + 1
    //     twoD[i] = make([]int, innerLen)
    //     for j := 0; j < innerLen; j++ {
    //         twoD[i][j] = i + j
    //     }
    // }
    // fmt.Println("2d: ", twoD)

    var ade [2]string
	ade[0] = "Hello"
	ade[1] = "World"
	fmt.Println(ade[0], ade[1])
	fmt.Println(ade)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

    names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aaa := names[0:2]
	bbb := names[1:3]
	fmt.Println(aaa, bbb)

	bbb[0] = "XXX"
	fmt.Println(aaa, bbb)
	fmt.Println(names)
}


func changeLastElement(slice [] string, newName string){
  slice[len(slice) - 1] = newName
  fmt.Println(slice)
}

func printFirstLastArray(array [4]int) {
    fmt.Println("First", array[0])
    fmt.Println("Last", array[3])
}


func changeFirst(slice []int, value int) {
    if (len(slice) > 0) {
        slice[0] = value
    }
}