package common

type EndPoint struct {
	name   string
	input  string
	output string
	info   string
}

type Meta struct {
	user_id   string
	device_id string
	endpoints []EndPoint
}
