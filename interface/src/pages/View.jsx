import Avatar from "../components/Avatar";

const GridAvatar = ({ name, hash }) => {
  return (
    <div className="flex flex-col text-center p-4 bg-neutral-100 rounded hover:shadow">
      <Avatar />
      <span className="my-2 main-text-color font-medium">{name}</span>
    </div>
  );
};

const View = () => {
  return (
    <div className="w-full">
      <div className="flex flex-row space-x-4 justify-center">
        <GridAvatar name="Default" />
        <GridAvatar name="Etherscan" />
      </div>
    </div>
  );
};

export default View;
