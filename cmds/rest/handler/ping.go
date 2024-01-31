package handler

type Ping struct {}

func (Ping) Pong(w http.ResponseWriter, r *http.Request) {
	httputil.WriteJson(w, r, http.StatusOK, &response.Ping{
		Ping: "pong",
	})
}