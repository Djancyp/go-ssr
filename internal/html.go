package internal

import "html/template"

const templateHTML = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="description" content="Vite + React + TS" />
    <link rel="icon" type="image/svg+xml" href="/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Title}}</title>
    <!--app-head-->
    <style>
    {{ .CSS }}
    </style>
  </head>
  <body>
    <div id="root">{{ .RenderedContent }}</div>
    <script type="module">
      {{ .JS }}
    </script>
    <script id="props">window.APP_PROPS = {{.InitialProps }}</script>
  </body>
</html>
`

func RenderHTML() (*template.Template, error) {
	templ, err := template.New("html").Parse(templateHTML)
	if err != nil {
		return nil, err
	}
	return templ, nil
}
