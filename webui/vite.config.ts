// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
// Plugins
import vue from '@vitejs/plugin-vue'
import vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'
import ViteFonts from 'unplugin-fonts/vite'
import yaml from 'js-yaml';
import fs from 'fs';

// Utilities
import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'
import path from 'path';


const configYaml = fs.readFileSync('./config.yaml', 'utf8');
const parsedConfig = yaml.load(configYaml);

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      template: { transformAssetUrls }
    }),
    // https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vite-plugin
    vuetify({
      autoImport: true,
      styles: {
        configFile: 'src/styles/settings.scss',
      },
    }),
    ViteFonts({
      google: {
        families: [{
          name: 'Roboto',
          styles: 'wght@100;300;400;500;700;900',
        }],
      },
    }),
  ],
  define: {
    'process.env': {
      BASE_PATH: parsedConfig.api.base_path,
    }
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: [
      '.js',
      '.json',
      '.jsx',
      '.mjs',
      '.ts',
      '.tsx',
      '.vue',
    ],
  },
  server: {
    host: '0.0.0.0',
    port: 3000,
    
    // Switch to https protocol with self-signed certificate and private key
    https: {
      key: fs.existsSync(path.resolve(__dirname, 'incoming/key.pem')) ? fs.readFileSync(path.resolve(__dirname, 'incoming/key.pem')) : undefined,
      cert: fs.existsSync(path.resolve(__dirname, 'incoming/cert.pem')) ? fs.readFileSync(path.resolve(__dirname, 'incoming/cert.pem')) : undefined
    },

    // Proxy is used to redirect certain requests from webui server to cfm-service server
    proxy: {
      '/api': { // Requests to /api will be proxied to the target URL
        target: parsedConfig.api.base_path,
        secure: false, // This will ignore the self-signed certificate
        changeOrigin: true, // Ensures the host header of the request is changed to the target URL
        rewrite: (path) => path.replace(/^\/api/, ''),
        configure: (proxy) => {
          proxy.on('proxyRes', (proxyRes) => {
            // Remove the strict-transport-security header
            delete proxyRes.headers['strict-transport-security'];
          });
        },
      },
    },
  },
})
