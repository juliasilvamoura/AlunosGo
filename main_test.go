package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juliasilvamoura/gin-api-rest/controller"
	"github.com/juliasilvamoura/gin-api-rest/database"
	"github.com/juliasilvamoura/gin-api-rest/model"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // modo mais simplificado de realise
	rotas := gin.Default()
	return rotas
}

// Criando mocks para os testes
func CriaAlunoMock() {
	aluno := model.Aluno{Nome: "Teste", RG: "123456789", CPF: "12345678912"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}
func DeletaAlunoMock() {
	var aluno model.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	r := SetupDasRotasTeste()
	r.GET("/:nome", controller.Saudacao)
	req, _ := http.NewRequest("GET", "/julia", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `"mgs": "Seja Bem-vindo(a) Julia ao gin API Rest"`
	var respostaBody string
	json.Unmarshal(resposta.Body.Bytes(), respostaBody)
	assert.Equal(t, mockDaResposta, respostaBody)
}

func TestListaTodosAlunosHandler(t *testing.T) {
	database.ConectDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock() // deleta no final o aluno mock
	r := SetupDasRotasTeste()
	r.GET("/alunos", controller.GetAlunosAll)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaPorCPF(t *testing.T) {
	database.ConectDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasTeste()
	r.GET("/alunos/cpf/:cpf", controller.SearchAlunoCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/785.154.125-21", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasTeste()
	r.GET("/alunos/:id", controller.GetAluno)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock model.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Teste", alunoMock)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectDatabase()
	CriaAlunoMock()
	r := SetupDasRotasTeste()
	r.DELETE("/alunos/:id", controller.DeleteAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasTeste()
	r.PATCH("/alunos/:id", controller.PatchAluno)
	aluno := model.Aluno{Nome: "Teste", RG: "473456789", CPF: "12345678700"}
	valorJson, _ := json.Marshal(aluno)
	url := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado model.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "12345678700", alunoMockAtualizado.CPF)
	assert.Equal(t, "473456789", alunoMockAtualizado.RG)
	assert.Equal(t, "Teste", alunoMockAtualizado.Nome)
}
