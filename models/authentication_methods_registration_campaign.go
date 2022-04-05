package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodsRegistrationCampaign 
type AuthenticationMethodsRegistrationCampaign struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Users and groups of users that are excluded from being prompted to set up the authentication method.
    excludeTargets []ExcludeTargetable;
    // Users and groups of users that are prompted to set up the authentication method.
    includeTargets []AuthenticationMethodsRegistrationCampaignIncludeTargetable;
    // Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum: 0 days. Maximum: 14 days. If the value is '0', the user is prompted during every MFA attempt.
    snoozeDurationInDays *int32;
    // Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure Active Directory for the setting. The default value is disabled.
    state *AdvancedConfigState;
}
// NewAuthenticationMethodsRegistrationCampaign instantiates a new authenticationMethodsRegistrationCampaign and sets the default values.
func NewAuthenticationMethodsRegistrationCampaign()(*AuthenticationMethodsRegistrationCampaign) {
    m := &AuthenticationMethodsRegistrationCampaign{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationMethodsRegistrationCampaignFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodsRegistrationCampaignFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodsRegistrationCampaign(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationMethodsRegistrationCampaign) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeTargets gets the excludeTargets property value. Users and groups of users that are excluded from being prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) GetExcludeTargets()([]ExcludeTargetable) {
    if m == nil {
        return nil
    } else {
        return m.excludeTargets
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodsRegistrationCampaign) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeTargets"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExcludeTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExcludeTargetable, len(val))
            for i, v := range val {
                res[i] = v.(ExcludeTargetable)
            }
            m.SetExcludeTargets(res)
        }
        return nil
    }
    res["includeTargets"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodsRegistrationCampaignIncludeTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodsRegistrationCampaignIncludeTargetable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodsRegistrationCampaignIncludeTargetable)
            }
            m.SetIncludeTargets(res)
        }
        return nil
    }
    res["snoozeDurationInDays"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnoozeDurationInDays(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AdvancedConfigState))
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. Users and groups of users that are prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) GetIncludeTargets()([]AuthenticationMethodsRegistrationCampaignIncludeTargetable) {
    if m == nil {
        return nil
    } else {
        return m.includeTargets
    }
}
// GetSnoozeDurationInDays gets the snoozeDurationInDays property value. Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum: 0 days. Maximum: 14 days. If the value is '0', the user is prompted during every MFA attempt.
func (m *AuthenticationMethodsRegistrationCampaign) GetSnoozeDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.snoozeDurationInDays
    }
}
// GetState gets the state property value. Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure Active Directory for the setting. The default value is disabled.
func (m *AuthenticationMethodsRegistrationCampaign) GetState()(*AdvancedConfigState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsRegistrationCampaign) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExcludeTargets()))
        for i, v := range m.GetExcludeTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("excludeTargets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := (*m.GetState()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationMethodsRegistrationCampaign) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeTargets sets the excludeTargets property value. Users and groups of users that are excluded from being prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) SetExcludeTargets(value []ExcludeTargetable)() {
    if m != nil {
        m.excludeTargets = value
    }
}
// SetIncludeTargets sets the includeTargets property value. Users and groups of users that are prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) SetIncludeTargets(value []AuthenticationMethodsRegistrationCampaignIncludeTargetable)() {
    if m != nil {
        m.includeTargets = value
    }
}
// SetSnoozeDurationInDays sets the snoozeDurationInDays property value. Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum: 0 days. Maximum: 14 days. If the value is '0', the user is prompted during every MFA attempt.
func (m *AuthenticationMethodsRegistrationCampaign) SetSnoozeDurationInDays(value *int32)() {
    if m != nil {
        m.snoozeDurationInDays = value
    }
}
// SetState sets the state property value. Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure Active Directory for the setting. The default value is disabled.
func (m *AuthenticationMethodsRegistrationCampaign) SetState(value *AdvancedConfigState)() {
    if m != nil {
        m.state = value
    }
}
