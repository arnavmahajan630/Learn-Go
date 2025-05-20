package main

import (
	"fmt"
	"log"

	data "github.com/Ocean-Whisperer/Learn-Go/Go-CLI/PDF-Tool/Data"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	BuildHeader(m)
	BuildPdfList(m)
	err := m.OutputFileAndClose("pdfs/example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PDF generated successfully!")

}

func BuildHeader(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/logo_pdf.png", props.Rect{Center: true, Percent: 75})
				if err != nil {
					fmt.Println("Error:", err)
				}
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("With Love By Ocean Whisperer", props.Text{
				Style: consts.Bold,
				Align: consts.Center,
				Color: getColor(),
			})
		})
	})
}

func getColor() color.Color {
	return color.Color{
		Red:   25,
		Green: 25,
		Blue:  112,
	}
}

func BuildPdfList(m pdf.Maroto) {
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Simple Products Table", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	lightpurp := getLightPurpleColor()
	tableheadings := []string{"Product", "Description", "Price"}
	contents := data.FruitList(20) 
	m.TableList(tableheadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightpurp,
	})
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 128,
		Blue:  128,
	}
}
