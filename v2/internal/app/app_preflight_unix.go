//go:build (linux || darwin) && !bindings

package app

import (
	"github.com/gcfguo/wails/v2/internal/logger"
	"github.com/gcfguo/wails/v2/pkg/options"
)

func PreflightChecks(_ *options.App, _ *logger.Logger) error {
	return nil
}
