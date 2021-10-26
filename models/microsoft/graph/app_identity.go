package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AppIdentity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Refers to the Unique GUID representing Application Id in the Azure Active Directory.
    appId *string;
    // Refers to the Application Name displayed in the Azure Portal.
    displayName *string;
    // Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
    servicePrincipalId *string;
    // Refers to the Service Principal Name is the Application name in the tenant.
    servicePrincipalName *string;
}
// Instantiates a new appIdentity and sets the default values.
func NewAppIdentity()(*AppIdentity) {
    m := &AppIdentity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appId property value. Refers to the Unique GUID representing Application Id in the Azure Active Directory.
func (m *AppIdentity) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the displayName property value. Refers to the Application Name displayed in the Azure Portal.
func (m *AppIdentity) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the servicePrincipalId property value. Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
func (m *AppIdentity) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
// Gets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
func (m *AppIdentity) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
// The deserialization information for the current model
func (m *AppIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
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
    res["servicePrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalId(val)
        return nil
    }
    res["servicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalName(val)
        return nil
    }
    return res
}
func (m *AppIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AppIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
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
func (m *AppIdentity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appId property value. Refers to the Unique GUID representing Application Id in the Azure Active Directory.
// Parameters:
//  - value : Value to set for the appId property.
func (m *AppIdentity) SetAppId(value *string)() {
    m.appId = value
}
// Sets the displayName property value. Refers to the Application Name displayed in the Azure Portal.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AppIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the servicePrincipalId property value. Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
// Parameters:
//  - value : Value to set for the servicePrincipalId property.
func (m *AppIdentity) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
// Sets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
// Parameters:
//  - value : Value to set for the servicePrincipalName property.
func (m *AppIdentity) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
