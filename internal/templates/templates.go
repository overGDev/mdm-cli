package templates

import _ "embed"

//go:embed schema.yaml
var SampleSchema string

//go:embed update_readme.yaml
var GitHubAction string

//go:embed script.py
var Script string

//go:embed ignore
var GitIgnore string
