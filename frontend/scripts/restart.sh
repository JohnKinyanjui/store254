APP_NAME="store254_site" # Replace with your app's name

rm -rf ./build
rm -rf ./dist

bun install
bun run build

sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl start $APP_NAME