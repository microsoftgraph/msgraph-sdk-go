package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy singleton.
type AuthenticationMethodsPolicy struct {
    Entity
    // Represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
    authenticationMethodConfigurations []AuthenticationMethodConfigurationable
    // A description of the policy.
    description *string
    // The name of the policy.
    displayName *string
    // The date and time of the last update to the policy.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The version of the policy in use.
    policyVersion *string
    // The reconfirmationInDays property
    reconfirmationInDays *int32
    // Enforce registration at sign-in time. This property can be used to remind users to set up targeted authentication methods.
    registrationEnforcement RegistrationEnforcementable
}
// NewAuthenticationMethodsPolicy instantiates a new authenticationMethodsPolicy and sets the default values.
func NewAuthenticationMethodsPolicy()(*AuthenticationMethodsPolicy) {
    m := &AuthenticationMethodsPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationMethodsPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationMethodsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodsPolicy(), nil
}
// GetAuthenticationMethodConfigurations gets the authenticationMethodConfigurations property value. Represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
func (m *AuthenticationMethodsPolicy) GetAuthenticationMethodConfigurations()([]AuthenticationMethodConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodConfigurations
    }
}
// GetDescription gets the description property value. A description of the policy.
func (m *AuthenticationMethodsPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of the policy.
func (m *AuthenticationMethodsPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationMethodConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodConfigurationable)
            }
            m.SetAuthenticationMethodConfigurations(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["policyVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyVersion(val)
        }
        return nil
    }
    res["reconfirmationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReconfirmationInDays(val)
        }
        return nil
    }
    res["registrationEnforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRegistrationEnforcementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationEnforcement(val.(RegistrationEnforcementable))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time of the last update to the policy.
func (m *AuthenticationMethodsPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPolicyVersion gets the policyVersion property value. The version of the policy in use.
func (m *AuthenticationMethodsPolicy) GetPolicyVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyVersion
    }
}
// GetReconfirmationInDays gets the reconfirmationInDays property value. The reconfirmationInDays property
func (m *AuthenticationMethodsPolicy) GetReconfirmationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.reconfirmationInDays
    }
}
// GetRegistrationEnforcement gets the registrationEnforcement property value. Enforce registration at sign-in time. This property can be used to remind users to set up targeted authentication methods.
func (m *AuthenticationMethodsPolicy) GetRegistrationEnforcement()(RegistrationEnforcementable) {
    if m == nil {
        return nil
    } else {
        return m.registrationEnforcement
    }
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMethodConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationMethodConfigurations()))
        for i, v := range m.GetAuthenticationMethodConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationMethodConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyVersion", m.GetPolicyVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("reconfirmationInDays", m.GetReconfirmationInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registrationEnforcement", m.GetRegistrationEnforcement())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethodConfigurations sets the authenticationMethodConfigurations property value. Represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
func (m *AuthenticationMethodsPolicy) SetAuthenticationMethodConfigurations(value []AuthenticationMethodConfigurationable)() {
    if m != nil {
        m.authenticationMethodConfigurations = value
    }
}
// SetDescription sets the description property value. A description of the policy.
func (m *AuthenticationMethodsPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The name of the policy.
func (m *AuthenticationMethodsPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time of the last update to the policy.
func (m *AuthenticationMethodsPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPolicyVersion sets the policyVersion property value. The version of the policy in use.
func (m *AuthenticationMethodsPolicy) SetPolicyVersion(value *string)() {
    if m != nil {
        m.policyVersion = value
    }
}
// SetReconfirmationInDays sets the reconfirmationInDays property value. The reconfirmationInDays property
func (m *AuthenticationMethodsPolicy) SetReconfirmationInDays(value *int32)() {
    if m != nil {
        m.reconfirmationInDays = value
    }
}
// SetRegistrationEnforcement sets the registrationEnforcement property value. Enforce registration at sign-in time. This property can be used to remind users to set up targeted authentication methods.
func (m *AuthenticationMethodsPolicy) SetRegistrationEnforcement(value RegistrationEnforcementable)() {
    if m != nil {
        m.registrationEnforcement = value
    }
}
