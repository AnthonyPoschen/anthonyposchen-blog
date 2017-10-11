# generate frontend production files.
cd ./frontend && npm run-script build && cd ../
#embed files within the binary.
packr build -o ./build/main main.go