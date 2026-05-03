package ocanvas

import (
	"github.com/dop251/goja"
	"github.com/lumifloat/tinyskia"
)

func CanvasModuleLoader(runtime *goja.Runtime, module *goja.Object) {
	exports := module.Get("exports").(*goja.Object)

	exports.Set("createCanvas", func(call goja.FunctionCall) goja.Value {
		width := call.Argument(0).ToInteger()
		height := call.Argument(1).ToInteger()

		canvas := runtime.NewObject()
		ctx := tinyskia.NewContext(int(width), int(height))
		canvas.Set("getContext", func(call goja.FunctionCall) goja.Value {
			contextType := call.Argument(0).String()
			if contextType == "2d" {
				return createRenderingContext(runtime, ctx)
			}
			return goja.Undefined()
		})

		return canvas
	})

}

func createRenderingContext(runtime *goja.Runtime, ctx *tinyskia.Context) *goja.Object {
	rCtx := runtime.NewObject()
	rCtx.DefineAccessorProperty("lineCap",
		runtime.ToValue(func() goja.Value {
			switch value := ctx.GetLineCap(); value {
			case tinyskia.LineCapButt:
				return runtime.ToValue("butt")
			case tinyskia.LineCapRound:
				return runtime.ToValue("round")
			case tinyskia.LineCapSquare:
				return runtime.ToValue("square")
			default:
				return runtime.ToValue("butt")
			}
		}),
		runtime.ToValue(func(lineCap string) {
			switch lineCap {
			case "butt":
				ctx.SetLineCap(tinyskia.LineCapButt)
			case "round":
				ctx.SetLineCap(tinyskia.LineCapRound)
			case "square":
				ctx.SetLineCap(tinyskia.LineCapSquare)
			default:
				// pass
			}
		}),
		goja.FLAG_TRUE, goja.FLAG_TRUE,
	)
	rCtx.DefineAccessorProperty("lineJoin",
		runtime.ToValue(func() goja.Value {
			switch value := ctx.GetLineJoin(); value {
			case tinyskia.LineJoinMiter:
				return runtime.ToValue("miter")
			case tinyskia.LineJoinRound:
				return runtime.ToValue("round")
			case tinyskia.LineJoinBevel:
				return runtime.ToValue("bevel")
			default:
				return runtime.ToValue("miter")
			}
		}),
		runtime.ToValue(func(lineJoin string) {
			switch lineJoin {
			case "miter":
				ctx.SetLineJoin(tinyskia.LineJoinMiter)
			case "round":
				ctx.SetLineJoin(tinyskia.LineJoinRound)
			case "bevel":
				ctx.SetLineJoin(tinyskia.LineJoinBevel)
			default:
				// pass
			}
		}),
		goja.FLAG_TRUE, goja.FLAG_TRUE,
	)
	rCtx.DefineAccessorProperty("textAlign",
		runtime.ToValue(func() goja.Value {
			switch value := ctx.GetTextAlign(); value {
			case tinyskia.TextAlignLeft:
				return runtime.ToValue("left")
			case tinyskia.TextAlignRight:
				return runtime.ToValue("right")
			case tinyskia.TextAlignCenter:
				return runtime.ToValue("center")
			case tinyskia.TextAlignStart:
				return runtime.ToValue("start")
			case tinyskia.TextAlignEnd:
				return runtime.ToValue("end")
			default:
				return runtime.ToValue("left")
			}
		}),
		runtime.ToValue(func(textAlign string) {
			switch textAlign {
			case "left":
				ctx.SetTextAlign(tinyskia.TextAlignLeft)
			case "right":
				ctx.SetTextAlign(tinyskia.TextAlignRight)
			case "center":
				ctx.SetTextAlign(tinyskia.TextAlignCenter)
			case "start":
				ctx.SetTextAlign(tinyskia.TextAlignStart)
			case "end":
				ctx.SetTextAlign(tinyskia.TextAlignEnd)
			default:
				// pass
			}
		}),
		goja.FLAG_TRUE, goja.FLAG_TRUE,
	)
	rCtx.DefineAccessorProperty("fontKerning",
		runtime.ToValue(func() goja.Value {
			switch value := ctx.GetFontKerning(); value {
			case tinyskia.FontKerningAuto:
				return runtime.ToValue("auto")
			case tinyskia.FontKerningNormal:
				return runtime.ToValue("normal")
			case tinyskia.FontKerningNone:
				return runtime.ToValue("none")
			default:
				return runtime.ToValue("auto")
			}
		}),
		runtime.ToValue(func(fontKerning string) {
			switch fontKerning {
			case "auto":
				ctx.SetFontKerning(tinyskia.FontKerningAuto)
			case "normal":
				ctx.SetFontKerning(tinyskia.FontKerningNormal)
			case "none":
				ctx.SetFontKerning(tinyskia.FontKerningNone)
			default:
				// pass
			}
		}),
		goja.FLAG_TRUE, goja.FLAG_TRUE,
	)

	return rCtx
}
