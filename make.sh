env GOOS=js GOARCH=wasm go build -o trekdoa.wasm .
cp -r images/* ./dist/images/