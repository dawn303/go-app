# Go-APP开发手册 001

Go-APP项目是基于`go-kratos`进行开发的，要使用go-kratos，需要先了解两个前置知识：wire依赖注入 和 protobuf。

### wire依赖注入

wire<google/wire>的好处：1.编译时就能发现问题； 2.对mock测试友好

wire主要有两个概念：Provider 和 Injector

1. Provider创建需要注入的对象

`wire.NewSet` 
`wire.Struct`
`wire.Value`
`wire.Bind`

2. Injector根据对象的依赖关系，将所有Provider依次构建成最终的目标对象

`wire.Build`
`wireApp`

### protobuf

1. 定义protobuf文件说明接口
2. 利用protoc将定义好的protobuf生成代码
3. 将生成代码整合到项目中
4. 完善业务逻辑
