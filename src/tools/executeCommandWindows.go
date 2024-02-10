package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

func ExecuteCommandWindowsNoAnsi(c echo.Context, command string) {
	cmd := exec.Command(`cmd`, `/C`, command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	peo := cmd.Run()
	if peo != nil {
		fmt.Println(peo)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	json.NewEncoder(c.Response()).Encode(stdout.String() + stderr.String())

}

func ExecuteCommandWindowsColor(c echo.Context, command string) {
	cmd := exec.Command(`cmd`, `/C`, command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	peo := cmd.Run()
	if peo != nil {
		fmt.Println(peo)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	json.NewEncoder(c.Response()).Encode(stdout.String() + stderr.String())
}
