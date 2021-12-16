import axios from "axios";

const client = axios.create({
  baseURL: "https://ipfs.infura.io:5001/api/v0",
  timeout: 30000,
});

const useIPFS = () => {
  const add = async (file, onUploadProgress) => {
    const data = new FormData();
    data.append(file, file);
    return client.post("/add", data, {
      onUploadProgress: onUploadProgress,
      headers: {
        ContentType: "multipart/form-data",
        Accept: "application/json",
      },
    });
  };

  return {
    add,
  };
};

export default useIPFS;
