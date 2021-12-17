//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Web3Avatar {

  mapping(address => mapping(string => string)) private avatars;

  event AvatarUpdated(address indexed _key, string _name, string _avatar);

  function getAvatar(string calldata _name) public view returns(string memory) {
    return avatars[msg.sender][_name];
  }

  function getAvatar(address _addr, string calldata _name) public view returns(string memory) {
    return avatars[_addr][_name];
  }

  // setAvatar("default", "0x00000000000000");
  // setAvatar("etherscan", "0x00000000000000");
  function setAvatar(string calldata _name, string calldata _avatar) public {
    avatars[msg.sender][_name] = _avatar;
    emit AvatarUpdated(msg.sender, _name, _avatar);
  }

}