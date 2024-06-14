var assert = require('assert');
describe('Basic Mocha Test', function () {
    it('should deal with objects', function () {
        var obj = { name: 'John', gender: 'male' };
        var objB = { name: 'John', gender: 'male' };
        assert.deepEqual(obj, objB);
    });

    it('should allow testing null', function () {
        var iAmNull = null;
        assert.equal(iAmNull, null);
    });
});