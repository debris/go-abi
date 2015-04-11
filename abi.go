package abi

import (
    "strings"
    "errors"

    "github.com/robertkrimen/otto"
)

type Abi struct {
    jsvm *otto.Otto
}

func New() (self *Abi, err error) {
    self = new(Abi)
    self.jsvm = otto.New()

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

func (self *Abi) EncodeMethod(json string, method string, params []string) (result string, err error) {
    if _, err = self.jsvm.Run("var json = " + json + ";"); err != nil {
        err = errors.New("incorrect json file")
        return
    }

    if _, err = self.jsvm.Run("var method = \"" + method + "\";"); err != nil {
        err = errors.New("incorrect method name")
        return
    }

    arr := toJsonArray(params)

    if _, err = self.jsvm.Run("var params = " + arr + ";"); err != nil {
        err = errors.New("incorrect params")
        return
    }

    value, err := self.jsvm.Run("abi.encodeMethod(json, method, params)")

    if err != nil {
        return
    }

    result = value.String()

    return
}

func (self *Abi) DecodeMethod(json string, method string, param string) (result string, err error) {
    if _, err = self.jsvm.Run("var json = " + json + ";"); err != nil {
        err = errors.New("incorrect json file")
        return
    }

    if _, err = self.jsvm.Run("var method = \"" + method + "\";"); err != nil {
        err = errors.New("incorrect method name")
        return
    }

    if _, err = self.jsvm.Run("var param = \"" + param + "\";"); err != nil {
        err = errors.New("incorrect param")
        return
    }

    value, err := self.jsvm.Run("abi.decodeMethod(json, method, param)")

    if err != nil {
        return
    }

    result = value.String()

    return
}

func (self *Abi) EncodeParam(kind string, param string) (result string, err error) {

    if _, err = self.jsvm.Run("var kind = \"" + kind + "\";"); err != nil {
        err = errors.New("incorrect kind")
        return
    }

    if _, err = self.jsvm.Run("var param = \"" + param + "\";"); err != nil {
        err = errors.New("incorrect param")
        return
    }

    value, err := self.jsvm.Run("abi.encodeParam(kind, param)")

    if err != nil {
        return
    }

    result = value.String()

    return
}

func (self *Abi) DecodeParam(kind string, param string) (result string, err error) {
    if _, err = self.jsvm.Run("var kind = \"" + kind + "\";"); err != nil {
        err = errors.New("incorrect kind")
        return
    }

    if _, err = self.jsvm.Run("var param = \"" + param + "\";"); err != nil {
        err = errors.New("incorrect param")
        return
    }

    value, err := self.jsvm.Run("abi.decodeParam(kind, param)")

    if err != nil {
        return
    }

    result = value.String()

    return
}
