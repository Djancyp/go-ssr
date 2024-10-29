package luna

import (
	"fmt"
	"gobuild/internal"
	"html/template"
	"runtime"

	"github.com/labstack/echo/v4"
)

type RenderConfig struct {
	Title  string
	Ctx    echo.Context
	Routes []internal.ReactRoute
}

func (e *Engine) RenderFrontend(config RenderConfig) error {
	pc, _, _, _ := runtime.Caller(1)
	ID := fmt.Sprint(pc)

	task := internal.JobRunner{
		ID:     ID,
		Logger: e.Logger,
		Path:   config.Ctx.Request().URL.Path,
		Routes: e.Config.Routes,
	}

	// check if path is exist in cache
	//
	for _, c := range e.Cache {
		if c.Path == task.Path {
			return c.HTML.Execute(config.Ctx.Response().Writer,
				map[string]interface{}{
					"RenderedContent": template.HTML(c.Body),
					"JS":              template.JS(c.JS),
					"CSS":             template.CSS(c.CSS),
					"Title":           config.Title,
				})
		}
	}

	html, body, css, js, err := task.Start()
	if err != nil {
		e.Logger.Error().Err(err).Msg("failed to start render task")

	}

	e.Cache = append(e.Cache, Cache{
		ID:   ID,
		Path: task.Path,
		HTML: html,
		Body: body,
		CSS:  css,
		JS:   js,
	})

	return html.Execute(config.Ctx.Response().Writer,
		map[string]interface{}{
			"RenderedContent": template.HTML(body),
			"JS":              template.JS(js),
			"CSS":             template.CSS(css),
			"Title":           config.Title,
		})

}
