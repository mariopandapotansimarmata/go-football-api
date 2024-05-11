package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/mariopandapotansimarmata/go-api-responsi-mobile/data"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFile("/favicon.ico", "./favicon.ico")

	//List League
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, render.JSON{Data: data.ListLeague})
	})

	//List Teams by ID League
	r.GET("/:id_league", func(c *gin.Context) {
		paramIdLeague := c.Params.ByName("id_league")
		idLeague, err := strconv.Atoi(paramIdLeague)
		if err != nil {
			panic(err)
		}

		if idLeague <= len(data.ListLeague) {
			c.JSON(http.StatusOK, render.JSON{Data: data.ListLeague[idLeague-1].Teams})
		} else {
			c.JSON(http.StatusNotFound, render.JSON{Data: "NotFound"})
		}
	})

	//Detail Team by ID League and ID Team
	r.GET("/:id_league/:id_team", func(c *gin.Context) {
		paramIdLeague := c.Params.ByName("id_league")
		idLeague, err1 := strconv.Atoi(paramIdLeague)
		if err1 != nil {
			c.JSON(http.StatusNotFound, render.JSON{Data: "Missing ID League URL Param"})
			panic(err1)
		}
		paramITteam := c.Params.ByName("id_team")
		idTeam, err2 := strconv.Atoi(paramITteam)
		if err2 != nil {
			c.JSON(http.StatusNotFound, render.JSON{Data: "Missing ID Team URL Param"})
			panic(err2)
		}

		leagueLength := len(data.ListLeague)
		teamLength := len(data.ListLeague[idLeague-1].Teams)

		if idLeague <= leagueLength && idTeam <= teamLength {
			c.JSON(http.StatusOK, render.JSON{Data: data.ListLeague[idLeague-1].Teams[idTeam-1]})
		} else {
			c.JSON(http.StatusNotFound, render.JSON{Data: "NotFound"})
		}
	})

	r.Run()

}
