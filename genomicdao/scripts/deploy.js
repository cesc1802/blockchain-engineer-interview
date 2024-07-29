const hre = require("hardhat");
const { ethers } = require("hardhat")

async function main() {
  try {
    const [owner] = await ethers.getSigners();

    const GeneNFTContract = await ethers.getContractFactory("GeneNFT");
    const deployedGeneNFT = await GeneNFTContract.deploy()

    const PostCovidStrokePreventionContract = await ethers.getContractFactory("PostCovidStrokePrevention");
    const deployedPostCovidStrokePrevention = await PostCovidStrokePreventionContract.deploy()

    const ControllerContract = await ethers.getContractFactory("Controller");
    const deployedContract = await ControllerContract.deploy(deployedGeneNFT.getAddress(), deployedPostCovidStrokePrevention.getAddress());

    await deployedGeneNFT.transferOwnership(deployedContract.target)
    await deployedPostCovidStrokePrevention.transferOwnership(deployedContract.target)
    await deployedContract.waitForDeployment();
  } catch (error) {
    console.error(error);
    process.exit(1);
  }
}

main();
