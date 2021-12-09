import { useState } from "react";
import { createModel } from "hox";
import { useWeb3React } from "@web3-react/core";
import { injected } from "../wallet/connectors";

const useWallet = () => {
  const [status, setStatus] = useState("disconnected");
  const { active, account, library, connector, activate, deactivate } =
    useWeb3React();

  const connected = async () => {
    return activate(injected);
  };

  const disconnect = async () => {
    return deactivate();
  };

  return {
    status,
    account,
    connected,
    disconnect,
  };
};

export default createModel(useWallet);
