package gopublicfield

import (
	"go/types"

	"github.com/gobwas/glob"
)

type pkgsStrategy interface {
	IsPkgs(currentPkg *types.Package, identObj types.Object) bool
}

func sourceType(obj types.Object) (*types.Named, bool) {
	objType := obj.Type()

	ptr, ok := objType.(*types.Pointer)
	for ok {
		objType = ptr.Elem()

		ptr, ok = objType.(*types.Pointer)
	}

	alias, ok := objType.(*types.Alias)
	for ok {
		objType = types.Unalias(alias)

		alias, ok = objType.(*types.Alias)
	}

	objType, ok = objType.(*types.Named)
	if !ok {
		return nil, false
	}

	namedType, ok := objType.(*types.Named)

	return namedType, ok
}

type nilPkg struct{}

func newNilPkg() nilPkg {
	return nilPkg{}
}

func (nilPkg) IsPkgs(_ *types.Package, _ types.Object) bool {
	return false
}

type anotherPkg struct{}

func newAnotherPkg() anotherPkg {
	return anotherPkg{}
}

func (anotherPkg) IsPkgs(
	currentPkg *types.Package,
	identObj types.Object,
) bool {
	t, _ := sourceType(identObj)

	return currentPkg.Path() != t.Obj().Pkg().Path()
}

type blockedPkgs struct {
	pkgs            []glob.Glob
	defaultStrategy pkgsStrategy
}

func newBlockedPkgs(
	pkgs []glob.Glob,
	defaultStrategy pkgsStrategy,
) blockedPkgs {
	return blockedPkgs{
		pkgs:            pkgs,
		defaultStrategy: defaultStrategy,
	}
}

func (b blockedPkgs) IsPkgs(
	currentPkg *types.Package,
	identObj types.Object,
) bool {
	sourceType, _ := sourceType(identObj)

	currentPkgPath := currentPkg.Path() + "/"
	isIncludedInBlocked := containsMatchGlob(b.pkgs, currentPkgPath)

	if isIncludedInBlocked {
		return false
	}

	identPkgPath := sourceType.Obj().Pkg().Path() + "/"
	isBlocked := containsMatchGlob(b.pkgs, identPkgPath)

	if isBlocked {
		return true
	}

	if b.defaultStrategy.IsPkgs(currentPkg, identObj) {
		return true
	}

	return false
}
