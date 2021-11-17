package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationMethodsRegistrationCampaign struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Users and groups of users that are excluded from being prompted to set up the authentication method.
    excludeTargets []ExcludeTarget;
    // Users and groups of users that are prompted to set up the authentication method.
    includeTargets []AuthenticationMethodsRegistrationCampaignIncludeTarget;
    // Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum 0 days. Maximum: 14 days. If the value is '0' – The user is prompted during every MFA attempt.
    snoozeDurationInDays *int32;
    // Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure AD for the setting. The default value is disabled.
    state *AdvancedConfigState;
}
// Instantiates a new authenticationMethodsRegistrationCampaign and sets the default values.
func NewAuthenticationMethodsRegistrationCampaign()(*AuthenticationMethodsRegistrationCampaign) {
    m := &AuthenticationMethodsRegistrationCampaign{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationMethodsRegistrationCampaign) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the excludeTargets property value. Users and groups of users that are excluded from being prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) GetExcludeTargets()([]ExcludeTarget) {
    if m == nil {
        return nil
    } else {
        return m.excludeTargets
    }
}
// Gets the includeTargets property value. Users and groups of users that are prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) GetIncludeTargets()([]AuthenticationMethodsRegistrationCampaignIncludeTarget) {
    if m == nil {
        return nil
    } else {
        return m.includeTargets
    }
}
// Gets the snoozeDurationInDays property value. Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum 0 days. Maximum: 14 days. If the value is '0' – The user is prompted during every MFA attempt.
func (m *AuthenticationMethodsRegistrationCampaign) GetSnoozeDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.snoozeDurationInDays
    }
}
// Gets the state property value. Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure AD for the setting. The default value is disabled.
func (m *AuthenticationMethodsRegistrationCampaign) GetState()(*AdvancedConfigState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *AuthenticationMethodsRegistrationCampaign) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeTargets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExcludeTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExcludeTarget, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExcludeTarget))
            }
            m.SetExcludeTargets(res)
        }
        return nil
    }
    res["includeTargets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsRegistrationCampaignIncludeTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodsRegistrationCampaignIncludeTarget, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthenticationMethodsRegistrationCampaignIncludeTarget))
            }
            m.SetIncludeTargets(res)
        }
        return nil
    }
    res["snoozeDurationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnoozeDurationInDays(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AdvancedConfigState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *AuthenticationMethodsRegistrationCampaign) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuthenticationMethodsRegistrationCampaign) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludeTargets()))
        for i, v := range m.GetExcludeTargets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("excludeTargets", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("snoozeDurationInDays", m.GetSnoozeDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *AuthenticationMethodsRegistrationCampaign) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the excludeTargets property value. Users and groups of users that are excluded from being prompted to set up the authentication method.
// Parameters:
//  - value : Value to set for the excludeTargets property.
func (m *AuthenticationMethodsRegistrationCampaign) SetExcludeTargets(value []ExcludeTarget)() {
    m.excludeTargets = value
}
// Sets the includeTargets property value. Users and groups of users that are prompted to set up the authentication method.
// Parameters:
//  - value : Value to set for the includeTargets property.
func (m *AuthenticationMethodsRegistrationCampaign) SetIncludeTargets(value []AuthenticationMethodsRegistrationCampaignIncludeTarget)() {
    m.includeTargets = value
}
// Sets the snoozeDurationInDays property value. Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum 0 days. Maximum: 14 days. If the value is '0' – The user is prompted during every MFA attempt.
// Parameters:
//  - value : Value to set for the snoozeDurationInDays property.
func (m *AuthenticationMethodsRegistrationCampaign) SetSnoozeDurationInDays(value *int32)() {
    m.snoozeDurationInDays = value
}
// Sets the state property value. Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure AD for the setting. The default value is disabled.
// Parameters:
//  - value : Value to set for the state property.
func (m *AuthenticationMethodsRegistrationCampaign) SetState(value *AdvancedConfigState)() {
    m.state = value
}
