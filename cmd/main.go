package main

import (
	"github.com/i0li/super_shiharai_kun/internal/config"
	"github.com/i0li/super_shiharai_kun/internal/infra/db"
	"github.com/i0li/super_shiharai_kun/internal/server"
)

func main() {
	config.Init()

	conn := db.MustConnect()

	r := server.NewRouter(conn)

	r.Run(":" + config.GetPort())
}
