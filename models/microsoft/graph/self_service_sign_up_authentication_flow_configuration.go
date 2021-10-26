package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SelfServiceSignUpAuthenticationFlowConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
    isEnabled *bool;
}
// Instantiates a new selfServiceSignUpAuthenticationFlowConfiguration and sets the default values.
func NewSelfServiceSignUpAuthenticationFlowConfiguration()(*SelfServiceSignUpAuthenticationFlowConfiguration) {
    m := &SelfServiceSignUpAuthenticationFlowConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isEnabled property value. Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// The deserialization information for the current model
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    return res
}
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
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
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isEnabled property value. Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
