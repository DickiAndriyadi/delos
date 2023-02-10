package controller

import (
	"delos/helper"
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/echo"
)

var endpointCounter = make(map[string]map[string]int)
var users = make(map[string]int)
var mutex = &sync.Mutex{}

func Counter(c echo.Context) error {
	// Lock the mutex to safely access the endpointCounter map
	mutex.Lock()
	defer mutex.Unlock()

	counter := endpointCounter
	return c.JSON(http.StatusOK, counter)
}

func CountEndpointHits(c echo.Context) {

	var endpoint string

	getEndpoint := c.Request().URL.Path

	endpoint = getEndpoint
	if len(getEndpoint) > 9 {
		parts := strings.Split(getEndpoint, "/")
		strEndpoint := strings.Join(parts[:len(parts)-1], "/") + "/:id"
		endpoint = strEndpoint
	}

	getIP, _ := helper.LocalIP()

	user := getIP.String()
	method := c.Request().Method
	key := method + " - " + endpoint

	// Lock the mutex to safely update the endpointCounter map
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := endpointCounter[key]; !ok {
		endpointCounter[key] = make(map[string]int)
	}

	if _, ok := users[user]; !ok {
		users[user] = 1
	}

	// endpointCounter[key]["count"] = 1

	for k, value := range users {
		endpointCounter[key][k] += value
	}

}
