package pkg

type Level struct {
	Name  string
	Value int
	Help  Help
}

type Help struct {
	Title string
	Info  string
	Code  string
}
