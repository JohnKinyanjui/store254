#!/bin/bash
sudo apt update
sudo apt install -y supervisor nginx software-properties-common

# Supervisor configuration for your Go application
DIR="/root"

GO_APP_NAME="store_private_api"             # Replace with your app's name
GO_APP_DIR="$DIR/store254_backend/bin/eccommerce"        # Replace with your app's directory
GO_APP_COMMAND="main"  

# Create Supervisor configuration file for the Go app
sudo bash -c "cat > /etc/supervisor/conf.d/$GO_APP_NAME.conf" << EOF
[program:$GO_APP_NAME]
command=$GO_APP_DIR/main
directory=$DIR/store254_backend
autostart=true
autorestart=true
stderr_logfile=/var/log/backend_error.err.log
stdout_logfile=/var/log/backend_error.out.log
EOF

# Update Supervisor and start the Go app
sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl start $GO_APP_NAME


