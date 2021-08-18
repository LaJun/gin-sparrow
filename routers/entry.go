package routers

import "github.com/gin-gonic/gin"

type Option func (engine *gin.Engine)

var options []Option

func Include(optObjs ...Option){
	options = append(options, optObjs...)
}

func InitRouters() *gin.Engine{
	r := gin.Default()
	for _, opt := range options{
		opt(r)
	}
	return r
}

