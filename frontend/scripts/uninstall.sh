#!/bin/bash

# Define application specific variables
APP_NAME="store254_site" # Replace with your app's name
DOMAIN="fusionmarket.cloud"  

# Remove Supervisor configuration for the Svelte app
sudo supervisorctl stop $APP_NAME
sudo rm -f /etc/supervisor/conf.d/$APP_NAME.conf
sudo supervisorctl reread
sudo supervisorctl update

# Disable and remove the Nginx site configuration for your domain
sudo rm -f /etc/nginx/sites-available/$DOMAIN
sudo rm -f /etc/nginx/sites-enabled/$DOMAIN

# Optionally: If you want to also remove the SSL certificate (not generally recommended)
# sudo certbot delete --cert-name $DOMAIN

# Reload Nginx to apply changes
sudo nginx -t && sudo systemctl reload nginx

echo "Svelte app configurations have been removed."
