// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ETicketMetaData contains all meta data concerning the ETicket contract.
var ETicketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_info\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"rights_\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_validTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"rights_\",\"type\":\"string[]\"}],\"name\":\"addRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_right\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_right\",\"type\":\"uint16\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002730380380620027308339810160408190526200003491620004b8565b826000815181106200004a576200004a6200052b565b60200260200101516040518060400160405280600781526020016619551a58dad95d60ca1b8152508160009081620000839190620005d0565b506001620000928282620005d0565b505050620000af620000a9620001e460201b60201c565b620001f5565b82600181518110620000c557620000c56200052b565b6020026020010151600a9081620000dd9190620005d0565b5082600281518110620000f457620000f46200052b565b6020026020010151600b90816200010c9190620005d0565b50826003815181106200012357620001236200052b565b6020026020010151600c90816200013b9190620005d0565b50600754601011620001935760405162461bcd60e51b815260206004820152601a60248201527f4d617820726967687473206f766572204d41585f726967687473000000000000604482015260640160405180910390fd5b8151620001a89060079060208501906200027a565b5082600481518110620001bf57620001bf6200052b565b6020026020010151600d9081620001d79190620005d0565b50600e55506200069c9050565b6000620001f062000247565b905090565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600073c77fa7345679ccdbdb6204f0ca6a024be4abd00219330162000273575060131936013560601c90565b33620001f0565b828054828255906000526020600020908101928215620002c5579160200282015b82811115620002c55782518290620002b49082620005d0565b50916020019190600101906200029b565b50620002d3929150620002d7565b5090565b80821115620002d3576000620002ee8282620002f8565b50600101620002d7565b508054620003069062000541565b6000825580601f1062000317575050565b601f0160209004906000526020600020908101906200033791906200033a565b50565b5b80821115620002d357600081556001016200033b565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171562000392576200039262000351565b604052919050565b6000601f8381840112620003ad57600080fd5b825160206001600160401b0380831115620003cc57620003cc62000351565b8260051b620003dd83820162000367565b9384528681018301938381019089861115620003f857600080fd5b84890192505b85831015620004ab57825184811115620004185760008081fd5b8901603f81018b136200042b5760008081fd5b858101518581111562000442576200044262000351565b62000455818a01601f1916880162000367565b81815260408d818486010111156200046d5760008081fd5b60005b838110156200048d578481018201518382018b0152890162000470565b505060009181018801919091528352509184019190840190620003fe565b9998505050505050505050565b600080600060608486031215620004ce57600080fd5b83516001600160401b0380821115620004e657600080fd5b620004f4878388016200039a565b945060208601519150808211156200050b57600080fd5b506200051a868287016200039a565b925050604084015190509250925092565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806200055657607f821691505b6020821081036200057757634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620005cb57600081815260208120601f850160051c81016020861015620005a65750805b601f850160051c820191505b81811015620005c757828155600101620005b2565b5050505b505050565b81516001600160401b03811115620005ec57620005ec62000351565b6200060481620005fd845462000541565b846200057d565b602080601f8311600181146200063c5760008415620006235750858301515b600019600386901b1c1916600185901b178555620005c7565b600085815260208120601f198616915b828110156200066d578886015182559484019460019091019084016200064c565b50858210156200068c5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61208480620006ac6000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806370a08231116100b8578063aec0cf301161007c578063aec0cf3014610280578063b88d4fde14610293578063beabacc8146102a6578063c87b56dd146102b9578063e985e9c5146102cc578063f2fde38b1461030857600080fd5b806370a082311461022b578063715018a61461024c5780638da5cb5b1461025457806395d89b4114610265578063a22cb4651461026d57600080fd5b806340c10f19116100ff57806340c10f19146101cc578063415dd877146101df57806342842e0e146101f25780636352211e14610205578063687ffe071461021857600080fd5b806301ffc9a71461013c57806306fdde0314610164578063081812fc14610179578063095ea7b3146101a457806323b872dd146101b9575b600080fd5b61014f61014a36600461157a565b61031b565b60405190151581526020015b60405180910390f35b61016c61036d565b60405161015b91906115ee565b61018c610187366004611601565b6103ff565b6040516001600160a01b03909116815260200161015b565b6101b76101b2366004611636565b610426565b005b6101b76101c7366004611660565b610552565b6101b76101da366004611636565b61058a565b6101b76101ed366004611749565b6105a0565b6101b7610200366004611660565b610654565b61018c610213366004611601565b61066f565b6101b7610226366004611833565b6106cf565b61023e61023936600461186f565b6106e2565b60405190815260200161015b565b6101b7610768565b6006546001600160a01b031661018c565b61016c61077c565b6101b761027b36600461188a565b61078b565b6101b761028e3660046118c6565b61079d565b6101b76102a13660046118f2565b6107af565b6101b76102b4366004611660565b6107ee565b61016c6102c7366004611601565b6107f6565b61014f6102da36600461196e565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6101b761031636600461186f565b610afb565b60006001600160e01b031982166380ac58cd60e01b148061034c57506001600160e01b03198216635b5e139f60e01b145b8061036757506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606000805461037c90611998565b80601f01602080910402602001604051908101604052809291908181526020018280546103a890611998565b80156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b5050505050905090565b600061040a82610b74565b506000908152600460205260409020546001600160a01b031690565b60006104318261066f565b9050806001600160a01b0316836001600160a01b0316036104a35760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b806001600160a01b03166104b5610bd3565b6001600160a01b031614806104d157506104d1816102da610bd3565b6105435760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c0000606482015260840161049a565b61054d8383610be2565b505050565b61056361055d610bd3565b82610c50565b61057f5760405162461bcd60e51b815260040161049a906119d2565b61054d838383610ccf565b610592610e6b565b61059c8282610ee4565b5050565b6105a8610e6b565b80516007546010916105b991611a36565b106105f85760405162461bcd60e51b815260206004820152600f60248201526e6f766572204d61782072696768747360881b604482015260640161049a565b60005b815181101561059c57600782828151811061061857610618611a49565b602090810291909101810151825460018101845560009384529190922001906106419082611aad565b508061064c81611b6d565b9150506105fb565b61054d838383604051806020016040528060008152506107af565b6000818152600260205260408120546001600160a01b0316806103675760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161049a565b6106d7610e6b565b61054d838383611026565b60006001600160a01b03821661074c5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161049a565b506001600160a01b031660009081526003602052604090205490565b610770610e6b565b61077a6000611186565b565b60606001805461037c90611998565b61059c610796610bd3565b83836111d8565b61059c82826107aa610bd3565b611026565b6107c06107ba610bd3565b83610c50565b6107dc5760405162461bcd60e51b815260040161049a906119d2565b6107e8848484846112a6565b50505050565b61057f610e6b565b6000818152600260205260409020546060906001600160a01b03166108515760405162461bcd60e51b81526020600482015260116024820152703737b732bc34b9ba32b73a103a37b5b2b760791b604482015260640161049a565b600061085b61036d565b600a600c600b6040516020016108749493929190611bf9565b60408051601f1981840301815290829052600e54634f689bcb60e11b8352600483015291508190600d9073__$73785166eb94af412e1735ba7ee7be91af$__90639ed1379690602401600060405180830381865af41580156108da573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109029190810190611cd8565b60405160200161091493929190611d4f565b604051602081830303815290604052905060005b60075461ffff82161015610a6d578160078261ffff168154811061094e5761094e611a49565b6000918252602080832088845260088252604080852061ffff8816865290925292205491019061099f576040518060400160405280600981526020016801cd3955cd4171d32960bf1b815250610a37565b600086815260086020908152604080832061ffff8716845290915290819020549051634f689bcb60e11b8152600481019190915273__$73785166eb94af412e1735ba7ee7be91af$__90639ed1379690602401600060405180830381865af4158015610a0f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a379190810190611cd8565b604051602001610a4993929190611e0d565b60405160208183030381529060405291508080610a6590611e8b565b915050610928565b50605d60f81b8160018351610a829190611eac565b81518110610a9257610a92611a49565b60200101906001600160f81b031916908160001a90535080604051602001610aba9190611ebf565b6040516020818303038152906040529050610ad4816112d9565b604051602001610ae49190611ee4565b604051602081830303815290604052915050919050565b610b03610e6b565b6001600160a01b038116610b685760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161049a565b610b7181611186565b50565b6000818152600260205260409020546001600160a01b0316610b715760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161049a565b6000610bdd61142c565b905090565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610c178261066f565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610c5c8361066f565b9050806001600160a01b0316846001600160a01b03161480610ca357506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b80610cc75750836001600160a01b0316610cbc846103ff565b6001600160a01b0316145b949350505050565b826001600160a01b0316610ce28261066f565b6001600160a01b031614610d465760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b606482015260840161049a565b6001600160a01b038216610da85760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161049a565b610db3600082610be2565b6001600160a01b0383166000908152600360205260408120805460019290610ddc908490611eac565b90915550506001600160a01b0382166000908152600360205260408120805460019290610e0a908490611a36565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b610e73610bd3565b6001600160a01b0316610e8e6006546001600160a01b031690565b6001600160a01b03161461077a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161049a565b6001600160a01b038216610f3a5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161049a565b6000818152600260205260409020546001600160a01b031615610f9f5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161049a565b6001600160a01b0382166000908152600360205260408120805460019290610fc8908490611a36565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b806001600160a01b03166110398461066f565b6001600160a01b03161461107c5760405162461bcd60e51b815260206004820152600a60248201526927b7363c9037bbb732b960b11b604482015260640161049a565b60075461ffff8316106110c35760405162461bcd60e51b815260206004820152600f60248201526e726967687473206f766572666c6f7760881b604482015260640161049a565b600e5442106111055760405162461bcd60e51b815260206004820152600e60248201526d1d1a58dad95d08195e1c1a5c995960921b604482015260640161049a565b600083815260086020908152604080832061ffff86168452909152902054156111625760405162461bcd60e51b815260206004820152600f60248201526e436865636b656420416c726561647960881b604482015260640161049a565b50600091825260086020908152604080842061ffff90931684529190529020429055565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b0316036112395760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161049a565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6112b1848484610ccf565b6112bd8484848461145c565b6107e85760405162461bcd60e51b815260040161049a90611f29565b606081516000036112f857505060408051602081019091526000815290565b600060405180606001604052806040815260200161200f60409139905060006003845160026113279190611a36565b6113319190611f7b565b61133c906004611f9d565b67ffffffffffffffff8111156113545761135461169c565b6040519080825280601f01601f19166020018201604052801561137e576020820181803683370190505b509050600182016020820185865187015b808210156113ea576003820191508151603f8160121c168501518453600184019350603f81600c1c168501518453600184019350603f8160061c168501518453600184019350603f811685015184535060018301925061138f565b5050600386510660018114611406576002811461141957611421565b603d6001830353603d6002830353611421565b603d60018303535b509195945050505050565b600073c77fa7345679ccdbdb6204f0ca6a024be4abd002193301611457575060131936013560601c90565b503390565b60006001600160a01b0384163b1561155957836001600160a01b031663150b7a02611485610bd3565b8786866040518563ffffffff1660e01b81526004016114a79493929190611fb4565b6020604051808303816000875af19250505080156114e2575060408051601f3d908101601f191682019092526114df91810190611ff1565b60015b61153f573d808015611510576040519150601f19603f3d011682016040523d82523d6000602084013e611515565b606091505b5080516000036115375760405162461bcd60e51b815260040161049a90611f29565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610cc7565b506001949350505050565b6001600160e01b031981168114610b7157600080fd5b60006020828403121561158c57600080fd5b813561159781611564565b9392505050565b60005b838110156115b95781810151838201526020016115a1565b50506000910152565b600081518084526115da81602086016020860161159e565b601f01601f19169290920160200192915050565b60208152600061159760208301846115c2565b60006020828403121561161357600080fd5b5035919050565b80356001600160a01b038116811461163157600080fd5b919050565b6000806040838503121561164957600080fd5b6116528361161a565b946020939093013593505050565b60008060006060848603121561167557600080fd5b61167e8461161a565b925061168c6020850161161a565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156116db576116db61169c565b604052919050565b600067ffffffffffffffff8211156116fd576116fd61169c565b50601f01601f191660200190565b600061171e611719846116e3565b6116b2565b905082815283838301111561173257600080fd5b828260208301376000602084830101529392505050565b6000602080838503121561175c57600080fd5b823567ffffffffffffffff8082111561177457600080fd5b818501915085601f83011261178857600080fd5b81358181111561179a5761179a61169c565b8060051b6117a98582016116b2565b91825283810185019185810190898411156117c357600080fd5b86860192505b83831015611814578235858111156117e15760008081fd5b8601603f81018b136117f35760008081fd5b6118048b898301356040840161170b565b83525091860191908601906117c9565b9998505050505050505050565b803561ffff8116811461163157600080fd5b60008060006060848603121561184857600080fd5b8335925061185860208501611821565b91506118666040850161161a565b90509250925092565b60006020828403121561188157600080fd5b6115978261161a565b6000806040838503121561189d57600080fd5b6118a68361161a565b9150602083013580151581146118bb57600080fd5b809150509250929050565b600080604083850312156118d957600080fd5b823591506118e960208401611821565b90509250929050565b6000806000806080858703121561190857600080fd5b6119118561161a565b935061191f6020860161161a565b925060408501359150606085013567ffffffffffffffff81111561194257600080fd5b8501601f8101871361195357600080fd5b6119628782356020840161170b565b91505092959194509250565b6000806040838503121561198157600080fd5b61198a8361161a565b91506118e96020840161161a565b600181811c908216806119ac57607f821691505b6020821081036119cc57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561036757610367611a20565b634e487b7160e01b600052603260045260246000fd5b601f82111561054d57600081815260208120601f850160051c81016020861015611a865750805b601f850160051c820191505b81811015611aa557828155600101611a92565b505050505050565b815167ffffffffffffffff811115611ac757611ac761169c565b611adb81611ad58454611998565b84611a5f565b602080601f831160018114611b105760008415611af85750858301515b600019600386901b1c1916600185901b178555611aa5565b600085815260208120601f198616915b82811015611b3f57888601518255948401946001909101908401611b20565b5085821015611b5d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060018201611b7f57611b7f611a20565b5060010190565b60008154611b9381611998565b60018281168015611bab5760018114611bc057611bef565b60ff1984168752821515830287019450611bef565b8560005260208060002060005b85811015611be65781548a820152908401908201611bcd565b50505082870194505b5050505092915050565b693d913730b6b2911d101160b11b81528451600090611c1f81600a850160208a0161159e565b72111610113232b9b1b934b83a34b7b7111d101160691b600a91840191820152611c4c601d820187611b86565b6e111610113232ba30b4b639911d101160891b81529050611c70600f820186611b86565b6c1116101134b6b0b3b2911d101160991b81529050611c92600d820185611b86565b7f222c202264657369676e6572223a202269736f746f702e746f70222c2022617481526b7472696275746573223a205b60a01b6020820152602c01979650505050505050565b600060208284031215611cea57600080fd5b815167ffffffffffffffff811115611d0157600080fd5b8201601f81018413611d1257600080fd5b8051611d20611719826116e3565b818152856020838501011115611d3557600080fd5b611d4682602083016020860161159e565b95945050505050565b60008451611d6181846020890161159e565b7f7b2274726169745f74797065223a2022e58f91e8a18ce696b9444944222c22769083019081526730b63ab2911d101160c11b6020820152611da66028820186611b86565b90507f227d2c207b2274726169745f74797065223a2022e69c89e69588e69c9f222c228152683b30b63ab2911d101160b91b60208201528351611df081602984016020880161159e565b62089f4b60ea1b60299290910191820152602c0195945050505050565b60008451611e1f81846020890161159e565b6f3d913a3930b4ba2fba3cb832911d101160811b908301908152611e466010820186611b86565b6b1116113b30b63ab2911d101160a11b81528451909150611e6e81600c84016020880161159e565b62089f4b60ea1b600c9290910191820152600f0195945050505050565b600061ffff808316818103611ea257611ea2611a20565b6001019392505050565b8181038181111561036757610367611a20565b60008251611ed181846020870161159e565b607d60f81b920191825250600101919050565b7f646174613a6170706c69636174696f6e2f6a736f6e3b6261736536342c000000815260008251611f1c81601d85016020870161159e565b91909101601d0192915050565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b600082611f9857634e487b7160e01b600052601260045260246000fd5b500490565b808202811582820484141761036757610367611a20565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611fe7908301846115c2565b9695505050505050565b60006020828403121561200357600080fd5b81516115978161156456fe4142434445464748494a4b4c4d4e4f505152535455565758595a6162636465666768696a6b6c6d6e6f707172737475767778797a303132333435363738392b2fa264697066735822122043196623d87969f8bfdaa7607766b014f7d29932aab83d466da98b948e85dbf164736f6c63430008130033",
}

// ETicketABI is the input ABI used to generate the binding from.
// Deprecated: Use ETicketMetaData.ABI instead.
var ETicketABI = ETicketMetaData.ABI

// ETicketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ETicketMetaData.Bin instead.
var ETicketBin = ETicketMetaData.Bin

// DeployETicket deploys a new Ethereum contract, binding an instance of ETicket to it.
func DeployETicket(auth *bind.TransactOpts, backend bind.ContractBackend, _info []string, rights_ []string, _validTime *big.Int) (common.Address, *types.Transaction, *ETicket, error) {
	parsed, err := ETicketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ETicketBin), backend, _info, rights_, _validTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ETicket{ETicketCaller: ETicketCaller{contract: contract}, ETicketTransactor: ETicketTransactor{contract: contract}, ETicketFilterer: ETicketFilterer{contract: contract}}, nil
}

// ETicket is an auto generated Go binding around an Ethereum contract.
type ETicket struct {
	ETicketCaller     // Read-only binding to the contract
	ETicketTransactor // Write-only binding to the contract
	ETicketFilterer   // Log filterer for contract events
}

// ETicketCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETicketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETicketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETicketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETicketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETicketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETicketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETicketSession struct {
	Contract     *ETicket          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETicketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETicketCallerSession struct {
	Contract *ETicketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ETicketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETicketTransactorSession struct {
	Contract     *ETicketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ETicketRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETicketRaw struct {
	Contract *ETicket // Generic contract binding to access the raw methods on
}

// ETicketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETicketCallerRaw struct {
	Contract *ETicketCaller // Generic read-only contract binding to access the raw methods on
}

// ETicketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETicketTransactorRaw struct {
	Contract *ETicketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETicket creates a new instance of ETicket, bound to a specific deployed contract.
func NewETicket(address common.Address, backend bind.ContractBackend) (*ETicket, error) {
	contract, err := bindETicket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETicket{ETicketCaller: ETicketCaller{contract: contract}, ETicketTransactor: ETicketTransactor{contract: contract}, ETicketFilterer: ETicketFilterer{contract: contract}}, nil
}

// NewETicketCaller creates a new read-only instance of ETicket, bound to a specific deployed contract.
func NewETicketCaller(address common.Address, caller bind.ContractCaller) (*ETicketCaller, error) {
	contract, err := bindETicket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETicketCaller{contract: contract}, nil
}

// NewETicketTransactor creates a new write-only instance of ETicket, bound to a specific deployed contract.
func NewETicketTransactor(address common.Address, transactor bind.ContractTransactor) (*ETicketTransactor, error) {
	contract, err := bindETicket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETicketTransactor{contract: contract}, nil
}

// NewETicketFilterer creates a new log filterer instance of ETicket, bound to a specific deployed contract.
func NewETicketFilterer(address common.Address, filterer bind.ContractFilterer) (*ETicketFilterer, error) {
	contract, err := bindETicket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETicketFilterer{contract: contract}, nil
}

// bindETicket binds a generic wrapper to an already deployed contract.
func bindETicket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ETicketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETicket *ETicketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETicket.Contract.ETicketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETicket *ETicketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETicket.Contract.ETicketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETicket *ETicketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETicket.Contract.ETicketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETicket *ETicketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETicket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETicket *ETicketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETicket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETicket *ETicketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETicket.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ETicket *ETicketCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ETicket *ETicketSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ETicket.Contract.BalanceOf(&_ETicket.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ETicket *ETicketCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ETicket.Contract.BalanceOf(&_ETicket.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ETicket *ETicketCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ETicket *ETicketSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ETicket.Contract.GetApproved(&_ETicket.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ETicket *ETicketCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ETicket.Contract.GetApproved(&_ETicket.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ETicket *ETicketCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ETicket *ETicketSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ETicket.Contract.IsApprovedForAll(&_ETicket.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ETicket *ETicketCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ETicket.Contract.IsApprovedForAll(&_ETicket.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ETicket *ETicketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ETicket *ETicketSession) Name() (string, error) {
	return _ETicket.Contract.Name(&_ETicket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ETicket *ETicketCallerSession) Name() (string, error) {
	return _ETicket.Contract.Name(&_ETicket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETicket *ETicketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETicket *ETicketSession) Owner() (common.Address, error) {
	return _ETicket.Contract.Owner(&_ETicket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETicket *ETicketCallerSession) Owner() (common.Address, error) {
	return _ETicket.Contract.Owner(&_ETicket.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ETicket *ETicketCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ETicket *ETicketSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ETicket.Contract.OwnerOf(&_ETicket.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ETicket *ETicketCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ETicket.Contract.OwnerOf(&_ETicket.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ETicket *ETicketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ETicket *ETicketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ETicket.Contract.SupportsInterface(&_ETicket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ETicket *ETicketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ETicket.Contract.SupportsInterface(&_ETicket.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ETicket *ETicketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ETicket *ETicketSession) Symbol() (string, error) {
	return _ETicket.Contract.Symbol(&_ETicket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ETicket *ETicketCallerSession) Symbol() (string, error) {
	return _ETicket.Contract.Symbol(&_ETicket.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ETicket *ETicketCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ETicket.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ETicket *ETicketSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ETicket.Contract.TokenURI(&_ETicket.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ETicket *ETicketCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ETicket.Contract.TokenURI(&_ETicket.CallOpts, tokenId)
}

// AddRights is a paid mutator transaction binding the contract method 0x415dd877.
//
// Solidity: function addRights(string[] rights_) returns()
func (_ETicket *ETicketTransactor) AddRights(opts *bind.TransactOpts, rights_ []string) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "addRights", rights_)
}

// AddRights is a paid mutator transaction binding the contract method 0x415dd877.
//
// Solidity: function addRights(string[] rights_) returns()
func (_ETicket *ETicketSession) AddRights(rights_ []string) (*types.Transaction, error) {
	return _ETicket.Contract.AddRights(&_ETicket.TransactOpts, rights_)
}

// AddRights is a paid mutator transaction binding the contract method 0x415dd877.
//
// Solidity: function addRights(string[] rights_) returns()
func (_ETicket *ETicketTransactorSession) AddRights(rights_ []string) (*types.Transaction, error) {
	return _ETicket.Contract.AddRights(&_ETicket.TransactOpts, rights_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ETicket *ETicketSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.Approve(&_ETicket.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.Approve(&_ETicket.TransactOpts, to, tokenId)
}

// Check is a paid mutator transaction binding the contract method 0x687ffe07.
//
// Solidity: function check(uint256 tokenId, uint16 _right, address _who) returns()
func (_ETicket *ETicketTransactor) Check(opts *bind.TransactOpts, tokenId *big.Int, _right uint16, _who common.Address) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "check", tokenId, _right, _who)
}

// Check is a paid mutator transaction binding the contract method 0x687ffe07.
//
// Solidity: function check(uint256 tokenId, uint16 _right, address _who) returns()
func (_ETicket *ETicketSession) Check(tokenId *big.Int, _right uint16, _who common.Address) (*types.Transaction, error) {
	return _ETicket.Contract.Check(&_ETicket.TransactOpts, tokenId, _right, _who)
}

// Check is a paid mutator transaction binding the contract method 0x687ffe07.
//
// Solidity: function check(uint256 tokenId, uint16 _right, address _who) returns()
func (_ETicket *ETicketTransactorSession) Check(tokenId *big.Int, _right uint16, _who common.Address) (*types.Transaction, error) {
	return _ETicket.Contract.Check(&_ETicket.TransactOpts, tokenId, _right, _who)
}

// Check0 is a paid mutator transaction binding the contract method 0xaec0cf30.
//
// Solidity: function check(uint256 tokenId, uint16 _right) returns()
func (_ETicket *ETicketTransactor) Check0(opts *bind.TransactOpts, tokenId *big.Int, _right uint16) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "check0", tokenId, _right)
}

// Check0 is a paid mutator transaction binding the contract method 0xaec0cf30.
//
// Solidity: function check(uint256 tokenId, uint16 _right) returns()
func (_ETicket *ETicketSession) Check0(tokenId *big.Int, _right uint16) (*types.Transaction, error) {
	return _ETicket.Contract.Check0(&_ETicket.TransactOpts, tokenId, _right)
}

// Check0 is a paid mutator transaction binding the contract method 0xaec0cf30.
//
// Solidity: function check(uint256 tokenId, uint16 _right) returns()
func (_ETicket *ETicketTransactorSession) Check0(tokenId *big.Int, _right uint16) (*types.Transaction, error) {
	return _ETicket.Contract.Check0(&_ETicket.TransactOpts, tokenId, _right)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactor) Mint(opts *bind.TransactOpts, _to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "mint", _to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 tokenId) returns()
func (_ETicket *ETicketSession) Mint(_to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.Mint(&_ETicket.TransactOpts, _to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactorSession) Mint(_to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.Mint(&_ETicket.TransactOpts, _to, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETicket *ETicketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETicket *ETicketSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETicket.Contract.RenounceOwnership(&_ETicket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETicket *ETicketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETicket.Contract.RenounceOwnership(&_ETicket.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.SafeTransferFrom(&_ETicket.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.SafeTransferFrom(&_ETicket.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ETicket *ETicketTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ETicket *ETicketSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ETicket.Contract.SafeTransferFrom0(&_ETicket.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ETicket *ETicketTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ETicket.Contract.SafeTransferFrom0(&_ETicket.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ETicket *ETicketTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ETicket *ETicketSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ETicket.Contract.SetApprovalForAll(&_ETicket.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ETicket *ETicketTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ETicket.Contract.SetApprovalForAll(&_ETicket.TransactOpts, operator, approved)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactor) Transfer(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "transfer", from, to, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketSession) Transfer(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.Transfer(&_ETicket.TransactOpts, from, to, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactorSession) Transfer(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.Transfer(&_ETicket.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.TransferFrom(&_ETicket.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ETicket *ETicketTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ETicket.Contract.TransferFrom(&_ETicket.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETicket *ETicketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ETicket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETicket *ETicketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETicket.Contract.TransferOwnership(&_ETicket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETicket *ETicketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETicket.Contract.TransferOwnership(&_ETicket.TransactOpts, newOwner)
}

// ETicketApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ETicket contract.
type ETicketApprovalIterator struct {
	Event *ETicketApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ETicketApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETicketApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ETicketApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ETicketApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETicketApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETicketApproval represents a Approval event raised by the ETicket contract.
type ETicketApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ETicket *ETicketFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ETicketApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ETicket.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ETicketApprovalIterator{contract: _ETicket.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ETicket *ETicketFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ETicketApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ETicket.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETicketApproval)
				if err := _ETicket.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ETicket *ETicketFilterer) ParseApproval(log types.Log) (*ETicketApproval, error) {
	event := new(ETicketApproval)
	if err := _ETicket.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETicketApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ETicket contract.
type ETicketApprovalForAllIterator struct {
	Event *ETicketApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ETicketApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETicketApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ETicketApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ETicketApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETicketApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETicketApprovalForAll represents a ApprovalForAll event raised by the ETicket contract.
type ETicketApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ETicket *ETicketFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ETicketApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ETicket.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ETicketApprovalForAllIterator{contract: _ETicket.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ETicket *ETicketFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ETicketApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ETicket.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETicketApprovalForAll)
				if err := _ETicket.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ETicket *ETicketFilterer) ParseApprovalForAll(log types.Log) (*ETicketApprovalForAll, error) {
	event := new(ETicketApprovalForAll)
	if err := _ETicket.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETicketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ETicket contract.
type ETicketOwnershipTransferredIterator struct {
	Event *ETicketOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ETicketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETicketOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ETicketOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ETicketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETicketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETicketOwnershipTransferred represents a OwnershipTransferred event raised by the ETicket contract.
type ETicketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETicket *ETicketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ETicketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETicket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ETicketOwnershipTransferredIterator{contract: _ETicket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETicket *ETicketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ETicketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETicket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETicketOwnershipTransferred)
				if err := _ETicket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETicket *ETicketFilterer) ParseOwnershipTransferred(log types.Log) (*ETicketOwnershipTransferred, error) {
	event := new(ETicketOwnershipTransferred)
	if err := _ETicket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETicketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ETicket contract.
type ETicketTransferIterator struct {
	Event *ETicketTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ETicketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETicketTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ETicketTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ETicketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETicketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETicketTransfer represents a Transfer event raised by the ETicket contract.
type ETicketTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ETicket *ETicketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ETicketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ETicket.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ETicketTransferIterator{contract: _ETicket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ETicket *ETicketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ETicketTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ETicket.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETicketTransfer)
				if err := _ETicket.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ETicket *ETicketFilterer) ParseTransfer(log types.Log) (*ETicketTransfer, error) {
	event := new(ETicketTransfer)
	if err := _ETicket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
