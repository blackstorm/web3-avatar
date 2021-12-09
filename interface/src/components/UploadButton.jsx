import styled from "styled-components";

const UploadButton = styled.button`
  border-radius: 50px;
  background-color: #fff3db;
  color: #c84648;

  &:hover {
    background-color: #c84648;
    color: #fff3db;
    transition-duration: 50ms;
    transition-property: background-color;
  }
`;

export default UploadButton;