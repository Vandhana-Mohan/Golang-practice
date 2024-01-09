package main
import "fmt"

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