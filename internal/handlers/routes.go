package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

var sessionLifeTime = int((10 * time.Hour).Seconds()) // Время жизни сессии в секундах

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		api.POST("/create_service", h.СreateService)        // Создание заявки на услугу
		api.POST("/get_service", h.GetServiceByUUID)        // Получить данные по услуге
		api.POST("/set_status_service", h.SetServiceStatus) // Установить статус
		api.POST("/withdraw_service", h.WithDrawService)    // Отзыв заявки клиентом
		api.POST("/delete_service", h.DeleteService)        // Удаление заявки
	}

	return router
}
