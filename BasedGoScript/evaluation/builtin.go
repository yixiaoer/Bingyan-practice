package evaluation

import (
	"fmt"
	"project/BasedGoScript/object"
)

var builtins = map[string]*object.Builtin{
	"Puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
