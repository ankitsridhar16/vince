package plot

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const (
	AXIS_TICK_LENGTH = 6
	LABEL_MARGIN     = 4
	LABEL_WIDTH      = 25
	TOTAL_PADDING    = 120
	LABEL_MAX_CHARS  = 18
	FONT_SIZE        = 10
	BASE_LINE_COLOR  = "#E2E6E9"
)

type createOptions struct {
	inside, around *html.Node
	style          map[string]string
	innerHtml      string
	attr           []html.Attribute
}

func createSVG(tag string, o createOptions) *html.Node {
	e := &html.Node{
		Namespace: "http://www.w3.org/2000/svg",
	}
	if o.inside != nil {
		o.inside.AppendChild(e)
	}
	if o.around != nil {
		o.around.Parent.InsertBefore(e, o.around)
		e.AppendChild(o.around)
	}
	if len(o.style) > 0 {
		e.Attr = append(e.Attr, style(o.style))
	}
	if o.innerHtml != "" {
		e.Data = o.innerHtml
	}
	if len(o.attr) > 0 {
		e.Attr = append(e.Attr, o.attr...)
	}
	return e
}

func style(m map[string]string) html.Attribute {
	// make sure this is idempotent. Sort keys used for the attribute
	ls := make([]string, 0, len(m))
	for k := range m {
		ls = append(ls, k)
	}
	sort.Strings(ls)
	var s strings.Builder
	for i := range ls {
		s.WriteString(ls[i])
		s.WriteByte(':')
		s.WriteString(m[ls[i]])
		s.WriteByte(';')
	}
	return html.Attribute{
		Key: "style",
		Val: s.String(),
	}
}

func renderVerticalGradient(svgDefElem *html.Node, gradientId string) *html.Node {
	return createSVG("linearGradient", createOptions{
		inside: svgDefElem,
		attr: []html.Attribute{
			{Key: "id", Val: gradientId},
			{Key: "x1", Val: "0"},
			{Key: "x2", Val: "0"},
			{Key: "y1", Val: "0"},
			{Key: "y2", Val: "1"},
		},
	})
}

func setGradientStop(gradElem *html.Node, offset, color, opacity string) *html.Node {
	return createSVG("stop", createOptions{
		inside: gradElem,
		style: map[string]string{
			"stop-color": color,
		},
		attr: []html.Attribute{
			{Key: "offset", Val: offset},
			{Key: "stop-opacity", Val: opacity},
		},
	})
}

func makeSVGContainer(parent *html.Node, className, width, height string) *html.Node {
	return createSVG("svg", createOptions{
		inside: parent,
		attr: []html.Attribute{
			{Key: "class", Val: className},
			{Key: "width", Val: width},
			{Key: "height", Val: height},
		},
	})
}

func makeSVGDefs(svgContainer *html.Node) *html.Node {
	return createSVG("defs", createOptions{
		inside: svgContainer,
	})
}

func makeSVGGroup(className string, args ...any) *html.Node {
	o := createOptions{
		attr: []html.Attribute{
			{Key: "class", Val: className},
		},
	}
	var transform string
	for _, a := range args {
		switch e := a.(type) {
		case string:
			transform = e
		case *html.Node:
			o.inside = e
		}
	}
	o.attr = append(o.attr, html.Attribute{Key: "transform", Val: transform})
	return createSVG("g", o)
}

type pathOpts struct {
	className, stroke, fill, strokeWidth string
}

func makePath(pathStr string, o pathOpts) *html.Node {
	if o.stroke == "" {
		o.stroke = "none"
	}
	if o.fill == "" {
		o.fill = "none"
	}
	if o.strokeWidth == "" {
		o.strokeWidth = "2"
	}
	return createSVG("path", createOptions{
		attr: []html.Attribute{
			{Key: "class", Val: o.className},
			{Key: "d", Val: pathStr},
		},
		style: map[string]string{
			"stroke":       o.stroke,
			"fill":         o.fill,
			"stroke-width": o.strokeWidth,
		},
	})
}

func makeGradient(svgDefElem *html.Node, color string, lighter bool) string {
	gradientId :=
		"path-fill-gradient" + "-" + color + "-"
	if lighter {
		gradientId += "lighter"
	} else {
		gradientId += "default"
	}
	gradientDef := renderVerticalGradient(svgDefElem, gradientId)
	opacities := []string{"1", "0.6", "0.2"}
	if lighter {
		opacities = []string{"0.4", "0.2", "0"}
	}
	setGradientStop(gradientDef, "0%", color, opacities[0])
	setGradientStop(gradientDef, "50%", color, opacities[1])
	setGradientStop(gradientDef, "100%", color, opacities[2])
	return gradientId
}

func rightRoundedBar(x, width, height int) string {
	radius := height / 2
	xOffset := width - radius
	return fmt.Sprintf("M%d,0 h%d q%d,0 %d,%d q0,%d -%d,%d h-%d v%d",
		x, xOffset, radius, radius, radius, radius, radius, radius, xOffset, height)
}

func leftRoundedBar(x, width, height int) string {
	radius := height / 2
	xOffset := width - radius
	return fmt.Sprintf("M%d,0 h%d v%d h-%d q-%d, 0 -%d,-%d q0,-%d %d,-%dz",
		x+radius, xOffset, height, xOffset, radius, radius, radius, radius, radius, radius)
}

func percentageBar(x, y, width, height int, isFirst, isLast bool, fill string) *html.Node {
	if fill == "" {
		fill = "none"
	}
	if isLast {
		pathStr := rightRoundedBar(x, width, height)
		return makePath(pathStr, pathOpts{
			className: "percentage-bar",
			fill:      fill,
		})
	}
	if isFirst {
		pathStr := leftRoundedBar(x, width, height)
		return makePath(pathStr, pathOpts{
			className: "percentage-bar",
			fill:      fill,
		})
	}
	return createSVG("rect", createOptions{
		attr: []html.Attribute{
			{Key: "class", Val: "percentage-bar"},
			{Key: "x", Val: strconv.Itoa(x)},
			{Key: "y", Val: strconv.Itoa(y)},
			{Key: "width", Val: strconv.Itoa(width)},
			{Key: "height", Val: strconv.Itoa(height)},
			{Key: "fill", Val: fill},
		},
	})
}

func heatSquare(className string, x, y, size, radius int, fill string, data ...html.Attribute) *html.Node {
	if fill == "" {
		fill = "none"
	}
	o := createOptions{
		attr: []html.Attribute{
			{Key: "class", Val: className},
			{Key: "x", Val: strconv.Itoa(x)},
			{Key: "y", Val: strconv.Itoa(y)},
			{Key: "width", Val: strconv.Itoa(size)},
			{Key: "height", Val: strconv.Itoa(size)},
			{Key: "rx", Val: strconv.Itoa(radius)},
			{Key: "fill", Val: fill},
		},
	}
	o.attr = append(o.attr, data...)
	return createSVG("rect", o)
}

func legendDot(x, y, size, radius int, fill, label, value string, fontSize int, truncate bool) *html.Node {
	if fill == "" {
		fill = "none"
	}
	if truncate {
		label = truncateString(label, LABEL_MAX_CHARS)
	}
	if fontSize == 0 {
		fontSize = FONT_SIZE
	}
	o := createOptions{
		attr: []html.Attribute{
			{Key: "class", Val: "legend-dot"},
			{Key: "x", Val: "0"},
			{Key: "y", Val: strconv.Itoa(4 - size)},
			{Key: "height", Val: strconv.Itoa(size)},
			{Key: "width", Val: strconv.Itoa(size)},
			{Key: "rx", Val: strconv.Itoa(radius)},
			{Key: "fill", Val: fill},
		},
	}

	textLabel := createSVG("text", createOptions{
		attr: []html.Attribute{
			{Key: "class", Val: "legend-dataset-label"},
			{Key: "y", Val: "0"},
			{Key: "x", Val: strconv.Itoa(size)},
			{Key: "dx", Val: strconv.Itoa(fontSize) + "px"},
			{Key: "dy", Val: strconv.Itoa(fontSize/3) + "px"},
			{Key: "font-size", Val: strconv.FormatFloat(float64(fontSize)*1.6, 'f', -1, 64) + "px"},
			{Key: "text-anchor", Val: "start"},
		},
		innerHtml: label,
	})
	var textValue *html.Node
	if value != "" {
		textValue = createSVG("text", createOptions{
			attr: []html.Attribute{
				{Key: "class", Val: "legend-dataset-value"},
				{Key: "x", Val: strconv.Itoa(size)},
				{Key: "y", Val: strconv.Itoa(FONT_SIZE + 10)},
				{Key: "dx", Val: strconv.Itoa(FONT_SIZE) + "px"},
				{Key: "dy", Val: strconv.Itoa(FONT_SIZE/3) + "px"},
				{Key: "font-size", Val: strconv.FormatFloat(float64(fontSize)*1.2, 'f', -1, 64) + "px"},
				{Key: "text-anchor", Val: "start"},
			},
			innerHtml: value,
		})
	}
	group := createSVG("g", createOptions{
		attr: []html.Attribute{
			{Key: "transform", Val: fmt.Sprintf("translate(%d, %d)", x, y)},
		},
	})
	group.AppendChild(createSVG("rect", o))
	group.AppendChild(textLabel)
	if textValue != nil {
		group.AppendChild(textLabel)
	}
	return group
}

func truncateString(txt string, n int) string {
	if txt == "" {
		return ""
	}
	if len(txt) > n {
		return txt[:n-3] + "..."
	}
	return txt
}

type textOptions struct {
	fontSize, dy     int
	fill, textAnchor string
}

func makeText(className string, x, y int, content string, o textOptions) *html.Node {
	if o.fontSize == 0 {
		o.fontSize = FONT_SIZE
	}
	if o.dy == 0 {
		o.dy = o.fontSize / 2
	}
	if o.fill == "" {
		o.fill = "var(--charts-label-color)"
	}
	if o.textAnchor == "" {
		o.textAnchor = "start"
	}
	return createSVG("text", createOptions{
		innerHtml: content,
		attr: []html.Attribute{
			{Key: "x", Val: strconv.Itoa(x)},
			{Key: "y", Val: strconv.Itoa(y)},
			{Key: "dy", Val: strconv.Itoa(o.dy) + "px"},
			{Key: "font-size", Val: strconv.Itoa(o.fontSize) + "px"},
			{Key: "fill", Val: o.fill},
			{Key: "text-anchor", Val: o.textAnchor},
		},
	})
}

type verLineOptions struct {
	stroke    string
	className string
}

func makeVertLine(x int, label string, y1, y2 int, o verLineOptions) *html.Node {
	if o.stroke == "" {
		o.stroke = BASE_LINE_COLOR
	}
	l := createSVG("line", createOptions{
		style: map[string]string{
			"stroke": o.stroke,
		},
		attr: []html.Attribute{
			{Key: "class", Val: "line-vertical " + o.className},
			{Key: "x1", Val: "0"},
			{Key: "x2", Val: "0"},
			{Key: "y1", Val: strconv.Itoa(y1)},
			{Key: "y2", Val: strconv.Itoa(y2)},
		},
	})
	y := y1 - LABEL_MARGIN - FONT_SIZE
	if y1 > y2 {
		y = y1 + LABEL_MARGIN
	}
	text := createSVG("text", createOptions{
		innerHtml: label,
		attr: []html.Attribute{
			{Key: "x", Val: "0"},
			{Key: "y", Val: strconv.Itoa(y)},
			{Key: "dy", Val: strconv.Itoa(FONT_SIZE) + "px"},
			{Key: "font-size", Val: strconv.Itoa(FONT_SIZE) + "px"},
			{Key: "text-anchor", Val: "middle"},
		},
	})
	line := createSVG("g", createOptions{
		attr: []html.Attribute{
			{Key: "transform", Val: fmt.Sprintf("translate(%d, 0)", x)},
		},
	})
	line.AppendChild(l)
	line.AppendChild(text)
	return line
}

type horiLineOptions struct {
	className, title, stroke, lineType, alignment string
	shortenNumbers                                bool
}

func makeHoriLine(y int, label any, x1, x2 int, o horiLineOptions) *html.Node {
	if o.stroke == "" {
		o.stroke = BASE_LINE_COLOR
	}
	if o.alignment == "" {
		o.alignment = "left"
	}
	if o.shortenNumbers {
		label = shortenLargeNumber(label)
	}
	className := "line-horizontal " + o.className + o.lineType
	var textXPos int
	if o.alignment == "left" {
		if o.title != "" {
			textXPos = x1 - LABEL_MARGIN + LABEL_WIDTH
		} else {
			textXPos = x1 - LABEL_MARGIN
		}
	} else {
		if o.title != "" {
			textXPos = x2 + LABEL_MARGIN*4 - LABEL_WIDTH
		} else {
			textXPos = x2 + LABEL_MARGIN*4
		}
	}
	lineX1Post := x1
	if o.title != "" {
		lineX1Post = x1 + LABEL_WIDTH
	}
	lineX2Post := x2
	if o.title != "" {
		lineX2Post = x2 - LABEL_WIDTH
	}
	l := createSVG("line", createOptions{
		attr: []html.Attribute{
			{Key: "class", Val: className},
			{Key: "x1", Val: strconv.Itoa(lineX1Post)},
			{Key: "x2", Val: strconv.Itoa(lineX2Post)},
			{Key: "y1", Val: "0"},
			{Key: "y2", Val: "0"},
		},
		style: map[string]string{
			"stroke": o.stroke,
		},
	})
	a := "start"
	if x1 < x2 {
		a = "end"
	}
	text := createSVG("text", createOptions{
		attr: []html.Attribute{
			{Key: "x", Val: strconv.Itoa(textXPos)},
			{Key: "y", Val: "0"},
			{Key: "dy", Val: strconv.Itoa(FONT_SIZE/2-2) + "px"},
			{Key: "font-size", Val: strconv.Itoa(FONT_SIZE) + "px"},
			{Key: "text-anchor", Val: a},
		},
		innerHtml: fmt.Sprint(label),
	})
	line := createSVG("g", createOptions{
		attr: []html.Attribute{
			{Key: "transform", Val: fmt.Sprintf("translate(0, %d)", y)},
			{Key: "stroke-opacity", Val: "1"},
		},
	})
	if text.Data == "0" {
		line.Attr = append(line.Attr, style(map[string]string{
			"stroke": "rgba(27, 31, 35, 0.6)",
		}))
	}
	line.AppendChild(l)
	line.AppendChild(text)
	return line
}
func shortenLargeNumber(a any) string {
	var n float64
	switch e := a.(type) {
	case string:
		if e == "" {
			return ""
		}
		n, _ = strconv.ParseFloat(e, 64)
	case int:
		n = float64(e)
	case float64:
		n = float64(e)
	}
	p := math.Floor(math.Log10(math.Abs(n)))
	if p < 2 {
		return strconv.FormatFloat(n, 'f', -1, 64)
	}
	l := math.Floor(p / 3)
	s := math.Pow(10, p-l*3) * +(n / math.Pow(10, p))
	s = math.Round(s*100) / 100
	return strconv.FormatFloat(s, 'f', -1, 64) + []string{"", "K", "M", "B", "T"}[int(l)]
}
