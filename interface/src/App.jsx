import { useRef, useState, useEffect } from "react";
import Nav from "./components/Nav";
import styled from "styled-components";
import UploadSVG from "./components/UploadSVG";
import Menu from "./components/Menu";
import UploadButton from "./components/UploadButton";
import Avatar from "./components/Avatar";
import useIPFS from "./hooks/ipfs";
import { imageReader } from "./utils/reader";

const ImageContainer = styled.div`
  width: 250px;
  height: 250px;
`;

const App = () => {
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
      const added = await ipfs.add(file);
      console.log(added.data);
    }
  };

  return (
    <div className="flex flex-col">
      <Nav />
      <div className="flex flex-col p-4 mt-2">
        <Menu />
        <div className="w-full md:w-96 mx-auto my-4">
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
              Upload
            </UploadButton>
          </div>

          <div className="my-4 text-gray-400 text-center">
            <p>
              Your avatar will be upload to IPFS, And store the IPFS hash to you
              selected chain.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
