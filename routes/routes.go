package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juliasilvamoura/gin-api-rest/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // avisar gin as paginas html
	r.Static("/assets","./assets")
	r.GET("/alunos", controller.GetAlunosAll)
	r.GET("/alunos/:id", controller.GetAluno)
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos/cpf/:cpf", controller.SearchAlunoCpf)
	r.POST("/alunos", controller.CreateAluno)
	r.DELETE("/alunos/:id", controller.DeleteAluno)
	r.PATCH("/alunos/:id", controller.PatchAluno)
	r.GET("/index", controller.ExibePaginaIndex)
	r.NoRoute(controller.RotaNaoEncontrada)
	r.Run(":8000")
}
