import React from "react";
import AuthDialog from "../../components/AuthDialog";
import { Row } from "antd";

const Auth = () => {
  return (
    <Row justify={"center"}>
      <AuthDialog />
    </Row>
  );
};

export default Auth;
