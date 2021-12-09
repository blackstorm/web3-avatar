import { create } from 'ipfs-http-client'

const client = create()

const IPFS = {
  upload: async (file) => {
    return client.add(file);
  },
};

export default IPFS;
