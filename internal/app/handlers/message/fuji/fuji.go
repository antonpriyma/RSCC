package fuji

import (
	"github.com/antonpriyma/RSCC/internal/app/handlers"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	tele "gopkg.in/telebot.v3"
	"io/ioutil"
	"math/rand"
	"path"
	"strings"
	"time"
)

type Config struct {
	// path to folder with Wilee photos
	PhotosPath string `mapstructure:"photos_path"`
	// description for wilee photo
	PhotoDesc string `mapstructure:"photo_desc"`

	// words which trigger photo sending
	TriggerWords []string `mapstructure:"trigger_words"`

	// limit per minute
	Limit float64 `mapstructure:"limit"`
}

type fujiHandler struct {
	Config
	Logger  log.Logger
	Limiter *rate.Limiter
}

func New(c Config, logger log.Logger, limiter *rate.Limiter) handlers.Handler {
	return fujiHandler{
		Config:  c,
		Logger:  logger,
		Limiter: limiter,
	}
}

func (h fujiHandler) Handle(ctx tele.Context) error {
	for _, word := range h.TriggerWords {
		if strings.Contains(strings.ToLower(ctx.Message().Text), word) {
			// TODO: log with ctx
			if h.Limiter.Allow() {
				err := h.sendWileePhoto(ctx)
				if err != nil {
					return errors.Wrap(err, "failed to send Wilee photo")
				}

				break
			}
		}
	}

	return nil
}

func (h fujiHandler) sendWileePhoto(ctx tele.Context) error {
	rand.Seed(time.Now().UnixNano())
	files, err := ioutil.ReadDir(h.PhotosPath)
	if err != nil {
		return errors.Wrap(err, "failed to read photos dir")
	}

	photo := &tele.Photo{
		File:    tele.FromDisk(path.Join(h.PhotosPath, files[rand.Intn(len(files))].Name())),
		Caption: h.PhotoDesc,
	}

	err = ctx.Reply(photo)
	return errors.Wrap(err, "failed to send photo")
}
