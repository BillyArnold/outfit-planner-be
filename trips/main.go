package trips_server

func Start() {
	s := NewServer(ServerOpts{port: ":3001"})

	s.Start()
}
