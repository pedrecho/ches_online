package app

import (
	"chess/internal/config"
	"chess/pkg/zaplogger"
	"fmt"
)

func Run(configPath string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("config initialization: %w", err)
	}
	zapsync, err := zaplogger.ReplaceZap(cfg.Logger)
	if err != nil {
		return fmt.Errorf("zaplogger initialization: %w", err)
	}
	defer zapsync()

	return nil
}
