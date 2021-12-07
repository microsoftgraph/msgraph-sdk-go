package changescreensharingrole

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// ChangeScreenSharingRoleRequestBody 
type ChangeScreenSharingRoleRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    role *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole;
}
// NewChangeScreenSharingRoleRequestBody instantiates a new changeScreenSharingRoleRequestBody and sets the default values.
func NewChangeScreenSharingRoleRequestBody()(*ChangeScreenSharingRoleRequestBody) {
    m := &ChangeScreenSharingRoleRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeScreenSharingRoleRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRole gets the role property value. 
func (m *ChangeScreenSharingRoleRequestBody) GetRole()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeScreenSharingRoleRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseScreenSharingRole)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole)
            m.SetRole(&cast)
        }
        return nil
    }
    return res
}
func (m *ChangeScreenSharingRoleRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChangeScreenSharingRoleRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRole() != nil {
        cast := m.GetRole().String()
        err := writer.WriteStringValue("role", &cast)
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
func (m *ChangeScreenSharingRoleRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRole sets the role property value. 
func (m *ChangeScreenSharingRoleRequestBody) SetRole(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole)() {
    if m != nil {
        m.role = value
    }
}
