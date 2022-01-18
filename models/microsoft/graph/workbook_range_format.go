package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookRangeFormat 
type WorkbookRangeFormat struct {
    Entity
    // Collection of border objects that apply to the overall range selected Read-only.
    borders []WorkbookRangeBorder;
    // Gets or sets the width of all colums within the range. If the column widths are not uniform, null will be returned.
    columnWidth *float64;
    // Returns the fill object defined on the overall range. Read-only.
    fill *WorkbookRangeFill;
    // Returns the font object defined on the overall range selected Read-only.
    font *WorkbookRangeFont;
    // Represents the horizontal alignment for the specified object. Possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed.
    horizontalAlignment *string;
    // Returns the format protection object for a range. Read-only.
    protection *WorkbookFormatProtection;
    // Gets or sets the height of all rows in the range. If the row heights are not uniform null will be returned.
    rowHeight *float64;
    // Represents the vertical alignment for the specified object. Possible values are: Top, Center, Bottom, Justify, Distributed.
    verticalAlignment *string;
    // Indicates if Excel wraps the text in the object. A null value indicates that the entire range doesn't have uniform wrap setting
    wrapText *bool;
}
// NewWorkbookRangeFormat instantiates a new workbookRangeFormat and sets the default values.
func NewWorkbookRangeFormat()(*WorkbookRangeFormat) {
    m := &WorkbookRangeFormat{
        Entity: *NewEntity(),
    }
    return m
}
// GetBorders gets the borders property value. Collection of border objects that apply to the overall range selected Read-only.
func (m *WorkbookRangeFormat) GetBorders()([]WorkbookRangeBorder) {
    if m == nil {
        return nil
    } else {
        return m.borders
    }
}
// GetColumnWidth gets the columnWidth property value. Gets or sets the width of all colums within the range. If the column widths are not uniform, null will be returned.
func (m *WorkbookRangeFormat) GetColumnWidth()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.columnWidth
    }
}
// GetFill gets the fill property value. Returns the fill object defined on the overall range. Read-only.
func (m *WorkbookRangeFormat) GetFill()(*WorkbookRangeFill) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
// GetFont gets the font property value. Returns the font object defined on the overall range selected Read-only.
func (m *WorkbookRangeFormat) GetFont()(*WorkbookRangeFont) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
// GetHorizontalAlignment gets the horizontalAlignment property value. Represents the horizontal alignment for the specified object. Possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed.
func (m *WorkbookRangeFormat) GetHorizontalAlignment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.horizontalAlignment
    }
}
// GetProtection gets the protection property value. Returns the format protection object for a range. Read-only.
func (m *WorkbookRangeFormat) GetProtection()(*WorkbookFormatProtection) {
    if m == nil {
        return nil
    } else {
        return m.protection
    }
}
// GetRowHeight gets the rowHeight property value. Gets or sets the height of all rows in the range. If the row heights are not uniform null will be returned.
func (m *WorkbookRangeFormat) GetRowHeight()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.rowHeight
    }
}
// GetVerticalAlignment gets the verticalAlignment property value. Represents the vertical alignment for the specified object. Possible values are: Top, Center, Bottom, Justify, Distributed.
func (m *WorkbookRangeFormat) GetVerticalAlignment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verticalAlignment
    }
}
// GetWrapText gets the wrapText property value. Indicates if Excel wraps the text in the object. A null value indicates that the entire range doesn't have uniform wrap setting
func (m *WorkbookRangeFormat) GetWrapText()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wrapText
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["borders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeBorder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookRangeBorder, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookRangeBorder))
            }
            m.SetBorders(res)
        }
        return nil
    }
    res["columnWidth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnWidth(val)
        }
        return nil
    }
    res["fill"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeFill() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(*WorkbookRangeFill))
        }
        return nil
    }
    res["font"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeFont() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(*WorkbookRangeFont))
        }
        return nil
    }
    res["horizontalAlignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHorizontalAlignment(val)
        }
        return nil
    }
    res["protection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFormatProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtection(val.(*WorkbookFormatProtection))
        }
        return nil
    }
    res["rowHeight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowHeight(val)
        }
        return nil
    }
    res["verticalAlignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerticalAlignment(val)
        }
        return nil
    }
    res["wrapText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWrapText(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookRangeFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetBorders sets the borders property value. Collection of border objects that apply to the overall range selected Read-only.
func (m *WorkbookRangeFormat) SetBorders(value []WorkbookRangeBorder)() {
    if m != nil {
        m.borders = value
    }
}
// SetColumnWidth sets the columnWidth property value. Gets or sets the width of all colums within the range. If the column widths are not uniform, null will be returned.
func (m *WorkbookRangeFormat) SetColumnWidth(value *float64)() {
    if m != nil {
        m.columnWidth = value
    }
}
// SetFill sets the fill property value. Returns the fill object defined on the overall range. Read-only.
func (m *WorkbookRangeFormat) SetFill(value *WorkbookRangeFill)() {
    if m != nil {
        m.fill = value
    }
}
// SetFont sets the font property value. Returns the font object defined on the overall range selected Read-only.
func (m *WorkbookRangeFormat) SetFont(value *WorkbookRangeFont)() {
    if m != nil {
        m.font = value
    }
}
// SetHorizontalAlignment sets the horizontalAlignment property value. Represents the horizontal alignment for the specified object. Possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed.
func (m *WorkbookRangeFormat) SetHorizontalAlignment(value *string)() {
    if m != nil {
        m.horizontalAlignment = value
    }
}
// SetProtection sets the protection property value. Returns the format protection object for a range. Read-only.
func (m *WorkbookRangeFormat) SetProtection(value *WorkbookFormatProtection)() {
    if m != nil {
        m.protection = value
    }
}
// SetRowHeight sets the rowHeight property value. Gets or sets the height of all rows in the range. If the row heights are not uniform null will be returned.
func (m *WorkbookRangeFormat) SetRowHeight(value *float64)() {
    if m != nil {
        m.rowHeight = value
    }
}
// SetVerticalAlignment sets the verticalAlignment property value. Represents the vertical alignment for the specified object. Possible values are: Top, Center, Bottom, Justify, Distributed.
func (m *WorkbookRangeFormat) SetVerticalAlignment(value *string)() {
    if m != nil {
        m.verticalAlignment = value
    }
}
// SetWrapText sets the wrapText property value. Indicates if Excel wraps the text in the object. A null value indicates that the entire range doesn't have uniform wrap setting
func (m *WorkbookRangeFormat) SetWrapText(value *bool)() {
    if m != nil {
        m.wrapText = value
    }
}
