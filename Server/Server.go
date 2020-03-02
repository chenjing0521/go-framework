package Server

type Server interface {
	ListenAndServe() error
}
