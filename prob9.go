package main

import (
    "fmt"
    "math"
)

func VerifyTriplet(c float64) (res bool) {
    // If there are no decimals, it is an integer root
    if c-math.Floor(c) == 0.0 {
        fmt.Printf("The dif is: %f\n", c-math.Floor(c))
        return true
    }
    return false
}

func FindTriplets() (a, b, c int){
    broken := false
    for a := 1; a < 1000; a++ { // For each A < 1000
        for b :=1; !broken && b < 1000; b++ { // To B until loop breaks

            c := math.Sqrt(float64((a*a)+(b*b)))
            fmt.Printf("Going a=%d, b=%d, c=%f\n",a,b,c)
            if (a+b+int(c)) > 1000 {
                broken = true
                fmt.Printf("Broken! a=%d, b=%d, c=%f \n",a,b,c)
            } else if (a+b+int(c)) == 1000 && VerifyTriplet(c) {
                return a,b,int(c)
            }
        }
        broken = false
    }

    return 0,0,0
}

func main() {
    a,b,c := FindTriplets()
    fmt.Printf("a=%d, b=%d, and c=%d, abc=%d\n",a,b,c,(a*b*c))
}
