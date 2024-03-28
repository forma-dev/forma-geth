// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IStrings {

  // Comparison functions
  function equal(string calldata _a, string calldata _b) view external returns (bool);
  function equalCaseFold(string calldata _a, string calldata _b) view external returns (bool);
  function contains(string calldata _str, string calldata _substr) view external returns (bool);
  function startsWith(string calldata _str, string calldata _substr) view external returns (bool);
  function endsWith(string calldata _str, string calldata _substr) view external returns (bool);

  // ...
  function indexOf(string calldata _str, string calldata _substr) view external returns (uint256);
  function toUpperCase(string calldata _str) view external returns (string memory);
  function toLowerCase(string calldata _str) view external returns (string memory);
  function padStart(string calldata _str, uint16 _len, string calldata _pad) view external returns (string memory);
  function padEnd(string calldata _str, uint16 _len, string calldata _pad) view external returns (string memory);
  function repeat(string calldata _str, uint16 _count) view external returns (string memory);
  function replace(string calldata _str, string calldata _old, string calldata _new, uint16 _n) view external returns (string memory);
  function replaceAll(string calldata _str, string calldata _old, string calldata _new) view external returns (string memory);
  function split(string calldata _str, string calldata _delim) view external returns (string[] memory);
  function trim(string calldata _str) view external returns (string memory);

}
