package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookRangeFormat struct {
    Entity
    borders []WorkbookRangeBorder;
    columnWidth *float64;
    fill *WorkbookRangeFill;
    font *WorkbookRangeFont;
    horizontalAlignment *string;
    protection *WorkbookFormatProtection;
    rowHeight *float64;
    verticalAlignment *string;
    wrapText *bool;
}
func NewWorkbookRangeFormat()(*WorkbookRangeFormat) {
    m := &WorkbookRangeFormat{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookRangeFormat) GetBorders()([]WorkbookRangeBorder) {
    if m == nil {
        return nil
    } else {
        return m.borders
    }
}
func (m *WorkbookRangeFormat) GetColumnWidth()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.columnWidth
    }
}
func (m *WorkbookRangeFormat) GetFill()(*WorkbookRangeFill) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
func (m *WorkbookRangeFormat) GetFont()(*WorkbookRangeFont) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
func (m *WorkbookRangeFormat) GetHorizontalAlignment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.horizontalAlignment
    }
}
func (m *WorkbookRangeFormat) GetProtection()(*WorkbookFormatProtection) {
    if m == nil {
        return nil
    } else {
        return m.protection
    }
}
func (m *WorkbookRangeFormat) GetRowHeight()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.rowHeight
    }
}
func (m *WorkbookRangeFormat) GetVerticalAlignment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verticalAlignment
    }
}
func (m *WorkbookRangeFormat) GetWrapText()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wrapText
    }
}
func (m *WorkbookRangeFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["borders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeBorder() })
        if err != nil {
            return err
        }
        res := make([]WorkbookRangeBorder, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookRangeBorder))
        }
        m.SetBorders(res)
        return nil
    }
    res["columnWidth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetColumnWidth(val)
        return nil
    }
    res["fill"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeFill() })
        if err != nil {
            return err
        }
        m.SetFill(val.(*WorkbookRangeFill))
        return nil
    }
    res["font"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeFont() })
        if err != nil {
            return err
        }
        m.SetFont(val.(*WorkbookRangeFont))
        return nil
    }
    res["horizontalAlignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHorizontalAlignment(val)
        return nil
    }
    res["protection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFormatProtection() })
        if err != nil {
            return err
        }
        m.SetProtection(val.(*WorkbookFormatProtection))
        return nil
    }
    res["rowHeight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetRowHeight(val)
        return nil
    }
    res["verticalAlignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVerticalAlignment(val)
        return nil
    }
    res["wrapText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWrapText(val)
        return nil
    }
    return res
}
func (m *WorkbookRangeFormat) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookRangeFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBorders()))
        for i, v := range m.GetBorders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("borders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("columnWidth", m.GetColumnWidth())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteStringValue("horizontalAlignment", m.GetHorizontalAlignment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("protection", m.GetProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("rowHeight", m.GetRowHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("verticalAlignment", m.GetVerticalAlignment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wrapText", m.GetWrapText())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookRangeFormat) SetBorders(value []WorkbookRangeBorder)() {
    m.borders = value
}
func (m *WorkbookRangeFormat) SetColumnWidth(value *float64)() {
    m.columnWidth = value
}
func (m *WorkbookRangeFormat) SetFill(value *WorkbookRangeFill)() {
    m.fill = value
}
func (m *WorkbookRangeFormat) SetFont(value *WorkbookRangeFont)() {
    m.font = value
}
func (m *WorkbookRangeFormat) SetHorizontalAlignment(value *string)() {
    m.horizontalAlignment = value
}
func (m *WorkbookRangeFormat) SetProtection(value *WorkbookFormatProtection)() {
    m.protection = value
}
func (m *WorkbookRangeFormat) SetRowHeight(value *float64)() {
    m.rowHeight = value
}
func (m *WorkbookRangeFormat) SetVerticalAlignment(value *string)() {
    m.verticalAlignment = value
}
func (m *WorkbookRangeFormat) SetWrapText(value *bool)() {
    m.wrapText = value
}
