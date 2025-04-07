package main

import (
	"errors"
	"strings"

	"github.com/zrcoder/agg/internal"
	_ "github.com/zrcoder/agg/internal/exported/github.com/zrcoder/agg/pkg/api/hanoi"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/igop/gopbuild/pkg"
)

const (
	gopfileName = "main.gop"

	hanoiPreCodes = `
	import (
		. "github.com/zrcoder/agg/pkg/api/hanoi"
	)

	var (
		a = A
		b = B
		c = C
	)

	type Pile = func()
	`
)

func hanoiCodeAction(code string, preAction func() error) error {
	if strings.TrimSpace(code) == "" {
		return errors.New("code is empty")
	}
	err := preAction()
	if err != nil {
		return err
	}
	code = hanoiPreCodes + code
	_, err = igop.RunFile(gopfileName, code, nil, 0)
	return err
}

func main() {
	internal.Run(hanoiCodeAction)
}
