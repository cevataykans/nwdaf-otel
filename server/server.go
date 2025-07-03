package server

type Server interface {
	Setup()
	Start() chan error
}
