require("@nomicfoundation/hardhat-toolbox");
require('hardhat-abi-exporter'); // 引入 abiExporter 插件

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  abiExporter: {
    path: './abi',
    runOnCompile: true,
    clear: true,
    spacing: 2
  }
};
