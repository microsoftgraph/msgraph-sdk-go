package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartDataLabelsable 
type WorkbookChartDataLabelsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetFormat()(WorkbookChartDataLabelFormatable)
    GetPosition()(*string)
    GetSeparator()(*string)
    GetShowBubbleSize()(*bool)
    GetShowCategoryName()(*bool)
    GetShowLegendKey()(*bool)
    GetShowPercentage()(*bool)
    GetShowSeriesName()(*bool)
    GetShowValue()(*bool)
    SetFormat(value WorkbookChartDataLabelFormatable)()
    SetPosition(value *string)()
    SetSeparator(value *string)()
    SetShowBubbleSize(value *bool)()
    SetShowCategoryName(value *bool)()
    SetShowLegendKey(value *bool)()
    SetShowPercentage(value *bool)()
    SetShowSeriesName(value *bool)()
    SetShowValue(value *bool)()
}
