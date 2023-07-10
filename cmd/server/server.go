package server

/*
 * Project: I-wish-you
 * Created Date: Sunday, July 9th 2023, 10:45:22 am
 * Author: Olimpiev Y. Y.
 * -----
 * Last Modified:  yr.olimpiev@gmail.com
 * Modified By: Olimpiev Y. Y.
 * -----
 * Copyright (c) 2023 NSU
 *
 * -----
 */

import (
	"strconv"
)

type Server struct {
	addr   string
	port   int
	router *Router
}

type Route struct {
	Path    string `json:"path"`
	Method  string `json:"method"`
	Handler string `json:"handler"`
}

func NewServer(addr string, port int) *Server {
	server := &Server{
		addr:   addr,
		port:   port,
		router: &Router{},
	}
	server.configure()
	return server
}

func (s *Server) Run() error {
	return s.router.Run(s.addr + ":" + strconv.Itoa(s.port))
}

func (s *Server) configure() {
	s.router.configure()
}
