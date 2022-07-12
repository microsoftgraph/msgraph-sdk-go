package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyInboundTrust 
type CrossTenantAccessPolicyInboundTrust struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies whether compliant devices from external Azure AD organizations are trusted.
    isCompliantDeviceAccepted *bool
    // Specifies whether hybrid Azure AD joined devices from external Azure AD organizations are trusted.
    isHybridAzureADJoinedDeviceAccepted *bool
    // Specifies whether MFA from external Azure AD organizations is trusted.
    isMfaAccepted *bool
}
// NewCrossTenantAccessPolicyInboundTrust instantiates a new crossTenantAccessPolicyInboundTrust and sets the default values.
func NewCrossTenantAccessPolicyInboundTrust()(*CrossTenantAccessPolicyInboundTrust) {
    m := &CrossTenantAccessPolicyInboundTrust{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCrossTenantAccessPolicyInboundTrustFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyInboundTrustFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyInboundTrust(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyInboundTrust) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyInboundTrust) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isCompliantDeviceAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompliantDeviceAccepted(val)
        }
        return nil
    }
    res["isHybridAzureADJoinedDeviceAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHybridAzureADJoinedDeviceAccepted(val)
        }
        return nil
    }
    res["isMfaAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaAccepted(val)
        }
        return nil
    }
    return res
}
// GetIsCompliantDeviceAccepted gets the isCompliantDeviceAccepted property value. Specifies whether compliant devices from external Azure AD organizations are trusted.
func (m *CrossTenantAccessPolicyInboundTrust) GetIsCompliantDeviceAccepted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCompliantDeviceAccepted
    }
}
// GetIsHybridAzureADJoinedDeviceAccepted gets the isHybridAzureADJoinedDeviceAccepted property value. Specifies whether hybrid Azure AD joined devices from external Azure AD organizations are trusted.
func (m *CrossTenantAccessPolicyInboundTrust) GetIsHybridAzureADJoinedDeviceAccepted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHybridAzureADJoinedDeviceAccepted
    }
}
// GetIsMfaAccepted gets the isMfaAccepted property value. Specifies whether MFA from external Azure AD organizations is trusted.
func (m *CrossTenantAccessPolicyInboundTrust) GetIsMfaAccepted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaAccepted
    }
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyInboundTrust) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isCompliantDeviceAccepted", m.GetIsCompliantDeviceAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHybridAzureADJoinedDeviceAccepted", m.GetIsHybridAzureADJoinedDeviceAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMfaAccepted", m.GetIsMfaAccepted())
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
func (m *CrossTenantAccessPolicyInboundTrust) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsCompliantDeviceAccepted sets the isCompliantDeviceAccepted property value. Specifies whether compliant devices from external Azure AD organizations are trusted.
func (m *CrossTenantAccessPolicyInboundTrust) SetIsCompliantDeviceAccepted(value *bool)() {
    if m != nil {
        m.isCompliantDeviceAccepted = value
    }
}
// SetIsHybridAzureADJoinedDeviceAccepted sets the isHybridAzureADJoinedDeviceAccepted property value. Specifies whether hybrid Azure AD joined devices from external Azure AD organizations are trusted.
func (m *CrossTenantAccessPolicyInboundTrust) SetIsHybridAzureADJoinedDeviceAccepted(value *bool)() {
    if m != nil {
        m.isHybridAzureADJoinedDeviceAccepted = value
    }
}
// SetIsMfaAccepted sets the isMfaAccepted property value. Specifies whether MFA from external Azure AD organizations is trusted.
func (m *CrossTenantAccessPolicyInboundTrust) SetIsMfaAccepted(value *bool)() {
    if m != nil {
        m.isMfaAccepted = value
    }
}
