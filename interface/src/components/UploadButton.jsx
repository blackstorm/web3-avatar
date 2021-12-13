import styled from "styled-components";

const UploadButton = styled.button`
  border-radius: 50px;
  background-color: #fff3db;
  color: #c84648;
  user-select: none;

  &:hover {
    background-color: #c84648;
    color: #fff3db;
    transition-duration: 50ms;
    transition-property: background-color;
  }
}

  &:disabled {
    color: #9e9e9e;
    cursor: not-allowed;
    &:hover {
      background-color: #fff3db;
    }
  }
`;

export default UploadButton;
