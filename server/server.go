package main

import (
	"database/sql"
	"flag"
	"log"

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

func (s *Server) Init(dbconf, env string) error {
	db, err := sql.Open("sqlite3", "./sample.db")
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Server) Close() error {
	return s.db.Close()
}

func (s *Server) Run() {
	s.Engine.Run()
}

func main() {
	var (
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	)
	flag.Parse()

	s := NewServer()
	if err := s.Init(*dbconf, *env); err != nil {
		log.Fatalf("fail to start server: ", err)
	}
	defer s.Close()

	s.Run()
}
