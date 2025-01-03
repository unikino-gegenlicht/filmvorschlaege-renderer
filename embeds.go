package main

import "embed"

//go:embed templates/sommer.html
//go:embed templates/winter.html
var content embed.FS

//go:embed style.css
var stylesheet string

//go:embed distributors.txt
var distributorBlacklist string
