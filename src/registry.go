package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"os"
	"regexp"
)

type Profile struct {
	Accounts []string `json:"accounts" binding:"required"`
	Address  string   `json:"address" binding:"required"`
	Alias    string
	Id       string
	Links    []string `json:"links" binding:"required"`
	Logo     string   `json:"logo" binding:"required"`
	Name     string   `json:"name" binding:"required"`
	Status   string
}

type LogLevel string

const (
	Info      = LogLevel("Info")
	Error     = LogLevel("Error")
	Fatal     = LogLevel("Fatal")
	Threshold = 10
)

var (
	sha     = "unknown"
	version = "dev"
	date    = "unknown"
	pod     string
	profile Profile
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
	log(Info, traceId(c), "Company: "+company, nil)

	var alpha = regexp.MustCompile(`^[[:alpha:]]+$`).MatchString
	if !alpha(company) {
		var msg = "company name can contain only letters"
		log(Info, traceId(c), msg, nil)
		c.JSON(http.StatusBadRequest, gin.H{
			"trace_id": *traceId(c),
			"status":   "validation",
			"error":    msg,
		})
	} else {
		if company == profile.Alias {
			c.JSON(http.StatusOK, gin.H{
				"accounts": profile.Accounts,
				"address":  profile.Address,
				"alias":    profile.Alias,
				"id":       profile.Id,
				"links":    profile.Links,
				"logo":     profile.Logo,
				"name":     profile.Name,
				"status":   profile.Status,
			})
		} else {
			log(Info, traceId(c), "company not found", nil)
			c.JSON(http.StatusNotFound, gin.H{
				"trace_id": *traceId(c),
				"error":    "company not found",
			})
		}
	}
}

func putCompany(c *gin.Context) {
	var pro Profile
	company := c.Param("company")
	log(Info, traceId(c), "Company: "+company, nil)

	if genericValidation(c) {
		var alpha = regexp.MustCompile(`^[[:alpha:]]+$`).MatchString
		if !alpha(company) {
			var msg = "company name can contain only letters"
			log(Info, traceId(c), msg, nil)
			c.JSON(http.StatusBadRequest, gin.H{
				"trace_id": *traceId(c),
				"status":   "validation",
				"error":    msg,
			})
		} else {
			c.Header("Content-Type", "application/json; charset=utf-8")
			err := c.BindJSON(&pro)
			if err != nil {
				log(Info, traceId(c), "Invalid payload", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"trace_id": *traceId(c),
					"status":   "validation",
					"error":    err.Error(),
				})
			} else {
				profile = pro
				profile.Alias = company
				profile.Id = "07652979"
				profile.Status = "active"
				var sim = 0
				if sim <= Threshold {
					log(Info, traceId(c), "profile created", nil)
					c.JSON(http.StatusCreated, gin.H{
						"trace_id":   *traceId(c),
						"alias":      company,
						"similarity": 0,
						"message":    "profile created",
					})
				} else {
					var msg = "logo similarity has exceeded threshold"
					log(Info, traceId(c), msg, nil)
					c.JSON(http.StatusBadRequest, gin.H{
						"trace_id":   *traceId(c),
						"alias":      company,
						"similarity": 64,
						"error":      msg,
					})
				}
			}
		}
	}
}

func genericValidation(c *gin.Context) bool {
	var cType string

	if c.Request.Header["Content-Type"] != nil {
		cType = c.Request.Header["Content-Type"][0]
	}

	if cType != "application/json" {
		var msg = "Content type '" + cType + "' is not supported"
		log(Info, traceId(c), msg, nil)
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"trace_id": *traceId(c),
			"status":   "validation",
			"error":    msg,
		})
		return false
	}
	return true
}

func traceId(c *gin.Context) *string {
	var h = c.Request.Header["X-Request-Id"]
	var id string
	if h != nil {
		id = h[0]
	} else {
		id = "n/a"
	}
	return &id
}

func log(level LogLevel, traceId *string, message string, logErr error) {

	if logErr != nil {
		switch level {
		case Info:
			glog.Infof("%s: %v", message, logErr)
		case Error:
			glog.Errorf("%s: %v", message, logErr)
		case Fatal:
			glog.Fatalf("%s: %v", message, logErr)
		}
	} else {
		switch level {
		case Info:
			glog.Infof("%s", message)
		case Error:
			glog.Errorf("%s", message)
		case Fatal:
			glog.Fatalf("%s", message)
		}
	}
}
