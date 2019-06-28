package app

import (
	"fmt"
	"net/http"
)

// GetHome - handler for "/"
func (app *App) GetHome(w http.ResponseWriter, r *http.Request) {
	var test string
	logStatus := r.Context().Value("LogStatus")
	fmt.Println(logStatus)
	userInfo := r.Context().Value("UserInfo")
	fmt.Println(userInfo)
	test = fmt.Sprintf("Log Status : %v \nUser Info : %v", logStatus, userInfo)
	fmt.Println(app.AppDI.Sample())
	fmt.Fprintf(w, test)
}
