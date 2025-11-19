// Package gormcnmstub tests validate code generation operations used in stub wrapper creation
// Auto verifies generated code correctness and AST manipulation patterns
// Tests cover stub generation, reflection-based code creation, and syntax validation
//
// gormcnmstub 测试包验证存根包装创建中使用的代码生成操作
// 自动验证生成代码的正确性和 AST 操作模式
// 测试涵盖存根生成、基于反射的代码创建和语法验证
package gormcnmstub

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
