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

    function batchRedeem(address[] memory players, string[] memory tokenURIs)
        public
        returns (uint256[] memory tokenIds)
    {
        require(players.length > 0, "players.length must be greater than 0");
        require(players.length < 1000, "players.length must be less than 1000");
        require(players.length == tokenURIs.length, "players.length not equal tokenURIs.length");

        // Initialize the tokenIds array with the correct length
        tokenIds = new uint256[](players.length);

        for (uint256 i = 0; i < players.length; i++) {
            tokenIds[i] = redeem(players[i], tokenURIs[i]);
        }
        return tokenIds;
    }
}