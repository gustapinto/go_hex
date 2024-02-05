package ping

import (
	"net/http"

	"github.com/gustapinto/go_hex/pkg/httputil"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	httputil.WriteJson(w, r, http.StatusOK, &PongResponse{
		Ping: "pong",
	})
}
