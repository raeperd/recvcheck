package recvcheck

import (
	"go/ast"

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

	structs := map[string]*structType{}
	inspector.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) != 1 {
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

		var st *structType
		st, ok = structs[recv.Name]
		if !ok {
			structs[recv.Name] = &structType{starMethods: []*ast.FuncDecl{}, typeMethods: []*ast.FuncDecl{}}
			st = structs[recv.Name]
		}

		if isStar {
			st.starMethods = append(st.starMethods, funcDecl)
		} else {
			st.typeMethods = append(st.typeMethods, funcDecl)
		}
	})

	for _, st := range structs {
		if len(st.starMethods) > 0 && len(st.typeMethods) > 0 {
			for _, typeMethod := range st.typeMethods {
				pass.Reportf(typeMethod.Pos(), "receiver type should be a pointer")
			}
		}
	}

	return nil, nil
}

type structType struct {
	starMethods []*ast.FuncDecl
	typeMethods []*ast.FuncDecl
}
