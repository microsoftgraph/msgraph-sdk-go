package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppIdentity provides operations to manage the auditLogRoot singleton.
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
// NewAppIdentity instantiates a new appIdentity and sets the default values.
func NewAppIdentity()(*AppIdentity) {
    m := &AppIdentity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppIdentityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAppIdentity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppId gets the appId property value. Refers to the Unique GUID representing Application Id in the Azure Active Directory.
func (m *AppIdentity) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetDisplayName gets the displayName property value. Refers to the Application Name displayed in the Azure Portal.
func (m *AppIdentity) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    return res
}
// GetServicePrincipalId gets the servicePrincipalId property value. Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
func (m *AppIdentity) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
// GetServicePrincipalName gets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
func (m *AppIdentity) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
func (m *AppIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppIdentity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppId sets the appId property value. Refers to the Unique GUID representing Application Id in the Azure Active Directory.
func (m *AppIdentity) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetDisplayName sets the displayName property value. Refers to the Application Name displayed in the Azure Portal.
func (m *AppIdentity) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetServicePrincipalId sets the servicePrincipalId property value. Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
func (m *AppIdentity) SetServicePrincipalId(value *string)() {
    if m != nil {
        m.servicePrincipalId = value
    }
}
// SetServicePrincipalName sets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
func (m *AppIdentity) SetServicePrincipalName(value *string)() {
    if m != nil {
        m.servicePrincipalName = value
    }
}
