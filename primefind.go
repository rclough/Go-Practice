package main

import (
    "math"
    "container/list"
    "fmt"
    "flag"
    "strconv"
)

var printTrail = flag.Bool("p",false,"Print the trail of primes")

func IsPrime(number int, primes *list.List) (isPrime bool) {

    for i := primes.Front(); 
            i != nil && i.Value.(int) <= int(math.Sqrt(float64(number))); 
            i = i.Next() {

        if ((number % i.Value.(int)) == 0) { return false }
    }

    return true

}

func DiscoverPrime(index int) (result int){
    // Set up variables
    primes := list.New()
    primes.PushBack(2)

    // For 
    for i := 1; primes.Len()< index; i++ {
        current := primes.Back().Value.(int) + i
        if IsPrime(current, primes) {
            primes.PushBack(current)
            if *printTrail {
                fmt.Printf("%d is prime \n", current)
            }
            i = 1
        }
    }

    return primes.Back().Value.(int)
}

func main() {
    // Deal with Flags
    flag.Parse()
    if flag.NArg() == 1 && (flag.NFlag() == 0 || flag.NFlag() == 1) {
        number,err:= strconv.Atoi(flag.Arg(0))

        if err == nil {
            result := DiscoverPrime(number)
            fmt.Printf("Answer is: %d\n", result);
        } else {
            fmt.Printf("Encountered an error parsing the number you gave\n")
        }
    } else {
        fmt.Printf("You messed up your comand line args\n")
    }
}
