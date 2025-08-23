// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleWallet {
    receive() external payable {}

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
function getaddr() public view returns (address) {
    return msg.sender;
}
}