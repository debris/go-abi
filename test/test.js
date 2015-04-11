var chai = require('chai');
var assert = chai.assert;
var abi = require('../abi');
var clone = function (object) { return JSON.parse(JSON.stringify(object)); };

var desc = [{
    "constant" : false,
    "inputs" : [{
        "name" : "a",
        "type" : "uint256"
    }],
    "name" : "hello",
    "outputs" : [],
    "type" : "function"
}, {
    "anonymous" : false,
    "inputs" : [],
    "name" : "Hello",
    "type" : "event"
}];

describe('testing encode', function () {
    it('should do simple eval', function () {
        
        // given
        var d = clone(desc); 
        var method = 'hello';
        
        // encode
        var result = abi.encode(d, method, [1]);

        assert.equal('0xb0f0c96a0000000000000000000000000000000000000000000000000000000000000001', result);

    });
});

