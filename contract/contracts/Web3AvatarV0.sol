//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Web3AvatarV0 {

  mapping(address => string) hashes;

  function setHash(string calldata _hash) public {
    hashes[msg.sender] = _hash;
  }

  function getHash() public view returns (string memory) {
    return hashes[msg.sender];
  }

  function getHash(address _addr) public view returns (string memory) {
    return hashes[_addr];
  }

}