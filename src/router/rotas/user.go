package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUser = []Rota{
	{
		URI: "/",
		Metodo: http.MethodGet,
		Funcao: controllers.Home,
		Auth: false,
	},
	{
		URI: "/about",
		Metodo: http.MethodGet,
		Funcao: controllers.About,
		Auth: false,
	},
}
