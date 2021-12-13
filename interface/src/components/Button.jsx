import styled from "styled-components";

const Button = styled.button`
  margin: 0px;
  height: 38px;
  background-color: #fff3db;
  border: 1px solid rgb(255, 255, 255);
  padding: 0.15rem 0.5rem;
  border-radius: 12px;
  font-weight: 500;
  
  &:hover {
    border: solid 1px #c84648;
  }
`;

export const FlexButton = (props) => {
  const { children, className, ...rest } = props;
  return (
    <Button
      className={`flex flex-row items-center ${className ? className : null}`}
      {...rest}
    >
      {children}
    </Button>
  );
};

export const LoadingButton = (props) => {
  const { loading, ...rest } = props;
  return <FlexButton {...rest}></FlexButton>;
};

export default Button;
