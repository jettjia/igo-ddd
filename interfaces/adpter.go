package interfaces

import "fmt"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RunApp() {
	fmt.Println("开始运行项目")
}
