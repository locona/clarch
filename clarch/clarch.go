package clarch

//
// import (
// "fmt"
// "html/template"
// "log"
// "os"
// "os/exec"
//
// "github.com/k0kubun/pp"
// "github.com/locona/clarch/bindata"
// )
//
// func (this *Clarch) Run() {
// if err := this.genEntity(); err != nil {
// log.Fatal(err)
// }
// if err := this.genRepo(); err != nil {
// log.Fatal(err)
// }
// if err := this.genUseCase(); err != nil {
// log.Fatal(err)
// }
// if err := this.genHttpHandler(); err != nil {
// log.Fatal(err)
// }
// log.Println("Complete ...")
// }
//
// func (this *Clarch) genHttpHandler() error {
// filePath := fmt.Sprintf("%s/handler/http", this.Config.Pkg)
// if err := mkdir(filePath); err != nil {
// return err
// }
//
// tplFile, err := bindata.Asset("clarch/templates/handler/http/http.tpl")
// t := template.New("handler")
// t, _ = t.Parse(string(tplFile))
// fp, err := os.OpenFile(fmt.Sprintf("%s/%s_handler.go", filePath, this.Config.Pkg), os.O_RDWR|os.O_CREATE, 0666)
// if err != nil {
// return err
// }
// defer fp.Close()
//
// value := struct {
// CurrentUser       string
// CurrentRepo string
// CamelPkg         string
// Pkg              string
// }{
// CurrentUser:       this.Config.CurrentUser,
// CurrentRepo: this.Config.CurrentRepo,
// CamelPkg:         camelCase(this.Config.Pkg),
// Pkg:              this.Config.Pkg,
// }
//
// return t.Execute(fp, value)
// }
//
// func (this *Clarch) genUseCase() error {
// filePath := fmt.Sprintf("%s/%s", this.Config.Pkg, "usecase")
// if err := mkdir(filePath); err != nil {
// return err
// }
//
// tplFile, err := bindata.Asset("clarch/templates/usecase/usecase.tpl")
// t := template.New("usecase")
// t, _ = t.Parse(string(tplFile))
// t = t.Funcs(template.FuncMap{
// "safeHTML": func(s interface{}) template.HTML {
// return template.HTML(fmt.Sprint(s))
// },
// })
// fp, err := os.OpenFile(fmt.Sprintf("%s/%s_ucase.go", filePath, this.Config.Pkg), os.O_RDWR|os.O_CREATE, 0666)
// defer fp.Close()
// if err != nil {
// return err
// }
//
// value := struct {
// CurrentUser       string
// CurrentRepo string
// CamelPkg         string
// Pkg              string
// }{
// CurrentUser:       this.Config.CurrentUser,
// CurrentRepo: this.Config.CurrentRepo,
// CamelPkg:         camelCase(this.Config.Pkg),
// Pkg:              this.Config.Pkg,
// }
// if err := t.Execute(fp, value); err != nil {
// return err
// }
//
// cmd := exec.Command("mockery", "-dir", fmt.Sprintf("%s/usecase", this.Config.Pkg), "-name", fmt.Sprintf("%sUsecase", camelCase(this.Config.Pkg)), "-outpkg", "mocks", "-output", fmt.Sprintf("%s/usecase/mocks", this.Config.Pkg))
// if err := cmd.Run(); err != nil {
// pp.Println(cmd)
// return err
// }
// return nil
// }
//
// func (this *Clarch) genRepo() error {
// filePath := fmt.Sprintf("%s/%s", this.Config.Pkg, "repository")
// if err := mkdir(filePath); err != nil {
// return err
// }
//
// tplFile, err := bindata.Asset("clarch/templates/repository/repository.tpl")
// t := template.New("repository")
// t, _ = t.Parse(string(tplFile))
// t = t.Funcs(template.FuncMap{
// "safeHTML": func(s interface{}) template.HTML {
// return template.HTML(fmt.Sprint(s))
// },
// })
// fp, err := os.OpenFile(fmt.Sprintf("%s/repository.go", filePath), os.O_RDWR|os.O_CREATE, 0666)
// defer fp.Close()
// if err != nil {
// return err
// }
//
// value := struct {
// CurrentUser       string
// CurrentRepo string
// CamelPkg         string
// Pkg              string
// CamelDB          string
// DB               string
// }{
// CurrentUser:       this.Config.CurrentUser,
// CurrentRepo: this.Config.CurrentRepo,
// CamelPkg:         camelCase(this.Config.Pkg),
// Pkg:              this.Config.Pkg,
// CamelDB:          camelCase(this.Config.DB),
// DB:               this.Config.DB,
// }
// if err := t.Execute(fp, value); err != nil {
// return err
// }
//
// cmd := exec.Command("mockery", "-dir", fmt.Sprintf("%s/repository", this.Config.Pkg), "-name", fmt.Sprintf("%sRepository", camelCase(this.Config.Pkg)), "-outpkg", "mocks", "-output", fmt.Sprintf("%s/repository/mocks", this.Config.Pkg))
// if err := cmd.Run(); err != nil {
// return err
// }
// tplFile2, err := bindata.Asset("clarch/templates/repository/gorm_repository.tpl")
// t2 := template.New("gorm_repository")
// t2, _ = t2.Parse(string(tplFile2))
// t2 = t2.Funcs(template.FuncMap{
// "safeHTML": func(s interface{}) template.HTML {
// return template.HTML(fmt.Sprint(s))
// },
// })
//
// fp2, err := os.OpenFile(fmt.Sprintf("%s/%s_%s.go", filePath, this.Config.DB, this.Config.Pkg), os.O_RDWR|os.O_CREATE, 0666)
// defer fp2.Close()
// if err != nil {
// return err
// }
//
// if err := t2.Execute(fp2, value); err != nil {
// return err
// }
//
// return nil
// }
//
// func (this *Clarch) genEntity() error {
// filePath := this.Config.Pkg
// if err := mkdir(filePath); err != nil {
// return err
// }
//
// tplFile, err := bindata.Asset("clarch/templates/entity.tpl")
// t := template.New("entity")
// t, _ = t.Parse(string(tplFile))
// fp, err := os.OpenFile(fmt.Sprintf("%s/entity.go", filePath), os.O_RDWR|os.O_CREATE, 0666)
// defer fp.Close()
// if err != nil {
// return err
// }
//
// value := struct {
// CurrentUser       string
// CurrentRepo string
// CamelPkg         string
// Pkg              string
// }{
// CurrentUser:       this.Config.CurrentUser,
// CurrentRepo: this.Config.CurrentRepo,
// CamelPkg:         camelCase(this.Config.Pkg),
// Pkg:              this.Config.Pkg,
// }
// return t.Execute(fp, value)
// }
