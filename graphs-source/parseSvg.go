package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

// Most of these "omitempty" declarations probably aren't needed
// because I changed the values to be pointers.

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	NS      string   `xml:"xmlns,attr"`
	// This doesn't quite work, but it's good enough. (It adds
	// an empty ns declaration, but that doesn't seem to affect anything.)
	LinkNS  string `xml:"http://www.w3.org/1999/xlink xlink,attr"`
	Height  string `xml:"height,attr"`
	Version string `xml:"version,attr"`
	ViewBox string `xml:"viewBox,attr"`
	Width   string `xml:"width,attr"`
	Defs    *Defs  `xml:"defs,omitempty"`
	G       []*G   `xml:"g,omitempty"`
}
type Defs struct {
	Paths []Path `xml:"path"`
}

type Path struct {
	D     string `xml:"d,attr"`
	ID    string `xml:"id,attr"`
	Style string `xml:"style,attr,omitempty"`
}

type Polyline struct {
	XMLName xml.Name `xml:"polyline,omitempty"`
	Fill    string   `xml:"fill,attr,omitempty"`
	Stroke  string   `xml:"stroke,attr,omitempty"`
	Points  string   `xml:"points,attr,omitempty"`
}

type G struct {
	ID         string  `xml:"id,attr,omitempty"`
	Transform  string  `xml:"transform,attr,omitempty"`
	Rects      []*Rect `xml:"rect,omitempty"`
	Uses       []Use   `xml:"use,omitempty"`
	OutputUses []OutputUse
	Polyline   *Polyline `xml:"polyline,omitempty"`
	G          []*G      `xml:"g,omitempty"`
	Path       []*Path   `xml:"path,omitempty"`
}

type Use struct {
	X    string `xml:"x,attr"`
	Y    string `xml:"y,attr"`
	Href string `xml:"href,attr"`
}

// Used for writing "use" elements XML because xlink:href gets lost on
// import with the above struct.
type OutputUse struct {
	XMLName xml.Name `xml:"use"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
	Href    string   `xml:"xlink:href,attr"`
}

type Rect struct {
	Height string `xml:"height,attr"`
	Width  string `xml:"width,attr"`
	X      string `xml:"x,attr"`
	Y      string `xml:"y,attr"`
	Style  string `xml:"style,attr,omitempty"`
}

func main() {
	xmlFile, err := os.Open("../AD/thvol-adjusted.svg")

	if err != nil {
		log.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var graph Svg
	xml.Unmarshal(byteValue, &graph)

	graphGroup, err := findGraphGroup(graph.G)

	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(graph.G[graphGroup].Rects, func(i, j int) bool {
		a, _ := strconv.ParseFloat(graph.G[graphGroup].Rects[i].X, 32)
		b, _ := strconv.ParseFloat(graph.G[graphGroup].Rects[j].X, 32)

		return a < b
	})
	convertToPolyline(graph.G[graphGroup])

	byteValue, _ = xml.MarshalIndent(graph, "    ", "    ")

	err = os.WriteFile("../AD/thvol-polyline.svg", byteValue, 0666)
	if err != nil {
		log.Fatal(err)
	}

}

// Converts the colection of gnuplot rectangles to a polyline by using the
// center of each rectangle. It's hacky, but it works better than other
// solutions tried including using rectangle size changes to find slope
// changes and plot points there.

func convertToPolyline(group *G) {

	var polyline Polyline

	polyline.Fill = "none"
	polyline.Stroke = "red"

	for _, rect := range group.Rects {

		yFloat, _ := strconv.ParseFloat(rect.Y, 32)
		xFloat, _ := strconv.ParseFloat(rect.X, 32)

		heightFloat, _ := strconv.ParseFloat(rect.Height, 32)
		widthFloat, _ := strconv.ParseFloat(rect.Width, 32)

		centerX := xFloat + widthFloat/2.0
		centerY := yFloat + heightFloat/2.0
		polyline.Points += fmt.Sprintf("%.5f,%.5f ", centerX, centerY)
	}

	group.Polyline = &polyline
}

// Find the group in the .svg file called "graph" because that's the
// group that contains the rectangles that make up the plotted line
// from gnuplot.

func findGraphGroup(groups []*G) (foundIndex int, err error) {
	var index int
	var group *G

	for index, group = range groups {
		if (*group).ID == "graph" {
			return index, nil
		}
	}
	return 0, errors.New("could not find graph group")
}
