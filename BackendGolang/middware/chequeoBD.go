package middware

import (
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
)

/*ChequeoDB permite conocer el estado de la Base de datos */
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion Perdida ", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
