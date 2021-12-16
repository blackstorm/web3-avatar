import { NavLink } from "react-router-dom";
import { Tooltip } from "antd";
import { useWeb3React } from "@web3-react/core";

const MenuButton = (props) => {
  const { disabled, ...rest } = props;

  const handleClick = (e) => {
    if (disabled) e.preventDefault();
  };

  return (
    <NavLink
      className={({ isActive }) =>
        `py-2 px-3 inline-block rounded-full font-medium text-base hover:text-black ${
          isActive ? "bg-gray-200 text-black" : "text-gray-400"
        }`
      }
      onClick={handleClick}
      {...rest}
    />
  );
};

const Menu = () => {
  const { active, account } = useWeb3React();
  return (
    <div className="bg-gray-100 p-1 rounded-full ">
      <div className="flex flex-row space-x-1 items-center">
        <MenuButton to="/">Upload</MenuButton>
        {active ? (
          <MenuButton to={`/view/${account}`}>My Avatar</MenuButton>
        ) : (
          <Tooltip title="Please connect your wallet first." placement="bottom">
            <div>
              <MenuButton to="/view" disabled>
                View
              </MenuButton>
            </div>
          </Tooltip>
        )}
      </div>
    </div>
  );
};

export default Menu;
