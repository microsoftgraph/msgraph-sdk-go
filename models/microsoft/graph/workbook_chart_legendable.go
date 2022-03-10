package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartLegendable 
type WorkbookChartLegendable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFormat()(WorkbookChartLegendFormatable)
    GetOverlay()(*bool)
    GetPosition()(*string)
    GetVisible()(*bool)
    SetFormat(value WorkbookChartLegendFormatable)()
    SetOverlay(value *bool)()
    SetPosition(value *string)()
    SetVisible(value *bool)()
}
