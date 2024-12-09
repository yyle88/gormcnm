package gormcnqs

import (
	"testing"

	"github.com/yyle88/must"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/cls_stub_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

func TestGen(t *testing.T) {
	packageName := syntaxgo_reflect.GetPkgNameV4(stub)
	must.Equals("gormcnm", packageName)

	cfg := &cls_stub_gen.StubGenConfig{
		SourceRootPath:    runpath.PARENT.UpTo(2, packageName),
		TargetPackageName: syntaxgo.CurrentPackageName(),
		ImportOptions:     syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:        runtestpath.SrcPath(t),
		AllowFileCreation: false,
	}

	param := cls_stub_gen.NewStubParam(stub, "stub")
	cls_stub_gen.GenerateStubs(cfg, param)
}
