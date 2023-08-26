package web

import (
	"github.com/cherrypeel/pkg/image"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Refresh 刷新分类
func Refresh(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(405, gin.H{
			"code": 1,
			"msg":  "error",
			"data": "Method Not Allowed",
		})
		return
	}

	err := image.GenerateClassifyJsonFile("./static/images")
	if err != nil {
		c.JSON(500, gin.H{
			"code": 1,
			"msg":  "error",
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})

}

// AllClassify 获取所有分类
func AllClassify(c *gin.Context) {
	allClassify, err := image.GenerateAllClassify()
	if err != nil {
		c.JSON(500, gin.H{
			"code": 1,
			"msg":  "error",
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": allClassify,
	})
}

// RandomImgByReq 根据请求返回一张图片
func RandomImgByReq(c *gin.Context) {
	classify := c.Query("classify")
	imgPath, err := image.RandomImgByReq(classify)

	if c.Request.Method == http.MethodPost {
		if err != nil {
			c.JSON(500, gin.H{
				"code": 1,
				"msg":  "error",
				"data": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": imgPath,
		})
	} else if c.Request.Method == http.MethodGet {
		accept := c.Request.Header.Get("Accept")
		isBrowserRequest := strings.Contains(accept, "text/html")

		if isBrowserRequest {
			c.HTML(http.StatusOK, "random_show.html", gin.H{
				"page":      "随机图片展示",
				"ApiUri":    "/v1/random?classify=其他&",
				"ImagePath": imgPath,
			})
		} else {
			c.Redirect(http.StatusFound, imgPath)
		}
	} else {
		c.JSON(405, gin.H{
			"code": 1,
			"msg":  "error",
			"data": "Method Not Allowed",
		})
		return
	}
}

// RandomExcitement 随机一张图片
func RandomExcitement(c *gin.Context) {
	imgPath, err := image.RandomImg()

	if c.Request.Method == http.MethodPost {
		if err != nil {
			c.JSON(500, gin.H{
				"code": 1,
				"msg":  "error",
				"data": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": imgPath,
		})
	} else if c.Request.Method == http.MethodGet {
		accept := c.Request.Header.Get("Accept")
		isBrowserRequest := strings.Contains(accept, "text/html")

		if isBrowserRequest {
			c.HTML(http.StatusOK, "random_show.html", gin.H{
				"page":      "随机图片展示",
				"ApiUri":    "/v1/randomExcitement?",
				"ImagePath": imgPath,
			})
		} else {
			c.Redirect(http.StatusFound, imgPath)
		}
	} else {
		c.JSON(405, gin.H{
			"code": 1,
			"msg":  "error",
			"data": "Method Not Allowed",
		})
		return
	}
}
