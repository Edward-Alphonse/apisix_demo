mkdir -p ./output/config

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./output/main
cp ./config/config.yaml ./output/config/config.yaml