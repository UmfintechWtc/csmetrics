package logx

import (
	"github.com/fatih/color"
)

// global colors.
var (
	ColorFuncMapWithFormat = map[string]func(format string, a ...interface{}) string{
		"Debugf": CyanBoldf,
		"Infof":  GreenBoldf,
		"Warnf":  YellowBoldf,
		"Errorf": RedBoldf,
		"Fatalf": RedBoldf,
	}
	ColorFuncMap = map[string]func(a ...interface{}) string{
		"Debug": CyanBold,
		"Info":  GreenBold,
		"Warn":  YellowBold,
		"Error": RedBold,
		"Fatal": RedBold,
	}
	// 红体加粗
	RedBold = func() func(a ...interface{}) string {
		return color.New(color.FgRed, color.Bold).SprintFunc()
	}()

	RedBoldf = func() func(format string, a ...interface{}) string {
		return color.New(color.FgRed, color.Bold).SprintfFunc()
	}()

	// 红色前景
	BgRed = func() func(format string, a ...interface{}) string {
		return color.New(color.BgRed).SprintfFunc()
	}()

	// 浅蓝体加粗
	CyanBold = func() func(a ...interface{}) string {
		return color.New(color.FgCyan, color.Bold).SprintFunc()
	}()

	CyanBoldf = func() func(format string, a ...interface{}) string {
		return color.New(color.FgCyan, color.Bold).SprintfFunc()
	}()

	// 浅蓝色前景
	BgCyan = func() func(format string, a ...interface{}) string {
		return color.New(color.BgCyan).SprintfFunc()
	}()

	// 蓝体加粗
	BlueBold = func() func(a ...interface{}) string {
		return color.New(color.FgBlue, color.Bold).SprintFunc()
	}()

	BlueBoldf = func() func(format string, a ...interface{}) string {
		return color.New(color.FgBlue, color.Bold).SprintfFunc()
	}()

	// 蓝色前景
	BgBlue = func() func(format string, a ...interface{}) string {
		return color.New(color.BgBlue).SprintfFunc()
	}()

	// 黄体加粗
	YellowBold = func() func(a ...interface{}) string {
		return color.New(color.FgYellow, color.Bold).SprintFunc()
	}()

	YellowBoldf = func() func(format string, a ...interface{}) string {
		return color.New(color.FgYellow, color.Bold).SprintfFunc()
	}()

	// 黄色前景
	BgYellow = func() func(format string, a ...interface{}) string {
		return color.New(color.BgYellow).SprintfFunc()
	}()

	// 绿体加粗
	GreenBold = func() func(a ...interface{}) string {
		return color.New(color.FgGreen, color.Bold).SprintFunc()
	}()

	GreenBoldf = func() func(format string, a ...interface{}) string {
		return color.New(color.FgGreen, color.Bold).SprintfFunc()
	}()

	// 绿色前景
	BgGreen = func() func(format string, a ...interface{}) string {
		return color.New(color.BgGreen).SprintfFunc()
	}()
)
