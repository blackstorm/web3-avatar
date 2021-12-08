const { expect } = require("chai");
const { ethers } = require("hardhat");

const W3AV1 = "Web3AvatarV0"

describe(W3AV1, function () {

  it("getHash", async function() {
    const [owner, addr1] = await ethers.getSigners();
    const W3A = await ethers.getContractFactory(W3AV1);
    const w3a = await W3A.deploy();
    await w3a.deployed();

    expect(await w3a['getHash()']()).to.be.equal('');

    const hash = "0x000000000000000";
    const tx = await w3a.setHash(hash);
    await tx.wait();

    expect(await w3a['getHash(address)'](owner.address)).to.be.equal(hash);
    expect(await w3a['getHash(address)'](addr1.address)).to.be.equal('');
  });
  
  it("setHash", async function () {
    const W3A = await ethers.getContractFactory(W3AV1);
    const w3a = await W3A.deploy();
    await w3a.deployed();

    const hash = "0x000000000000000";

    const tx = await w3a.setHash(hash);
    await tx.wait();
  });
});
