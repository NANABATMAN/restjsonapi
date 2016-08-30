package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func createUser(c echo.Context) error {
	user := &User{
		ID: seq,
	}
	if err := c.Bind(user); err != nil {
		return err
	}
	users[user.ID] = user
	seq++
	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user.ID = id
	users[id] = user
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
