package main

import (
   "fmt"
   "io"
   "os"
)

type logWriter struct {}

func main() {
   lw := logWriter{}

   r, err := os.Open(os.Args[1])

   if err != nil {
      fmt.Println("oh something go wrong...")
      os.Exit(1)
   }

   io.Copy(lw, r)
}

func (logWriter) Write(bs []byte) (int, error) {

   fmt.Println(string(bs))

   return len(bs), nil
}

