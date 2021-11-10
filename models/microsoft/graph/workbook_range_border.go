package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookRangeBorder struct {
    Entity
    // HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange').
    color *string;
    // Constant value that indicates the specific side of the border. The possible values are: EdgeTop, EdgeBottom, EdgeLeft, EdgeRight, InsideVertical, InsideHorizontal, DiagonalDown, DiagonalUp. Read-only.
    sideIndex *string;
    // One of the constants of line style specifying the line style for the border. The possible values are: None, Continuous, Dash, DashDot, DashDotDot, Dot, Double, SlantDashDot.
    style *string;
    // Specifies the weight of the border around a range. The possible values are: Hairline, Thin, Medium, Thick.
    weight *string;
}
// Instantiates a new workbookRangeBorder and sets the default values.
func NewWorkbookRangeBorder()(*WorkbookRangeBorder) {
    m := &WorkbookRangeBorder{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange').
func (m *WorkbookRangeBorder) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// Gets the sideIndex property value. Constant value that indicates the specific side of the border. The possible values are: EdgeTop, EdgeBottom, EdgeLeft, EdgeRight, InsideVertical, InsideHorizontal, DiagonalDown, DiagonalUp. Read-only.
func (m *WorkbookRangeBorder) GetSideIndex()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sideIndex
    }
}
// Gets the style property value. One of the constants of line style specifying the line style for the border. The possible values are: None, Continuous, Dash, DashDot, DashDotDot, Dot, Double, SlantDashDot.
func (m *WorkbookRangeBorder) GetStyle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.style
    }
}
// Gets the weight property value. Specifies the weight of the border around a range. The possible values are: Hairline, Thin, Medium, Thick.
func (m *WorkbookRangeBorder) GetWeight()(*string) {
    if m == nil {
        return nil
    } else {
        return m.weight
    }
}
// The deserialization information for the current model
func (m *WorkbookRangeBorder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["sideIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSideIndex(val)
        }
        return nil
    }
    res["style"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStyle(val)
        }
        return nil
    }
    res["weight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookRangeBorder) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookRangeBorder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sideIndex", m.GetSideIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("style", m.GetStyle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("weight", m.GetWeight())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange').
// Parameters:
//  - value : Value to set for the color property.
func (m *WorkbookRangeBorder) SetColor(value *string)() {
    m.color = value
}
// Sets the sideIndex property value. Constant value that indicates the specific side of the border. The possible values are: EdgeTop, EdgeBottom, EdgeLeft, EdgeRight, InsideVertical, InsideHorizontal, DiagonalDown, DiagonalUp. Read-only.
// Parameters:
//  - value : Value to set for the sideIndex property.
func (m *WorkbookRangeBorder) SetSideIndex(value *string)() {
    m.sideIndex = value
}
// Sets the style property value. One of the constants of line style specifying the line style for the border. The possible values are: None, Continuous, Dash, DashDot, DashDotDot, Dot, Double, SlantDashDot.
// Parameters:
//  - value : Value to set for the style property.
func (m *WorkbookRangeBorder) SetStyle(value *string)() {
    m.style = value
}
// Sets the weight property value. Specifies the weight of the border around a range. The possible values are: Hairline, Thin, Medium, Thick.
// Parameters:
//  - value : Value to set for the weight property.
func (m *WorkbookRangeBorder) SetWeight(value *string)() {
    m.weight = value
}
