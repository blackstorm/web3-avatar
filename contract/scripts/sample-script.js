const hre = require("hardhat");

async function main() {
  const W3A = await hre.ethers.getContractFactory("Web3AvatarV0")
  const w3a = await W3A.deploy();

  await w3a.deployed();

  console.log("w3a deployed to:", w3a.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
