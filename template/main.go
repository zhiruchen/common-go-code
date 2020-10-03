package main

import (
	"encoding/json"
	"os"
	"text/template"
)

const tmpl = `
"resources": {{ .res }}
`

func main() {
	data := map[string]map[string]interface{}{
		"limits": {
			"cpu":    "4",
			"memory": "4Gi",
		},
		"requests": {
			"cpu":    "4",
			"memory": "4Gi",
		},
	}

	bs, _ := json.Marshal(data)

	res := map[string]string{
		"res": string(bs),
	}

	t := template.Must(template.New("").Parse(tmpl))
	if err := t.Execute(os.Stdout, res); err != nil {
		panic(err)
	}
}
