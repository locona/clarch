package clarch

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ini struct{}

func (this *ini) mapTplAndOutPath() []map[string]string {
	dirPath := "project"
	maps := make([]map[string]string, 0)
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/main.tpl",
		"out": "main.go",
	})

	// cmd
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/cmd/root.tpl",
		"out": "cmd/root.go",
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/cmd/api.tpl",
		"out": "cmd/api.go",
	})

	// project
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/handler.tpl",
		"out": fmt.Sprintf("%s/handler.go", dirPath),
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/repository.tpl",
		"out": fmt.Sprintf("%s/repository.go", dirPath),
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/config.tpl",
		"out": fmt.Sprintf("%s/config.go", dirPath),
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/validator/validator.tpl",
		"out": fmt.Sprintf("%s/validator/validator.go", dirPath),
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/logger/logger.tpl",
		"out": fmt.Sprintf("%s/logger/logger.go", dirPath),
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/middleware/middleware.tpl",
		"out": fmt.Sprintf("%s/middleware/middleware.go", dirPath),
	})

	// infra
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/infra/mysql.tpl",
		"out": "infra/mysql.go",
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/infra/redis.tpl",
		"out": "infra/redis.go",
	})

	return maps
}

func (cmd *cmd) Init() {
	var o ini
	cmd.gen(o.mapTplAndOutPath())

	if err := exec.Command("dep", "init").Run(); err != nil {
		panic(err)
	}
}

func NewInit() *cmd {
	dir, _ := os.Getwd()
	dirNames := strings.Split(dir, "/")
	return &cmd{
		value: value{
			CurrentDir:  dir,
			CurrentUser: dirNames[len(dirNames)-2],
			CurrentRepo: dirNames[len(dirNames)-1],
		},
	}
}
