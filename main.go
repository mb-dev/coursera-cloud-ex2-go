package main

import (
  "os"
)

const FAILURE int = -1
const ARGS_COUNT int = 2

func main() {
  if len(os.Args) != ARGS_COUNT {
    os.Exit(FAILURE)
  }

  var application = NewApplication(os.Args[1])
  application.run()
}
