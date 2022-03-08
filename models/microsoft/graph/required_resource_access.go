package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RequiredResourceAccess provides operations to manage the collection of application entities.
type RequiredResourceAccess struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
    resourceAccess []ResourceAccessable;
    // The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
    resourceAppId *string;
}
// NewRequiredResourceAccess instantiates a new requiredResourceAccess and sets the default values.
func NewRequiredResourceAccess()(*RequiredResourceAccess) {
    m := &RequiredResourceAccess{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRequiredResourceAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequiredResourceAccessFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRequiredResourceAccess(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequiredResourceAccess) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequiredResourceAccess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resourceAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceAccessable, len(val))
            for i, v := range val {
                res[i] = v.(ResourceAccessable)
            }
            m.SetResourceAccess(res)
        }
        return nil
    }
    res["resourceAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAppId(val)
        }
        return nil
    }
    return res
}
// GetResourceAccess gets the resourceAccess property value. The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
func (m *RequiredResourceAccess) GetResourceAccess()([]ResourceAccessable) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccess
    }
}
// GetResourceAppId gets the resourceAppId property value. The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
func (m *RequiredResourceAccess) GetResourceAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceAppId
    }
}
func (m *RequiredResourceAccess) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RequiredResourceAccess) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetResourceAccess() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceAccess()))
        for i, v := range m.GetResourceAccess() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequiredResourceAccess) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResourceAccess sets the resourceAccess property value. The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
func (m *RequiredResourceAccess) SetResourceAccess(value []ResourceAccessable)() {
    if m != nil {
        m.resourceAccess = value
    }
}
// SetResourceAppId sets the resourceAppId property value. The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
func (m *RequiredResourceAccess) SetResourceAppId(value *string)() {
    if m != nil {
        m.resourceAppId = value
    }
}
