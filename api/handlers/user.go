package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/javiorfo/go-microservice-lib/auditory"
	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-lib/response"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-lib/tracing"
	"github.com/javiorfo/go-microservice-users/api/request"
	"github.com/javiorfo/go-microservice-users/domain/model"
	"github.com/javiorfo/go-microservice-users/domain/service"
	"github.com/javiorfo/steams"
)

const (
	USER_FIND_ERROR   = "USER_FIND_ERROR"
	USER_CREATE_ERROR = "DUMMY_CREATE_ERROR"
	USER_LOGIN_ERROR  = "DUMMY_LOGIN_ERROR"
)

// @Summary		Find a user by ID
// @Description	Get user details by ID
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	model.User
// @Failure		400	{object}	response.restResponseError	"Invalid ID"
// @Failure		404	{object}	response.restResponseError	"Internal Error"
// @Router			/{id} [get]
// @Security		BearerAuth
func FindById(us service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := c.Params("id")
		log.Infof("%s Find user by ID: %v", tracing.LogTraceAndSpan(c), param)

		id, err := strconv.Atoi(param)
		if err != nil {
			log.Error("Invalid ID")
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, USER_FIND_ERROR, "Invalid ID"))
		}

		if user, err := us.FindById(uint(id)); err != nil {
			return c.Status(http.StatusNotFound).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, USER_FIND_ERROR, err.Error()))
		} else {
			return c.JSON(user)
		}
	}
}

// @Summary		List all users
// @Description	Get a list of users with pagination
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			page		query		int											false	"Page number"
// @Param			size		query		int											false	"Size per page"
// @Param			sortBy		query		string										false	"Sort by field"
// @Param			sortOrder	query		string										false	"Sort order (asc or desc)"
// @Success		200			{object}	response.RestResponsePagination[model.User]	"Paginated list of users"
// @Failure		400			{object}	response.restResponseError					"Invalid query parameters"
// @Failure		500			{object}	response.restResponseError					"Internal server error"
// @Router			/ [get]
// @Security		BearerAuth
func FindAll(us service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := c.Query("page", "1")
		s := c.Query("size", "10")
		sb := c.Query("sortBy", "id")
		so := c.Query("sortOrder", "asc")

		log.Infof("%s Listing users...", tracing.LogTraceAndSpan(c))
		log.Infof("%s page %s, size %s, sortBy %s, sortOrder %s ", tracing.LogTraceAndSpan(c), p, s, sb, so)

		page, err := pagination.ValidateAndGetPage(p, s, sb, so)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, USER_FIND_ERROR, err.Error()))
		}

		users, err := us.FindAll(*page)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(response.InternalServerError(c, err.Error()))
		}

		return c.JSON(response.RestResponsePagination[model.User]{
			Pagination: pagination.Paginator(*page, len(users)),
			Elements:   users,
		})
	}
}

// @Summary		Create a new user
// @Description	Create a new user with the provided information
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			user	body		request.User	true	"User information"
// @Success		201		{object}	interface{}
// @Failure		400		{object}	response.restResponseError	"Invalid request body or validation errors"
// @Failure		500		{object}	response.restResponseError	"Internal server error"
// @Router			/ [post]
// @Security		BearerAuth
func Create(us service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRequest := new(request.User)

		if err := c.BodyParser(userRequest); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, USER_CREATE_ERROR, "Invalid request body"))
		}
		validate := validator.New()
		if err := validate.Struct(userRequest); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, USER_CREATE_ERROR, validationErrors.Error()))
		}

		log.Infof("%s Received user: %+v", tracing.LogTraceAndSpan(c), userRequest)

		roles := steams.Mapping(steams.OfSlice(userRequest.Permission.Roles), func(v string) model.Role {
			return model.Role{Name: v}
		}).Collect()

		user := model.User{
			Username:   userRequest.Username,
			Email:      userRequest.Email,
			Permission: model.Permission{Name: userRequest.Permission.Name, Roles: roles},
			Temporary:  true,
			Auditable: auditory.Auditable{
				CreatedBy: security.GetTokenUsername(c),
			},
		}
		password, err := us.Create(&user)

		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(response.InternalServerError(c, err.Error()))
		}

        return c.Status(fiber.StatusCreated).JSON(fiber.Map{"username": user.Username, "password": password})
	}
}

// @Summary		Login user
// @Description	Login a user and return a JWT token
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			user	body		request.LoginUser	true	"Username and password"
// @Success		201		{object}	interface{}
// @Failure		400		{object}	response.restResponseError	"Invalid request body or validation errors"
// @Failure		500		{object}	response.restResponseError	"Internal server error"
// @Router			/login [post]
// @Security		BearerAuth
func Login(us service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRequest := new(request.LoginUser)

		if err := c.BodyParser(userRequest); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, USER_LOGIN_ERROR, "Invalid request body"))
		}
		validate := validator.New()
		if err := validate.Struct(userRequest); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, USER_LOGIN_ERROR, validationErrors.Error()))
		}

		log.Infof("%s Received credentials: %+v", tracing.LogTraceAndSpan(c), userRequest)

		token, err := us.Login(userRequest.Username, userRequest.Password)

		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(response.InternalServerError(c, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
	}
}
