GOOS=linux go build -o main
sam local invoke BallColorsProcessor --event event.json