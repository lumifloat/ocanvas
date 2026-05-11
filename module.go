package ocanvas

import (
	"github.com/dop251/goja"
	"github.com/lumifloat/tinyskia"
)

var (
	tinyskiaPath2d = goja.NewSymbol("tinyskia_path2d")
)

func CanvasModuleLoader(runtime *goja.Runtime, module *goja.Object) {
	exports := module.Get("exports").(*goja.Object)

	exports.Set("createCanvas", func(call goja.FunctionCall) goja.Value {
		width := call.Argument(0).ToInteger()
		height := call.Argument(1).ToInteger()

		canvas := runtime.NewObject()
		canvas.Set("getContext", func(call goja.FunctionCall) goja.Value {
			contextType := call.Argument(0).String()
			if contextType == "2d" {
				return createRenderingContext(runtime, int(width), int(height))
			}
			return goja.Undefined()
		})

		return canvas
	})

}

func createRenderingContext(runtime *goja.Runtime, width int, height int) *goja.Object {
	ctx := tinyskia.NewContext(width, height)
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
	rCtx.Set("fill", func(call goja.FunctionCall) goja.Value {
		var path2d *tinyskia.Path2D
		var fillRule string
		switch len(call.Arguments) {
		case 0:
			break
		case 1:
			call.Argument(0).ExportType()
			var p0 *tinyskia.Path2D
			t := call.Argument(0).ToObject(runtime).GetSymbol(tinyskiaPath2d)
			if err := runtime.ExportTo(t, &p0); err != nil {
				fillRule = call.Argument(0).String()
			} else {
				path2d = p0
			}
		default:
			var p0 *tinyskia.Path2D
			t := call.Argument(0).ToObject(runtime).GetSymbol(tinyskiaPath2d)
			if err := runtime.ExportTo(t, &p0); err != nil {
				fillRule = call.Argument(0).String()
			} else {
				path2d = p0
				fillRule = call.Argument(1).String()
			}
		}
		if fillRule == "evenodd" {
			if path2d != nil {
				ctx.FillPathWithFillRule(path2d, tinyskia.FillRuleEvenOdd)
			} else {
				ctx.FillWithFillRule(tinyskia.FillRuleEvenOdd)
			}
		} else {
			if path2d != nil {
				ctx.FillPath(path2d)
			} else {
				ctx.Fill()
			}
		}
		return goja.Undefined()
	})
	rCtx.Set("stroke", func(call goja.FunctionCall) goja.Value {
		var path2d *tinyskia.Path2D
		switch len(call.Arguments) {
		case 0:
			break
		default:
			call.Argument(0).ExportType()
			var p0 *tinyskia.Path2D
			t := call.Argument(0).ToObject(runtime).GetSymbol(tinyskiaPath2d)
			if err := runtime.ExportTo(t, &p0); err == nil {
				path2d = p0
			}
		}
		if path2d != nil {
			ctx.StrokePath(path2d)
		} else {
			ctx.Stroke()
		}
		return goja.Undefined()
	})

	return rCtx
}

func newPath2D(runtime *goja.Runtime) *goja.Object {
	p := tinyskia.NewPath2D()
	rP := runtime.NewObject()
	rP.DefineDataPropertySymbol(
		tinyskiaPath2d,
		runtime.ToValue(p),
		goja.FLAG_TRUE, goja.FLAG_TRUE, goja.FLAG_TRUE,
	)

	rP.Set("moveTo", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}
		x := call.Argument(0).ToFloat()
		y := call.Argument(1).ToFloat()
		p.MoveTo(x, y)
		return goja.Undefined()
	})
	rP.Set("lineTo", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}
		x := call.Argument(0).ToFloat()
		y := call.Argument(1).ToFloat()
		p.LineTo(x, y)
		return goja.Undefined()
	})
	rP.Set("closePath", func(call goja.FunctionCall) goja.Value {
		p.ClosePath()
		return goja.Undefined()
	})
	rP.Set("quadraticCurveTo", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 4 {
			return goja.Undefined()
		}
		x1 := call.Argument(0).ToFloat()
		y1 := call.Argument(1).ToFloat()
		x2 := call.Argument(2).ToFloat()
		y2 := call.Argument(3).ToFloat()
		p.QuadraticCurveTo(x1, y1, x2, y2)
		return goja.Undefined()
	})
	rP.Set("bezierCurveTo", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 6 {
			return goja.Undefined()
		}
		x1 := call.Argument(0).ToFloat()
		y1 := call.Argument(1).ToFloat()
		x2 := call.Argument(2).ToFloat()
		y2 := call.Argument(3).ToFloat()
		x3 := call.Argument(4).ToFloat()
		y3 := call.Argument(5).ToFloat()
		p.BezierCurveTo(x1, y1, x2, y2, x3, y3)
		return goja.Undefined()
	})
	rP.Set("arcTo", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 4 {
			return goja.Undefined()
		}
		x1 := call.Argument(0).ToFloat()
		y1 := call.Argument(1).ToFloat()
		x2 := call.Argument(2).ToFloat()
		y2 := call.Argument(3).ToFloat()
		radius := call.Argument(4).ToFloat()
		p.ArcTo(x1, y1, x2, y2, radius)
		return goja.Undefined()
	})
	rP.Set("rect", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 4 {
			return goja.Undefined()
		}
		x := call.Argument(0).ToFloat()
		y := call.Argument(1).ToFloat()
		w := call.Argument(2).ToFloat()
		h := call.Argument(3).ToFloat()
		p.Rect(x, y, w, h)
		return goja.Undefined()
	})
	rP.Set("roundRect", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 5 {
			return goja.Undefined()
		}
		x := call.Argument(0).ToFloat()
		y := call.Argument(1).ToFloat()
		w := call.Argument(2).ToFloat()
		h := call.Argument(3).ToFloat()
		var radius []float64
		if err := runtime.ExportTo(call.Argument(4), &radius); err != nil {
			p.RoundRect(x, y, w, h, radius)
			return goja.Undefined()
		}
		p.RoundRect(x, y, w, h, nil)
		return goja.Undefined()
	})
	rP.Set("arc", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 5 {
			return goja.Undefined()
		}
		x := call.Argument(0).ToFloat()
		y := call.Argument(1).ToFloat()
		radius := call.Argument(2).ToFloat()
		startAngle := call.Argument(3).ToFloat()
		endAngle := call.Argument(4).ToFloat()
		counterclockwise := call.Argument(5).ToBoolean()
		p.Arc(x, y, radius, startAngle, endAngle, counterclockwise)
		return goja.Undefined()
	})
	rP.Set("ellipse", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 7 {
			return goja.Undefined()
		}
		x := call.Argument(0).ToFloat()
		y := call.Argument(1).ToFloat()
		radiusX := call.Argument(2).ToFloat()
		radiusY := call.Argument(3).ToFloat()
		rotation := call.Argument(4).ToFloat()
		startAngle := call.Argument(5).ToFloat()
		endAngle := call.Argument(6).ToFloat()
		counterclockwise := call.Argument(7).ToBoolean()
		p.Ellipse(x, y, radiusX, radiusY, rotation, startAngle, endAngle, counterclockwise)
		return goja.Undefined()
	})
	rP.Set("addPath", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		var p0 *tinyskia.Path2D
		t := call.Argument(0).ToObject(runtime).GetSymbol(tinyskiaPath2d)
		if err := runtime.ExportTo(t, &p0); err != nil {
			return goja.Undefined()
		}
		p.AddPath(p0)
		return goja.Undefined()
	})
	return rP
}
