package web

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thakkarkaushal/http-service-clean-architecture/services"
)

type Web interface {
	Run()
}

type webGin struct {
	service services.UserService
}

func New(c services.UserService) Web {
	return webGin{service: c}
}

func (w webGin) Run() {
	server := gin.Default()
	server.GET("/users", func(c *gin.Context) {
		users, err := w.service.FindAll()
		if err != nil {
			panic(err)
		}
		if len(users) == 0 {
			c.JSON(200, gin.H{
				"data":   nil,
				"count":  0,
				"limit":  20,
				"offset": 0,
			})
			return
		}
		urlData := c.Request.URL.Query()
		limit := urlData["page"]
		if limit == nil {
			offset := 0
			if len(users) >= 20 {
				c.JSON(200, gin.H{
					"data":   users[offset:20],
					"count":  len(users[offset:20]),
					"limit":  20,
					"offset": offset,
				})
				return
			} else if len(users) < 20 {
				fmt.Println(len(users))
				c.JSON(200, gin.H{
					"data":   users[offset:len(users)],
					"count":  len(users[offset:len(users)]),
					"limit":  20,
					"offset": offset,
				})
				return
			}
		} else {
			page, err := strconv.Atoi(limit[0])
			if err != nil {
				panic(err)
			}
			offset := (page - 1) * 20
			if len(users) >= offset {
				c.JSON(200, gin.H{
					"data":   users[offset : offset+20],
					"count":  len(users[offset : offset+20]),
					"limit":  20,
					"offset": offset,
				})
				return
			} else if len(users) > offset {
				value := len(users) - offset
				c.JSON(200, gin.H{
					"data":   users[offset:(offset + value)],
					"count":  len(users[offset:(offset + value)]),
					"limit":  20,
					"offset": offset,
				})
				return
			} else if len(users) < offset {
				c.JSON(200, gin.H{
					"data":   nil,
					"count":  nil,
					"limit":  20,
					"offset": nil,
				})
				return
			}
		}
	})

	server.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		userid, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		user, err := w.service.FindByID(userid)
		c.JSON(200, gin.H{
			"data" : user,
		})
	})

	server.Run()
}
