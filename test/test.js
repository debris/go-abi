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
    "outputs" : [{
        "name" : "b",
        "type" : "uint256"
    }],
    "type" : "function"
}, {
    "anonymous" : false,
    "inputs" : [],
    "name" : "Hello",
    "type" : "event"
}];

describe('testing encode', function () {
    it('should do simple encode', function () {
        // given
        var d = clone(desc); 
        var method = 'hello';
        
        // when
        var result = abi.encodeMethod(d, method, [1]);

        // then
        assert.equal('0xb0f0c96a0000000000000000000000000000000000000000000000000000000000000001', result);
    });
});

describe('testing decode', function () {
    it('should do simple decode', function () {
        // given
        var d = clone(desc); 
        var method = 'hello';

        // when
        var result = abi.decodeMethod(d, method, '0x0000000000000000000000000000000000000000000000000000000000000001');
        
        // then
        assert.deepEqual(result, '1');
    });
});


