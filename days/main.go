package days_server

func Start() {
	s := NewServer(ServerOpts{port: ":3004"})

	s.Start()
}
