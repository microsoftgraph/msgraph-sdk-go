package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RequiredResourceAccess struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
    resourceAccess []ResourceAccess;
    // The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
    resourceAppId *string;
}
// Instantiates a new requiredResourceAccess and sets the default values.
func NewRequiredResourceAccess()(*RequiredResourceAccess) {
    m := &RequiredResourceAccess{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequiredResourceAccess) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the resourceAccess property value. The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
func (m *RequiredResourceAccess) GetResourceAccess()([]ResourceAccess) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccess
    }
}
// Gets the resourceAppId property value. The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
func (m *RequiredResourceAccess) GetResourceAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceAppId
    }
}
// The deserialization information for the current model
func (m *RequiredResourceAccess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resourceAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceAccess() })
        if err != nil {
            return err
        }
        res := make([]ResourceAccess, len(val))
        for i, v := range val {
            res[i] = *(v.(*ResourceAccess))
        }
        m.SetResourceAccess(res)
        return nil
    }
    res["resourceAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceAppId(val)
        return nil
    }
    return res
}
func (m *RequiredResourceAccess) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RequiredResourceAccess) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceAccess()))
        for i, v := range m.GetResourceAccess() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("resourceAccess", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceAppId", m.GetResourceAppId())
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
func (m *RequiredResourceAccess) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the resourceAccess property value. The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
// Parameters:
//  - value : Value to set for the resourceAccess property.
func (m *RequiredResourceAccess) SetResourceAccess(value []ResourceAccess)() {
    m.resourceAccess = value
}
// Sets the resourceAppId property value. The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
// Parameters:
//  - value : Value to set for the resourceAppId property.
func (m *RequiredResourceAccess) SetResourceAppId(value *string)() {
    m.resourceAppId = value
}
