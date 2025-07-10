package handler

import (
	_ "embed"
	"net/http"

	"github.com/maihb/sub2clash/common"
	"github.com/maihb/sub2clash/config"
	M "github.com/maihb/sub2clash/model"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func SubHandler(model M.ClashType, template string) func(c *gin.Context) {
	return func(c *gin.Context) {
		query, err := M.ParseSubQuery(c)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		sub, err := common.BuildSub(model, query, template, config.GlobalConfig.CacheExpire, config.GlobalConfig.RequestRetryTimes)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if len(query.Subs) == 1 {
			userInfoHeader, err := common.FetchSubscriptionUserInfo(query.Subs[0], "clash", config.GlobalConfig.RequestRetryTimes)
			if err == nil {
				c.Header("subscription-userinfo", userInfoHeader)
			}
		}

		if query.NodeListMode {
			nodelist := M.NodeList{}
			nodelist.Proxy = sub.Proxy
			marshal, err := yaml.Marshal(nodelist)
			if err != nil {
				c.String(http.StatusInternalServerError, "YAML序列化失败: "+err.Error())
				return
			}
			c.String(http.StatusOK, string(marshal))
			return
		}
		marshal, err := yaml.Marshal(sub)
		if err != nil {
			c.String(http.StatusInternalServerError, "YAML序列化失败: "+err.Error())
			return
		}
		c.String(http.StatusOK, string(marshal))
	}
}
