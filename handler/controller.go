package handler

import (
	"gozzafadillah/handler/response"
	"gozzafadillah/models"
	"gozzafadillah/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FileUpload(c echo.Context) error {
	//upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.ResponseMedia{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.ResponseMedia{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formFile})
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.ResponseMedia{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	return c.JSON(
		http.StatusOK,
		response.ResponseMedia{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}

func RemoteUpload(c echo.Context) error {
	var url models.Url

	//validate the request body
	if err := c.Bind(&url); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.ResponseMedia{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.ResponseMedia{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Error uploading file"},
			})
	}

	return c.JSON(
		http.StatusOK,
		response.ResponseMedia{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}
