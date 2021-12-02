package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SelfServiceSignUpAuthenticationFlowConfiguration 
type SelfServiceSignUpAuthenticationFlowConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
    isEnabled *bool;
}
// NewSelfServiceSignUpAuthenticationFlowConfiguration instantiates a new selfServiceSignUpAuthenticationFlowConfiguration and sets the default values.
func NewSelfServiceSignUpAuthenticationFlowConfiguration()(*SelfServiceSignUpAuthenticationFlowConfiguration) {
    m := &SelfServiceSignUpAuthenticationFlowConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    return res
}
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
