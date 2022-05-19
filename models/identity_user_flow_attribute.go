package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityUserFlowAttribute provides operations to manage the identityContainer singleton.
type IdentityUserFlowAttribute struct {
    Entity
    // The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string , boolean , int64 , stringCollection , dateTime.
    dataType *IdentityUserFlowAttributeDataType
    // The description of the user flow attribute that's shown to the user at the time of sign-up.
    description *string
    // The display name of the user flow attribute.
    displayName *string
    // The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required.
    userFlowAttributeType *IdentityUserFlowAttributeType
}
// NewIdentityUserFlowAttribute instantiates a new identityUserFlowAttribute and sets the default values.
func NewIdentityUserFlowAttribute()(*IdentityUserFlowAttribute) {
    m := &IdentityUserFlowAttribute{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityUserFlowAttributeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityUserFlowAttributeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.identityUserFlowAttribute":
                        return NewIdentityUserFlowAttribute(), nil
                }
            }
        }
    }
    return NewIdentityUserFlowAttribute(), nil
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
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityUserFlowAttribute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityUserFlowAttributeDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val.(*IdentityUserFlowAttributeDataType))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["userFlowAttributeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetUserFlowAttributeType gets the userFlowAttributeType property value. The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required.
func (m *IdentityUserFlowAttribute) GetUserFlowAttributeType()(*IdentityUserFlowAttributeType) {
    if m == nil {
        return nil
    } else {
        return m.userFlowAttributeType
    }
}
// Serialize serializes information the current object
func (m *IdentityUserFlowAttribute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
