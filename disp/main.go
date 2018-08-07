package main

import "github.com/lanzafame/kh"

var usage = `disp usage:
	list	-- list available displays
	config	-- configure xrandr based on available screens
`

type disp struct{}

func (d *disp) Help(fa *kh.FingerArgs, fr *kh.Response) error {
	fr.SetVerbose(fa.Flags.Verbose)
	fr.WriteStdout(usage)
	return nil
}

func (d *disp) Execute(fa *kh.FingerArgs, fr *kh.Response) error {
	fr.SetVerbose(fa.Flags.Verbose)

	args := fa.Args
	if len(args) < 1 {
		fr.WriteStderr(usage)
		return nil
	}

	switch args[0] {
	case "list":
		fr.WriteStdout("listing stuff")
	case "config":
		fr.WriteStdout("configing stuff")
	default:
		fr.WriteStderr(usage)
	}

	return nil
}
