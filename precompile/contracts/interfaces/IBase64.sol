// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IBase64 {

    function encode(bytes calldata data) view external returns (string memory);

    function encodeURL(bytes calldata data) view external returns (string memory);

    function decode(string calldata data) view external returns (bytes memory);

    function decodeURL(string calldata data) view external returns (bytes memory);

}
