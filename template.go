package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/schmorrison/simple-upload/assets"
)

var funcMap = template.FuncMap{
	// "fileTree": fileTreeHTML,
	"safeJS": func(s string) template.JS {
		return template.JS(s)
	},
	"safeCSS": func(s string) template.CSS {
		return template.CSS(s)
	},
}

// func fileTreeHTML(root FileNode) template.JS {
// 	s := "var treeValues = {"
// 	for name, node := range root.Children {
// 		s += fmt.Sprintf("%s", fileTreeChildren(name, node))
// 	}
// 	s += "};"

// 	return template.JS(s)
// }

// func fileTreeChildren(name string, node FileNode) string {
// 	s := ""
// 	for name, a := range node.Children {
// 		s += fmt.Sprintf(`"%s": {"children": {%s}},`, name, fileTreeChildren(name, a))
// 	}
// 	return s
// }

func indexPage(w http.ResponseWriter, r *http.Request) {
	templatePage := template.New("").Funcs(funcMap).Funcs(sprig.FuncMap())
	t := template.Must(templatePage.Parse(
		fmt.Sprint(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Simple Upload</title>

		<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
		{{ range .Styles }}
		<style>{{ . | safeCSS }}</style>
		{{ end }}
		{{ range .Scripts }}
		<script>{{ . | safeJS }}</script>
		{{ end }}

		<style>
			#droppable:hover {
				background-color: rgba(0,0,0,0.5);
			}
		</style>
		<script>
			document.addEventListener("DOMContentReady", () => {
				let drp = document.getElementById('droppable');
				
				drp.addEventListener('dragover', (e) => {
					e.preventDefault();
					this.addClass('hover');
				});
				drp.addEventListener('dragleave', (e) => {
					e.preventDefault();
					this.removeClass('hover');
				});
			});

			function iframeBack() {
				let ifr = document.getElementById("file-browser");
				let loc = ifr.contentWindow.location.pathname;
				if (loc !== "/files/") {
					ifr.contentWindow.history.back();
				}
			}
		</script>
	</head>
	
	<body>
		<div class="row">
			<div class="col s12 m4 l3">
				<div class="card">
					<div class="card-content">
						<span class="card-title">Upload Files to Remote</span>
						<p>
							Either drag and drop a file onto this card, or click on this card to open the file upload dialog.
						</p>
						<form id="droppable" action="/upload/" method="post" enctype="multipart/form-data" class="dropzone">
							<div class="fallback">
								<input name="file" type="file" />
								<br><br>
								<input type="submit" value="UPLOAD" />
							</div>
						</form>
					</div>
				</div>
			</div>
			<div class="col s12 m8 l9">
				<div class="card">
					<div class="card-content white-text">
						<a class="waves-effect waves-light btn" onclick="iframeBack()">Back</a>
						<span class="card-title">Download Files from Remote</span>
						<p>
							Click on a file below to download from the server.
						</p>
						<iframe id="file-browser" style="width: 100%; height: 100%;" src="/files/"></iframe>
					</div>
				</div>
			</div>
		</div>
	</body>
</html>
`)))

	m := map[string]interface{}{
		"Scripts": jsFiles,  // []string
		"Styles":  cssFiles, // []string
		// "Files":   getFileTree(), // map[string]FileTree
	}

	if err := t.Execute(w, m); err != nil {
		err = fmt.Errorf("Failed to execute template: %s", err)
		fmt.Println(err)
		return
	}
}

// type FileNode struct {
// 	Directory bool
// 	Size      int64
// 	Modified  time.Time
// 	FileMode  uint32
// 	Children  map[string]FileNode
// }

// func newFileNode() FileNode {
// 	return FileNode{
// 		Children: make(map[string]FileNode),
// 	}
// }

// func getFileTree() FileNode {
// 	root := newFileNode()
// 	root.Directory = true
// 	if err := filepath.Walk(".", func(fp string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			err = fmt.Errorf("stumbled on '%s': %s", fp, err)
// 			return err
// 		}

// 		node := newFileNode()
// 		node.Directory = info.IsDir()
// 		node.Size = info.Size()
// 		node.Modified = info.ModTime()
// 		node.FileMode = uint32(info.Mode())

// 		fp = path.Clean(fp)
// 		segments := strings.Split(fp, "\\")

// 		current := root
// 		for _, a := range segments {
// 			if v, ok := current.Children[a]; ok {
// 				current = v
// 			} else {
// 				current.Children[a] = node
// 			}
// 		}

// 		return nil
// 	}); err != nil {
// 		err = fmt.Errorf("Failed to walk directory tree: %s", err)
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	return root
// }

var cssFiles = getStaticFiles(map[string]string{
	"dropzoneCSS":    "/res/dropzone.min.css",
	"materializeCSS": "/res/materialize.min.css",
})
var jsFiles = getStaticFiles(map[string]string{
	"dropzoneJS":    "/res/dropzone.min.js",
	"materializeJS": "/res/materialize.min.js",
})

func getStaticFiles(list map[string]string) []string {
	m := []string{}

	for _, v := range list {
		body, err := readFile(v)
		if err != nil {
			err = fmt.Errorf("Failed to load CSS file: %s", err)
			fmt.Println(err)
			os.Exit(1)
		}

		m = append(m, body)
	}
	return m
}

func readFile(path string) (string, error) {
	file, ok := assets.FS.String(path)
	if !ok {
		return "", fmt.Errorf("File '%s' not found in assets", path)
	}
	return file, nil
}
