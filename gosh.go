package main

import (
  "fmt"
  "bufio"
  "os"
)

const prompt = ": "

func main() {
  inputReader := bufio.NewReader(os.Stdin)
  var input string

  for {
  
    fmt.Print(prompt)
    input, _ = inputReader.ReadString('\n')
    input = input[:len(input)-1]
    
    fmt.Print(parseCommand(input)[3])

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
