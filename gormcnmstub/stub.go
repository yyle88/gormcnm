// Package gormcnmstub provides stub implementations and shared instances used in gormcnm operations
// Auto exposes ColumnOperationClass instance with common operations like JOIN, COALESCE
// Supports testing and development with pre-configured column operation patterns
//
// gormcnmstub 包提供 gormcnm 操作中使用的存根实现和共享实例
// 自动暴露 ColumnOperationClass 实例，包含 JOIN、COALESCE 等常用操作
// 支持使用预配置的列操作模式进行测试和开发
package gormcnmstub

import "github.com/yyle88/gormcnm"

// stub provides a shared instance of ColumnOperationClass for testing purposes
// stub 提供 ColumnOperationClass 的共享实例，用于测试目的
var stub = &gormcnm.ColumnOperationClass{}
