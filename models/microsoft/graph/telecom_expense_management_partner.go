package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TelecomExpenseManagementPartner struct {
    Entity
    // Whether the partner's AAD app has been authorized to access Intune.
    appAuthorized *bool;
    // Display name of the TEM partner.
    displayName *string;
    // Whether Intune's connection to the TEM service is currently enabled or disabled.
    enabled *bool;
    // Timestamp of the last request sent to Intune by the TEM partner.
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
    url *string;
}
// Instantiates a new telecomExpenseManagementPartner and sets the default values.
func NewTelecomExpenseManagementPartner()(*TelecomExpenseManagementPartner) {
    m := &TelecomExpenseManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appAuthorized property value. Whether the partner's AAD app has been authorized to access Intune.
func (m *TelecomExpenseManagementPartner) GetAppAuthorized()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appAuthorized
    }
}
// Gets the displayName property value. Display name of the TEM partner.
func (m *TelecomExpenseManagementPartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enabled property value. Whether Intune's connection to the TEM service is currently enabled or disabled.
func (m *TelecomExpenseManagementPartner) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// Gets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *TelecomExpenseManagementPartner) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// Gets the url property value. URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
func (m *TelecomExpenseManagementPartner) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// The deserialization information for the current model
func (m *TelecomExpenseManagementPartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appAuthorized"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAppAuthorized(val)
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
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnabled(val)
        return nil
    }
    res["lastConnectionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastConnectionDateTime(val)
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
func (m *TelecomExpenseManagementPartner) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TelecomExpenseManagementPartner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appAuthorized", m.GetAppAuthorized())
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
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastConnectionDateTime", m.GetLastConnectionDateTime())
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
// Sets the appAuthorized property value. Whether the partner's AAD app has been authorized to access Intune.
// Parameters:
//  - value : Value to set for the appAuthorized property.
func (m *TelecomExpenseManagementPartner) SetAppAuthorized(value *bool)() {
    m.appAuthorized = value
}
// Sets the displayName property value. Display name of the TEM partner.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TelecomExpenseManagementPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enabled property value. Whether Intune's connection to the TEM service is currently enabled or disabled.
// Parameters:
//  - value : Value to set for the enabled property.
func (m *TelecomExpenseManagementPartner) SetEnabled(value *bool)() {
    m.enabled = value
}
// Sets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
// Parameters:
//  - value : Value to set for the lastConnectionDateTime property.
func (m *TelecomExpenseManagementPartner) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// Sets the url property value. URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
// Parameters:
//  - value : Value to set for the url property.
func (m *TelecomExpenseManagementPartner) SetUrl(value *string)() {
    m.url = value
}
