package main

import (
  "os"
  "fmt"
  "io"
  "syscall"
  "os/exec"
  "log"
)


type Cmd struct  {
  Path, Dir string
  Args[] string
  Env [] string
  Stdin io.Reader
  Stdout, Stderr io.Writer
  // ExtraFiles[] *os.file
  SysProcAttr *syscall.SysProcAttr
  Process *os.Process
  ProcessState *os.ProcessState
}

// var cmd Cmd
var output string
var err error
var counter int

func main(){
  output,err := exec.Command("docker", "ps").Output()
   // err := cmd.Run()
  fmt.Println("\n Command finished with error: ", err)
  fmt.Println("fmt - Output:",output)
  log.Printf("Command finished with errors: %v ", err)
  log.Printf("\n log - Output: %v", output)
  // var i int
  // i = 0
  for i :=0; i<100; i++ {
    if output[i] == 32 {
      counter ++
    }
  }
  log.Print(string(output))
  fmt.Println()
  output = string(output)
  fmt.Println(counter)
  fmt.Println(output)
}
