package store

type train struct {
	booked      int
	source      string
	destination string
}

func NewTrain(booked int, source, destination string) *train {
	return &train{booked, source, destination}
}
