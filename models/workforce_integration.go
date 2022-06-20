package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkforceIntegration 
type WorkforceIntegration struct {
    ChangeTrackedEntity
    // API version for the call back URL. Start with 1.
    apiVersion *int32
    // Name of the workforce integration.
    displayName *string
    // The workforce integration encryption resource.
    encryption WorkforceIntegrationEncryptionable
    // Indicates whether this workforce integration is currently active and available.
    isActive *bool
    // This property has replaced supports in v1.0. We recommend that you use this property instead of supports. The supports property is still supported in beta for the time being. The possible values are: none, shift, swapRequest, openshift, openShiftRequest, userShiftPreferences, offerShiftRequest, unknownFutureValue, timeCard, timeOffReason, timeOff, timeOffRequest. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: timeCard, timeOffReason, timeOff, timeOffRequest. If selecting more than one value, all values must start with the first letter in uppercase.
    supportedEntities *WorkforceIntegrationSupportedEntities
    // Workforce Integration URL for callbacks from the Shifts service.
    url *string
}
// NewWorkforceIntegration instantiates a new WorkforceIntegration and sets the default values.
func NewWorkforceIntegration()(*WorkforceIntegration) {
    m := &WorkforceIntegration{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// CreateWorkforceIntegrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkforceIntegrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkforceIntegration(), nil
}
// GetApiVersion gets the apiVersion property value. API version for the call back URL. Start with 1.
func (m *WorkforceIntegration) GetApiVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.apiVersion
    }
}
// GetDisplayName gets the displayName property value. Name of the workforce integration.
func (m *WorkforceIntegration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEncryption gets the encryption property value. The workforce integration encryption resource.
func (m *WorkforceIntegration) GetEncryption()(WorkforceIntegrationEncryptionable) {
    if m == nil {
        return nil
    } else {
        return m.encryption
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkforceIntegration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["apiVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiVersion(val)
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
    res["encryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkforceIntegrationEncryptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryption(val.(WorkforceIntegrationEncryptionable))
        }
        return nil
    }
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["supportedEntities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkforceIntegrationSupportedEntities)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedEntities(val.(*WorkforceIntegrationSupportedEntities))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetIsActive gets the isActive property value. Indicates whether this workforce integration is currently active and available.
func (m *WorkforceIntegration) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetSupportedEntities gets the supportedEntities property value. This property has replaced supports in v1.0. We recommend that you use this property instead of supports. The supports property is still supported in beta for the time being. The possible values are: none, shift, swapRequest, openshift, openShiftRequest, userShiftPreferences, offerShiftRequest, unknownFutureValue, timeCard, timeOffReason, timeOff, timeOffRequest. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: timeCard, timeOffReason, timeOff, timeOffRequest. If selecting more than one value, all values must start with the first letter in uppercase.
func (m *WorkforceIntegration) GetSupportedEntities()(*WorkforceIntegrationSupportedEntities) {
    if m == nil {
        return nil
    } else {
        return m.supportedEntities
    }
}
// GetUrl gets the url property value. Workforce Integration URL for callbacks from the Shifts service.
func (m *WorkforceIntegration) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// Serialize serializes information the current object
func (m *WorkforceIntegration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("apiVersion", m.GetApiVersion())
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
        err = writer.WriteObjectValue("encryption", m.GetEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedEntities() != nil {
        cast := (*m.GetSupportedEntities()).String()
        err = writer.WriteStringValue("supportedEntities", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiVersion sets the apiVersion property value. API version for the call back URL. Start with 1.
func (m *WorkforceIntegration) SetApiVersion(value *int32)() {
    if m != nil {
        m.apiVersion = value
    }
}
// SetDisplayName sets the displayName property value. Name of the workforce integration.
func (m *WorkforceIntegration) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEncryption sets the encryption property value. The workforce integration encryption resource.
func (m *WorkforceIntegration) SetEncryption(value WorkforceIntegrationEncryptionable)() {
    if m != nil {
        m.encryption = value
    }
}
// SetIsActive sets the isActive property value. Indicates whether this workforce integration is currently active and available.
func (m *WorkforceIntegration) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetSupportedEntities sets the supportedEntities property value. This property has replaced supports in v1.0. We recommend that you use this property instead of supports. The supports property is still supported in beta for the time being. The possible values are: none, shift, swapRequest, openshift, openShiftRequest, userShiftPreferences, offerShiftRequest, unknownFutureValue, timeCard, timeOffReason, timeOff, timeOffRequest. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: timeCard, timeOffReason, timeOff, timeOffRequest. If selecting more than one value, all values must start with the first letter in uppercase.
func (m *WorkforceIntegration) SetSupportedEntities(value *WorkforceIntegrationSupportedEntities)() {
    if m != nil {
        m.supportedEntities = value
    }
}
// SetUrl sets the url property value. Workforce Integration URL for callbacks from the Shifts service.
func (m *WorkforceIntegration) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
