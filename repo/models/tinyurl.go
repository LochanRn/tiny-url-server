package models

import (
	"time"
)

type URL struct {
	OriginalURL string    // The original long URL
	TinyURL     string    // The shortened URL
	CreatedAt   time.Time // Creation timestamp
}

type Counter struct {
	Count  int
	Domain string
}

type CounterList []Counter

func (c CounterList) Len() int           { return len(c) }
func (c CounterList) Less(i, j int) bool { return c[i].Count > c[j].Count }
func (c CounterList) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
