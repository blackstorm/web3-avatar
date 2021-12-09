import { useRef, useState } from "react";
import Nav from "./components/Nav";
import styled from "styled-components";
import UploadSVG from "./components/UploadSVG";
import Menu from "./components/Menu";
import UploadButton from "./components/UploadButton";
import Avatar from "./components/Avatar";

const ImageContainer = styled.div`
  width: 250px;
  height: 250px;
`;

const PreviewAvatar = (props) => {
  const { src, className, ...rest } = props;

  return (
    <div className={`w-10 h-10" ${className}`} {...rest}>
      {src ? (
        <img src={src} alt="" />
      ) : (
        <UploadSVG style={{ color: "#c84648" }} />
      )}
    </div>
  );
};

const App = () => {
  const [file, setFile] = useState();
  const fileInputRef = useRef();

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
      const added = await IPFS.upload(file)
      console.log(added)
    }
  };

  return (
    <div className="flex flex-col">
      <input
        type="file"
        accept="image/*"
        ref={fileInputRef}
        className="hidden"
        onChange={onFileInputChange}
      />
      <Nav />
      <div className="flex flex-col p-4 mt-2">
        <Menu />
        <div className="w-full md:w-96 mx-auto my-4">
          <ImageContainer className="mx-auto rounded-full relative group mt-10">
            <div
              onClick={onImageContainerClick}
              className="absolute flex bg-black bg-opacity-20 w-full h-full rounded-full z-50 invisible group-hover:visible cursor-pointer"
            >
              <PreviewAvatar className="mx-auto my-auto" />
            </div>
            <Avatar className="absolute w-full h-full rounded-full z-0" />
          </ImageContainer>

          <div className="mt-10">
            <UploadButton
              className="w-full py-4 text-xl"
              onClick={onUploadButtonClick}
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
