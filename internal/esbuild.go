package internal

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/buke/quickjs-go"
	"github.com/rs/zerolog"

	esbuild "github.com/evanw/esbuild/pkg/api"
)

var textEncoderPolyfill = `function TextEncoder(){}TextEncoder.prototype.encode=function(string){var octets=[];var length=string.length;var i=0;while(i<length){var codePoint=string.codePointAt(i);var c=0;var bits=0;if(codePoint<=0x0000007F){c=0;bits=0x00}else if(codePoint<=0x000007FF){c=6;bits=0xC0}else if(codePoint<=0x0000FFFF){c=12;bits=0xE0}else if(codePoint<=0x001FFFFF){c=18;bits=0xF0}octets.push(bits|(codePoint>>c));c-=6;while(c>=0){octets.push(0x80|((codePoint>>c)&0x3F));c-=6}i+=codePoint>=0x10000?2:1}return octets};function TextDecoder(){}TextDecoder.prototype.decode=function(octets){var string="";var i=0;while(i<octets.length){var octet=octets[i];var bytesNeeded=0;var codePoint=0;if(octet<=0x7F){bytesNeeded=0;codePoint=octet&0xFF}else if(octet<=0xDF){bytesNeeded=1;codePoint=octet&0x1F}else if(octet<=0xEF){bytesNeeded=2;codePoint=octet&0x0F}else if(octet<=0xF4){bytesNeeded=3;codePoint=octet&0x07}if(octets.length-i-bytesNeeded>0){var k=0;while(k<bytesNeeded){octet=octets[i+k+1];codePoint=(codePoint<<6)|(octet&0x3F);k+=1}}else{codePoint=0xFFFD;bytesNeeded=octets.length-i}string+=String.fromCodePoint(codePoint);i+=bytesNeeded+1}return string};`
var processPolyfill = `var process = {env: {NODE_ENV: "development"}};`
var consolePolyfill = `var console = {log: function(){}};`

type JobRunner struct {
	ID                 string
	Logger             zerolog.Logger
	Path               string
	serverRenderResult chan serverRenderResult
	clientRenderResult chan clientRenderResult
}

type serverRenderResult struct {
	html string
	css  string
	err  error
}

type clientRenderResult struct {
	js  string
	err error
}

func (j *JobRunner) Start() (html *template.Template, body string, css string, js string, err error) {
	j.serverRenderResult = make(chan serverRenderResult)
	j.clientRenderResult = make(chan clientRenderResult)

	serverJS, err := BuildServer()

	if err != nil {
		j.Logger.Error().Err(err).Msg("failed to build server")
		panic(err)
	}
	client, err := BuildClient()
	if err != nil {
		j.Logger.Error().Err(err).Msg("failed to build client")
		panic(err)
	}
	j.Logger.Info().Msg("server built")

	ht, err := RenderHTML()
	if err != nil {
		j.Logger.Error().Err(err).Msg("failed to render html")
		return nil, "", "", "", err
	}
	serverBody, err := renderReactToHTMLNew(serverJS.JS,j.Path)
	if err != nil {
		j.Logger.Error().Err(err).Msg("failed to render react to html")
		return nil, "", "", "", err
	}
	return ht, serverBody, client.CSS, client.JS, nil
}

type BuildResult struct {
	JS  string
	CSS string
}

var Loader = map[string]esbuild.Loader{
	".png":   esbuild.LoaderFile,
	".svg":   esbuild.LoaderFile,
	".jpg":   esbuild.LoaderFile,
	".jpeg":  esbuild.LoaderFile,
	".gif":   esbuild.LoaderFile,
	".bmp":   esbuild.LoaderFile,
	".woff2": esbuild.LoaderFile,
	".woff":  esbuild.LoaderFile,
	".ttf":   esbuild.LoaderFile,
	".eot":   esbuild.LoaderFile,
	".mp4":   esbuild.LoaderFile,
	".webm":  esbuild.LoaderFile,
	".wav":   esbuild.LoaderFile,
	".mp3":   esbuild.LoaderFile,
	".m4a":   esbuild.LoaderFile,
	".aac":   esbuild.LoaderFile,
	".oga":   esbuild.LoaderFile,
	".json":  esbuild.LoaderFile,
	".txt":   esbuild.LoaderFile,
	".xml":   esbuild.LoaderFile,
	".csv":   esbuild.LoaderFile,
	".ts":    esbuild.LoaderTS,
	".tsx":   esbuild.LoaderTSX,
	".js":    esbuild.LoaderJS,
	".jsx":   esbuild.LoaderJSX,
	".css":   esbuild.LoaderCSS,
	".html":  esbuild.LoaderFile,
}

func BuildClient() (BuildResult, error) {
	opt := esbuild.Build(esbuild.BuildOptions{
		EntryPoints:       []string{"./frontend/src/entry-client.tsx"},
		Outdir:            "/",
		Bundle:            true,
		Write:             false,
		Metafile:          false,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Loader:            Loader,
	})

	if len(opt.Errors) > 0 {
		return BuildResult{},
			fmt.Errorf("build error: %v", opt.Errors[0].Text)
	}

	result := BuildResult{}
	for _, file := range opt.OutputFiles {

		if strings.HasSuffix(file.Path, ".css") {
			result.CSS = string(file.Contents)
		} else if strings.HasSuffix(file.Path, ".js") {
			result.JS = string(file.Contents)
		}
	}

	return result, nil
}

func BuildServer() (BuildResult, error) {
	opt := esbuild.Build(esbuild.BuildOptions{
		EntryPoints:       []string{"frontend/src/entry-server.tsx"},
		Bundle:            true,
		Write:             false,
		Outdir:            "/",
		Format:            esbuild.FormatESModule,
		Platform:          esbuild.PlatformBrowser,
		Target:            esbuild.ES2015,
		AssetNames:        fmt.Sprintf("%s/[name]", strings.TrimPrefix("/assets/", "/")),
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Loader:            Loader,
		Banner: map[string]string{
			"js": textEncoderPolyfill + processPolyfill + consolePolyfill,
		},
	})

	if len(opt.Errors) > 0 {
		return BuildResult{},
			fmt.Errorf("build error: %v", opt.Errors[0].Text)
	}

	result := BuildResult{}
	for _, file := range opt.OutputFiles {

		if strings.HasSuffix(file.Path, ".css") {
			result.CSS = string(file.Contents)
		} else if strings.HasSuffix(file.Path, ".js") {
			result.JS = string(file.Contents)
		}
	}

	return result, nil
}

func renderReactToHTMLNew(js string,path string) (string, error) {
	// save the js to a OutputFiles
	//
	// os.WriteFile("server.mjs", []byte(js), 0644)
	// Create a new QuickJS runtime with module support
	rt := quickjs.NewRuntime(quickjs.WithModuleImport(true))

	defer rt.Close()

	ctx := rt.NewContext()
	defer ctx.Close()

	r1, err := ctx.LoadModule(js, "server")
	if err != nil {
		return "", fmt.Errorf("failed to evaluate module: %w", err)
	}
	defer r1.Free()

  // add path to the render function
	val, err := ctx.Eval(`
      import {render} from 'server';
     globalThis.result = render().html;
  `)

	val.Free()
	if err != nil {
		return "", fmt.Errorf("failed to evaluate module: %w", err)
	}

	result := ctx.Globals().Get("result").String()
	return result, nil
}
