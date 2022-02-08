package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityUserFlowAttribute 
type IdentityUserFlowAttribute struct {
    Entity
    // The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string , boolean , int64 , stringCollection , dateTime.
    dataType *IdentityUserFlowAttributeDataType;
    // The description of the user flow attribute that's shown to the user at the time of sign-up.
    description *string;
    // The display name of the user flow attribute.
    displayName *string;
    // The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required.
    userFlowAttributeType *IdentityUserFlowAttributeType;
}
// NewIdentityUserFlowAttribute instantiates a new identityUserFlowAttribute and sets the default values.
func NewIdentityUserFlowAttribute()(*IdentityUserFlowAttribute) {
    m := &IdentityUserFlowAttribute{
        Entity: *NewEntity(),
    }
    return m
}
// GetDataType gets the dataType property value. The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string , boolean , int64 , stringCollection , dateTime.
func (m *IdentityUserFlowAttribute) GetDataType()(*IdentityUserFlowAttributeDataType) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
// GetDescription gets the description property value. The description of the user flow attribute that's shown to the user at the time of sign-up.
func (m *IdentityUserFlowAttribute) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the user flow attribute.
func (m *IdentityUserFlowAttribute) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetUserFlowAttributeType gets the userFlowAttributeType property value. The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required.
func (m *IdentityUserFlowAttribute) GetUserFlowAttributeType()(*IdentityUserFlowAttributeType) {
    if m == nil {
        return nil
    } else {
        return m.userFlowAttributeType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityUserFlowAttribute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityUserFlowAttributeDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val.(*IdentityUserFlowAttributeDataType))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["userFlowAttributeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityUserFlowAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserFlowAttributeType(val.(*IdentityUserFlowAttributeType))
        }
        return nil
    }
    return res
}
func (m *IdentityUserFlowAttribute) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IdentityUserFlowAttribute) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDataType() != nil {
        cast := (*m.GetDataType()).String()
        err = writer.WriteStringValue("dataType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetUserFlowAttributeType() != nil {
        cast := (*m.GetUserFlowAttributeType()).String()
        err = writer.WriteStringValue("userFlowAttributeType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDataType sets the dataType property value. The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string , boolean , int64 , stringCollection , dateTime.
func (m *IdentityUserFlowAttribute) SetDataType(value *IdentityUserFlowAttributeDataType)() {
    if m != nil {
        m.dataType = value
    }
}
// SetDescription sets the description property value. The description of the user flow attribute that's shown to the user at the time of sign-up.
func (m *IdentityUserFlowAttribute) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the user flow attribute.
func (m *IdentityUserFlowAttribute) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetUserFlowAttributeType sets the userFlowAttributeType property value. The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required.
func (m *IdentityUserFlowAttribute) SetUserFlowAttributeType(value *IdentityUserFlowAttributeType)() {
    if m != nil {
        m.userFlowAttributeType = value
    }
}
