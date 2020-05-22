package main

import (
"fmt",
"time"
)

func nameFunction(name string)
{
  for i:= 0; i < 2; i ++ {
   fmt.Println(name, ":" ,i);
   }
}

func main() {
  nameFunciton("FromFunction")
  go nameFunction("FromRoutine")
  
  go func(msg string)
  {
    fmt.Println(msg)
  }("going")
  time.Sleep(time.Second)
  fmt.Println("done")
}
