package main

import (
	"math/rand"
	"time"
)

type Player struct {
	Name string `json:"name"`
	City string `json:"city"`
}

var players = []Player{
	{Name: "David Beckham"},
	{Name: "Steven Gerrard"},
	{Name: "Frank Lampard"},
	{Name: "Rio Ferdinand"},
	{Name: "John Terry"},
	{Name: "Paul Scholes"},
	{Name: "Michael Owen"},
	{Name: "Gary Neville"},
	{Name: "Raheem Sterling"},
	{Name: "Harry Kane"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := range players {
		players[i].City = cities[rand.Intn(len(cities))]
	}
}
