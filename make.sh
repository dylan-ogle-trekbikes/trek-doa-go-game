env GOOS=js GOARCH=wasm go build -o trekdoa.wasm .
cp trekdoa.wasm ./static
cp -r images/* ./static/images/
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./static