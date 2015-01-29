package ci

import "os/exec"

type Status int

const (
	Pending Status = 1 << iota
	Running
	Success
	Failure
)

type BuildStep interface {
	Build()
	Errors() []error
	Status() Status
}

type ShellCommand struct {
	BuildStep
	*exec.Cmd
	err    error
	status Status
}

func NewShellCommand(name string, arg ...string) *ShellCommand {
	shellCmd := &ShellCommand{
		Cmd:    exec.Command(name, arg...),
		status: Pending,
	}
	return shellCmd
}

func (s *ShellCommand) Build() {
	err := s.Cmd.Start()
	if err != nil {
		s.err = err
		s.status = Failure
		return
	}
	s.status = Running
	err = s.Cmd.Wait()
	if err != nil {
		s.err = err
		s.status = Failure
		return
	}
	s.status = Success
}

func (s *ShellCommand) Status() Status {
	return s.status
}
func (s *ShellCommand) Errors() []error {
	return []error{s.err}
}
