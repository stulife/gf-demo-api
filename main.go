package main

import (
	_ "gf-demo-api/internal/packed"

	"gf-demo-api/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
