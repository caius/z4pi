package app

type App struct {
	KbusPath string
	HttpPort int
}

func NewApp(kbus_path string, http_port int) (*App, error) {
	newApp := App{
		KbusPath: kbus_path,
		HttpPort: http_port,
	}

	return &newApp, nil
}

func (a *App) Run() error {
	return nil
}
