package routes

//internal/routes/routes.go
import (
	"log/slog"
	handler "url-shortener/internal/handlers"
	"url-shortener/internal/middleware"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(svc *service.URLService, logger *slog.Logger) *gin.Engine {
	r := gin.New()

	urlHandler := handler.NewURLHandler(svc)
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger(logger))

	r.POST("/shorten", urlHandler.CreateShortURL)
	r.GET("/:alias", urlHandler.RedirectToURL)
	r.DELETE("/:alias", urlHandler.DeleteShortURL)
	return r
}
