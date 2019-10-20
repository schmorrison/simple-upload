package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/schmorrison/simple-upload/assets"
)

func indexPage(w http.ResponseWriter, r *http.Request) {

	staticRes, err := getStaticFiles()
	if err != nil {
		err = fmt.Errorf("Failed to get static resources: %s", err)
		fmt.Println(err)
		return
	}

	templatePage := template.New("UploadFilePage")
	t := template.Must(templatePage.Parse(
		fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Simple Upload</title>

		<script>%s</script>
		<style>%s</style>
	</head>
	
	<body>
		<p>
		This is the most minimal example of Dropzone. The upload in this example
		doesn't work, because there is no actual server to handle the file upload.
		</p>
		
		<form action="/upload" method="post" enctype="multipart/form-data" class="dropzone">
			<div class="fallback">
				<input name="file" type="file" />
			</div>
		</form>
	</body>
</html>
`, staticRes["JS"], staticRes["CSS"])))

	if err := t.Execute(w, ""); err != nil {
		err = fmt.Errorf("Failed to execute template: %s", err)
		fmt.Println(err)
		return
	}
}

func getStaticFiles() (m map[string]string, err error) {
	cssFile, err := assets.FS.Open("/res/basic.min.css")
	if err != nil {
		err = fmt.Errorf("Failed to load CSS file: %s", err)
		fmt.Println(err)
		return
	}

	css, err := ioutil.ReadAll(cssFile)
	if err != nil {
		err = fmt.Errorf("Failed to load CSS file: %s", err)
		fmt.Println(err)
		return
	}

	m["css"] = string(css)

	jsFile, err := assets.FS.Open("/res/dropzone.min.js")
	if err != nil {
		err = fmt.Errorf("Failed to load JS file: %s", err)
		fmt.Println(err)
		return
	}

	js, err := ioutil.ReadAll(jsFile)
	if err != nil {
		err = fmt.Errorf("Failed to read JS file: %s", err)
		fmt.Println(err)
		return
	}

	m["js"] = string(js)

	return
}
