package main

import (
  "fmt"
)

type Application struct {
  infile string
  mp1 []MP1Node
  mp2 []MP2Node
  en EmulNet
  en1 EmulNet
}

func NewApplication(infile string) *Application {
  app := new (Application)
  app.infile = infile
  // TOOD: Load params file
  gpsz := 3
  app.en = EmulNet{}
  app.en1 = EmulNet{}

  for i:=0; i < gpsz; i++ {
    member := new (Member)
    address := app.en.enInit()
    app.mp1 = append(app.mp1, MP1Node{*member, app.en, address})
    app.mp2 = append(app.mp2, MP2Node{*member, app.en1, address})
  }

  return app
}

func (app Application) run() {
  fmt.Printf("Hello")
}
