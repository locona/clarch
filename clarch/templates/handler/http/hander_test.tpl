package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/{{.GithubUser}}/{{.GithubRepository}}/job"
	"github.com/{{.GithubUser}}/{{.GithubRepository}}/job/usecase/mock"
	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var m {{.Pkg}}.{{.CamelPkg}}
	err := faker.FakeData(&m)
	assert.NoError(t, err)

	mockUCase := new(mock.JobUsecase)
	mockJobList := make([]*{{.Pkg}}.{{.CamelPkg}}, 0)
	mockJobList = append(mockJobList, &m)
	mockUCase.On("FindAll").Return(mockJobList, nil)

	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	handler := HttpJobHandler{
		{{.Pkg}}UC: mockUCase,
	}
	handler.FindAll(c)

	pp.Println(res)
	assert.Equal(t, http.StatusOK, res.Code)
	mockUCase.AssertExpectations(t)
}

func TestFindById(t *testing.T) {
	var j {{.Pkg}}.{{.CamelPkg}}
	err := faker.FakeData(&j)
	assert.NoError(t, err)
}

func TestStore(t *testing.T) {
}

func TestUpdate(t *testing.T) {
}

func TestDelete(t *testing.T) {
}
