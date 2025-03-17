package api

import ("github.com/gin-gonic/gin")

type Server struct{
	router *gin.Engine
	listenAddress string
}

func NewServer()*Server{

	return &Server{
		router: gin.Default(),
		listenAddress: ":8080",
	}
}

func (s *Server) StartServer() error {

	err := s.router.Run(s.listenAddress)

	return err
}