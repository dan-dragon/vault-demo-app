package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "log"
)

func main() {
  contents, err := ioutil.ReadFile(os.Getenv("SECRETS_FILE"))
  if err != nil {
    log.Fatal(err)
  }

  fmt.Fprintf(os.Stdout, "%s", contents)
}
