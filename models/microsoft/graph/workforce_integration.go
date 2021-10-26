package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkforceIntegration struct {
    ChangeTrackedEntity
    // API version for the call back URL. Start with 1.
    apiVersion *int32;
    // Name of the workforce integration.
    displayName *string;
    // The workforce integration encryption resource.
    encryption *WorkforceIntegrationEncryption;
    // Indicates whether this workforce integration is currently active and available.
    isActive *bool;
    // The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue.
    supportedEntities *WorkforceIntegrationSupportedEntities;
    // Workforce Integration URL for callbacks from the Shifts service.
    url *string;
}
// Instantiates a new workforceIntegration and sets the default values.
func NewWorkforceIntegration()(*WorkforceIntegration) {
    m := &WorkforceIntegration{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// Gets the apiVersion property value. API version for the call back URL. Start with 1.
func (m *WorkforceIntegration) GetApiVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.apiVersion
    }
}
// Gets the displayName property value. Name of the workforce integration.
func (m *WorkforceIntegration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the encryption property value. The workforce integration encryption resource.
func (m *WorkforceIntegration) GetEncryption()(*WorkforceIntegrationEncryption) {
    if m == nil {
        return nil
    } else {
        return m.encryption
    }
}
// Gets the isActive property value. Indicates whether this workforce integration is currently active and available.
func (m *WorkforceIntegration) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// Gets the supportedEntities property value. The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue.
func (m *WorkforceIntegration) GetSupportedEntities()(*WorkforceIntegrationSupportedEntities) {
    if m == nil {
        return nil
    } else {
        return m.supportedEntities
    }
}
// Gets the url property value. Workforce Integration URL for callbacks from the Shifts service.
func (m *WorkforceIntegration) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// The deserialization information for the current model
func (m *WorkforceIntegration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["apiVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetApiVersion(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["encryption"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkforceIntegrationEncryption() })
        if err != nil {
            return err
        }
        m.SetEncryption(val.(*WorkforceIntegrationEncryption))
        return nil
    }
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsActive(val)
        return nil
    }
    res["supportedEntities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkforceIntegrationSupportedEntities)
        if err != nil {
            return err
        }
        cast := val.(WorkforceIntegrationSupportedEntities)
        m.SetSupportedEntities(&cast)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    return res
}
func (m *WorkforceIntegration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkforceIntegration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetSupportedEntities().String()
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
// Sets the apiVersion property value. API version for the call back URL. Start with 1.
// Parameters:
//  - value : Value to set for the apiVersion property.
func (m *WorkforceIntegration) SetApiVersion(value *int32)() {
    m.apiVersion = value
}
// Sets the displayName property value. Name of the workforce integration.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *WorkforceIntegration) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the encryption property value. The workforce integration encryption resource.
// Parameters:
//  - value : Value to set for the encryption property.
func (m *WorkforceIntegration) SetEncryption(value *WorkforceIntegrationEncryption)() {
    m.encryption = value
}
// Sets the isActive property value. Indicates whether this workforce integration is currently active and available.
// Parameters:
//  - value : Value to set for the isActive property.
func (m *WorkforceIntegration) SetIsActive(value *bool)() {
    m.isActive = value
}
// Sets the supportedEntities property value. The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue.
// Parameters:
//  - value : Value to set for the supportedEntities property.
func (m *WorkforceIntegration) SetSupportedEntities(value *WorkforceIntegrationSupportedEntities)() {
    m.supportedEntities = value
}
// Sets the url property value. Workforce Integration URL for callbacks from the Shifts service.
// Parameters:
//  - value : Value to set for the url property.
func (m *WorkforceIntegration) SetUrl(value *string)() {
    m.url = value
}
