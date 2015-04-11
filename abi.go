package abi

import (
    "github.com/robertkrimen/otto"
)

type Abi struct {
    jsvm *otto.Otto
    json string
}

func New(json string) (self *Abi, err error) {

    self = new(Abi)
    self.jsvm = otto.New()
    self.json = json

    data, err := Asset("web3.min.js");
    if err != nil {
        return
    }

    web3 := string(data[:])

    _, err = self.jsvm.Run(web3)
    if err != nil {
        return
    }

    return
}

func (self *Abi) Encode(method string, params []string) (result string, err error) {
    return
}

func (self *Abi) Decode(method string, params []string) (result string, err error) {
    return
}

