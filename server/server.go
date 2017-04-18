package main

import (
	"database/sql"
	"log"

	"github.com/VG-Tech-Dojo/vg-1day-2017/server/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Server is whole server implementation for this app
type Server struct {
	db     *sql.DB
	Engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Engine: gin.Default(),
	}
}

func (s *Server) Init() error {
	// open db connection
	db, err := sql.Open("sqlite3", "./sample.db")
	if err != nil {
		return err
	}
	s.db = db

	// routing
	message := &controller.Message{DB: db}
	s.Engine.GET("/", message.Root)

	return nil
}

func (s *Server) Close() error {
	return s.db.Close()
}

func (s *Server) Run() {
	s.Engine.Run()
}

func main() {
	// TODO: dbのconfigとかenvとか引数で受け取るようにする

	s := NewServer()
	if err := s.Init(); err != nil {
		log.Fatalf("fail to start server: ", err)
	}
	defer s.Close()

	s.Run()
}
