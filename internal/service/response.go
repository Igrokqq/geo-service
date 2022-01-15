package service

type Response struct {
	Error error
	Code  int
	Value interface{}
}
