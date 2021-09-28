package app

type App struct {
	SerialPort string
	Port       int
}

func NewApp(serial_path string, port int) (*App, error) {
	newApp := App{
		SerialPort: serial_path,
		Port:       port,
	}

	return &newApp, nil
}

func (a *App) Run() error {
	return nil
}
