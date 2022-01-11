package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wa1kman999/goblog/pkg/common/constants"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type bufferedWriter struct {
	gin.ResponseWriter
	out    *bufio.Writer
	Buffer bytes.Buffer
}

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = os.Stdout

	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 生产环境下，打印req body 但是不打印 res body
	if string(constants.ReleaseMode) == os.Getenv("GIN_MODE") {
		return func(c *gin.Context) {

			// 开始时间
			start := time.Now()

			// 读取req body，如果 request header 中 Content-Type 为 application/json
			reqBody := make(map[string]interface{})
			if c.Request.Header.Get("Content-Type") == "application/json" {
				bodyCopy := new(bytes.Buffer)
				// Read the whole body
				_, err := io.Copy(bodyCopy, c.Request.Body)
				if err != nil {
					logrus.Error(err)
				} else {
					bodyData := bodyCopy.Bytes()
					c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
					if len(bodyData) > 0 {
						err = json.Unmarshal(bodyData, &reqBody)
						if err != nil {
							logrus.Errorf("%s", string(bodyData))
							return
						}
					}
				}
			}

			// 处理请求
			c.Next()

			// 结束时间
			end := time.Now()

			// 执行时间
			rt := end.Sub(start)

			// 请求方式
			method := c.Request.Method

			// 请求路由
			url := c.Request.URL.String()

			// 状态码
			code := c.Writer.Status()

			// 请求IP
			client := c.ClientIP()

			// 协议
			proto := c.Request.Proto

			// 日志格式
			logger.WithFields(logrus.Fields{
				"type":            "http",
				"code":            code,
				"rt":              rt,
				"client":          client,
				"method":          method,
				"url":             url,
				"proto":           proto,
				"body_bytes_sent": c.Request.ContentLength,
			}).Info()

			if len(reqBody) > 0 {
				logger.WithFields(logrus.Fields{
					"prefix": "req_body",
				}).Info(reqBody)
			}
		}
	}

	// 测试环境，打印req body和res body
	if string(constants.TestMode) == os.Getenv("GIN_MODE") {
		return func(c *gin.Context) {

			// 忽略掉swagger路由
			if isSwaggerRoute(c) {
				return
			}

			// 开始时间
			start := time.Now()

			// 读取res body
			w := bufio.NewWriter(c.Writer)
			defer func() {
				err := w.Flush()
				if err != nil {
					logrus.Error(err)
				}
			}()
			resBuf := bytes.Buffer{}
			newWriter := &bufferedWriter{c.Writer, w, resBuf}
			c.Writer = newWriter

			// 处理请求
			c.Next()

			// 结束时间
			end := time.Now()

			// 执行时间
			rt := end.Sub(start)

			// 请求方式
			method := c.Request.Method

			// 请求路由
			url := c.Request.URL.String()

			// 状态码
			code := c.Writer.Status()

			// 请求IP
			client := c.ClientIP()

			// 协议
			proto := c.Request.Proto

			// 读取req body
			reqBuf, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				logrus.Error(err)
				return
			}
			reqBody := make(map[string]interface{})
			if len(reqBuf) > 0 {
				_ = json.Unmarshal(reqBuf, &reqBody)
			}

			resBody := make(map[string]interface{})
			resBuff := newWriter.Buffer.Bytes()
			if len(resBuff) > 0 {
				_ = json.Unmarshal(resBuff, &resBody)
			}

			// 日志格式
			logger.WithFields(logrus.Fields{
				"type":            "http",
				"code":            code,
				"rt":              rt,
				"client":          client,
				"method":          method,
				"url":             url,
				"proto":           proto,
				"req_body":        reqBody,
				"res_body":        resBody,
				"body_bytes_sent": c.Request.ContentLength,
			}).Info()

		}
	}

	// 开发环境下，打印req body和res body
	return func(c *gin.Context) {

		// 忽略掉swagger路由
		if isSwaggerRoute(c) {
			return
		}

		logger.SetFormatter(&prefixed.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// 读取req body，如果 request header 中 Content-Type 为 application/json
		reqBody := make(map[string]interface{})
		if c.Request.Header.Get("Content-Type") == "application/json" {
			bodyCopy := new(bytes.Buffer)
			// Read the whole body
			_, err := io.Copy(bodyCopy, c.Request.Body)
			if err != nil {
				logrus.Error(err)
			} else {
				bodyData := bodyCopy.Bytes()
				c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
				if len(bodyData) > 0 {
					err = json.Unmarshal(bodyData, &reqBody)
					if err != nil {
						logrus.Error(err)
						return
					}
				}
			}
		}

		// 读取res body
		w := bufio.NewWriter(c.Writer)
		defer func() {
			err := w.Flush()
			if err != nil {
				logrus.Error(err)
			}
		}()
		resBuf := bytes.Buffer{}
		newWriter := &bufferedWriter{c.Writer, w, resBuf}
		c.Writer = newWriter

		// 处理请求
		c.Next()

		// 解析res body
		resBody := make(map[string]interface{})
		resBuff := newWriter.Buffer.Bytes()
		if len(resBuff) > 0 {
			_ = json.Unmarshal(resBuff, &resBody)
		}

		// 日志格式
		logger.WithFields(logrus.Fields{
			"prefix":          "req_body",
			"body_bytes_sent": c.Request.ContentLength,
		}).Info(reqBody)
		logger.WithFields(logrus.Fields{
			"prefix": "res_body",
		}).Info(resBody)
	}
}

// isSwaggerRoute 判断是否是swagger路由，如果是则取消对
func isSwaggerRoute(c *gin.Context) bool {
	// 忽略掉swagger路由
	return strings.Contains(c.Request.RequestURI, "swagger")
}
