package main

import (
    "os"
    "fmt"
    "errors"
    "io/ioutil"

    "github.com/debris/abi"
)

type mode string

const (
    Encode mode = "enc"
    Decode = "dec"
)

type options struct {
    abi string
    mode
    method string
    params []string
}

func help() (str string) {
    str = `
Usage:
    abi <mode> <abi_json> <method_name> (<arg1> , (<arg2>, ... ))

    modes: enc, dec
    abi_json: path to json abi file
    method_name: method you want to encode
    args: method argsuments

Example:
    abi enc myContract.abi myMethod 12 0x123 hello
    `
    return
}

func parseArgs(args []string) (o *options, err error) {
    o = new(options)

    if len(args) < 5 {
        err = errors.New("Not enought params")
        return
    }

    bytes, err := ioutil.ReadFile(args[2])

    if err != nil {
        return
    }

    o.abi = string(bytes[:])

    if args[3] != string(Encode) && args[3] != string(Decode) {
        err = errors.New("mode must be 'enc' or 'dec'")
        return
    }

    o.mode = mode(os.Args[3])
    o.method = os.Args[4]
    o.params = os.Args[5:]

    return
}

func execute(o *options) (result string, err error) {
    coder, err := abi.New(o.abi);

    if err != nil {
        return
    }

    if o.mode == Encode {
        result, err = coder.Encode(o.method, o.params)
    } else if o.mode == Decode {
        result, err = coder.Decode(o.method, o.params)
    }

    return
}

func main() {
    o, err := parseArgs(os.Args)
    if err != nil {
        fmt.Println(err)
        fmt.Println(help())
        os.Exit(1)
    }

    result, err := execute(o)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(result)
}
