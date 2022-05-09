pragma solidity ^0.8.0;

contract dataTest {
    uint index;
    uint public txNum;
    uint64 public startTime;
    mapping(bytes => bool) data;
    mapping(uint64 => bytes) data2;

    function reset(uint64 _startTime) public {
        startTime = _startTime;
        txNum = 0;
    }

    function hasHugeCallData(bytes[] memory inputs) public {
        for (uint i=0; i< inputs.length; i++) {
            data[inputs[i]] = !data[inputs[i]];
        }
    }

    function costManyGas(bytes memory input, uint64 complexity) public {
        for (uint i=0; i< complexity; i++) {
            index += 1;
            data2[index] = input;
        }
        txNum += 1;
    }
}
