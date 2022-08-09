package serverProvider

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (srv *Server) SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	//	r.Use(middleware.Logger)

	r.Route("/api", func(api chi.Router) {
		api.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to chunks"))

		})

		////api.Post("/upload-chunk", srv.upload)
		//api.Post("/upload-chunk", srv.uploadV2)
		//api.Post("/upload_chunk_v3", srv.uploadV3)
		//api.Post("/upload-chunk-completed", srv.handleCompletedChunk)
	})

	return r
}
