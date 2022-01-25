package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RegistrationEnforcement 
type RegistrationEnforcement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Run campaigns to remind users to setup targeted authentication methods.
    authenticationMethodsRegistrationCampaign *AuthenticationMethodsRegistrationCampaign;
}
// NewRegistrationEnforcement instantiates a new registrationEnforcement and sets the default values.
func NewRegistrationEnforcement()(*RegistrationEnforcement) {
    m := &RegistrationEnforcement{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegistrationEnforcement) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthenticationMethodsRegistrationCampaign gets the authenticationMethodsRegistrationCampaign property value. Run campaigns to remind users to setup targeted authentication methods.
func (m *RegistrationEnforcement) GetAuthenticationMethodsRegistrationCampaign()(*AuthenticationMethodsRegistrationCampaign) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsRegistrationCampaign
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegistrationEnforcement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationMethodsRegistrationCampaign"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsRegistrationCampaign() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodsRegistrationCampaign(val.(*AuthenticationMethodsRegistrationCampaign))
        }
        return nil
    }
    return res
}
func (m *RegistrationEnforcement) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RegistrationEnforcement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationMethodsRegistrationCampaign", m.GetAuthenticationMethodsRegistrationCampaign())
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
func (m *RegistrationEnforcement) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthenticationMethodsRegistrationCampaign sets the authenticationMethodsRegistrationCampaign property value. Run campaigns to remind users to setup targeted authentication methods.
func (m *RegistrationEnforcement) SetAuthenticationMethodsRegistrationCampaign(value *AuthenticationMethodsRegistrationCampaign)() {
    if m != nil {
        m.authenticationMethodsRegistrationCampaign = value
    }
}
