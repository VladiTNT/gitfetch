package assets

import _ "embed"

//go:embed help.txt
var HelpMsg string

//go:embed ignores.json
var IgnoresJsonData []byte
