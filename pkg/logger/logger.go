package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maxexq/parksoi-shop/pkg/utils"
)

type ILogger interface {
	Print() ILogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(c *fiber.Ctx)
}

type Logger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitLogger(c *fiber.Ctx, res any) ILogger {
	log := &Logger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		StatusCode: c.Response().StatusCode(),
		Path:       c.Path(),
	}

	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(c)
	return log
}

func (l *Logger) Print() ILogger {
	utils.Debug(l)
	return l
}

func (l *Logger) Save() {

	data := utils.Output(l)

	fileName := fmt.Sprintf("./assets/logs/parksoil.%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("open file failed: %v", err)

	}

	defer file.Close()

	file.WriteString(string(data) + "\n")
}

func (l *Logger) SetQuery(c *fiber.Ctx) {

}

func (l *Logger) SetBody(c *fiber.Ctx) {
	var body any

	if err := c.BodyParser(&body); err != nil {
		log.Printf("parse body failed: %v", err)
	}

	switch l.Path {
	case "api/v1/signup":
		l.Body = "test signup"
	default:
		l.Body = body
	}

}

func (l *Logger) SetResponse(c *fiber.Ctx) {

}
