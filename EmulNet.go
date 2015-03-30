package main

import (
  "strconv"
)

type EmulNet struct {
  nextid int
}

func NewEmulNet() *EmulNet {
  em := &EmulNet{
    nextid: 1,
  }

  return em
}

func (emul EmulNet) enInit() string {
  emul.nextid += 1
  return strconv.Itoa(emul.nextid)
}
