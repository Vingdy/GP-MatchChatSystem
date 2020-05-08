package router

import (
	"GP/controller"
	"github.com/gorilla/mux"
	"net/http"
	"GP/constant"
	"GP/utils"
	"fmt"
)

func SetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/",AllowOrigin(TokenCheck(keep))).Methods("GET")
	router.HandleFunc("/api/register", AllowOrigin(controller.Register)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/login", AllowOrigin(controller.Login)).Methods("POST", "OPTIONS")

	router.HandleFunc("/ws/chat", controller.WsMain)

	router.HandleFunc("/api/getuser", AllowOrigin(controller.GetUser)).Methods("GET")
	router.HandleFunc("/api/updateuser", AllowOrigin(controller.UpdateUser)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/updatepassword", AllowOrigin(controller.UpdatePassword)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/banuser", AllowOrigin(controller.BanUser)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/cancelbanuser", AllowOrigin(controller.CancelBanUser)).Methods("POST", "OPTIONS")

	return router
}

func keep(w http.ResponseWriter, r *http.Request) {
	fb := utils.NewFeedBack(w)
	fmt.Println(r.Header)
	token := r.Header.Get("AccessToken")
	fmt.Println(token)
	info, err := utils.GetTokenInfo(token)
	fmt.Println(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fb.FbCode(constant.STATUS_INTERNAL_SERVER_ERROR).FbData(info).Response()
}

func AllowOrigin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	})
}


