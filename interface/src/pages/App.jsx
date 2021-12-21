import { useRef, useState, useEffect } from "react";
import UploadSVG from "../components/UploadSVG";
import UploadButton from "../components/UploadButton";
import Avatar from "../components/Avatar";
import useIPFS from "../hooks/ipfs";
import { imageReader } from "../utils/reader";
import styled from "styled-components";
import { useWeb3React } from "@web3-react/core";
import ABI from "../abi/Web3Avatar.abi.json";
import { ethers } from "ethers";

const ImageContainer = styled.div`
  width: 250px;
  height: 250px;
`;

const App = () => {
  const { active, account, library } = useWeb3React();
  const [file, setFile] = useState();
  const [isDisableUploadButton, setIsDisableUploadButton] = useState(true);
  const [isUploading, setIsUploading] = useState(false);
  const [avatarSrc, setAvatarSrc] = useState();

  const fileInputRef = useRef();
  const ipfs = useIPFS();

  useEffect(() => {
    setIsDisableUploadButton(!file);
    if (file) {
      imageReader(file, (res) => {
        setAvatarSrc(res);
      });
    }
  }, [file]);

  const onImageContainerClick = (e) => {
    e.preventDefault();
    fileInputRef.current.click();
  };

  const onFileInputChange = (e) => {
    const input = e.target;
    if (input.files && input.files.length > 0) {
      const file = input.files[0];
      setFile(file);
    }
  };

  const onUploadButtonClick = async (e) => {
    e.preventDefault();
    if (file) {
      try {
        setIsUploading(true);
        const added = await ipfs.add(file);

        const signer = library.getSigner();
        const contract = new ethers.Contract(
          "0xEBFFe5EEe4a3a195cC726B0Aa3D7988Ed480679a",
          ABI,
          signer
        );
        contract.setAvatar("default", added.data.Hash).then((res) => {
          setIsUploading(false);
          // TODO handle res
          // console.log(res);
        });
      } catch (err) {
        setIsUploading(false);
        console.log(err);
      }
    }
  };

  return (
    <>
      <input
        type="file"
        accept="image/*"
        ref={fileInputRef}
        className="hidden"
        onChange={onFileInputChange}
      />
      <ImageContainer className="mx-auto rounded-full relative group mt-10">
        <div
          onClick={onImageContainerClick}
          className="absolute flex bg-black bg-opacity-20 w-full h-full rounded-full z-50 invisible group-hover:visible cursor-pointer"
        >
          <div className="w-10 h-10 m-auto">
            <UploadSVG style={{ color: "#c84648" }} />
          </div>
        </div>
        <Avatar
          src={avatarSrc}
          className="absolute w-full h-full rounded-full z-0"
        />
      </ImageContainer>

      <div className="mt-10">
        <UploadButton
          className={`w-full py-4 text-xl ${
            isUploading && "uploading-animation"
          }`}
          onClick={onUploadButtonClick}
          disabled={isDisableUploadButton || isUploading}
        >
          {active ? "Upload" : "Please connect your wallet"}
        </UploadButton>
      </div>

      <div className="my-4 text-gray-400 text-center">
        <p>Your avatar will be upload to IPFS.</p>
      </div>
    </>
  );
};

export default App;
