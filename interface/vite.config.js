import { defineConfig } from "vite";
import inject from '@rollup/plugin-inject';
import react from "@vitejs/plugin-react";
import vitePluginImp from "vite-plugin-imp";

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      process: "process/browser",
      stream: "stream-browserify",
      zlib: "browserify-zlib",
      util: "util",
    },
  },
  plugins: [
    react(),
    vitePluginImp({
      libList: [
        {
          libName: "antd",
          libDirectory: "es",
          style: (name) => `antd/es/${name}/style`,
        },
      ],
    }),
  ],
  css: {
    preprocessorOptions: {
      less: {
        javascriptEnabled: true,
      },
    },
  },
  build: {
    rollupOptions: {
      // external: ['web3'],
      plugins: [inject({ Buffer: ['buffer', 'Buffer'] })],
      commonjsOptions: {
        transformMixedEsModules: true,
      },
      output: {
        manualChunks: {
          lib: ["antd", "react"],
          web3: ["web3"],
        },
      },
    },
  },
});
