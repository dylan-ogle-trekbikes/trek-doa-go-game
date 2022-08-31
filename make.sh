env GOOS=js GOARCH=wasm go build -o trekdoa.wasm .
mv trekdoa.wasm ./dist/
cp -r images/* ./dist/images/