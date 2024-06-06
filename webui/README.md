# cfm-webui

A Composer & Fabric Manager (CFM) single-page application presenting a web UI using Vue.js 3.

## vue project

The `webui` project was created using `npm create vuetify`

- ✔ Project name: … webui
- ✔ Which preset would you like to install? › Base (Vuetify, VueRouter)
- ✔ Use TypeScript? … Yes
- ✔ Would you like to install dependencies with yarn, npm, pnpm, or bun? › npm

## Build Setup on Ubuntu Linux

- Install dependencies

```bash
sudo apt update
sudo apt install nodejs npm
npm --version
```

- Install project dependencies

```bash
npm install
```

- CFM service connection configuration
  - The cfm-webui defaults to looking for a cfm-service connection at 127.0.0.1:8080.
  - If a different IP is desired, change the YAML configuration file `config.yaml` in the root folder.

## Run Webui

- Run `cfm-webui`

```bash
# Go to the `webui` project folder
cd webui
npm run dev
```

- Open browser
  - Press 'o' to open a browser window (or copy the URL from the terminal and paste into your browser)
  - Press 'h' for help
