package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/routes"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain/constants"
	"github.com/axel-andrade/secret-gift-api/internal/infra/bootstrap"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *http.Server
}

// NewServer cria e retorna uma nova instância do servidor Gin com configurações padrão de middleware.
func NewServer(port string) Server {
	r := gin.New()

	// Adiciona o middleware de controle de CORS, permitindo qualquer origem
	r.Use(middlewares.Cors())

	// Adiciona middlewares padrão
	// gzip: Comprime as respostas HTTP com GZIP.
	// cors: Define as configurações padrão de CORS para permitir todas as solicitações de origem cruzada.
	// requestid: Gera um ID de solicitação exclusivo para cada solicitação.
	// definições de cabeçalho de segurança padrão para proteger contra ataques comuns.
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(requestid.New())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://apis.google.com; style-src 'self' 'unsafe-inline'; img-src 'self' data:;")
	})

	// Configura o cache das respostas
	r.Use(middlewares.Cache(time.Minute))

	if os.Getenv("ENV") == constants.PROD_ENV {
		gin.SetMode(gin.ReleaseMode)
	}

	// Cria um novo servidor HTTP com a configuração TLS e o roteador Gin como handler.
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return Server{
		port:   port,
		server: srv,
	}
}

func (s *Server) AddRoutes(dependencies *bootstrap.Dependencies) {
	router := routes.ConfigRoutes(s.server.Handler.(*gin.Engine), dependencies)
	router.SetTrustedProxies([]string{"127.0.0.1"})
}

func (s *Server) Run() {
	log.Printf("Server starting on port %s", s.port)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

func (s *Server) Shutdown() {
	if s.server == nil {
		return
	}

	if err := s.server.Shutdown(nil); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	}

	log.Println("Server shutdown completed")
}
