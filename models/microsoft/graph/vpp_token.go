package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type VppToken struct {
    Entity
    // The apple Id associated with the given Apple Volume Purchase Program Token.
    appleId *string;
    // Whether or not apps for the VPP token will be automatically updated.
    automaticallyUpdateApps *bool;
    // Whether or not apps for the VPP token will be automatically updated.
    countryOrRegion *string;
    // The expiration date time of the Apple Volume Purchase Program Token.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last modification date time associated with the Apple Volume Purchase Program Token.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed.
    lastSyncStatus *VppTokenSyncStatus;
    // The organization associated with the Apple Volume Purchase Program Token
    organizationName *string;
    // Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM.
    state *VppTokenState;
    // The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
    token *string;
    // The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education.
    vppTokenAccountType *VppTokenAccountType;
}
// Instantiates a new vppToken and sets the default values.
func NewVppToken()(*VppToken) {
    m := &VppToken{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appleId property value. The apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppToken) GetAppleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleId
    }
}
// Gets the automaticallyUpdateApps property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) GetAutomaticallyUpdateApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.automaticallyUpdateApps
    }
}
// Gets the countryOrRegion property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// Gets the expirationDateTime property value. The expiration date time of the Apple Volume Purchase Program Token.
func (m *VppToken) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the lastModifiedDateTime property value. Last modification date time associated with the Apple Volume Purchase Program Token.
func (m *VppToken) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the lastSyncDateTime property value. The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
func (m *VppToken) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// Gets the lastSyncStatus property value. Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed.
func (m *VppToken) GetLastSyncStatus()(*VppTokenSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncStatus
    }
}
// Gets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *VppToken) GetOrganizationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationName
    }
}
// Gets the state property value. Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM.
func (m *VppToken) GetState()(*VppTokenState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the token property value. The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
func (m *VppToken) GetToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.token
    }
}
// Gets the vppTokenAccountType property value. The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education.
func (m *VppToken) GetVppTokenAccountType()(*VppTokenAccountType) {
    if m == nil {
        return nil
    } else {
        return m.vppTokenAccountType
    }
}
// The deserialization information for the current model
func (m *VppToken) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppleId(val)
        return nil
    }
    res["automaticallyUpdateApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutomaticallyUpdateApps(val)
        return nil
    }
    res["countryOrRegion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryOrRegion(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["lastSyncStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenSyncStatus)
        if err != nil {
            return err
        }
        cast := val.(VppTokenSyncStatus)
        m.SetLastSyncStatus(&cast)
        return nil
    }
    res["organizationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrganizationName(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenState)
        if err != nil {
            return err
        }
        cast := val.(VppTokenState)
        m.SetState(&cast)
        return nil
    }
    res["token"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetToken(val)
        return nil
    }
    res["vppTokenAccountType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenAccountType)
        if err != nil {
            return err
        }
        cast := val.(VppTokenAccountType)
        m.SetVppTokenAccountType(&cast)
        return nil
    }
    return res
}
func (m *VppToken) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *VppToken) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("automaticallyUpdateApps", m.GetAutomaticallyUpdateApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
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
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLastSyncStatus() != nil {
        cast := m.GetLastSyncStatus().String()
        err = writer.WriteStringValue("lastSyncStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationName", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("token", m.GetToken())
        if err != nil {
            return err
        }
    }
    if m.GetVppTokenAccountType() != nil {
        cast := m.GetVppTokenAccountType().String()
        err = writer.WriteStringValue("vppTokenAccountType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appleId property value. The apple Id associated with the given Apple Volume Purchase Program Token.
// Parameters:
//  - value : Value to set for the appleId property.
func (m *VppToken) SetAppleId(value *string)() {
    m.appleId = value
}
// Sets the automaticallyUpdateApps property value. Whether or not apps for the VPP token will be automatically updated.
// Parameters:
//  - value : Value to set for the automaticallyUpdateApps property.
func (m *VppToken) SetAutomaticallyUpdateApps(value *bool)() {
    m.automaticallyUpdateApps = value
}
// Sets the countryOrRegion property value. Whether or not apps for the VPP token will be automatically updated.
// Parameters:
//  - value : Value to set for the countryOrRegion property.
func (m *VppToken) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// Sets the expirationDateTime property value. The expiration date time of the Apple Volume Purchase Program Token.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *VppToken) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the lastModifiedDateTime property value. Last modification date time associated with the Apple Volume Purchase Program Token.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *VppToken) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the lastSyncDateTime property value. The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
// Parameters:
//  - value : Value to set for the lastSyncDateTime property.
func (m *VppToken) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// Sets the lastSyncStatus property value. Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed.
// Parameters:
//  - value : Value to set for the lastSyncStatus property.
func (m *VppToken) SetLastSyncStatus(value *VppTokenSyncStatus)() {
    m.lastSyncStatus = value
}
// Sets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token
// Parameters:
//  - value : Value to set for the organizationName property.
func (m *VppToken) SetOrganizationName(value *string)() {
    m.organizationName = value
}
// Sets the state property value. Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM.
// Parameters:
//  - value : Value to set for the state property.
func (m *VppToken) SetState(value *VppTokenState)() {
    m.state = value
}
// Sets the token property value. The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
// Parameters:
//  - value : Value to set for the token property.
func (m *VppToken) SetToken(value *string)() {
    m.token = value
}
// Sets the vppTokenAccountType property value. The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education.
// Parameters:
//  - value : Value to set for the vppTokenAccountType property.
func (m *VppToken) SetVppTokenAccountType(value *VppTokenAccountType)() {
    m.vppTokenAccountType = value
}
