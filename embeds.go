package main

import "embed"

//go:embed style.css
//go:embed templates/sommer.html
//go:embed templates/winter.html
var content embed.FS
