package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {

	//当用户使用 http://localhost:8080/Hello-web/username url时，处理的命令
	mx.HandleFunc("/Hello-web/{id}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		id := vars["id"]
		//Content中的Data
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"当前测试文件： Hello-web/" + id})

	}

}
