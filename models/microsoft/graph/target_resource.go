package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new targetResource and sets the default values.
func NewTargetResource()(*TargetResource) {
    m := &TargetResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. Indicates the visible name defined for the resource. Typically specified when the resource is created.
func (m *TargetResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the groupType property value. When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
func (m *TargetResource) GetGroupType()(*GroupType) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
// Gets the id property value. Indicates the unique ID of the resource.
func (m *TargetResource) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the modifiedProperties property value. Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
func (m *TargetResource) GetModifiedProperties()([]ModifiedProperty) {
    if m == nil {
        return nil
    } else {
        return m.modifiedProperties
    }
}
// Gets the type_escaped property value. Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
func (m *TargetResource) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the userPrincipalName property value. When type is set to User, this includes the user name that initiated the action; null for other types.
func (m *TargetResource) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *TargetResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["groupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupType)
        if err != nil {
            return err
        }
        cast := val.(GroupType)
        m.SetGroupType(&cast)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["modifiedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewModifiedProperty() })
        if err != nil {
            return err
        }
        res := make([]ModifiedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*ModifiedProperty))
        }
        m.SetModifiedProperties(res)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *TargetResource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TargetResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. Indicates the visible name defined for the resource. Typically specified when the resource is created.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TargetResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the groupType property value. When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
// Parameters:
//  - value : Value to set for the groupType property.
func (m *TargetResource) SetGroupType(value *GroupType)() {
    m.groupType = value
}
// Sets the id property value. Indicates the unique ID of the resource.
// Parameters:
//  - value : Value to set for the id property.
func (m *TargetResource) SetId(value *string)() {
    m.id = value
}
// Sets the modifiedProperties property value. Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
// Parameters:
//  - value : Value to set for the modifiedProperties property.
func (m *TargetResource) SetModifiedProperties(value []ModifiedProperty)() {
    m.modifiedProperties = value
}
// Sets the type_escaped property value. Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *TargetResource) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the userPrincipalName property value. When type is set to User, this includes the user name that initiated the action; null for other types.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *TargetResource) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
