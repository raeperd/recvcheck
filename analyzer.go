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
	excludedMethods := map[string]struct{}{
		"MarshalText":   {},
		"MarshalJSON":   {},
		"MarshalYAML":   {},
		"MarshalXML":    {},
		"MarshalBinary": {},
		"GobEncode":     {},
	}

	if s.DisableBuiltin {
		excludedMethods = map[string]struct{}{}
	}

	a := &analyzer{excludedMethods: excludedMethods}

	return &analysis.Analyzer{
		Name:     "recvcheck",
		Doc:      "checks for receiver type consistency",
		Run:      a.run,
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

type analyzer struct {
	excludedMethods map[string]struct{}
}

func (r *analyzer) run(pass *analysis.Pass) (any, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	structs := map[string]*structType{}
	inspector.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) != 1 {
			return
		}

		if r.isExcluded(funcDecl) {
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

func (r *analyzer) isExcluded(f *ast.FuncDecl) bool {
	if f.Name == nil || f.Name.Name == "" {
		return true
	}

	_, found := r.excludedMethods[f.Name.Name]
	return found
}

type structType struct {
	starUsed bool
	typeUsed bool
}
