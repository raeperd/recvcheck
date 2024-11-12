package recvcheck

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "recvcheck",
	Doc:      "checks for receiver type consistency",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	methodSignatureFilter := newMethodSignatureFilter(false)

	structs := map[string]*structType{}
	inspector.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) != 1 {
			return
		}
		if methodSignatureFilter(funcDecl) {
			return
		}

		var recv *ast.Ident
		var isStar bool
		switch recvType := funcDecl.Recv.List[0].Type.(type) {
		case *ast.StarExpr:
			isStar = true
			if recv, ok = recvType.X.(*ast.Ident); !ok {
				return
			}
		case *ast.Ident:
			recv = recvType
		default:
			return
		}

		st, ok := structs[recv.Name]
		if !ok {
			structs[recv.Name] = &structType{}
			st = structs[recv.Name]
		}

		if isStar {
			st.starUsed = true
		} else {
			st.typeUsed = true
		}
	})

	for recv, st := range structs {
		if st.starUsed && st.typeUsed {
			pass.Reportf(pass.Pkg.Scope().Lookup(recv).Pos(), "the methods of %q use pointer receiver and non-pointer receiver.", recv)
		}
	}

	return nil, nil
}

type structType struct {
	starUsed bool
	typeUsed bool
}

func newMethodSignatureFilter(noBuiltin bool, signatures ...string) func(*ast.FuncDecl) bool {
	filter := map[string]struct{}{
		"MarshalText() ([]byte, error)":                    {},
		"MarshalJSON() ([]byte, error)":                    {},
		"MarshalXML(*xml.Encoder, xml.StartElement) error": {},
		"MarshalBinary() ([]byte, error)":                  {},
		"GobEncode() ([]byte, error)":                      {},
		"Value() (driver.Value, error)":                    {},
		"Value() (any, error)":                             {},
		"Value() (interface{}, error)":                     {},
	}
	if noBuiltin {
		filter = map[string]struct{}{}
	}
	for _, sig := range signatures {
		filter[sig] = struct{}{}
	}

	return func(f *ast.FuncDecl) bool {
		// Build signature string: "FuncName(type1, type2) (retType1, retType2), FuncName2() error"
		var sig strings.Builder
		sig.WriteString(f.Name.Name)
		sig.WriteString("(")
		if f.Type.Params != nil {
			for i, param := range f.Type.Params.List {
				if i > 0 {
					sig.WriteString(", ")
				}
				sig.WriteString(types.ExprString(param.Type))
			}
		}
		sig.WriteString(")")

		if f.Type.Results != nil && len(f.Type.Results.List) > 0 {
			sig.WriteString(" ")
			if len(f.Type.Results.List) > 1 {
				sig.WriteString("(")
			}
			for i, result := range f.Type.Results.List {
				if i > 0 {
					sig.WriteString(", ")
				}
				sig.WriteString(types.ExprString(result.Type))
			}
			if len(f.Type.Results.List) > 1 {
				sig.WriteString(")")
			}
		}

		_, exists := filter[sig.String()]
		return exists
	}
}
