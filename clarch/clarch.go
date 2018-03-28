package clarch

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/k0kubun/pp"
)

type Clarch struct {
	Config Config
}

func (this *Clarch) Init() {
	if err := this.genProject(); err != nil {
		log.Fatal(err)
	}
	log.Println("Complete ...")
}

func (this *Clarch) Run() {
	if err := this.genEntity(); err != nil {
		log.Fatal(err)
	}
	if err := this.genRepo(); err != nil {
		log.Fatal(err)
	}
	if err := this.genUseCase(); err != nil {
		log.Fatal(err)
	}
	if err := this.genHttpHandler(); err != nil {
		log.Fatal(err)
	}
	log.Println("Complete ...")
}

func (this *Clarch) genProject() error {
	dirPath := "project"
	if err := mkdir(dirPath); err != nil {
		return err
	}

	t := template.Must(template.ParseFiles("./clarch/templates/project/handler.tpl"))
	fp, err := os.OpenFile(fmt.Sprintf("%s/handler.go", dirPath), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()

	value := struct{}{}
	return t.Execute(fp, value)
}

func (this *Clarch) genHttpHandler() error {
	filePath := fmt.Sprintf("%s/handler/http", this.Config.Pkg)
	if err := mkdir(filePath); err != nil {
		return err
	}

	t := template.Must(template.ParseFiles("./clarch/templates/handler/http/http.tpl"))
	fp, err := os.OpenFile(fmt.Sprintf("%s/%s_handler.go", filePath, this.Config.Pkg), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()

	value := struct {
		GithubUser       string
		GithubRepository string
		CamelPkg         string
		Pkg              string
	}{
		GithubUser:       this.Config.GithubUser,
		GithubRepository: this.Config.GithubRepository,
		CamelPkg:         camelCase(this.Config.Pkg),
		Pkg:              this.Config.Pkg,
	}

	return t.Execute(fp, value)
}

func (this *Clarch) genUseCase() error {
	filePath := fmt.Sprintf("%s/%s", this.Config.Pkg, "usecase")
	if err := mkdir(filePath); err != nil {
		return err
	}

	t := template.Must(template.ParseFiles("./clarch/templates/usecase/usecase.tpl"))
	t = t.Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	})
	fp, err := os.OpenFile(fmt.Sprintf("%s/%s_ucase.go", filePath, this.Config.Pkg), os.O_RDWR|os.O_CREATE, 0666)
	defer fp.Close()
	if err != nil {
		return err
	}

	value := struct {
		GithubUser       string
		GithubRepository string
		CamelPkg         string
		Pkg              string
	}{
		GithubUser:       this.Config.GithubUser,
		GithubRepository: this.Config.GithubRepository,
		CamelPkg:         camelCase(this.Config.Pkg),
		Pkg:              this.Config.Pkg,
	}
	if err := t.Execute(fp, value); err != nil {
		return err
	}

	cmd := exec.Command("mockery", "-dir", fmt.Sprintf("%s/usecase", this.Config.Pkg), "-name", fmt.Sprintf("%sUsecase", camelCase(this.Config.Pkg)), "-output", fmt.Sprintf("%s/usecase/mock", this.Config.Pkg))
	if err := cmd.Run(); err != nil {
		pp.Println(cmd)
		return err
	}
	return nil
}

func (this *Clarch) genRepo() error {
	filePath := fmt.Sprintf("%s/%s", this.Config.Pkg, "repository")
	if err := mkdir(filePath); err != nil {
		return err
	}

	t := template.Must(template.ParseFiles("./clarch/templates/repository/repository.tpl"))
	t = t.Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	})
	fp, err := os.OpenFile(fmt.Sprintf("%s/repository.go", filePath), os.O_RDWR|os.O_CREATE, 0666)
	defer fp.Close()
	if err != nil {
		return err
	}

	value := struct {
		GithubUser       string
		GithubRepository string
		CamelPkg         string
		Pkg              string
		CamelDB          string
		DB               string
	}{
		GithubUser:       this.Config.GithubUser,
		GithubRepository: this.Config.GithubRepository,
		CamelPkg:         camelCase(this.Config.Pkg),
		Pkg:              this.Config.Pkg,
		CamelDB:          camelCase(this.Config.DB),
		DB:               this.Config.DB,
	}
	if err := t.Execute(fp, value); err != nil {
		return err
	}

	cmd := exec.Command("mockery", "-dir", fmt.Sprintf("%s/repository", this.Config.Pkg), "-name", fmt.Sprintf("%sRepository", camelCase(this.Config.Pkg)), "-output", fmt.Sprintf("%s/repository/mock", this.Config.Pkg))
	if err := cmd.Run(); err != nil {
		return err
	}

	t2 := template.Must(template.ParseFiles("./clarch/templates/repository/gorm_repository.tpl"))
	t2 = t2.Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	})

	fp2, err := os.OpenFile(fmt.Sprintf("%s/%s_%s.go", filePath, this.Config.DB, this.Config.Pkg), os.O_RDWR|os.O_CREATE, 0666)
	defer fp2.Close()
	if err != nil {
		return err
	}

	if err := t2.Execute(fp2, value); err != nil {
		pp.Println(err)
		return err
	}

	return nil
}

func (this *Clarch) genEntity() error {
	filePath := this.Config.Pkg
	if err := mkdir(filePath); err != nil {
		return err
	}
	t := template.Must(template.ParseFiles(fmt.Sprintf("./clarch/templates/%s.tpl", "entity")))
	fp, err := os.OpenFile(fmt.Sprintf("%s/entity.go", filePath), os.O_RDWR|os.O_CREATE, 0666)
	defer fp.Close()
	if err != nil {
		return err
	}

	value := struct {
		GithubUser       string
		GithubRepository string
		CamelPkg         string
		Pkg              string
	}{
		GithubUser:       this.Config.GithubUser,
		GithubRepository: this.Config.GithubRepository,
		CamelPkg:         camelCase(this.Config.Pkg),
		Pkg:              this.Config.Pkg,
	}
	return t.Execute(fp, value)
}

func mkdir(filePath string) error {
	return os.MkdirAll(filePath, os.ModePerm)
}

func camelCase(s string) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s))

	var prev rune
	for _, curr := range s {
		if !isDelimiter(curr) {
			if isDelimiter(prev) || (prev == 0) {
				buffer = append(buffer, toUpper(curr))
			} else {
				buffer = append(buffer, toLower(curr))
			}
		}
		prev = curr
	}

	return string(buffer)
}

func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}

func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
