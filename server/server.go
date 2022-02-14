package server

func Initialize() {
	// config := config.GetConfig()

	r := NewRouter()
	// r.Run(config)
	r.Run()
}
