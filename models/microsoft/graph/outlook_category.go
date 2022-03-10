package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookCategory provides operations to manage the collection of drive entities.
type OutlookCategory struct {
    Entity
    // A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
    color *CategoryColor;
    // A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
    displayName *string;
}
// NewOutlookCategory instantiates a new outlookCategory and sets the default values.
func NewOutlookCategory()(*OutlookCategory) {
    m := &OutlookCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOutlookCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookCategoryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOutlookCategory(), nil
}
// GetColor gets the color property value. A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
func (m *OutlookCategory) GetColor()(*CategoryColor) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetDisplayName gets the displayName property value. A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
func (m *OutlookCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCategoryColor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*CategoryColor))
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
// Serialize serializes information the current object
func (m *OutlookCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
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
// SetColor sets the color property value. A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
func (m *OutlookCategory) SetColor(value *CategoryColor)() {
    if m != nil {
        m.color = value
    }
}
// SetDisplayName sets the displayName property value. A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
func (m *OutlookCategory) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
