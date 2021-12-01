package getmembergroups

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetMemberGroupsRequestBody 
type GetMemberGroupsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    securityEnabledOnly *bool;
}
// NewGetMemberGroupsRequestBody instantiates a new getMemberGroupsRequestBody and sets the default values.
func NewGetMemberGroupsRequestBody()(*GetMemberGroupsRequestBody) {
    m := &GetMemberGroupsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetMemberGroupsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetSecurityEnabledOnly gets the securityEnabledOnly property value. 
func (m *GetMemberGroupsRequestBody) GetSecurityEnabledOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityEnabledOnly
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetMemberGroupsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["securityEnabledOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabledOnly(val)
        }
        return nil
    }
    return res
}
func (m *GetMemberGroupsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetMemberGroupsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("securityEnabledOnly", m.GetSecurityEnabledOnly())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetMemberGroupsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSecurityEnabledOnly sets the securityEnabledOnly property value. 
func (m *GetMemberGroupsRequestBody) SetSecurityEnabledOnly(value *bool)() {
    if m != nil {
        m.securityEnabledOnly = value
    }
}
