package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/schmorrison/simple-upload/assets"
)

func indexPage(w http.ResponseWriter, r *http.Request) {

	jsRes, err := getStaticFiles(jsFiles)
	if err != nil {
		err = fmt.Errorf("Failed to get JS resources: %s", err)
		fmt.Println(err)
		return
	}

	cssRes, err := getStaticFiles(cssFiles)
	if err != nil {
		err = fmt.Errorf("Failed to get CSS resources: %s", err)
		fmt.Println(err)
		return
	}

	m := map[string]map[string]string{
		"Scripts": jsRes,
		"Styles":  cssRes,
	}

	templatePage := template.New("UploadFilePage")
	t := template.Must(templatePage.Parse(
		fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Simple Upload</title>

		{{ range $key, $value := .Scripts }}
		<script>{{ . }}</script>
		{{ end}}
		{{ range $key, $value := .Styles }}
		<style>{{ . }}</style>
		{{ end }}
	</head>
	
	<body>
		<p>
		This is the most minimal example of Dropzone. The upload in this example
		doesn't work, because there is no actual server to handle the file upload.
		</p>
		
		<form action="/upload" method="post" enctype="multipart/form-data" class="dropzone">
			<div class="fallback">
				<input name="file" type="file" />
				<br><br>
				<input type="submit" value="UPLOAD">
			</div>
		</form>
	</body>
</html>
`, m)))

	if err := t.Execute(w, ""); err != nil {
		err = fmt.Errorf("Failed to execute template: %s", err)
		fmt.Println(err)
		return
	}
}

var jsFiles = map[string]string{
	"dropzoneCSS":    "/res/dropzone.min.css",
	"materializeCSS": "/res/materialize.min.css",
}
var cssFiles = map[string]string{
	"dropzoneJS":    "/res/dropzone.min.js",
	"materializeJS": "/res/materialize.min.js"}

func getStaticFiles(list map[string]string) (map[string]string, error) {
	m := make(map[string]string)

	for k, v := range list {
		body, err := readFile(v)
		if err != nil {
			err = fmt.Errorf("Failed to load CSS file: %s", err)
			fmt.Println(err)
			return m, err
		}

		m[k] = body
	}
	return m, nil
}

func readFile(path string) (string, error) {
	file, err := assets.FS.Open(path)
	if err != nil {
		err = fmt.Errorf("Failed to load file [%s]: %s", path, err)
		fmt.Println(err)
		return "", err
	}

	body, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("Failed to load CSS file [%s]: %s", path, err)
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}
