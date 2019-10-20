package main

import (
	"html/template"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
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

	}
}
