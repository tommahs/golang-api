package filehandler

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
  data += writestring
  data += "\n"
  err := ioutil.WriteFile(filename, []byte(data), 0644)
  check(err)
  return err
  }


func Main(filename string, err error) {
  WriteFile(filename, fmt.Sprintf("%s", err)) // // TODO: testing filehandlers
  check(err)
  fmt.Println("done!")
}
