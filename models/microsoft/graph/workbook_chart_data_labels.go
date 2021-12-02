package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartDataLabels 
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
// NewWorkbookChartDataLabels instantiates a new workbookChartDataLabels and sets the default values.
func NewWorkbookChartDataLabels()(*WorkbookChartDataLabels) {
    m := &WorkbookChartDataLabels{
        Entity: *NewEntity(),
    }
    return m
}
// GetFormat gets the format property value. Represents the format of chart data labels, which includes fill and font formatting. Read-only.
func (m *WorkbookChartDataLabels) GetFormat()(*WorkbookChartDataLabelFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetPosition gets the position property value. DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
func (m *WorkbookChartDataLabels) GetPosition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// GetSeparator gets the separator property value. String representing the separator used for the data labels on a chart.
func (m *WorkbookChartDataLabels) GetSeparator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.separator
    }
}
// GetShowBubbleSize gets the showBubbleSize property value. Boolean value representing if the data label bubble size is visible or not.
func (m *WorkbookChartDataLabels) GetShowBubbleSize()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showBubbleSize
    }
}
// GetShowCategoryName gets the showCategoryName property value. Boolean value representing if the data label category name is visible or not.
func (m *WorkbookChartDataLabels) GetShowCategoryName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showCategoryName
    }
}
// GetShowLegendKey gets the showLegendKey property value. Boolean value representing if the data label legend key is visible or not.
func (m *WorkbookChartDataLabels) GetShowLegendKey()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLegendKey
    }
}
// GetShowPercentage gets the showPercentage property value. Boolean value representing if the data label percentage is visible or not.
func (m *WorkbookChartDataLabels) GetShowPercentage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showPercentage
    }
}
// GetShowSeriesName gets the showSeriesName property value. Boolean value representing if the data label series name is visible or not.
func (m *WorkbookChartDataLabels) GetShowSeriesName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showSeriesName
    }
}
// GetShowValue gets the showValue property value. Boolean value representing if the data label value is visible or not.
func (m *WorkbookChartDataLabels) GetShowValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetFormat sets the format property value. Represents the format of chart data labels, which includes fill and font formatting. Read-only.
func (m *WorkbookChartDataLabels) SetFormat(value *WorkbookChartDataLabelFormat)() {
    if m != nil {
        m.format = value
    }
}
// SetPosition sets the position property value. DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
func (m *WorkbookChartDataLabels) SetPosition(value *string)() {
    if m != nil {
        m.position = value
    }
}
// SetSeparator sets the separator property value. String representing the separator used for the data labels on a chart.
func (m *WorkbookChartDataLabels) SetSeparator(value *string)() {
    if m != nil {
        m.separator = value
    }
}
// SetShowBubbleSize sets the showBubbleSize property value. Boolean value representing if the data label bubble size is visible or not.
func (m *WorkbookChartDataLabels) SetShowBubbleSize(value *bool)() {
    if m != nil {
        m.showBubbleSize = value
    }
}
// SetShowCategoryName sets the showCategoryName property value. Boolean value representing if the data label category name is visible or not.
func (m *WorkbookChartDataLabels) SetShowCategoryName(value *bool)() {
    if m != nil {
        m.showCategoryName = value
    }
}
// SetShowLegendKey sets the showLegendKey property value. Boolean value representing if the data label legend key is visible or not.
func (m *WorkbookChartDataLabels) SetShowLegendKey(value *bool)() {
    if m != nil {
        m.showLegendKey = value
    }
}
// SetShowPercentage sets the showPercentage property value. Boolean value representing if the data label percentage is visible or not.
func (m *WorkbookChartDataLabels) SetShowPercentage(value *bool)() {
    if m != nil {
        m.showPercentage = value
    }
}
// SetShowSeriesName sets the showSeriesName property value. Boolean value representing if the data label series name is visible or not.
func (m *WorkbookChartDataLabels) SetShowSeriesName(value *bool)() {
    if m != nil {
        m.showSeriesName = value
    }
}
// SetShowValue sets the showValue property value. Boolean value representing if the data label value is visible or not.
func (m *WorkbookChartDataLabels) SetShowValue(value *bool)() {
    if m != nil {
        m.showValue = value
    }
}
