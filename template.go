package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/schmorrison/simple-upload/assets"
)

func indexPage(w http.ResponseWriter, r *http.Request) {

	data, err := assets.FS.Open("basic.min.css")
	data, err := assets.FS.Open("dropzone.min.js")

	templatePage := template.New("UploadFilePage")
	t := template.Must(templatePage.Parse(`<!DOCTYPE html>

	<html>
		<head>
		<meta charset="utf-8">
		<title>Simple Upload</title>

		<script src="/res/dropzone.min.js"></script>
		<link rel="stylesheet" href="/res/dropzone.min.css">
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
	`))

	if err := t.Execute(w, ""); err != nil {
		err = fmt.Errorf("Failed to execute template: %s", err)
		fmt.Println(err)
		return
	}
}

var css = `.dropzone,.dropzone *{box-sizing:border-box}.dropzone{position:relative}.dropzone .dz-preview{position:relative;display:inline-block;width:120px;margin:0.5em}.dropzone .dz-preview .dz-progress{display:block;height:15px;border:1px solid #aaa}.dropzone .dz-preview .dz-progress .dz-upload{display:block;height:100%;width:0;background:green}.dropzone .dz-preview .dz-error-message{color:red;display:none}.dropzone .dz-preview.dz-error .dz-error-message,.dropzone .dz-preview.dz-error .dz-error-mark{display:block}.dropzone .dz-preview.dz-success .dz-success-mark{display:block}.dropzone .dz-preview .dz-error-mark,.dropzone .dz-preview .dz-success-mark{position:absolute;display:none;left:30px;top:30px;width:54px;height:58px;left:50%;margin-left:-27px}`
