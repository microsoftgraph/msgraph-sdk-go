package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TelecomExpenseManagementPartner 
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
// NewTelecomExpenseManagementPartner instantiates a new telecomExpenseManagementPartner and sets the default values.
func NewTelecomExpenseManagementPartner()(*TelecomExpenseManagementPartner) {
    m := &TelecomExpenseManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppAuthorized gets the appAuthorized property value. Whether the partner's AAD app has been authorized to access Intune.
func (m *TelecomExpenseManagementPartner) GetAppAuthorized()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appAuthorized
    }
}
// GetDisplayName gets the displayName property value. Display name of the TEM partner.
func (m *TelecomExpenseManagementPartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnabled gets the enabled property value. Whether Intune's connection to the TEM service is currently enabled or disabled.
func (m *TelecomExpenseManagementPartner) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *TelecomExpenseManagementPartner) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// GetUrl gets the url property value. URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
func (m *TelecomExpenseManagementPartner) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TelecomExpenseManagementPartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appAuthorized"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppAuthorized(val)
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
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["lastConnectionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConnectionDateTime(val)
        }
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *TelecomExpenseManagementPartner) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAppAuthorized sets the appAuthorized property value. Whether the partner's AAD app has been authorized to access Intune.
func (m *TelecomExpenseManagementPartner) SetAppAuthorized(value *bool)() {
    if m != nil {
        m.appAuthorized = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the TEM partner.
func (m *TelecomExpenseManagementPartner) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnabled sets the enabled property value. Whether Intune's connection to the TEM service is currently enabled or disabled.
func (m *TelecomExpenseManagementPartner) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *TelecomExpenseManagementPartner) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastConnectionDateTime = value
    }
}
// SetUrl sets the url property value. URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
func (m *TelecomExpenseManagementPartner) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
