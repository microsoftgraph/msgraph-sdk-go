package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartDataLabels struct {
    Entity
    // Represents the format of chart data labels, which includes fill and font formatting. Read-only.
    format *WorkbookChartDataLabelFormat;
    // DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
    position *string;
    // String representing the separator used for the data labels on a chart.
    separator *string;
    // Boolean value representing if the data label bubble size is visible or not.
    showBubbleSize *bool;
    // Boolean value representing if the data label category name is visible or not.
    showCategoryName *bool;
    // Boolean value representing if the data label legend key is visible or not.
    showLegendKey *bool;
    // Boolean value representing if the data label percentage is visible or not.
    showPercentage *bool;
    // Boolean value representing if the data label series name is visible or not.
    showSeriesName *bool;
    // Boolean value representing if the data label value is visible or not.
    showValue *bool;
}
// Instantiates a new workbookChartDataLabels and sets the default values.
func NewWorkbookChartDataLabels()(*WorkbookChartDataLabels) {
    m := &WorkbookChartDataLabels{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the format property value. Represents the format of chart data labels, which includes fill and font formatting. Read-only.
func (m *WorkbookChartDataLabels) GetFormat()(*WorkbookChartDataLabelFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the position property value. DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
func (m *WorkbookChartDataLabels) GetPosition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// Gets the separator property value. String representing the separator used for the data labels on a chart.
func (m *WorkbookChartDataLabels) GetSeparator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.separator
    }
}
// Gets the showBubbleSize property value. Boolean value representing if the data label bubble size is visible or not.
func (m *WorkbookChartDataLabels) GetShowBubbleSize()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showBubbleSize
    }
}
// Gets the showCategoryName property value. Boolean value representing if the data label category name is visible or not.
func (m *WorkbookChartDataLabels) GetShowCategoryName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showCategoryName
    }
}
// Gets the showLegendKey property value. Boolean value representing if the data label legend key is visible or not.
func (m *WorkbookChartDataLabels) GetShowLegendKey()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLegendKey
    }
}
// Gets the showPercentage property value. Boolean value representing if the data label percentage is visible or not.
func (m *WorkbookChartDataLabels) GetShowPercentage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showPercentage
    }
}
// Gets the showSeriesName property value. Boolean value representing if the data label series name is visible or not.
func (m *WorkbookChartDataLabels) GetShowSeriesName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showSeriesName
    }
}
// Gets the showValue property value. Boolean value representing if the data label value is visible or not.
func (m *WorkbookChartDataLabels) GetShowValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showValue
    }
}
// The deserialization information for the current model
func (m *WorkbookChartDataLabels) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartDataLabelFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*WorkbookChartDataLabelFormat))
        }
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["separator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeparator(val)
        }
        return nil
    }
    res["showBubbleSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowBubbleSize(val)
        }
        return nil
    }
    res["showCategoryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowCategoryName(val)
        }
        return nil
    }
    res["showLegendKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLegendKey(val)
        }
        return nil
    }
    res["showPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowPercentage(val)
        }
        return nil
    }
    res["showSeriesName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowSeriesName(val)
        }
        return nil
    }
    res["showValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowValue(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookChartDataLabels) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartDataLabels) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("separator", m.GetSeparator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showBubbleSize", m.GetShowBubbleSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showCategoryName", m.GetShowCategoryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showLegendKey", m.GetShowLegendKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showPercentage", m.GetShowPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showSeriesName", m.GetShowSeriesName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showValue", m.GetShowValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the format property value. Represents the format of chart data labels, which includes fill and font formatting. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChartDataLabels) SetFormat(value *WorkbookChartDataLabelFormat)() {
    m.format = value
}
// Sets the position property value. DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
// Parameters:
//  - value : Value to set for the position property.
func (m *WorkbookChartDataLabels) SetPosition(value *string)() {
    m.position = value
}
// Sets the separator property value. String representing the separator used for the data labels on a chart.
// Parameters:
//  - value : Value to set for the separator property.
func (m *WorkbookChartDataLabels) SetSeparator(value *string)() {
    m.separator = value
}
// Sets the showBubbleSize property value. Boolean value representing if the data label bubble size is visible or not.
// Parameters:
//  - value : Value to set for the showBubbleSize property.
func (m *WorkbookChartDataLabels) SetShowBubbleSize(value *bool)() {
    m.showBubbleSize = value
}
// Sets the showCategoryName property value. Boolean value representing if the data label category name is visible or not.
// Parameters:
//  - value : Value to set for the showCategoryName property.
func (m *WorkbookChartDataLabels) SetShowCategoryName(value *bool)() {
    m.showCategoryName = value
}
// Sets the showLegendKey property value. Boolean value representing if the data label legend key is visible or not.
// Parameters:
//  - value : Value to set for the showLegendKey property.
func (m *WorkbookChartDataLabels) SetShowLegendKey(value *bool)() {
    m.showLegendKey = value
}
// Sets the showPercentage property value. Boolean value representing if the data label percentage is visible or not.
// Parameters:
//  - value : Value to set for the showPercentage property.
func (m *WorkbookChartDataLabels) SetShowPercentage(value *bool)() {
    m.showPercentage = value
}
// Sets the showSeriesName property value. Boolean value representing if the data label series name is visible or not.
// Parameters:
//  - value : Value to set for the showSeriesName property.
func (m *WorkbookChartDataLabels) SetShowSeriesName(value *bool)() {
    m.showSeriesName = value
}
// Sets the showValue property value. Boolean value representing if the data label value is visible or not.
// Parameters:
//  - value : Value to set for the showValue property.
func (m *WorkbookChartDataLabels) SetShowValue(value *bool)() {
    m.showValue = value
}
