# Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

APP_NAME := cxl-host
INSTALL_PATH := /usr/local/bin/
SYSTEMD_PATH := /etc/systemd/system/


help:
	@echo ""
	@echo "-----------------------------------------------------------------------------------"
	@echo "make clean        - remove all"
	@echo "make local        - build a local executable"
	@echo "make run          - build a local executable and run it"
	@echo "make install      - Install $(APP_NAME) as a system daemon"
	@echo "make uninstall    - Uninstall $(APP_NAME) from system"
	@echo ""

clean:
	@echo "Clean up..."
	go clean
	rm -f $(APP_NAME)

local: clean
	@echo "Build local executable..."
	go build -o $(APP_NAME) -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" main.go
	ls -lh $(APP_NAME)

run: local
	@echo "Running $(APP_NAME)..."
	./$(APP_NAME) -verbosity=0

install: 
	@echo "Install $(APP_NAME) to $(INSTALL_PATH)"
	cp $(APP_NAME) $(INSTALL_PATH)
	@echo "[Unit]" > $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "Description= Seagate Composer & Fabric Manager CXL host Redfish service" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "Wants=network.target" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "After=syslog.target network-online.target" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "[Service]" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "Type=idle" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "ExecStart=$(INSTALL_PATH)$(APP_NAME)" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "Restart=on-failure" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "RestartSec=10" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "KillMode=process" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "[Install]" >> $(SYSTEMD_PATH)$(APP_NAME).service
	@echo "WantedBy=multi-user.target" >> $(SYSTEMD_PATH)$(APP_NAME).service
	chmod 640 $(SYSTEMD_PATH)$(APP_NAME).service
	systemctl daemon-reload
	systemctl enable $(APP_NAME)
	systemctl start $(APP_NAME)

uninstall:
	systemctl stop $(APP_NAME).service
	systemctl disable $(APP_NAME).service
	rm $(SYSTEMD_PATH)$(APP_NAME).service
	systemctl daemon-reload
	rm $(INSTALL_PATH)$(APP_NAME)