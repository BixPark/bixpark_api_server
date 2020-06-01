mkdir _build
go build -o _build
cd spa_web
npm install
npm react-scripts build
cd ../
chmod +x ./_build/bixpark_server
./_build/bixpark_server  -config ./_build/app.config.yaml