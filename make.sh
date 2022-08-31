env GOOS=js GOARCH=wasm go build -o trekdoa.wasm ./dist
cp -r images/* ./dist/images/