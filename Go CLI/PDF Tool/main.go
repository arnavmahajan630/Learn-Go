package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
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
			m.Text("Prepared with Love by Ocean Whisperer", props.Text{
				Top:    3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getColor(),
			})
		})
	})
}

func getColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
