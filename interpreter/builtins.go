package interpreter

import (
	"fmt"
	"strings"
)

var builtins = map[string]*Builtin{}

func init() {
	builtins["print"] = &Builtin{
		Fn: func(args ...Object) Object {
			var out strings.Builder
			for i, arg := range args {
				if i > 0 {
					out.WriteString(" ")
				}
				out.WriteString(arg.String())
			}
			fmt.Println(out.String())
			return NULL
		},
	}
	builtins["type"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError(
					"type() takes exactly 1 argument, got %d",
					len(args),
				)
			}
			return &String{Value: string(args[0].Type())}
		},
	}
	builtins["len"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError(
					"len() takes exactly 1 argument, got %d",
					len(args),
				)
			}
			switch arg := args[0].(type) {
			case *String:
				return &Integar{Value: int64(len(arg.Value))}
			case *Array:
				return &Integar{Value: int64(len(arg.Elements))}
				// todo add table
			default:
				return newError(
					"len() not supported for %s",
					args[0].Type(),
				)
			}
		},
	}
}
