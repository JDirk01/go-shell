package cmd

import(
  "fmt"
  "os"
)

func GoshPwd(args []string) {
  currdir, err := os.Getwd()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(currdir)
  return
}
