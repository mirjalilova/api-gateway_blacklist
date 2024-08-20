package api

import (
	"github.com/mirjalilova/api-gateway_blacklist/api/handler"
	middlerware "github.com/mirjalilova/api-gateway_blacklist/api/middleware"
	_ "github.com/mirjalilova/api-gateway_blacklist/docs"
	"golang.org/x/exp/slog"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Black List
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath  /router
// @description					Description for what is this security definition being used
func NewGin(h *handler.HandlerStruct) *gin.Engine {
	router := gin.Default()

	enforcer, err := casbin.NewEnforcer("gateway/casbin/model.conf", "gateway/casbin/policy.csv")
	if err != nil {
		slog.Error("Error while initializing casbin enforcer: %v", err)
	}

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	// url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	admin := router.Group("/admin").Use(middlerware.NewAuth(enforcer))
	{
		admin.POST("/approve/:user_id", h.Approve)
		admin.GET("/hr_list", h.ListHR)
		admin.DELETE("/delete_hr/:hr_id", h.DeleteHR)
	}

	hr := router.Group("/employee").Use(middlerware.NewAuth(enforcer))
	{
		hr.POST("/create", h.CreateEmployee)
		hr.GET("/:employee_id", h.GetEmployee)
		hr.GET("/all", h.GetAllEmployees)
		hr.PUT("/update/:employee_id", h.UpdateEmployee)
		hr.DELETE("/:employee_id", h.DeleteEmployee)
	}

	HR := router.Group("/blacklist").Use(middlerware.NewAuth(enforcer))
	{
		HR.POST("/add", h.AddEmployee)
		HR.GET("/all", h.GetAll)
		HR.DELETE("/remove/:employee_id", h.RemoveEmployee)

	}

	return router
}