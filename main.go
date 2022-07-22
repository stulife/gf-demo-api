package main

import (
	"gf-demo-api/internal/cmd"
	_ "gf-demo-api/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func main() {
	genv.Set("GF_GCFG_FILE", "config.test.yaml")
	cmd.Main.Run(gctx.New())
}
