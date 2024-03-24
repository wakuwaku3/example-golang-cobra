package echo

import "github.com/wakuwaku3/example-golang-cobra/lib/command"

type usecase struct {
	args []string
}

func NewUsecase(args []string) *usecase {
	return &usecase{args}
}

func (u *usecase) Execute() error {
	return command.Execute("echo", u.args...)
}
