// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IIntegers {

  function toString(uint256 _i) view external returns (string memory);
  function toString(int256 _i) view external returns (string memory);
  function toHexString(uint256 _i) view external returns (string memory);
  function fromHexString(string calldata _str) view external returns (uint256);

}
