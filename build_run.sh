mkdir _build
mkdir _build/media
#cp config.yaml _build/app.config.yaml
go build -o _build
cd spa_web
npm install
rm -R buid
npm react-scripts build
cd ../
chmod +x ./_build/bixpark_server
./_build/bixpark_server  -config ./_build/app.config.yaml