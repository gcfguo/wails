package main

import (
	"github.com/gcfguo/wails/v2/cmd/wails/flags"
	"github.com/gcfguo/wails/v2/cmd/wails/internal"
	"github.com/gcfguo/wails/v2/internal/colour"
	"github.com/gcfguo/wails/v2/internal/github"
	"github.com/pterm/pterm"
)

func showReleaseNotes(f *flags.ShowReleaseNotes) error {
	if f.NoColour {
		pterm.DisableColor()
		colour.ColourEnabled = false
	}

	version := internal.Version
	if f.Version != "" {
		version = f.Version
	}

	app.PrintBanner()
	releaseNotes := github.GetReleaseNotes(version, f.NoColour)
	pterm.Println(releaseNotes)

	return nil
}
