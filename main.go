package main

import (
	"gin/routers"
	"net/http"
	"time"
)

func main(){
	router := routers.InitRouters()
	routers.Include(routers.LoadApiRouters)
	s := &http.Server{
		Addr: "9100",
		Handler: router,
		ReadTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}