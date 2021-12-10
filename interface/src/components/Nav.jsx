import { useEffect, useState } from "react";
import Logo from "../asserts/svg/logo.svg";
import { Menu, Dropdown, Button, Tooltip } from "antd";
import { MoreOutlined, GlobalOutlined } from "@ant-design/icons";
import styled from "styled-components";
import Avatar from "./Avatar";
import { useWeb3React } from "@web3-react/core";
import { injected } from "../wallet/connectors";
import { shortAddress } from "../utils/account";

const netIcon = (name) => {
  return `https://www.gemini.com/images/currencies/icons/default/${name}.svg`;
};

const NETS = [
  {
    name: "Ethereum",
    address: "eth",
    icon: netIcon("eth"),
  },
  {
    name: "Binance Smart Chain",
    address: "bsc",
    icon: netIcon("bnb"),
  },
];

const NetSelector = () => {
  const menu = (
    <Menu>
      <div className="text-gray-400 py-2 px-3">Select NET</div>
      {NETS.map((net) => (
        <Menu.Item key={net.address}>
          <div className="flex flex-row items-center space-x-2">
            <img
              src={net.icon}
              alt={net.icon}
              className="w-4 h-4 rounded-full"
            />
            <span>{net.name}</span>
          </div>
        </Menu.Item>
      ))}
    </Menu>
  );

  return (
    <Dropdown overlay={menu} placement="bottomRight" arrow>
      <Button shape="round">Ethereum</Button>
    </Dropdown>
  );
};

const More = () => {
  const menu = (
    <Menu>
      <Menu.Item>
        <a href="https://github.com/blackstorm/web3-avatar" target="_blank">
          Github
        </a>
      </Menu.Item>
      <Menu.Item>
        <a href="https://github.com/blackstorm/web3-avatar" target="_blank">
          About
        </a>
      </Menu.Item>
    </Menu>
  );

  return (
    <Dropdown overlay={menu} placement="bottomRight" arrow>
      <Button shape="circle" type="text" icon={<MoreOutlined />}></Button>
    </Dropdown>
  );
};

const Lang = () => {
  const menu = (
    <Menu>
      <Menu.Item>English</Menu.Item>
      <Menu.Item>Chinese</Menu.Item>
    </Menu>
  );

  return (
    <Dropdown overlay={menu} placement="bottomRight" arrow>
      <Button type="text" shape="circle" icon={<GlobalOutlined />}></Button>
    </Dropdown>
  );
};

const CuteLogo = styled.img`
  &:hover {
    transform: rotate(-20deg);
    transition-duration: 100ms;
  }
`;

const ConnectWalletButton = () => {
  const [isConnecting, setIsConnecting] = useState(false);
  const wallet = useWeb3React();

  const connectWallet = async () => {
    try {
      setIsConnecting(true);
      await wallet.activate(injected);
    } finally {
      setIsConnecting(false);
    }
  };

  if (wallet.active) {
    return (
      <Tooltip title={`Click to disconnect`}>
        <Button
          shape="round"
          icon={<Avatar className="w-4 h-4 inline-block mr-2" />}
        >
          {shortAddress(wallet.account)}
        </Button>
      </Tooltip>
    );
  }

  if (isConnecting) {
    return <Button type="text" shape="circle" loading={true} disabled={true} />;
  }

  return (
    <Button shape="round" onClick={connectWallet}>
      Connect Wallet
    </Button>
  );
};

const Nav = () => {
  return (
    <nav className="flex items-center justify-between p-4">
      <div>
        <a href="/">
          <CuteLogo src={Logo} alt="logo" className="w-8" />
        </a>
      </div>

      <div className="flex flex-row space-x-1 md:space-x-2">
        <NetSelector />
        <ConnectWalletButton />
        <Lang />
        <More />
      </div>
    </nav>
  );
};

export default Nav;
