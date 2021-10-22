package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartAxisTitle struct {
    Entity
    format *WorkbookChartAxisTitleFormat;
    text *string;
    visible *bool;
}
func NewWorkbookChartAxisTitle()(*WorkbookChartAxisTitle) {
    m := &WorkbookChartAxisTitle{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartAxisTitle) GetFormat()(*WorkbookChartAxisTitleFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *WorkbookChartAxisTitle) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
func (m *WorkbookChartAxisTitle) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
func (m *WorkbookChartAxisTitle) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxisTitleFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartAxisTitleFormat))
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetText(val)
        return nil
    }
    res["visible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetVisible(val)
        return nil
    }
    return res
}
func (m *WorkbookChartAxisTitle) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookChartAxisTitle) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visible", m.GetVisible())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookChartAxisTitle) SetFormat(value *WorkbookChartAxisTitleFormat)() {
    m.format = value
}
func (m *WorkbookChartAxisTitle) SetText(value *string)() {
    m.text = value
}
func (m *WorkbookChartAxisTitle) SetVisible(value *bool)() {
    m.visible = value
}
