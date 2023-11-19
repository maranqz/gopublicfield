package gopublicfield

import (
	"go/types"
	"strings"
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

type pkgs struct {
	pkgs            []string
	defaultStrategy pkgsStrategy
}

func newBlockedPkgs(
	pkgsSlice []string,
	defaultStrategy pkgsStrategy,
) pkgs {
	return pkgs{
		pkgs:            pkgsSlice,
		defaultStrategy: defaultStrategy,
	}
}

func (b pkgs) IsPkgs(
	currentPkg *types.Package,
	identObj types.Object,
) bool {
	sourceType, _ := sourceType(identObj)

	identPkgPath := sourceType.Obj().Pkg().Path() + "/"
	currentPkgPath := currentPkg.Path() + "/"

	for _, blockedPkg := range b.pkgs {
		isBlocked := strings.HasPrefix(identPkgPath, blockedPkg)
		isIncludedInBlocked := strings.HasPrefix(currentPkgPath, blockedPkg)

		if isIncludedInBlocked {
			continue
		}

		if isBlocked {
			return true
		}

		if b.defaultStrategy.IsPkgs(currentPkg, identObj) {
			return true
		}
	}

	return false
}
