package constants

// RunTimeMode 运行环境
type RunTimeMode string

// 运行模式相关配置
const (
	DebugMode   RunTimeMode = "debug"   // 调试模式
	ReleaseMode RunTimeMode = "release" // 发布模式
	TestMode    RunTimeMode = "test"    // 测试模式
)
