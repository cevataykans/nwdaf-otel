package nwdaf

type NWDAF interface {
}

type server struct {
}

func New() NWDAF {
	return server{}
}
