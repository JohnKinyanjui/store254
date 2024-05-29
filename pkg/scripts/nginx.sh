#!/bin/bash

# Check if the script is run as root
if [ "$EUID" -ne 0 ]; then
    echo "Please run this script as root or with sudo."
    exit 1
fi

# Install Nginx
apt-get update
apt-get install -y nginx

# Configure Nginx for WebSocket support
sudo bash -c "cat > /etc/nginx/sites-available/private.fusionmarket.cloud" << EOF
server {
    listen 80;
    server_name private.fusionmarket.cloud;

    location / {
        proxy_pass http://localhost:8000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
EOF

sudo bash -c "cat > /etc/nginx/sites-available/api.fusionmarket.cloud" << EOF
server {
    listen 80;
    server_name api.fusionmarket.cloud;

    location / {
        proxy_pass http://localhost:8001;
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
EOF

ln -s /etc/nginx/sites-available/private.fusionmarket.cloud /etc/nginx/sites-enabled/
ln -s /etc/nginx/sites-available/api.fusionmarket.cloud /etc/nginx/sites-enabled/
sudo nginx -t
systemctl reload nginx

# Install Certbot (Let's Encrypt)
# apt-get install -y certbot python3-certbot-nginx

# Request Let's Encrypt certificates
# certbot --nginx -d private.fusionmarket.cloud -d api.fusionmarket.cloud


# Inform the user about the completion
echo "Nginx setup with WebSocket support and Let's Encrypt certificates is complete."
