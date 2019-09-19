package main
// package filehandler

import (
  "fmt"
  "io/ioutil"
)

func check (e error) {
  if e !=nil {
    panic(e)
  }

}

func ReadFile(filename string) (string){
    data, err := ioutil.ReadFile(filename)
    check(err)
    return string(data)
}

func WriteFile(filename string, writestring string) (error) {
  data := ReadFile(filename)
  fmt.Print(data)
  data += writestring
  data += "\n"
  err := ioutil.WriteFile(filename, []byte(data), 0644)
  check(err)
  return err
  }



func main(){
  err := WriteFile("../test.txt", "Hallo")
  check(err)
  fmt.Println("done!")
}
