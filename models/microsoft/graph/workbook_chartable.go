package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartable 
type WorkbookChartable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAxes()(WorkbookChartAxesable)
    GetDataLabels()(WorkbookChartDataLabelsable)
    GetFormat()(WorkbookChartAreaFormatable)
    GetHeight()(*float64)
    GetLeft()(*float64)
    GetLegend()(WorkbookChartLegendable)
    GetName()(*string)
    GetSeries()([]WorkbookChartSeriesable)
    GetTitle()(WorkbookChartTitleable)
    GetTop()(*float64)
    GetWidth()(*float64)
    GetWorksheet()(WorkbookWorksheetable)
    SetAxes(value WorkbookChartAxesable)()
    SetDataLabels(value WorkbookChartDataLabelsable)()
    SetFormat(value WorkbookChartAreaFormatable)()
    SetHeight(value *float64)()
    SetLeft(value *float64)()
    SetLegend(value WorkbookChartLegendable)()
    SetName(value *string)()
    SetSeries(value []WorkbookChartSeriesable)()
    SetTitle(value WorkbookChartTitleable)()
    SetTop(value *float64)()
    SetWidth(value *float64)()
    SetWorksheet(value WorkbookWorksheetable)()
}
