package clarch

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/locona/clarch/bindata"
)

type handler struct {
	http http
}

type http struct {
	cmdadd
}

type uc struct {
	cmdadd
}
type repo struct {
	cmdadd
}
type entity struct {
	cmdadd
}

type value struct {
	CurrentDir  string
	CurrentRepo string
	CurrentUser string
	CamelPkg    string
	Pkg         string
}

type cmdadd struct {
	value value
}

func (this *http) gen() error {
	fp, err := this.openFile()
	defer fp.Close()
	if err != nil {
		return err
	}
	b, err := bindata.Asset(this.tplPath())
	if err != nil {
		return err
	}

	t := template.New("http")
	t, _ = t.Parse(string(b))
	t = t.Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	})
	return t.Execute(fp, this.value)
}

func (this *http) openFile() (*os.File, error) {
	filePath := fmt.Sprintf("%s/handler/http/%s_handler.go", this.value.Pkg, this.value.Pkg)
	return openOrCreate(filePath)
}

func (this *http) tplPath() string {
	return "clarch/templates/handler/http/http.tpl"
}

func (this *http) genPath() string {
	return fmt.Sprint("%s/entity.go")
}

func (this *uc) gen() error {
	fp, err := this.openFile()
	defer fp.Close()
	if err != nil {
		return err
	}
	b, err := bindata.Asset(this.tplPath())
	if err != nil {
		return err
	}

	t := template.New("http")
	t, _ = t.Parse(string(b))
	t = t.Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	})
	return t.Execute(fp, this.value)
}

func (this *uc) openFile() (*os.File, error) {
	filePath := fmt.Sprintf("%s/usecase/%s_uc.go", this.value.Pkg, this.value.Pkg)
	return openOrCreate(filePath)
}
func (this *uc) tplPath() string {
	return "clarch/templates/usecase/usecase.tpl"
}

func (this *repo) gen() error {
	fp, err := this.openFile()
	defer fp.Close()
	if err != nil {
		return err
	}
	b, err := bindata.Asset(this.tplPath())
	if err != nil {
		return err
	}

	t := template.New("repo")
	t, _ = t.Parse(string(b))
	t = t.Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	})
	return t.Execute(fp, this.value)
}

func (this *repo) openFile() (*os.File, error) {
	filePath := fmt.Sprintf("%s/repository/repository.go", this.value.Pkg)
	return openOrCreate(filePath)
}

func (this *repo) tplPath() string {
	return "clarch/templates/repository/repository.tpl"
}

func (this *entity) gen() error {
	fp, err := this.openFile()
	defer fp.Close()
	if err != nil {
		return err
	}
	b, err := bindata.Asset(this.tplPath())
	if err != nil {
		return err
	}

	t := template.New("entity")
	t, _ = t.Parse(string(b))
	t = t.Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	})
	return t.Execute(fp, this.value)
}

func (this *entity) openFile() (*os.File, error) {
	filePath := fmt.Sprintf("%s/entity.go", this.value.Pkg)
	return openOrCreate(filePath)
}

func (this *entity) tplPath() string {
	return "clarch/templates/entity.tpl"
}

func (this *cmdadd) Add() error {
	if err := GenHttp(this); err != nil {
		return err
	}
	if err := GenUC(this); err != nil {
		return err
	}
	if err := GenRepo(this); err != nil {
		return err
	}
	if err := GenEntity(this); err != nil {
		return err
	}

	return nil
}

func NewCmdAdd(pkg string) *cmdadd {
	dir, _ := os.Getwd()
	dirNames := strings.Split(dir, "/")
	return &cmdadd{
		value: value{
			CurrentDir:  dir,
			CurrentUser: dirNames[len(dirNames)-2],
			CurrentRepo: dirNames[len(dirNames)-1],
			CamelPkg:    camelCase(pkg),
			Pkg:         pkg,
		},
	}
}

func GenHttp(cmd *cmdadd) error {
	o := http{*cmd}
	return o.gen()
}

func GenUC(cmd *cmdadd) error {
	o := uc{*cmd}
	return o.gen()
}

func GenRepo(cmd *cmdadd) error {
	o := repo{*cmd}
	return o.gen()
}

func GenEntity(cmd *cmdadd) error {
	o := entity{*cmd}
	return o.gen()
}
