package transport

import (
	"cmd/main.go/models"
	"cmd/main.go/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
	"time"
)

// MyHandler handles HTTP requests
type MyHandler struct {
	service *service.Service
}

// NewMyHandler creates a new instance of MyHandler
func NewMyHandler(service *service.Service) *MyHandler {
	return &MyHandler{
		service,
	}
}

// InitRoutes initializes the routes
func (h *MyHandler) InitRoutes() *gin.Engine {
	router := gin.Default()

	// Register handlers
	router.POST("/createuser", h.CreateUserHandler)
	router.GET("/users", h.GetUsersListByParamsHandler)
	router.GET("/worklogs", h.GetWorkLogsHandler)
	router.POST("/startwork", h.StartWorkHandler)
	router.POST("/stopwork", h.StopWorkHandler)
	router.PUT("/updateuser", h.UpdateUserHandler)
	router.DELETE("/deleteuser", h.DeleteUserHandler)

	// Register Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.People true "User to create"
// @Success 201 {object} models.People
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /createuser [post]
func (h *MyHandler) CreateUserHandler(c *gin.Context) {
	var user models.People
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsersListByParamsHandler godoc
// @Summary Get list of users
// @Description Get list of users based on filter parameters
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.People true "User filter"
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {array} models.People
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *MyHandler) GetUsersListByParamsHandler(c *gin.Context) {
	var user models.People
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var pagination models.Pagination
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))

	users, err := h.service.GetUsersListByParams(&user, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users list"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetWorkLogsHandler godoc
// @Summary Get work logs
// @Description Get work logs for a user within a specified date range
// @Tags WorkLogs
// @Produce  json
// @Param userID query int true "User ID"
// @Param start query string true "Start date" Format(date-time)
// @Param end query string true "End date" Format(date-time)
// @Success 200 {array} models.WorkLog
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /worklogs [get]
func (h *MyHandler) GetWorkLogsHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	start, err := time.Parse(time.RFC3339, c.Query("start"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	end, err := time.Parse(time.RFC3339, c.Query("end"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	workLogs, err := h.service.GetWorkLogs(userID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get work logs"})
		return
	}

	c.JSON(http.StatusOK, workLogs)
}

// StartWorkHandler godoc
// @Summary Start work
// @Description Start work on a task for a user
// @Tags Work
// @Accept  json
// @Produce  json
// @Param input body object{UserID int `json:"userID"`; TaskID int `json:"taskID"`} true "Work start input"
// @Success 200 {string} string "Work started successfully"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /startwork [post]
func (h *MyHandler) StartWorkHandler(c *gin.Context) {
	var input struct {
		UserID int `json:"userID"`
		TaskID int `json:"taskID"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.StartWork(input.UserID, input.TaskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start work"})
		return
	}

	c.JSON(http.StatusOK, "Work started successfully")
}

// StopWorkHandler godoc
// @Summary Stop work
// @Description Stop work on a task for a user
// @Tags Work
// @Accept  json
// @Produce  json
// @Param input body object{UserID int `json:"userID"`; TaskID int `json:"taskID"`} true "Work stop input"
// @Success 200 {string} string "Work stopped successfully"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stopwork [post]
func (h *MyHandler) StopWorkHandler(c *gin.Context) {
	var input struct {
		UserID int `json:"userID"`
		TaskID int `json:"taskID"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.StopWork(input.UserID, input.TaskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop work"})
		return
	}

	c.JSON(http.StatusOK, "Work stopped successfully")
}

// UpdateUserHandler godoc
// @Summary Update a user
// @Description Update user details with the input payload
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.People true "User to update"
// @Success 200 {object} models.People
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /updateuser [put]
func (h *MyHandler) UpdateUserHandler(c *gin.Context) {
	var user models.People
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUserHandler godoc
// @Summary Delete a user
// @Description Delete a user with the provided details
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.People true "User to delete"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /deleteuser [delete]
func (h *MyHandler) DeleteUserHandler(c *gin.Context) {
	var user models.People
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.DeleteUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
