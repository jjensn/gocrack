package server

import (
	"errors"

	"github.com/jjensn/gocrack/server/authentication"
	"github.com/jjensn/gocrack/server/filemanager"
	"github.com/jjensn/gocrack/server/notifications"
	"github.com/jjensn/gocrack/server/rpc"
	"github.com/jjensn/gocrack/server/storage"
	"github.com/jjensn/gocrack/server/web"
)

// Config describes all the configuration values of the GoCrack server
type Config struct {
	// Debug, if true will enable verbose logging and will overwrite the flag passed into the server
	Debug          bool                        `yaml:"debug"`
	WebServer      web.Config                  `yaml:"web_server"`
	RPC            rpc.Config                  `yaml:"rpc_server"`
	Database       storage.Config              `yaml:"database"`
	FileManager    filemanager.Config          `yaml:"file_manager"`
	Authentication authentication.AuthSettings `yaml:"authentication"`
	Notification   notifications.Config        `yaml:"notifications"`
}

func (s *Config) validate() error {
	if err := s.WebServer.Validate(); err != nil {
		return err
	}

	if err := s.FileManager.Validate(); err != nil {
		return err
	}

	if err := s.RPC.Validate(); err != nil {
		return err
	}

	if err := s.Database.Validate(); err != nil {
		return err
	}

	if s.Database.ConnectionString == "" {
		return errors.New("database.connection_string must not be empty")
	}

	if err := s.Authentication.Validate(); err != nil {
		return err
	}

	return s.Notification.Validate()
}
