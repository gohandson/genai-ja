package server

import (
	"context"
	_ "embed"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/gohandson/genai-ja/solution/section02/step04/genaichat"
)

var (
	//go:embed index.html
	indexhtml string
	indexTmpl = template.Must(template.New("index").Parse(indexhtml))
)

type History struct {
	Bot   string
	User  string
	Reply string
}

type Server struct {
	mux        *http.ServeMux
	httpserver *http.Server
	template   *template.Template
	bot        *genaichat.Bot
	history    []*History
}

func New(addr string, bot *genaichat.Bot) *Server {
	mux := http.NewServeMux()
	s := &Server{
		mux: mux,
		httpserver: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
		template: indexTmpl,
		bot:      bot,
		history: []*History{{
			Bot:   bot.Name,
			User:  "",
			Reply: bot.FirstMessage,
		}},
	}
	s.init()
	return s
}

func (s *Server) init() {
	s.mux.HandleFunc("/", s.handleIndex)
	s.mux.HandleFunc("/bot", s.handleBot)
}

func (s *Server) ListenAndServe() error {
	return s.httpserver.ListenAndServe()
}

func (s *Server) RecentHistory(limit int) []*History {
	if limit > len(s.history) {
		return s.history
	}
	return s.history[len(s.history)-limit:]
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		History []*History
	}{
		History: s.RecentHistory(10),
	}

	if err := s.template.Execute(w, data); err != nil {
		slog.Error("error", "err", err)
		status := http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
		return
	}
}

func (s *Server) handleBot(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	msg := r.FormValue("msg")
	resp, err := s.bot.Send(ctx, msg)
	if err != nil {
		status := http.StatusInternalServerError
		slog.Error("error", "err", err)
		http.Error(w, http.StatusText(status), status)
		return
	}
	s.history = append(s.history, &History{
		Bot:   s.bot.Name,
		User:  msg,
		Reply: resp,
	})

	http.Redirect(w, r, "/", http.StatusFound)
}
