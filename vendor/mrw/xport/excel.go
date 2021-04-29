package xport

import (
	"github.com/tealeg/xlsx"
)

type ExcelWorkbook struct {
	*xlsx.File
}

type ExcelSheet struct {
	*xlsx.Sheet
}

func NewExcelWorkbook() *ExcelWorkbook {
	wb := &ExcelWorkbook{
		File: xlsx.NewFile(),
	}
	return wb
}

func (e *ExcelWorkbook) MustGetSheet(name string) *ExcelSheet {
	var err error
	s, ok := e.Sheet[name]
	if !ok {
		s, err = e.AddSheet(name)
		if err != nil {
			panic(err)
		}
	}
	return &ExcelSheet{Sheet: s}
}

type Cell struct {
	Value  interface{}
	Format string
}

func (s *ExcelSheet) AddVariedRow(cells []Cell) {
	r := s.Sheet.AddRow()
	for _, c := range cells {
		cell := r.AddCell()
		v := c.Value
		cell.SetValue(v)
	}
}

func (s *ExcelSheet) AddStringRow(values []string) {
	r := s.Sheet.AddRow()
	for _, v := range values {
		r.AddCell().SetString(v)
	}
}

func (s *ExcelSheet) AddHeader(values []string) {
	r := s.Sheet.AddRow()
	style := xlsx.NewStyle()
	style.Border.Top = "thin"
	style.Border.TopColor = "00000000"
	style.Border.Bottom = "thin"
	style.Border.BottomColor = "00000000"
	style.Border.Left = "thin"
	style.Border.LeftColor = "00000000"
	style.Border.Right = "thin"
	style.Border.RightColor = "00000000"
	style.ApplyBorder = true
	for _, v := range values {
		c := r.AddCell()
		c.SetValue(v)
		c.SetStyle(style)
	}
}

func (s *ExcelSheet) AddTitle(title string, fontsize int) {
	r := s.Sheet.AddRow()
	style := xlsx.NewStyle()
	style.Font.Bold = true
	style.Font.Size = fontsize
	style.ApplyFont = true
	c := r.AddCell()
	c.SetValue(title)
	c.SetStyle(style)
}
