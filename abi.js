
var sha3 = require('crypto-js/sha3');
var abi = require('./node_modules/web3/lib/solidity/abi');

module.exports = (function () {
    var encode = function () {
        return sha3('hello', {
            outputLength: 256
        });
    };

    var decode = function () {
    };

    return {
        encode: encode,
        decode: decode
    };
})();

