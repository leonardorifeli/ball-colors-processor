GOOS=linux go build -o main
zip main main
aws lambda update-function-code --function-name ball-colors-processor --zip-file fileb://main.zip