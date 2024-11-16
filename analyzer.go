package recvcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// NewAnalyzer returns a new analyzer to check for receiver type consistency.
func NewAnalyzer(s Setting) *analysis.Analyzer {
	// Default excludes for Marshal/Encode methods https://github.com/raeperd/recvcheck/issues/7
	excludeMethods := map[string]struct{}{
		"MarshalText":   {},
		"MarshalJSON":   {},
		"MarshalYAML":   {},
		"MarshalXML":    {},
		"MarshalBinary": {},
		"GobEncode":     {},
	}

	if s.DisableBuiltin {
		excludeMethods = map[string]struct{}{}
	}

	return &analysis.Analyzer{
		Name:     "recvcheck",
		Doc:      "checks for receiver type consistency",
		Run:      run(newFilter(excludeMethods)),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

// Setting is the configuration for the analyzer.
type Setting struct {
	// DisableBuiltin if true, disables the built-in method excludes.
	// Built-in excluded methods:
	//   - "MarshalText"
	//   - "MarshalJSON"
	//   - "MarshalYAML"
	//   - "MarshalXML"
	//   - "MarshalBinary"
	//   - "GobEncode"
	DisableBuiltin bool
}

func run(filter func(*ast.FuncDecl) bool) func(*analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

		structs := map[string]*structType{}
		inspector.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) != 1 {
				return
			}

			if filter(funcDecl) {
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
}

type structType struct {
	starUsed bool
	typeUsed bool
}

func newFilter(excludes map[string]struct{}) func(*ast.FuncDecl) bool {
	return func(f *ast.FuncDecl) bool {
		if f.Name == nil || f.Name.Name == "" {
			return true
		}

		_, found := excludes[f.Name.Name]
		return found
	}
}
