package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"

	tools "github.com/drpaij0se/pairat/src/tools"
	"github.com/labstack/echo"
)

func AnsiOn(c echo.Context) error {
	switch runtime.GOOS {
	case "linux", "darwin":
		var inputCommand command
		reqBody, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		json.Unmarshal([]byte(reqBody), &inputCommand)
		if inputCommand.Command == "" {
			json.NewEncoder(c.Response()).Encode("No command provided.")
		} else {
			tools.ExecuteCommandUnixColor(c, inputCommand.Command)
		}
	case "windows":
		var inputCommand command
		reqBody, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		json.Unmarshal([]byte(reqBody), &inputCommand)
		if inputCommand.Command == "" {
			json.NewEncoder(c.Response()).Encode("No command provided.")
		} else {
			tools.ExecuteCommandWindowsColor(c, inputCommand.Command)
		}
	default:
		fmt.Println("no os detected")
	}
	return nil
}
