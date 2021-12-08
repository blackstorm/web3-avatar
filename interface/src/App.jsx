import Nav from "./components/Nav";
import styled from "styled-components";
import Logo from "./asserts/svg/logo.svg";
import UploadSVG from "./components/UploadSVG";

const UploadButton = styled.button`
  border-radius: 50px;
  background-color: #fff3db;
  color: #c84648;

  &:hover {
    background-color: #c84648;
    color: #fff3db;
    transition-duration: 50ms;
    transition-property: background-color;
  }
`;

const MenuButton = (props) => {
  const { active, ...rest } = props;
  const activeClass = active ? "bg-gray-200 text-black" : "text-gray-400";

  return (
    <button
      className={`py-2 px-3 rounded-full font-medium text-base hover:text-black ${activeClass}`}
      {...rest}
    />
  );
};

const Menu = () => {
  return (
    <div className="bg-gray-100 p-1 rounded-full ">
      <div className="flex flex-row space-x-1 items-center">
        <MenuButton active>Upload</MenuButton>
        <MenuButton>View</MenuButton>
      </div>
    </div>
  );
};

const ImageContainer = styled.div`
  width: 250px;
  height: 250px;
`;

const App = () => {
  return (
    <div className="flex flex-col">
      <Nav />
      <div className="flex flex-col p-4 mt-2">
        <Menu />
        <div className="w-full md:w-96 mx-auto my-4">
          <ImageContainer className="mx-auto rounded-full relative group mt-10">
            <div className="absolute flex bg-black bg-opacity-20 w-full h-full rounded-full z-50 invisible group-hover:visible cursor-pointer">
              <UploadSVG
                className="w-10 h-10 mx-auto my-auto"
                style={{ color: "#c84648" }}
              />
            </div>
            <img
              src={Logo}
              alt="logo"
              className="absolute w-full h-full rounded-full z-0"
            />
          </ImageContainer>

          <div className="mt-10">
            <UploadButton className="w-full py-4 text-xl">Upload</UploadButton>
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
