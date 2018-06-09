package main

import (
	"context"
	"sort"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Balls []Ball `json:"balls"`
}

type Ball struct {
	Color string `json:"color"`
	Total int    `json:"total"`
}

type EventResult struct {
	Balls             []Ball `json:"ballColors"`
	QuantityUnitBalls int    `json:"quantityUnitBalls"`
}

func HandleEvent(ctx context.Context, event Event) (EventResult, error) {
	var eventResult EventResult
	balls := make(map[string]Ball)

	for _, ball := range event.Balls {
		parse(ball, balls)
	}

	eventResult.Balls = parseToEventResult(balls)
	eventResult.QuantityUnitBalls = len(parseToEventResult(balls))

	return eventResult, nil
}

func parse(ball Ball, balls map[string]Ball) {
	value, exist := balls[ball.Color]

	if !exist {
		balls[ball.Color] = build(ball.Color)
	} else {
		value.Total++
		balls[ball.Color] = value
	}
}

func parseToEventResult(balls map[string]Ball) []Ball {
	ballToEventResult := make([]Ball, 0)

	for _, ball := range balls {
		ballToEventResult = append(ballToEventResult, ball)
	}

	return sortBalls(ballToEventResult)
}

func sortBalls(balls []Ball) []Ball {
	sort.Slice(balls, func(i, j int) bool {
		return balls[j].Total < balls[i].Total
	})

	return balls
}

func build(color string) Ball {
	var ball Ball
	ball.Color = color
	ball.Total = 1

	return ball
}

func main() {
	lambda.Start(HandleEvent)
}
