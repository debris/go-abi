# solabi
command line tool &amp;&amp; library for encoding/decoding solidity input/output

### Build

```bash
go build -o abi ./solabi
```

### Usage

```bash
./abi <mode> <type> <abi_json> <method_name> (<arg1> , (<arg2>, ... ))

modes: enc, dec
type: method, int, uint, bytes, real, ureal, address, bool
abi_json: path to json abi file (only for type "method")
method_name: method you want to encode (only for type "method")
args: argsuments (can be multiple for type "method")
```

### Example


#### encode/decode single param
```bash
# encode/decode single param

./abi enc int 1
0x0000000000000000000000000000000000000000000000000000000000000001

./abi enc int -1
0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff

./abi enc int 0x1
0x0000000000000000000000000000000000000000000000000000000000000001

./abi enc bytes hello
0x000000000000000000000000000000000000000000000000000000000000000568656c6c6f000000000000000000000000000000000000000000000000000000

./abi dec bytes 0x000000000000000000000000000000000000000000000000000000000000000568656c6c6f000000000000000000000000000000000000000000000000000000
hello

./abi dec int 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
-1

./abi dec uint 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
115792089237316195423570985008687907853269984665640564039457584007913129639935
```

#### encode/decode methods && params

```bash
# encode/decode methods && params

./abi enc method example/Test0.abi hello 0x1231
0xb0f0c96a0000000000000000000000000000000000000000000000000000000000001231

./abi dec method example/Test0.abi hello 0x0000000000000000000000000000000000000000000000000000000000000001
1
```

