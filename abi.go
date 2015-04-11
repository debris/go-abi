package abi

import (
    "strings"
    "errors"

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

    data, err := Asset("abi.min.js");
    if err != nil {
        return
    }

    web3 := string(data[:])

    if _, err = self.jsvm.Run(web3); err != nil {
        return
    }

    if _, err = self.jsvm.Run("var abi = require('abi');"); err != nil {
        return
    }

    return
}

func toJsonArray(params []string) (arr string) {
    arr = "[]"
    if len(params) > 0 {
        arr = "[\"" + strings.Join(params, "\", \"") + "\"]"
    }
    return
}

func load(abi *Abi, method string, params []string) (error) {
    if _, err := abi.jsvm.Run("var json = " + abi.json + ";"); err != nil {
        return errors.New("incorrect json file")
    }

    if _, err := abi.jsvm.Run("var method = \"" + method + "\";"); err != nil {
        return errors.New("incorrect method name")
    }

    arr := toJsonArray(params)

    if _, err := abi.jsvm.Run("var params = " + arr + ";"); err != nil {
        return errors.New("incorrect method name")
    }

    return nil
}

func (self *Abi) Encode(method string, params []string) (result string, err error) {
    if err = load(self, method, params); err != nil {
        return
    }

    value, err := self.jsvm.Run("abi.encode(json, method, params)")

    if err != nil {
        return
    }

    result = value.String()

    return
}

func (self *Abi) Decode(method string, params []string) (result string, err error) {
    if err = load(self, method, params); err != nil {
        return
    }

    value, err := self.jsvm.Run("abi.decode(json, method, params)")

    if err != nil {
        return
    }

    result = value.String()

    return
}

