// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Web3Avatar is ERC20 {
    constructor() ERC20("Web3Avatar Token", "W3A") {
        _mint(msg.sender, 100000000 * 10**6);
    }
}
