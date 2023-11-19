package gopublicfield

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/gobwas/glob"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func NewAnalyzer() *analysis.Analyzer {
	analyzer := &analysis.Analyzer{
		Name:     "publicfield",
		Doc:      "Blocks using public fields",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	cfg := config{}

	analyzer.Flags.Var(&cfg.packageGlobs, "packageGlobs", packageGlobsDesc)
	analyzer.Flags.BoolVar(&cfg.onlyPackageGlobs, "onlyPackageGlobs", false, onlyPackageGlobsDesc)

	analyzer.Run = run(&cfg)

	return analyzer
}

func run(cfg *config) func(pass *analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		var blockedStrategy pkgsStrategy = newAnotherPkg()
		if len(cfg.packageGlobs) > 0 {
			defaultStrategy := blockedStrategy
			if cfg.onlyPackageGlobs {
				defaultStrategy = newNilPkg()
			}

			pkgGlobs, err := compileGlobs(cfg.packageGlobs.Value())
			if err != nil {
				return nil, err
			}

			blockedStrategy = newBlockedPkgs(
				pkgGlobs,
				defaultStrategy,
			)
		}

		for _, file := range pass.Files {
			v := &visiter{
				pass:            pass,
				blockedStrategy: blockedStrategy,
			}
			v.walk(file)
		}

		return nil, nil
	}
}

type visiter struct {
	pass            *analysis.Pass
	blockedStrategy pkgsStrategy
}

func (v *visiter) walk(n ast.Node) {
	if n != nil {
		ast.Walk(v, n)
	}
}

func (v *visiter) Visit(node ast.Node) ast.Visitor {
	incDecStmt, ok := node.(*ast.IncDecStmt)
	if ok {
		v.selectorExpr(incDecStmt.X)

		return v
	}

	assignStmt, ok := node.(*ast.AssignStmt)
	if !ok {
		return v
	}

	for _, lh := range assignStmt.Lhs {
		v.selectorExpr(lh)
	}

	return v
}

func (v *visiter) selectorExpr(expr ast.Expr) {
	selExpr, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return
	}

	structIdent, ok := selExpr.X.(*ast.Ident)
	fieldIdent := selExpr.Sel

	if !ok {
		structIdent = v.tryResolvePointers(selExpr)
	}

	if structIdent == nil {
		return
	}

	structObj := v.pass.TypesInfo.ObjectOf(structIdent)
	fieldObj := v.pass.TypesInfo.ObjectOf(fieldIdent)

	if structObj == nil || fieldObj == nil {
		return
	}

	if _, ok := sourceType(structObj); !ok {
		return
	}

	if v.blockedStrategy.IsPkgs(v.pass.Pkg, structObj) {
		v.report(structIdent, structObj, fieldObj)
	}
}

func (v *visiter) tryResolvePointers(selExpr *ast.SelectorExpr) *ast.Ident {
	parenExpr, ok := selExpr.X.(*ast.ParenExpr)
	if !ok {
		return nil
	}

	starExpr, ok := parenExpr.X.(*ast.StarExpr)
	if !ok {
		return nil
	}

	var res *ast.Ident

	ok = true
	for ok {
		res, ok = starExpr.X.(*ast.Ident)
		if ok {
			break
		}

		starExpr, ok = starExpr.X.(*ast.StarExpr)
	}

	return res
}

func (v *visiter) report(
	node ast.Node,
	structObj types.Object,
	fieldObj types.Object,
) {
	structType, _ := sourceType(structObj)

	v.pass.Reportf(
		node.Pos(),
		fmt.Sprintf(
			`Field '%s' in %s.%s can be changes only inside nested package.`,
			fieldObj.Name(),
			structType.Obj().Pkg().Name(),
			structType.Obj().Name(),
		),
	)
}

func containsMatchGlob(globs []glob.Glob, el string) bool {
	for _, g := range globs {
		if g.Match(el) {
			return true
		}
	}

	return false
}

func compileGlobs(globs []string) ([]glob.Glob, error) {
	compiledGlobs := make([]glob.Glob, len(globs))

	for idx, globString := range globs {
		glob, err := glob.Compile(globString)
		if err != nil {
			return nil, fmt.Errorf("unable to compile globs %s: %w", glob, err)
		}

		compiledGlobs[idx] = glob
	}

	return compiledGlobs, nil
}
