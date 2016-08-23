package main

import (
	"github.com/Lanzafame/kh"
)

type Finger struct{}

func (p *Finger) Help(fa *kh.FingerArgs, resp *kh.Response) error {
	resp.SetVerbose(fa.Flags.Verbose)
	resp.WriteStdout("mara makes deploying to marathon simple")
	return nil
}

func (p *Finger) Execute(fa *kh.FingerArgs, resp *kh.Response) error {
	resp.SetVerbose(fa.Flags.Verbose)
	resp.WriteStdout("make all the marathons easy")
	return nil
}

func main() {
	finger := &Finger{}
	kh.Register(finger)
	kh.Run()
}
