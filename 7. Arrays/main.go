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