package recvcheck

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// NewAnalyzer returns a new analyzer to check for receiver type consistency.
func NewAnalyzer(s Setting) *analysis.Analyzer {
	excludeMethods := []string{ // Default excludes for Marshal/Encode/Value methods #7
		"MarshalText() ([]byte, error)",
		"MarshalJSON() ([]byte, error)",
		"MarshalXML(*xml.Encoder, xml.StartElement) error",
		"MarshalBinary() ([]byte, error)",
		"GobEncode() ([]byte, error)",
		"Value() (driver.Value, error)",
		"Value() (any, error)",
		"Value() (interface{}, error)",
	}
	if s.NoBuiltinExcludeMethod {
		excludeMethods = []string{}
	}
	excludeMethods = append(excludeMethods, s.ExcludeMethod...)

	filter := newMethodSignatureFilter(excludeMethods...)

	return &analysis.Analyzer{
		Name:     "recvcheck",
		Doc:      "checks for receiver type consistency",
		Run:      runWithFilter(filter),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

// Setting is the configuration for the analyzer.
type Setting struct {
	// ExcludeMethod specifies method signatures to exclude from receiver type checking.
	// Each signature should be in the format "MethodName(paramTypes) returnTypes".
	//
	// Examples of valid signatures:
	//   - "MarshalJSON() ([]byte, error)"
	//   - "UnmarshalJSON([]byte) error"
	//   - "String() string"
	//
	// These signatures are merged with built-in excluded signatures unless
	// NoBuiltinExcludeMethod is set to true.
	//
	// Built-in excluded signatures:
	//   - "MarshalText() ([]byte, error)"
	//   - "MarshalJSON() ([]byte, error)"
	//   - "MarshalXML(*xml.Encoder, xml.StartElement) error"
	//   - "MarshalBinary() ([]byte, error)"
	//   - "GobEncode() ([]byte, error)"
	//   - "Value() (driver.Value, error)"
	//   - "Value() (any, error)"
	//   - "Value() (interface{}, error)"
	ExcludeMethod          []string
	NoBuiltinExcludeMethod bool // if true, disables the built-in excluded method signatures.
}

func runWithFilter(filter func(*ast.FuncDecl) bool) func(*analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
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

func newMethodSignatureFilter(signatures ...string) func(*ast.FuncDecl) bool {
	if len(signatures) == 0 {
		return func(*ast.FuncDecl) bool { return false }
	}

	excludes := make(map[string]struct{}, len(signatures))
	for _, sig := range signatures {
		excludes[sig] = struct{}{}
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

		_, exists := excludes[sig.String()]
		return exists
	}
}
