package server

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/maihb/sub2clash/config"
	"github.com/maihb/sub2clash/constant"
	"github.com/maihb/sub2clash/model"
	"github.com/maihb/sub2clash/server/handler"
	"github.com/maihb/sub2clash/server/middleware"

	"github.com/gin-gonic/gin"
)

//go:embed static
var staticFiles embed.FS

func SetRoute(r *gin.Engine) {
	tpl, err := template.ParseFS(staticFiles, "static/*")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	r.SetHTMLTemplate(tpl)

	r.GET(
		"/static/*filepath", func(c *gin.Context) {
			c.FileFromFS("static/"+c.Param("filepath"), http.FS(staticFiles))
		},
	)
	r.GET(
		"/", func(c *gin.Context) {
			version := constant.Version
			c.HTML(
				200, "index.html", gin.H{
					"Version": version,
				},
			)
		},
	)
	r.GET("/clash", middleware.ZapLogger(), handler.SubHandler(model.Clash, config.GlobalConfig.ClashTemplate))
	r.GET("/meta", middleware.ZapLogger(), handler.SubHandler(model.ClashMeta, config.GlobalConfig.MetaTemplate))
	r.GET("/s/:hash", middleware.ZapLogger(), handler.GetRawConfHandler)
	r.POST("/short", middleware.ZapLogger(), handler.GenerateLinkHandler)
	r.PUT("/short", middleware.ZapLogger(), handler.UpdateLinkHandler)
	r.GET("/short", middleware.ZapLogger(), handler.GetRawConfUriHandler)
}
