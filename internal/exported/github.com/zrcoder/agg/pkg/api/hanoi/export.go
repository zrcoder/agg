// export by github.com/goplus/igop/cmd/qexp

package hanoi

import (
	q "github.com/zrcoder/agg/pkg/api/hanoi"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "hanoi",
		Path: "github.com/zrcoder/agg/pkg/api/hanoi",
		Deps: map[string]string{
			"github.com/zrcoder/agg/internal/hanoi": "hanoi",
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
