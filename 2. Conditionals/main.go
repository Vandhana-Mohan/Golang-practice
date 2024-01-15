package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)
  if eludedGuards >= 50{
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  openedVault := rand.Intn(100)

  if isHeistOn && openedVault >= 70{
    fmt.Println("Grab and GO!")
  } else{
    isHeistOn = false
    fmt.Println("Vault cant be opened")
  }

  leftSafely := rand.Intn(5)
  amtStolen:= 10000 + rand.Intn(1000000)
  if isHeistOn {
    switch leftSafely{
      case  0:
          isHeistOn = false
          fmt.Println("heist failed")
      case 1: 
          isHeistOn = false 
          fmt.Println("Turns out vault doors don't open from the inside...")  
      case 2: 
          isHeistOn = false 
          fmt.Println("security camera..")  
      default:
          isHeistOn = true 
          fmt.Println("Start the getaway car!")  
          fmt.Println("Money stolen : ", amtStolen)
    }
  }
  fmt.Println("The status of isHeistOn is ", isHeistOn)
  fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

  defer fmt.Println("world")

	fmt.Println("hello")
  
}
