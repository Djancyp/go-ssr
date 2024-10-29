package luna

import (
	"gobuild/internal"
	"html/template"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type Engine struct {
	Logger zerolog.Logger
	Server *echo.Echo
	Config Config
	Cache  []Cache
}

type Cache struct {
	ID   string
	Path string
	HTML *template.Template
	Body string
	CSS  string
	JS   string
}

type Config struct {
	ENV         string `default:"development"`
	FrontendDir string `default:"frontend"`
	TailwindCSS bool   `default:"false"`
	Routes      []internal.ReactRoute
}

type PropsResponse struct {
	Path string `json:"path"`
}

func New(config Config) (*Engine, error) {
	server := echo.New()
	server.Static("/assets", "frontend/src/assets/")
	server.Use(middleware.CORS())

	server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	server.POST("/props", func(c echo.Context) error {
		// c.bind
		body := PropsResponse{}
		if err := c.Bind(&body); err != nil {
			return err
		}
		to := body.Path
		props := make(map[string]interface{})
		for _, route := range config.Routes {
			if matched, params := internal.MatchPath(route.Path, to); matched {
				if params != nil {
					props[to] = route.Props(params)
				} else {
					props[to] = route.Props()
				}
			}
		}
		return c.JSON(200, props)

	})
	return &Engine{
		Logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
		Server: server,
		Config: config,
	}, nil
}

func (e *Engine) Start(address string) error {
	return e.Server.Start(address)
}

func (e *Engine) GET(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.GET(p, h, m...)
}

func (e *Engine) POST(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.POST(p, h, m...)
}
func (e *Engine) PUT(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.PUT(p, h, m...)
}

func (e *Engine) DELETE(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.DELETE(p, h, m...)
}

func (e *Engine) PATCH(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.PATCH(p, h, m...)
}

func (e *Engine) OPTIONS(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.OPTIONS(p, h, m...)
}

func (e *Engine) HEAD(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.HEAD(p, h, m...)
}

func (e *Engine) CONNECT(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.CONNECT(p, h, m...)
}
func (e *Engine) TRACE(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Server.TRACE(p, h, m...)
}
func (e *Engine) Any(p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return e.Server.Any(p, h, m...)
}
func (e *Engine) Static(prefix, root string) {
	e.Server.Static(prefix, root)
}
func (e *Engine) File(path, file string) {
	e.Server.File(path, file)
}
func (e *Engine) Group(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return e.Server.Group(prefix, m...)
}
func (e *Engine) Use(m ...echo.MiddlewareFunc) {
	e.Server.Use(m...)
}

func (e *Engine) FrontEnd() {

	e.GET("/*", func(c echo.Context) error {
		// make sure this is not a static file
		return e.RenderFrontend(
			RenderConfig{
				Title: "Go + React + SSR",
				Ctx:   c,
			},
		)
	})
}

func (e *Engine) BuildRoutes() {

}
