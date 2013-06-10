
### Mixpanel API in Go

Here is the simple usage:

```go
package main

import (
	"fmt"
	"github.com/athom/gomixpanel"
)

func main() {
	gomixpanel.SetApiToken("76d89d8aeebb5053719915f18124dbf7")
	gomixpanel.Track("Simple Track Demo", map[string]interface{}{
		"lib":     "GoMixpanel",
		"authors": []string{"Rich Collins", "athom"},
		"age":     12,
		"version": 20130610,
	})

	gomixpanel.TrackWithCallback("Singleton API Event", map[string]interface{}{
		"lib":     "GoMixpanel",
		"authors": []string{"Rich Collins", "athom"},
		"age":     12,
		"version": 20130610,
	}, func() {
		fmt.Println("Put Singleton Event Success")
	})
}
```
