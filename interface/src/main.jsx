import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import "./index.css";
import { Web3ReactProvider } from "@web3-react/core";
import Web3 from "web3";
import { BrowserRouter } from "react-router-dom";

function getLibrary(provider) {
  return new Web3(provider);
}

ReactDOM.render(
    <Web3ReactProvider getLibrary={getLibrary}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Web3ReactProvider>,
  document.getElementById("root")
);
