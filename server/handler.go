package server

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"

	"google.golang.org/grpc"
	_ "swag-grpc-crud/docs"
	pb "swag-grpc-crud/proto"
)

// Initialize gRPC client
func InitGRPCClient() pb.UserServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewUserServiceClient(conn)
}

var grpcClient = InitGRPCClient()

// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body pb.CreateUserRequest true "User to create"
// @Success 200 {object} pb.CreateUserResponse
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req pb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	res, err := grpcClient.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, HTTPError{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pb.GetUserResponse
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetUserRequest{Id: id}
	res, err := grpcClient.GetUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusNotFound, HTTPError{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update a user
// @Description Update user details
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body pb.UpdateUserRequest true "User data to update"
// @Success 200 {object} pb.UpdateUserResponse
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	req.Id = id
	res, err := grpcClient.UpdateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, HTTPError{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pb.DeleteUserResponse
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req := &pb.DeleteUserRequest{Id: id}
	res, err := grpcClient.DeleteUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusNotFound, HTTPError{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// HTTPError represents an HTTP error response.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SetupRouter sets up the Gin router with API endpoints.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger UI endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User endpoints
	users := r.Group("/users")
	{
		users.POST("", CreateUser)
		users.GET("/:id", GetUser)
		users.PUT("/:id", UpdateUser)
		users.DELETE("/:id", DeleteUser)
	}

	return r
}

// StartHTTPServer starts an HTTP server on port 8080.
func StartHTTPServer() {
	r := SetupRouter()

	// Log that the server is starting
	log.Println("HTTP server listening on port 8080...")

	// Serve the HTTP server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
