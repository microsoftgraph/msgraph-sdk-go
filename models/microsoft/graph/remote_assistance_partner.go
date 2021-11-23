package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RemoteAssistancePartner 
type RemoteAssistancePartner struct {
    Entity
    // Display name of the partner.
    displayName *string;
    // Timestamp of the last request sent to Intune by the TEM partner.
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A friendly description of the current TeamViewer connector status. Possible values are: notOnboarded, onboarding, onboarded.
    onboardingStatus *RemoteAssistanceOnboardingStatus;
    // URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
    onboardingUrl *string;
}
// NewRemoteAssistancePartner instantiates a new remoteAssistancePartner and sets the default values.
func NewRemoteAssistancePartner()(*RemoteAssistancePartner) {
    m := &RemoteAssistancePartner{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Display name of the partner.
func (m *RemoteAssistancePartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *RemoteAssistancePartner) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// GetOnboardingStatus gets the onboardingStatus property value. A friendly description of the current TeamViewer connector status. Possible values are: notOnboarded, onboarding, onboarded.
func (m *RemoteAssistancePartner) GetOnboardingStatus()(*RemoteAssistanceOnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// GetOnboardingUrl gets the onboardingUrl property value. URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
func (m *RemoteAssistancePartner) GetOnboardingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onboardingUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoteAssistancePartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["onboardingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAssistanceOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RemoteAssistanceOnboardingStatus)
            m.SetOnboardingStatus(&cast)
        }
        return nil
    }
    res["onboardingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingUrl(val)
        }
        return nil
    }
    return res
}
func (m *RemoteAssistancePartner) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RemoteAssistancePartner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetOnboardingStatus() != nil {
        cast := m.GetOnboardingStatus().String()
        err = writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onboardingUrl", m.GetOnboardingUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Display name of the partner.
func (m *RemoteAssistancePartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *RemoteAssistancePartner) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// SetOnboardingStatus sets the onboardingStatus property value. A friendly description of the current TeamViewer connector status. Possible values are: notOnboarded, onboarding, onboarded.
func (m *RemoteAssistancePartner) SetOnboardingStatus(value *RemoteAssistanceOnboardingStatus)() {
    m.onboardingStatus = value
}
// SetOnboardingUrl sets the onboardingUrl property value. URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
func (m *RemoteAssistancePartner) SetOnboardingUrl(value *string)() {
    m.onboardingUrl = value
}
