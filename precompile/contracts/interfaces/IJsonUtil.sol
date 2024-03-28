// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IJsonUtil {

  /////////////////////////////////////// READ METHODS (on jsonBlob) //////////////////////////////////////////

  function get(string calldata jsonBlob, string calldata path) view external returns (string memory);

  function getInt(string calldata jsonBlob, string calldata path) view external returns (int256);

  function getUint(string calldata jsonBlob, string calldata path) view external returns (uint256);

  function getBool(string calldata jsonBlob, string calldata path) view external returns (bool);

  function dataURI(string calldata jsonBlob) view external returns (string memory);

  function exists(string calldata jsonBlob, string calldata path) view external returns (bool);

  function validate(string calldata jsonBlob) view external returns (bool);

  ////////////////////////////////////// MODIFICATION VIEWS (on jsonBlob) //////////////////////////////////////////

  function compact(string calldata jsonBlob) view external returns (string memory);

  function set(string calldata jsonBlob, string calldata path, string calldata value) view external returns (string memory);

  function set(string calldata jsonBlob, string[] memory paths, string[] memory values) view external returns (string memory);

  function setRaw(string calldata jsonBlob, string calldata path, string calldata rawBlob) view external returns (string memory);

  function setRaw(string calldata jsonBlob, string[] memory paths, string[] memory rawBlobs) view external returns (string memory);

  function setInt(string calldata jsonBlob, string calldata path, int256 value) view external returns (string memory);

  function setInt(string calldata jsonBlob, string[] memory paths, int256[] memory values) view external returns (string memory);

  function setUint(string calldata jsonBlob, string calldata path, uint256 value) view external returns (string memory);

  function setUint(string calldata jsonBlob, string[] memory paths, uint256[] memory values) view external returns (string memory);

  function setBool(string calldata jsonBlob, string calldata path, bool value) view external returns (string memory);

  function setBool(string calldata jsonBlob, string[] memory paths, bool[] memory values) view external returns (string memory);

  function subReplace(string calldata jsonBlob, string calldata searchPath, string calldata replacePath, string calldata value) view external returns (string memory);

  function subReplace(string calldata jsonBlob, string calldata searchPath, string[] memory replacePaths, string[] memory values) view external returns (string memory);

  function subReplaceInt(string calldata jsonBlob, string calldata searchPath, string calldata replacePath, int256 value) view external returns (string memory);

  function subReplaceInt(string calldata jsonBlob, string calldata searchPath, string[] memory replacePaths, int256[] memory values) view external returns (string memory);

  function subReplaceUint(string calldata jsonBlob, string calldata searchPath, string calldata replacePath, uint256 value) view external returns (string memory);

  function subReplaceUint(string calldata jsonBlob, string calldata searchPath, string[] memory replacePaths, uint256[] memory values) view external returns (string memory);

  function subReplaceBool(string calldata jsonBlob, string calldata searchPath, string calldata replacePath, bool value) view external returns (string memory);

  function subReplaceBool(string calldata jsonBlob, string calldata searchPath, string[] memory replacePaths, bool[] memory values) view external returns (string memory);

  function remove(string calldata jsonBlob, string calldata path) view external returns (string memory);

}
