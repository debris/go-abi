
var sha3 = require('crypto-js/sha3');
var abi = require('./node_modules/web3/lib/solidity/abi');

module.exports = (function () {

    var dumbFill = function (json) {
        // TODO: we may not want to modify input params, maybe use copy instead?
        json.forEach(function (method) {
            if (method.name.indexOf('(') === -1) {
                var displayName = method.name;
                var typeName = method.inputs.map(function(i){return i.type; }).join();
                method.name = displayName + '(' + typeName + ')';
            }
        });
        return json;
    }

    var getFullMethodName = function (json, name) {
        dumbFill(json);
        return json.filter(function (method) {
            return method.name.indexOf(name + '(') === 0;
        })[0].name;
    };

    var functionSignature = function (name) {
        return sha3(name, {
            outputLength: 256
        }).toString().slice(0, 8);
    };

    var eventSignature = function (name) {
        return sha3(name, {
            outputLength: 256
        }).toString();
    };

    var encodeMethod = function (json, method, params) {

        var parser = abi.inputParser(json);
        var name = getFullMethodName(json, method);
        var signature = "";
        var encoded = "";

        try {
            signature = functionSignature(name);
            encoded = parser[method].apply(null, params);
        } catch (e) {
            throw new Error("Incorrect method name: " + method);
        }

        return "0x" + signature + encoded;
    };

    var decodeMethod = function (json, method, param) {

        var parser = abi.outputParser(json);
        var decoded = "";
        
        // TODO: that's dumb, fix this in web3.js
        if (param.indexOf('0x') !== 0) {
            param = '0x' + param;
        }

        try {
            decoded = parser[method].apply(null, [param]);
        } catch (e) {
            throw new Error("Incorrect method name: " + method);
        }

        return decoded[0].toString(10);
    };

    var encodeParam = function (kind, param) {

    };


    var decodeParam = function (kind, param) {

    };

    return {
        encodeMethod: encodeMethod,
        decodeMethod: decodeMethod,
        encodeParam: encodeParam,
        decodeParam: decodeParam
    };
})();

