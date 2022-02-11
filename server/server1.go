package server
import (
	"source"
    "net/http"
    "github.com/gorilla/mux"
)

type Server struct {
	mux.Router
}

func Service_one() *Server {
	myRouter := mux.NewRouter().StrictSlash(true)
	router.Route("/api", func(api chi.Router) {
		api.Get("/", func(responseWriter http.ResponseWriter, request *http.Request) {
			_, err := responseWriter.Write([]byte("server server server server deployed 1st time server server server server deployed 1st time server server server server deployed 1st time server server server server deployed 1st time"))
			if err != nil {
				return
			}
		})
		api.Post("/text", handler.CountFrequency)
	})
	return &Server{router}
}

func Service_two() *Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Post("/call_service", handler.CallService)
	})
	return &Server{router}
}
func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}