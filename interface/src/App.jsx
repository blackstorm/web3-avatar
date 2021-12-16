import { Routes, Route } from "react-router-dom";
import Nav from "./components/Nav";
import Menu from "./components/Menu";
import Home from "./pages/App";
import View from "./pages/View";
import NotFound from "./pages/NotFound";

const App = () => {
  return (
    <div className="flex flex-col">
      <Nav />
      <div className="flex flex-col p-4 mt-2">
        <Menu />
        <div className="w-full md:w-96 mx-auto my-4">
          <Routes>
            <Route path="*" element={<NotFound />} />
            <Route path="/" element={<Home />} />
            <Route path="/view/:address" element={<View />} />
          </Routes>
        </div>
      </div>
    </div>
  );
};

export default App;
