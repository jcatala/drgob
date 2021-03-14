package banner

import (
	"bytes"
	"fmt"
	"github.com/dimiro1/banner"
	"os"
)

func Banner(version string){
	templ := fmt.Sprintf(`{{ .Title "(D)iscord (R)eddit (G)o(B)ot" "" 4 }}
{{ .AnsiColor.BrightCyan }}Thanks for using Drgob ! {{ .AnsiColor.Default }}
{{ .AnsiColor.BrightGreen}}Drgob version: %s {{.AnsiColor.Default}}
Go version: {{ .GoVersion }}
Now: {{ .Now "Monday, 2 Jan 2006" }}
How to use:
`,version)
	isEnabled := true
	isColorEnabled := true
	banner.Init(os.Stdout, isEnabled, isColorEnabled,bytes.NewBufferString(templ) )
}