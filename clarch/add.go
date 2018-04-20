package clarch

import (
	"fmt"
	"os"
	"strings"
)

type handler struct{}
type http struct{}
type uc struct{}
type repo struct{}
type model struct{}
type api struct{}

func (this *handler) mapTplAndOutPath(pkg string) []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/handler/http/http.tpl",
		"out": fmt.Sprintf("%s/handler/http/%s_handler.go", pkg, pkg),
	})

	m = append(m, map[string]string{
		"tpl": "clarch/templates/handler/http/params.tpl",
		"out": fmt.Sprintf("%s/handler/http/params.go", pkg),
	})

	return m
}

func (this *uc) mapTplAndOutPath(pkg string) []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/usecase/usecase.tpl",
		"out": fmt.Sprintf("%s/usecase/%s_uc.go", pkg, pkg),
	})
	return m
}

func (this *repo) mapTplAndOutPath(pkg string) []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/repository/repository.tpl",
		"out": fmt.Sprintf("%s/repository/repository.go", pkg),
	})
	return m
}

func (this *model) mapTplAndOutPath(pkg string) []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/model.tpl",
		"out": fmt.Sprintf("model/%s.go", pkg),
	})
	return m
}

func (this *api) mapTplAndOutPath(pkg string) []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/add.tpl",
		"out": "",
	})
	return m
}

func (this *cmd) Add() {
	GenHttp(this)
	GenUC(this)
	GenRepo(this)
	GenModel(this)
	OutApi(this)
}

func NewAdd(pkg string) *cmd {
	dir, _ := os.Getwd()
	dirNames := strings.Split(dir, "/")
	return &cmd{
		value: value{
			CurrentDir:  dir,
			CurrentUser: dirNames[len(dirNames)-2],
			CurrentRepo: dirNames[len(dirNames)-1],
			CamelPkg:    camelCase(pkg),
			Pkg:         pkg,
		},
	}
}

func GenHttp(cmd *cmd) {
	var o handler
	cmd.gen(o.mapTplAndOutPath(cmd.value.Pkg))
}

func GenUC(cmd *cmd) {
	var o uc
	cmd.gen(o.mapTplAndOutPath(cmd.value.Pkg))
}

func GenRepo(cmd *cmd) {
	var o repo
	cmd.gen(o.mapTplAndOutPath(cmd.value.Pkg))
}

func GenModel(cmd *cmd) {
	var o model
	cmd.gen(o.mapTplAndOutPath(cmd.value.Pkg))
}

func OutApi(cmd *cmd) {
	var o api
	cmd.stdout(o.mapTplAndOutPath(cmd.value.Pkg))
}
