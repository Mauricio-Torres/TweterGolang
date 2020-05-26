package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Mauricio-Torres/TweterGolang/middware"
	"github.com/Mauricio-Torres/TweterGolang/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejador seteo el puerto y se pone en modo escucha a ese puerto creado */
func Manejador() {
	router := mux.NewRouter()

	// adicion de rutas
	router.HandleFunc("/registro", middware.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middware.ChequeoDB(routers.Loguer)).Methods("POST")
	router.HandleFunc("/verperfil", middware.ChequeoDB(middware.ValidatorJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middware.ChequeoDB(middware.ValidatorJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middware.ChequeoDB(middware.ValidatorJWT(routers.GrabarTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middware.ChequeoDB(middware.ValidatorJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middware.ChequeoDB(middware.ValidatorJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/relacionarTweets", middware.ChequeoDB(middware.ValidatorJWT(routers.RelacionandoTweets))).Methods("POST")
	router.HandleFunc("/eliminarRelacionTweets", middware.ChequeoDB(middware.ValidatorJWT(routers.EliminarRelacionTweeter))).Methods("DELETE")
	router.HandleFunc("/consultarRelacionTweets", middware.ChequeoDB(middware.ValidatorJWT(routers.ConsultarRelacionEntreUsuarios))).Methods("GET")

	// manejo de imagenes
	router.HandleFunc("/subirAvatar", middware.ChequeoDB(middware.ValidatorJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middware.ChequeoDB(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middware.ChequeoDB(middware.ValidatorJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middware.ChequeoDB(routers.ObtenerBanner)).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
