import styled from "styled-components";
import Logo from "../asserts/svg/logo.svg";

const Root = styled.img`
  &:hover {
    transform: rotate(-20deg);
    transition-duration: 100ms;
  }
`;

const LogoComponent = (props) => {
  return <Root src={Logo} {...props} />;
};

export default LogoComponent;
