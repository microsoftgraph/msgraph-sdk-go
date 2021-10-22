package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartLegendFormat struct {
    Entity
    fill *WorkbookChartFill;
    font *WorkbookChartFont;
}
func NewWorkbookChartLegendFormat()(*WorkbookChartLegendFormat) {
    m := &WorkbookChartLegendFormat{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartLegendFormat) GetFill()(*WorkbookChartFill) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
func (m *WorkbookChartLegendFormat) GetFont()(*WorkbookChartFont) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
func (m *WorkbookChartLegendFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartFill() })
        if err != nil {
            return err
        }
        m.SetFill(val.(*WorkbookChartFill))
        return nil
    }
    res["font"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartFont() })
        if err != nil {
            return err
        }
        m.SetFont(val.(*WorkbookChartFont))
        return nil
    }
    return res
}
func (m *WorkbookChartLegendFormat) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookChartLegendFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fill", m.GetFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookChartLegendFormat) SetFill(value *WorkbookChartFill)() {
    m.fill = value
}
func (m *WorkbookChartLegendFormat) SetFont(value *WorkbookChartFont)() {
    m.font = value
}
