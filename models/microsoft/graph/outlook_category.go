package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OutlookCategory struct {
    Entity
    // A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
    color *CategoryColor;
    // A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
    displayName *string;
}
// Instantiates a new outlookCategory and sets the default values.
func NewOutlookCategory()(*OutlookCategory) {
    m := &OutlookCategory{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the color property value. A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
func (m *OutlookCategory) GetColor()(*CategoryColor) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// Gets the displayName property value. A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
func (m *OutlookCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// The deserialization information for the current model
func (m *OutlookCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCategoryColor)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CategoryColor)
            m.SetColor(&cast)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *OutlookCategory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OutlookCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColor() != nil {
        cast := m.GetColor().String()
        err = writer.WriteStringValue("color", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the color property value. A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
// Parameters:
//  - value : Value to set for the color property.
func (m *OutlookCategory) SetColor(value *CategoryColor)() {
    m.color = value
}
// Sets the displayName property value. A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *OutlookCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
