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
    mode
    kind string
    abi string
    method string
    params []string
}

func help() (str string) {
    str = `
Usage:
    abi <mode> <type> <abi_json> <method_name> (<arg1> , (<arg2>, ... ))

    modes: enc, dec
    type: method, int, uint, bytes, real, ureal, address, bool
    abi_json: path to json abi file (only for type "method")
    method_name: method you want to encode (only for type "method")
    args: argsuments (can be multiple for type "method")

Example:
    abi enc myContract.abi myMethod 12 0x123 hello
    `
    return
}

func parseArgs(args []string) (o *options, err error) {
    o = new(options)

    if len(args) < 3 {
        err = errors.New("Not enought params")
        return
    }

    switch args[1] {
        case string(Encode), string(Decode):
        o.mode = mode(args[1]); break
        default:
        err = errors.New("mode must be 'enc' or 'dec'"); return

    }

    o.kind = args[2];
    if o.kind != "method" {
        o.params = os.Args[3:]
        return
    }

    if len(args) < 5 {
        err = errors.New("Not enought params to encode/decode method")
        return
    }

    bytes, err := ioutil.ReadFile(args[3])

    if err != nil {
        return
    }

    o.abi = string(bytes[:])
    o.method = os.Args[4]
    o.params = os.Args[5:]

    return
}

func execute(o *options) (result string, err error) {
    coder, err := abi.New();

    if err != nil {
        return
    }

    if o.kind == "method" {
        if o.mode == Encode {
            result, err = coder.EncodeMethod(o.abi, o.method, o.params)
        } else {
            result, err = coder.DecodeMethod(o.abi, o.method, o.params[0])
        }
    } else {
        if o.mode == Encode {
            result, err = coder.EncodeParam(string(o.kind), o.params[0])
        } else {
            result, err = coder.DecodeParam(string(o.kind), o.params[0])
        }
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
