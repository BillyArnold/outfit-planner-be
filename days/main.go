package days_server

func Start() {
	s := NewServer(ServerOpts{port: ":3002"})

	s.Start()
}
