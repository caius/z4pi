package app

import (
	"log"

	"github.com/qcasey/gokbus"
)

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

	kbus, err := gokbus.New(a.KbusPath, 9600)
	if err != nil {
		return err
	}

	go kbus.Start()

	// Read everything through to stdout
	for packet := range kbus.ReadChannel {
		log.Printf("<- %s", packet.Pretty())
	}

	return nil
}
