package clarch

import (
	"fmt"
	"html/template"
	"os"

	"github.com/locona/clarch/bindata"
)

type value struct {
	CurrentDir  string
	CurrentRepo string
	CurrentUser string
	CamelPkg    string
	Pkg         string
}

type cmd struct {
	value value
}

func (this *cmd) gen(maps []map[string]string) {
	for idx := range maps {
		fp, err := openOrCreate(maps[idx]["out"])
		defer fp.Close()
		if err != nil {
			panic(err)
		}
		b, err := bindata.Asset(maps[idx]["tpl"])
		if err != nil {
			panic(err)
		}

		t := template.New("gen")
		t, _ = t.Parse(string(b))
		t = t.Funcs(template.FuncMap{
			"safeHTML": func(s interface{}) template.HTML {
				return template.HTML(fmt.Sprint(s))
			},
		})

		if err := t.Execute(fp, this.value); err != nil {
			panic(err)
		}
	}
}

func (this *cmd) stdout(maps []map[string]string) {
	for idx := range maps {
		b, err := bindata.Asset(maps[idx]["tpl"])
		if err != nil {
			panic(err)
		}

		t := template.New("gen")
		t, _ = t.Parse(string(b))
		t = t.Funcs(template.FuncMap{
			"safeHTML": func(s interface{}) template.HTML {
				return template.HTML(fmt.Sprint(s))
			},
		})

		if err := t.Execute(os.Stdout, this.value); err != nil {
			panic(err)
		}
	}
}
