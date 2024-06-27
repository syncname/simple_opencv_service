package api

import (
	"cvrecognizer/internal/config"
	"cvrecognizer/internal/recognizers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	rec       *recognizers.Recognizer
	routerApi *fiber.App
	cfg       *config.Config
}

func NewApp(cfg *config.Config, routerApi *fiber.App) (*App, error) {

	r, err := recognizers.New(cfg)
	if err != nil {
		return nil, err
	}

	app := &App{
		rec:       r,
		routerApi: routerApi,
		cfg:       cfg,
	}

	app.InitApi()

	return app, nil

}

func (a *App) InitApi() {
	a.routerApi.Post("/facepos", a.GetFacePos)
	a.routerApi.Post("/emotion/onnx", a.GetEmotionONNX)
	a.routerApi.Post("/emotion/caffe", a.GetCaffeEmotion)
	a.routerApi.Post("/age", a.GetAge)
	a.routerApi.Post("/gender", a.GetGender)
	a.routerApi.Post("/full/info", a.GetFullInfo)
}

func (a *App) Start() error {
	return a.routerApi.Listen(fmt.Sprintf("%s:%d", a.cfg.Server.Host, a.cfg.Server.Port))
}
