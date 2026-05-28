package handler

import (
	"crud-gin/internal/dto"
	app_errors "crud-gin/internal/errors"
	"crud-gin/internal/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{service: service}
}

// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with the provided information
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateUserDTO  true  "User data"
// @Success      201   {object}  models.User
// @Failure      400   {object}  app_errors.ErrorResponse
// @Failure      500   {object}  app_errors.ErrorResponse
// @Router       /users [post]
func (handler *UserHandler) Create(c *gin.Context) {
	var data dto.CreateUserDTO

	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": app_errors.ErrBadRequest})
		return
	}

	user, err := handler.service.Create(c.Request.Context(), data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetAll godoc
// @Summary      Get all users
// @Description  Returns a list of all users
// @Tags         Users
// @Produce      json
// @Success      200  {array} dto.UserDTO
// @Failure      500  {object}  app_errors.ErrorResponse
// @Router       /users [get]
func (handler *UserHandler) GetAll(c *gin.Context) {
	users, err := handler.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})

}

// GetByID godoc
// @Summary      Get user by ID
// @Description  Returns a single user by their ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  dto.UserDTO
// @Failure      400  {object}  app_errors.ErrorResponse
// @Failure      404  {object}  app_errors.ErrorResponse
// @Failure      500  {object}  app_errors.ErrorResponse
// @Router       /users/{id} [get]
func (handler *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": app_errors.ErrBadRequest})
		return
	}

	user, err := handler.service.GetById(c.Request.Context(), id)
	if err != nil {
		// FIX: Return the correct HTTP status for not-found, bad-request, and internal errors.
		if errors.Is(err, app_errors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, app_errors.ErrBadRequest) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update godoc
// @Summary      Update user
// @Description  Update an existing user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int                true  "User ID"
// @Param        body  body      dto.UpdateUserDTO  true  "Updated user data"
// @Success      200   {object}  dto.UserDTO
// @Failure      400   {object}  app_errors.ErrorResponse
// @Failure      404   {object}  app_errors.ErrorResponse
// @Failure      500   {object}  app_errors.ErrorResponse
// @Router       /users/{id} [put]
func (handler *UserHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": app_errors.ErrBadRequest})
		return
	}

	var data dto.UpdateUserDTO
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.service.Update(c.Request.Context(), id, data)
	if err != nil {
		// FIX: Ensure Update returns normal HTTP statuses instead of always returning 500 or ignoring errors.
		if errors.Is(err, app_errors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, app_errors.ErrBadRequest) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary      Delete user
// @Description  Delete a user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  dto.UserDTO
// @Failure      400  {object}  app_errors.ErrorResponse
// @Failure      404  {object}  app_errors.ErrorResponse
// @Failure      500  {object}  app_errors.ErrorResponse
// @Router       /users/{id} [delete]
func (handler *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": app_errors.ErrBadRequest})
		return
	}

	user, err := handler.service.Delete(c.Request.Context(), id)
	if err != nil {
		// FIX: Distinguish not-found from internal delete errors.
		if errors.Is(err, app_errors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
