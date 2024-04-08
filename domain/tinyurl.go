package domain

type TinyURL struct {
	URL               string
	TinyURL           string
	CreationTimeStamp Time
}

type Counter struct {
	Domain string
	Count  int
}

type CounterList []Counter
