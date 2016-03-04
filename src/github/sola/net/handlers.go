package net

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github/sola/repository"
	"net/http"
)

func Login(render render.Render, param repository.LoginRequest) {
	render.JSON(200, repository.LoginResponse{
		repository.DefaultResponse{
			"200", "成功",
		},
		repository.LoginResponseData{
			"",
			"",
			"",
			"",
			"",
			"",
			"",
		},
	})
}

func FormTest(w http.ResponseWriter, r *http.Request) {
	result := r.MultipartForm.Value
	fmt.Println(result)
}
