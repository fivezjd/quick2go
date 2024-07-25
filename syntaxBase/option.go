package syntaxBase

/**
option 模式的优势是Server 新增了字段之后，不需要需改NewServer,只需要新增一个Option即可，只不过需要注意，
option模式需要使用指针，否则无法修改值

*/

type Server struct {
	Addr string
	Port int
	Name string
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	s := &Server{
		Addr: "0.0.0.0",
		Port: 8080,
		Name: "server",
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithAddr(addr string) Option {
	return func(s *Server) {
		s.Addr = addr
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithName(name string) Option {
	return func(s *Server) {
		s.Name = name
	}
}
