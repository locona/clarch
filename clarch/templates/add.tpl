// import {{.Pkg}}s
{{.Pkg}}Handler "github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}/handler/http"
{{.Pkg}}UC "github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}/usecase"
{{.Pkg}}Repo "github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}/repository"

// {{.Pkg}}
{
	{{.Pkg}}Repo := {{.Pkg}}Repo.New{{.CamelPkg}}Repository(infra.Mysql)
	{{.Pkg}}UC := {{.Pkg}}UC.New{{.CamelPkg}}Usecase({{.Pkg}}Repo)
	{{.Pkg}}Handler.New{{.CamelPkg}}HttpHandler(api, {{.Pkg}}UC)
}
