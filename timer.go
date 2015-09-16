package main

import (
    "fmt"
    "time"
)

func sleep(secs int) {
    select {
          case <-time.After(time.Duration(secs) * time.Second):
            return
    }
}

func main() {
    
    fmt.Print("Enter a number: ")
    var x int
    fmt.Scanf("%d", &x)
    fmt.Println(time.Now())
    sleep(x)
    fmt.Println(time.Now())
}