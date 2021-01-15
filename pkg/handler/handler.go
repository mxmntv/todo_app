package handler

import "github.com/gin-gonic/gin"

/*Handler -*/
type Handler struct {
}

/*InitRoutes - */
func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getLists)
			lists.GET("/:id", h.getList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getItems)
			}
		}
	}

	return router
}
