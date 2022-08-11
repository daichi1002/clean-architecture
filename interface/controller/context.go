package controller

// 今回利用するginのメソッドのインターフェースを定義
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
