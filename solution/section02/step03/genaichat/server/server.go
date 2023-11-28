package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gohandson/genai-ja/solution/section02/step03/genaichat"
)

type Server struct {
	mux        *http.ServeMux
	httpserver *http.Server
	bot        *genaichat.Bot
}

func New(addr string, bot *genaichat.Bot) *Server {
	mux := http.NewServeMux()
	s := &Server{
		mux: mux,
		httpserver: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
		bot: bot,
	}
	s.init()
	return s
}

func (s *Server) init() {
	s.mux.HandleFunc("/", s.handleIndex)
}

func (s *Server) ListenAndServe() error {
	return s.httpserver.ListenAndServe()
}

type Response struct {
	Messge string `json:"message"`
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	msg := r.FormValue("msg")

	slog.Info(msg)

	botresp, err := s.bot.Send(ctx, msg)
	if err != nil {
		status := http.StatusInternalServerError
		slog.Error("error", "err", err)
		http.Error(w, http.StatusText(status), status)
		return
	}

	var resp Response
	respReader := strings.NewReader(botresp)
	decoder := json.NewDecoder(respReader)
	if err := decoder.Decode(&resp); err != nil {
		status := http.StatusInternalServerError
		slog.Error("error", "err", err)
		http.Error(w, http.StatusText(status), status)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(resp); err != nil {
		status := http.StatusInternalServerError
		slog.Error("error", "err", err)
		http.Error(w, http.StatusText(status), status)
		return
	}
}
