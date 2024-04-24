// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ICompress {

    function compress(bytes calldata data) view external returns (bytes memory);

    function decompress(bytes calldata data) view external returns (bytes memory);

}
