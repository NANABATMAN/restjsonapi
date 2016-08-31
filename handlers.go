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
		return c.NoContent(500)
	}
	users.RLock()
	users.m[user.ID] = user
	seq++
	users.RUnlock()
	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users.RLock()
	user := users.m[id]
	users.RUnlock()
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.NoContent(500)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user.ID = id
	users.RLock()
	users.m[id] = user
	u := users.m[id]
	users.RUnlock()
	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users.RLock()
	delete(users.m, id)
	users.RUnlock()
	return c.NoContent(http.StatusNoContent)
}
