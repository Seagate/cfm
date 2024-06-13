FROM golang:1.22

RUN apt update
RUN apt install -y nodejs npm

# Copy whole repo into the image
COPY . /cfm

# setup web UI and build the dist package
WORKDIR /cfm/webui
RUN npm install
RUN npm run build

# Build CFM SERVICE & CLI
WORKDIR /cfm
RUN rm cfmdatastore.json
RUN ln -s /local/cfmdatastore.json cfmdatastore.json 
RUN make local


# Start the service
ENTRYPOINT [  "/cfm/cfm-service" ] 
