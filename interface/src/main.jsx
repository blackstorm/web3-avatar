import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import "./index.css";
import { UseWalletProvider } from "use-wallet";

ReactDOM.render(
  <React.StrictMode>
    <UseWalletProvider
      chainId={5}
      connectors={{
        portis: { dAppId: "Web3A" },
      }}
    >
      <App />
    </UseWalletProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
