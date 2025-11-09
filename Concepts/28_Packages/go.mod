// If have to create module, then in cmd type: go mod init github.com/soumik171/podcast(demo project name)

module github.com/soumik171/podcast

go 1.25.3

// If wanna use aditional third party packages, then in cmd type(only for color): go get github.com/fatih/color

// If use any of the third party packages, after using them then, in cmd type for autodetect automatically: go mod tidy

require github.com/fatih/color v1.18.0

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
