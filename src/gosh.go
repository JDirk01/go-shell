package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "goshell/cmd"
)

const prompt = ": "

func main() {
  inputReader := bufio.NewReader(os.Stdin)
  var input string

  for {
  
    fmt.Print(prompt)
    input, _ = inputReader.ReadString('\n')
    input = input[:len(input)-1]
    
    executeCommand(parseCommand(input))
  }
 return  
}

func parseCommand(command string) []string{
  //Tokenize command into array of strings [0] == command, [..] == args
  command += " "
  tokenized := []string{}
  var left int = 0
  var curr_arg string
  
  for i := 0; i < len(command); i++{
    if command[i] == ' '{
      curr_arg = command[left:i]
      left = i + 1
      tokenized = append(tokenized, curr_arg)
    }
  }

  return tokenized
}


func executeCommand(command []string) {
  //check if first index of command is cd, or exit
  if command[0] == "pwd"{
    cmd.GoshPwd(command)
    return
  }
  //combine tokenized array into string to execute command using exec library function
  joined_cmd := strings.Join(command, " ") 
  fmt.Print("Executing command: ", joined_cmd)
  return
}

/*func goshPwd(args []string) {
  currdir, err := os.Getwd()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(currdir)
  return
}*/
