package main

import (
	"errors"
	"fmt"

	_ "github.com/zrcoder/agg/internal/exported/github.com/zrcoder/agg/pkg/api/hanoi"
	"github.com/zrcoder/agg/internal/hanoi"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/igop/pkg/sync"
)

const (
	gopfileName = "main.gop"

	preCodes = `
	import (
		. "github.com/zrcoder/agg/pkg/api/hanoi"
	)
	`
)

func runCode(code string) error {
	if code == "" {
		return errors.New("code is empty")
	}
	hanoi.Hanoi.Reset()
	code = preCodes + code
	_, err := igop.RunFile(gopfileName, code, nil, 0)
	fmt.Println(err)
	return err
}

func main() {
	hanoi.Run(runCode)
}
