// export by github.com/goplus/ixgo/cmd/qexp

package hanoi

import (
	q "github.com/zrcoder/agg/pkg/export/hanoi"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "hanoi",
		Path: "github.com/zrcoder/agg/pkg/export/hanoi",
		Deps: map[string]string{
			"github.com/zrcoder/agg/internal":             "internal",
			"github.com/zrcoder/agg/internal/games/hanoi": "hanoi",
			"time": "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"A": reflect.ValueOf(q.A),
			"B": reflect.ValueOf(q.B),
			"C": reflect.ValueOf(q.C),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
