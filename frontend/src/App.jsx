import React from "react";

import { Row } from "antd";
import "./App.css";
import AuthDialog from "./components/AuthDialog";

function App() {
  return (
    <div className="App">
      <Row justify={"center"}>
        <AuthDialog />
      </Row>
    </div>
  );
}

export default App;
