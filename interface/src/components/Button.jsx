import styled from "styled-components";
import { Spin } from "antd";
import { LoadingOutlined } from "@ant-design/icons";

const Button = styled.button`
  margin: 0px;
  height: 38px;
  background-color: #fff3db;
  border: 1px solid rgb(255, 255, 255);
  padding: 0.15rem 0.5rem;
  border-radius: 12px;
  font-weight: 500;

  &:hover {
    color: black;
    border: solid 1px #c84648;
  }
`;

const loadingIcon = <LoadingOutlined className="main-text-color" spin />;

export const FlexButton = (props) => {
  const { children, loading, className, ...rest } = props;
  return (
    <Button
      className={`flex flex-row items-center ${className ? className : null}`}
      {...rest}
    >
      {loading && <Spin className="mr-1" indicator={loadingIcon} />}
      {children}
    </Button>
  );
};

export const LoadingButton = (props) => {
  const { loading, ...rest } = props;
  return <FlexButton {...rest}></FlexButton>;
};

export default Button;
