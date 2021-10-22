package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartAxis struct {
    Entity
    format *WorkbookChartAxisFormat;
    majorGridlines *WorkbookChartGridlines;
    majorUnit *Json;
    maximum *Json;
    minimum *Json;
    minorGridlines *WorkbookChartGridlines;
    minorUnit *Json;
    title *WorkbookChartAxisTitle;
}
func NewWorkbookChartAxis()(*WorkbookChartAxis) {
    m := &WorkbookChartAxis{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartAxis) GetFormat()(*WorkbookChartAxisFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *WorkbookChartAxis) GetMajorGridlines()(*WorkbookChartGridlines) {
    if m == nil {
        return nil
    } else {
        return m.majorGridlines
    }
}
func (m *WorkbookChartAxis) GetMajorUnit()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.majorUnit
    }
}
func (m *WorkbookChartAxis) GetMaximum()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.maximum
    }
}
func (m *WorkbookChartAxis) GetMinimum()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.minimum
    }
}
func (m *WorkbookChartAxis) GetMinorGridlines()(*WorkbookChartGridlines) {
    if m == nil {
        return nil
    } else {
        return m.minorGridlines
    }
}
func (m *WorkbookChartAxis) GetMinorUnit()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.minorUnit
    }
}
func (m *WorkbookChartAxis) GetTitle()(*WorkbookChartAxisTitle) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *WorkbookChartAxis) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxisFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartAxisFormat))
        return nil
    }
    res["majorGridlines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartGridlines() })
        if err != nil {
            return err
        }
        m.SetMajorGridlines(val.(*WorkbookChartGridlines))
        return nil
    }
    res["majorUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetMajorUnit(val.(*Json))
        return nil
    }
    res["maximum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetMaximum(val.(*Json))
        return nil
    }
    res["minimum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetMinimum(val.(*Json))
        return nil
    }
    res["minorGridlines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartGridlines() })
        if err != nil {
            return err
        }
        m.SetMinorGridlines(val.(*WorkbookChartGridlines))
        return nil
    }
    res["minorUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetMinorUnit(val.(*Json))
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxisTitle() })
        if err != nil {
            return err
        }
        m.SetTitle(val.(*WorkbookChartAxisTitle))
        return nil
    }
    return res
}
func (m *WorkbookChartAxis) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookChartAxis) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("majorGridlines", m.GetMajorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorUnit", m.GetMajorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimum", m.GetMinimum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorGridlines", m.GetMinorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorUnit", m.GetMinorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookChartAxis) SetFormat(value *WorkbookChartAxisFormat)() {
    m.format = value
}
func (m *WorkbookChartAxis) SetMajorGridlines(value *WorkbookChartGridlines)() {
    m.majorGridlines = value
}
func (m *WorkbookChartAxis) SetMajorUnit(value *Json)() {
    m.majorUnit = value
}
func (m *WorkbookChartAxis) SetMaximum(value *Json)() {
    m.maximum = value
}
func (m *WorkbookChartAxis) SetMinimum(value *Json)() {
    m.minimum = value
}
func (m *WorkbookChartAxis) SetMinorGridlines(value *WorkbookChartGridlines)() {
    m.minorGridlines = value
}
func (m *WorkbookChartAxis) SetMinorUnit(value *Json)() {
    m.minorUnit = value
}
func (m *WorkbookChartAxis) SetTitle(value *WorkbookChartAxisTitle)() {
    m.title = value
}
