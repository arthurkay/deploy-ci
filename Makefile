# Add missing dependencies to the go.mod file
tidy:
	go mod tidy
	
# Build the single binary of the app
build:
	go build

# Clean up the old binary builds
clean:
	rm ./deploy-ci

# Start the deploy-ci daemon
start:
	sudo systemctl start deploy-ci

# Restart the deploy-ci daemon
restart:
	sudo systemctl restart deploy-ci

# Stop the deploy-ci daemon
stop:
	sudo systemctl start deploy-ci

# Clean up old builds, get dependecies, build and restart the app daemon
update: tidy build restart

# Configure systemd and start the vieo daemon
setup: tidy build
	cat deploy-ci.conf > deploy-ci.temp
	sed -i 's+User=.*+User='"$$(whoami)"'+g' ./deploy-ci.temp
	sed -i 's+EnvironmentFile=.*+EnvironmentFile='"$$(pwd)"/environment.conf'+g' ./deploy-ci.temp
	sed -i 's+WorkingDirectory=.*+WorkingDirectory='"$$(pwd)"'+g' ./deploy-ci.temp
	sed -i 's+ExecStart=.*+ExecStart='"$$(pwd)"'/deploy-ci ${DEPLOY_PORT} ${DEPLOY_ANSIBLE_FILE} ${DEPLOY_USER}+g' ./deploy-ci.temp
	sudo mv ./deploy-ci.temp /etc/systemd/system/deploy-ci.service
	sudo systemctl daemon-reload
	sudo systemctl restart deploy-ci
