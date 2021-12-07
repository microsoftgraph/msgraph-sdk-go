package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TargetResource 
type TargetResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates the visible name defined for the resource. Typically specified when the resource is created.
    displayName *string;
    // When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
    groupType *GroupType;
    // Indicates the unique ID of the resource.
    id *string;
    // Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
    modifiedProperties []ModifiedProperty;
    // Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
    type_escaped *string;
    // When type is set to User, this includes the user name that initiated the action; null for other types.
    userPrincipalName *string;
}
// NewTargetResource instantiates a new targetResource and sets the default values.
func NewTargetResource()(*TargetResource) {
    m := &TargetResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Indicates the visible name defined for the resource. Typically specified when the resource is created.
func (m *TargetResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGroupType gets the groupType property value. When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
func (m *TargetResource) GetGroupType()(*GroupType) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
// GetId gets the id property value. Indicates the unique ID of the resource.
func (m *TargetResource) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetModifiedProperties gets the modifiedProperties property value. Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
func (m *TargetResource) GetModifiedProperties()([]ModifiedProperty) {
    if m == nil {
        return nil
    } else {
        return m.modifiedProperties
    }
}
// GetType_escaped gets the type property value. Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
func (m *TargetResource) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. When type is set to User, this includes the user name that initiated the action; null for other types.
func (m *TargetResource) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["groupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(GroupType)
            m.SetGroupType(&cast)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["modifiedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewModifiedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ModifiedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*ModifiedProperty))
            }
            m.SetModifiedProperties(res)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *TargetResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TargetResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetGroupType() != nil {
        cast := m.GetGroupType().String()
        err := writer.WriteStringValue("groupType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetModifiedProperties()))
        for i, v := range m.GetModifiedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("modifiedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetResource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Indicates the visible name defined for the resource. Typically specified when the resource is created.
func (m *TargetResource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGroupType sets the groupType property value. When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
func (m *TargetResource) SetGroupType(value *GroupType)() {
    if m != nil {
        m.groupType = value
    }
}
// SetId sets the id property value. Indicates the unique ID of the resource.
func (m *TargetResource) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetModifiedProperties sets the modifiedProperties property value. Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
func (m *TargetResource) SetModifiedProperties(value []ModifiedProperty)() {
    if m != nil {
        m.modifiedProperties = value
    }
}
// SetType_escaped sets the type property value. Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
func (m *TargetResource) SetType_escaped(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. When type is set to User, this includes the user name that initiated the action; null for other types.
func (m *TargetResource) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
