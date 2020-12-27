package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"os"
	"regexp"
	//	"strconv"
)

var (
	sha     = "unknown"
	version = "dev"
	date    = "unknown"
	pod     string
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", statusOk)
	r.GET("/version", getVersion)
	g := r.Group("/company")
	{
		g.GET("/:company", getCompany)
		g.PUT("/:company", putCompany)
	}
	return r
}

func main() {
	var port string
	flag.StringVar(&port, "port", "5000", "server listening port")
	flag.Parse()

	pod, _ = os.Hostname()

	router := setupRouter()
	router.Run(":" + port)
}

func getVersion(c *gin.Context) {
	body := gin.H{
		"sha":      sha,
		"version":  version,
		"date":     date,
		"hostname": pod,
		"ginmode":  gin.Mode(),
		"lang":     "golang",
	}
	c.JSON(http.StatusOK, body)
}

func statusOk(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getCompany(c *gin.Context) {
	company := c.Param("company")
	glog.Info("getCompany: " + company)

	var alpha = regexp.MustCompile(`^[[:alpha:]]+$`).MatchString
	if !alpha(company) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company name can contain only letters"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":       "07652979",
			"alias":    "tineye",
			"name":     "TinEye Ltd",
			"address":  "223 Queen Street East, Toronto, Ontario, Canada M5A 1S2",
			"accounts": []string{"NL91ABNA0417164300"},
			"links": []string{
				"https://services.tineye.com/MatchEngine",
				"https://services.tineye.com/developers/matchengine/what_is_matchengine",
			},
			"status": "active",
			"logo":   "iVBORw0KGgoAAAANSUhEU...AASUVORK5CYII=",
		})
	}
	glog.Infof("getCompany: %d", http.StatusOK)
}

func putCompany(c *gin.Context) {
	company := c.Param("company")
	glog.Info("getCompany: " + company)

	var alpha = regexp.MustCompile(`^[[:alpha:]]+$`).MatchString
	if !alpha(company) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company name can contain only letters"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"alias":      company,
			"similarity": 64,
			"error":      "logo similarity has exceeded threshold",
		})
	}
}
