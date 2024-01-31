package ping

import (
	"github.com/gustapinto/go_hex/pkg/httputil"
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	httputil.WriteJson(w, r, http.StatusOK, &PongResponse{
		Ping: "pong",
	})
}
