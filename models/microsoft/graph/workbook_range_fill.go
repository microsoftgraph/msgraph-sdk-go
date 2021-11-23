package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workbookRangeFill 
type WorkbookRangeFill struct {
    Entity
    // HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange')
    color *string;
}
// NewWorkbookRangeFill instantiates a new workbookRangeFill and sets the default values.
func NewWorkbookRangeFill()(*WorkbookRangeFill) {
    m := &WorkbookRangeFill{
        Entity: *NewEntity(),
    }
    return m
}
// GetColor gets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange')
func (m *WorkbookRangeFill) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeFill) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
func (m *WorkbookRangeFill) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookRangeFill) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// SetColor sets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange')
func (m *WorkbookRangeFill) SetColor(value *string)() {
    m.color = value
}
