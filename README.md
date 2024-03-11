# gormcnm

## gormcnm means: gorm column name. can help you use enum column name and enum type define. not use raw string.

English doc:
This tool allows you to enumerate the field names and column types of Gorm. By doing so, it helps to reduce the usage of raw strings in your project, minimizing the chances of making mistakes. When you need to modify Gorm fields, this tool can assist you in identifying issues during the coding or compilation phase, making the refactoring of Gorm types easier, simpler, and more reliable.

In addition to field enumeration, this tool provides some simple logic to simplify query operations. However, it assumes that you already have a good understanding of the basic CRUD operations using Gorm.

I developed this tool for my own projects and found it extremely useful. I decided to open source it to benefit future projects and the wider community. I hope you find it helpful, and if you do, please consider giving it a star. Your support is greatly appreciated.

Thank you!

Chinese doc:
该工具能枚举gorm的字段名和列类型，这样能避免gorm使用者在项目中写太多原始字符串，也就能避免写错，当您需要修改gorm字段时，该工具也能帮助您在编码阶段或编译阶段就发现问题，使得重构gorm类型更简单更轻松更可靠。当然该工具还提供一些简单的逻辑，使得查询操作变得更容易，但这都是建立在您已经充分掌握gorm基本增删改查操作的基础上的。

我自己开发出来以后觉得非常好用，就顺带把它开源出来，以方便以后我做别的项目，也方便大家，希望大家能给星星哦，谢谢大家。

English doc:
This tool is not a framework but rather a utility. As a result, it is non-intrusive, meaning you can use it alongside Gorm at any time. The tool's functionality is not comprehensive; it only addresses common CRUD operations. However, I believe this is sufficient for most use cases. In situations where you encounter special scenarios, you can still rely on Gorm's specific logic.

Chinese doc:
这个工具不是框架，而是一个工具，因此不具备侵入性，即，你可以随时用gorm，也可以随时用这个工具。 因此该工具的功能不是完整的，而是只能解决常用的增删改查逻辑，我相信这已经是足够用的，当遇到特殊场景时再使用gorm的特殊逻辑就行。

How to use:
1. define models
2. write enum vars
3. use gorm to select update delete

Demo:
[Example 文件](/demo/main/main.go)
