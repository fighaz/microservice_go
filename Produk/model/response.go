package model

type ResponseProduct struct {
	Message string
	Data    Product
}
type Response struct {
	Message string
	Data    interface{}
}
