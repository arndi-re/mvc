package main

import (
"github.com/andie-re/mvc/initializer" 
"fmt")

func init(){
  fmt.Println("Starting app...")
  initializer.ConnectToDatabase()
}

func main(){
  fmt.Println("Hello")
}

