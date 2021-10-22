package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartDataLabels struct {
    Entity
    format *WorkbookChartDataLabelFormat;
    position *string;
    separator *string;
    showBubbleSize *bool;
    showCategoryName *bool;
    showLegendKey *bool;
    showPercentage *bool;
    showSeriesName *bool;
    showValue *bool;
}
func NewWorkbookChartDataLabels()(*WorkbookChartDataLabels) {
    m := &WorkbookChartDataLabels{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartDataLabels) GetFormat()(*WorkbookChartDataLabelFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *WorkbookChartDataLabels) GetPosition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
func (m *WorkbookChartDataLabels) GetSeparator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.separator
    }
}
func (m *WorkbookChartDataLabels) GetShowBubbleSize()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showBubbleSize
    }
}
func (m *WorkbookChartDataLabels) GetShowCategoryName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showCategoryName
    }
}
func (m *WorkbookChartDataLabels) GetShowLegendKey()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLegendKey
    }
}
func (m *WorkbookChartDataLabels) GetShowPercentage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showPercentage
    }
}
func (m *WorkbookChartDataLabels) GetShowSeriesName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showSeriesName
    }
}
func (m *WorkbookChartDataLabels) GetShowValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showValue
    }
}
func (m *WorkbookChartDataLabels) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartDataLabelFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartDataLabelFormat))
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPosition(val)
        return nil
    }
    res["separator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSeparator(val)
        return nil
    }
    res["showBubbleSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowBubbleSize(val)
        return nil
    }
    res["showCategoryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowCategoryName(val)
        return nil
    }
    res["showLegendKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowLegendKey(val)
        return nil
    }
    res["showPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowPercentage(val)
        return nil
    }
    res["showSeriesName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowSeriesName(val)
        return nil
    }
    res["showValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowValue(val)
        return nil
    }
    return res
}
func (m *WorkbookChartDataLabels) IsNil()(bool) {
    return m == nil
}
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
func (m *WorkbookChartDataLabels) SetFormat(value *WorkbookChartDataLabelFormat)() {
    m.format = value
}
func (m *WorkbookChartDataLabels) SetPosition(value *string)() {
    m.position = value
}
func (m *WorkbookChartDataLabels) SetSeparator(value *string)() {
    m.separator = value
}
func (m *WorkbookChartDataLabels) SetShowBubbleSize(value *bool)() {
    m.showBubbleSize = value
}
func (m *WorkbookChartDataLabels) SetShowCategoryName(value *bool)() {
    m.showCategoryName = value
}
func (m *WorkbookChartDataLabels) SetShowLegendKey(value *bool)() {
    m.showLegendKey = value
}
func (m *WorkbookChartDataLabels) SetShowPercentage(value *bool)() {
    m.showPercentage = value
}
func (m *WorkbookChartDataLabels) SetShowSeriesName(value *bool)() {
    m.showSeriesName = value
}
func (m *WorkbookChartDataLabels) SetShowValue(value *bool)() {
    m.showValue = value
}
