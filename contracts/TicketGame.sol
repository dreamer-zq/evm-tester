// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract TicketGame is ERC721URIStorage {
    uint256 private _tokenIdCounter;

    constructor() ERC721("TicketGame", "ticket") {}

    function redeem(address player, string memory tokenURI)
        public
        returns (uint256)
    {
        uint256 tokenId = _tokenIdCounter;
        _mint(player, tokenId);
        _setTokenURI(tokenId, tokenURI);

        _tokenIdCounter += 1;
        return tokenId;
    }

    function batchRedeem(address[] memory player, string[] memory tokenURI)
        public
        returns (uint256[] memory tokenIds)
    {
       require(player.length >0,"player.length must be greater than 0");
       require(player.length <1000,"player.length must be less than 1000");
       require(player.length == tokenURI.length,"player.length not equal tokenURI.length");
       for (uint256 i = 0; i < player.length; i++) {
           tokenIds[i] = redeem(player[i], tokenURI[i]);
       }
       return tokenIds;
    }
}