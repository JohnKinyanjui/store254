#!/bin/bash

# Update and install necessary packages
sudo apt update
sudo apt install -y nginx certbot python3-certbot-nginx supervisor
sudo apt install unzip
# Install Bun (check the latest installation method from Bun documentation)
curl -fsSL https://bun.sh/install | bash
source /root/.bashrc
# Define application variables
DIR="/root"
APP_NAME="store254_site" # Replace with your app's name
DOMAIN="fusionmarket.cloud"  # Replace with your domain
PORT=3000 # Port for your app to listen on

bun install
bun run build

# Supervisor configuration for the Svelte app
sudo bash -c "cat > /etc/supervisor/conf.d/$APP_NAME.conf" << EOF
[program:$APP_NAME]
command= $DIR/.bun/bin/bun $DIR/store254_site/build/index.js
directory=$DIR/store254_site
autostart=true
autorestart=true
stderr_logfile=/var/log/$APP_NAME.err.log
stdout_logfile=/var/log/$APP_NAME.out.log
EOF

# Update Supervisor and start the app
sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl start $APP_NAME

# Nginx configuration as a reverse proxy for the Svelte app
sudo bash -c "cat > /etc/nginx/sites-available/$DOMAIN" << EOF
server {
    listen 80;
    server_name $DOMAIN;

    location / {
        proxy_pass http://localhost:$PORT;
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host \$host;
        proxy_cache_bypass \$http_upgrade;
    }
}
EOF

# Enable the site and reload Nginx
sudo ln -s /etc/nginx/sites-available/$DOMAIN /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx

# Add Certbot PPA and install Certbot
sudo add-apt-repository ppa:certbot/certbot -y
sudo apt update
sudo apt install -y certbot python3-certbot-nginx

# Obtain and install SSL certificate
sudo certbot --nginx -d $DOMAIN

echo "Svelte app deployment complete!"
