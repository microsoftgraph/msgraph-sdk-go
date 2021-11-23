package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// authenticationMethodsRegistrationCampaignIncludeTarget 
type AuthenticationMethodsRegistrationCampaignIncludeTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The object identifier of an Azure Active Directory user or group.
    id *string;
    // The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
    targetedAuthenticationMethod *string;
    // The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
    targetType *AuthenticationMethodTargetType;
}
// NewAuthenticationMethodsRegistrationCampaignIncludeTarget instantiates a new authenticationMethodsRegistrationCampaignIncludeTarget and sets the default values.
func NewAuthenticationMethodsRegistrationCampaignIncludeTarget()(*AuthenticationMethodsRegistrationCampaignIncludeTarget) {
    m := &AuthenticationMethodsRegistrationCampaignIncludeTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetId gets the id property value. The object identifier of an Azure Active Directory user or group.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetTargetedAuthenticationMethod gets the targetedAuthenticationMethod property value. The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetedAuthenticationMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedAuthenticationMethod
    }
}
// GetTargetType gets the targetType property value. The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetType()(*AuthenticationMethodTargetType) {
    if m == nil {
        return nil
    } else {
        return m.targetType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["targetedAuthenticationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedAuthenticationMethod(val)
        }
        return nil
    }
    res["targetType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodTargetType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AuthenticationMethodTargetType)
            m.SetTargetType(&cast)
        }
        return nil
    }
    return res
}
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetedAuthenticationMethod", m.GetTargetedAuthenticationMethod())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := m.GetTargetType().String()
        err := writer.WriteStringValue("targetType", &cast)
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
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetId sets the id property value. The object identifier of an Azure Active Directory user or group.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetId(value *string)() {
    m.id = value
}
// SetTargetedAuthenticationMethod sets the targetedAuthenticationMethod property value. The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetedAuthenticationMethod(value *string)() {
    m.targetedAuthenticationMethod = value
}
// SetTargetType sets the targetType property value. The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetType(value *AuthenticationMethodTargetType)() {
    m.targetType = value
}
