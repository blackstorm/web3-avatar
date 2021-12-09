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

export default Menu;