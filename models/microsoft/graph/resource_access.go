package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ResourceAccess struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier for one of the oauth2PermissionScopes or appRole instances that the resource application exposes.
    id *string;
    // Specifies whether the id property references an oauth2PermissionScopes or an appRole. The possible values are: Scope (for OAuth 2.0 permission scopes) or Role (for app roles).
    type_escaped *string;
}
// Instantiates a new resourceAccess and sets the default values.
func NewResourceAccess()(*ResourceAccess) {
    m := &ResourceAccess{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceAccess) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the id property value. The unique identifier for one of the oauth2PermissionScopes or appRole instances that the resource application exposes.
func (m *ResourceAccess) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the type_escaped property value. Specifies whether the id property references an oauth2PermissionScopes or an appRole. The possible values are: Scope (for OAuth 2.0 permission scopes) or Role (for app roles).
func (m *ResourceAccess) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *ResourceAccess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
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
    return res
}
func (m *ResourceAccess) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ResourceAccess) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *ResourceAccess) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the id property value. The unique identifier for one of the oauth2PermissionScopes or appRole instances that the resource application exposes.
// Parameters:
//  - value : Value to set for the id property.
func (m *ResourceAccess) SetId(value *string)() {
    m.id = value
}
// Sets the type_escaped property value. Specifies whether the id property references an oauth2PermissionScopes or an appRole. The possible values are: Scope (for OAuth 2.0 permission scopes) or Role (for app roles).
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *ResourceAccess) SetType_escaped(value *string)() {
    m.type_escaped = value
}
