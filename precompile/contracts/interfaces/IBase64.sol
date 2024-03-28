// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IBase64 {

    function encode(bytes memory data) view external returns (string memory);
    function encodeURL(bytes memory data) view external returns (string memory);

}
