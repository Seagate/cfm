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
    host:'0.0.0.0',
    port: 3000,
  },
})
