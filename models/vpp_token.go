package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppToken 
type VppToken struct {
    Entity
    // The apple Id associated with the given Apple Volume Purchase Program Token.
    appleId *string
    // Whether or not apps for the VPP token will be automatically updated.
    automaticallyUpdateApps *bool
    // Whether or not apps for the VPP token will be automatically updated.
    countryOrRegion *string
    // The expiration date time of the Apple Volume Purchase Program Token.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last modification date time associated with the Apple Volume Purchase Program Token.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed.
    lastSyncStatus *VppTokenSyncStatus
    // The organization associated with the Apple Volume Purchase Program Token
    organizationName *string
    // Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM, duplicateLocationId.
    state *VppTokenState
    // The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
    token *string
    // The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education.
    vppTokenAccountType *VppTokenAccountType
}
// NewVppToken instantiates a new vppToken and sets the default values.
func NewVppToken()(*VppToken) {
    m := &VppToken{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVppTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVppTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppToken(), nil
}
// GetAppleId gets the appleId property value. The apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppToken) GetAppleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleId
    }
}
// GetAutomaticallyUpdateApps gets the automaticallyUpdateApps property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) GetAutomaticallyUpdateApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.automaticallyUpdateApps
    }
}
// GetCountryOrRegion gets the countryOrRegion property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date time of the Apple Volume Purchase Program Token.
func (m *VppToken) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VppToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleId(val)
        }
        return nil
    }
    res["automaticallyUpdateApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticallyUpdateApps(val)
        }
        return nil
    }
    res["countryOrRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegion(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
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
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["lastSyncStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncStatus(val.(*VppTokenSyncStatus))
        }
        return nil
    }
    res["organizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*VppTokenState))
        }
        return nil
    }
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    res["vppTokenAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenAccountType(val.(*VppTokenAccountType))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification date time associated with the Apple Volume Purchase Program Token.
func (m *VppToken) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
func (m *VppToken) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetLastSyncStatus gets the lastSyncStatus property value. Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed.
func (m *VppToken) GetLastSyncStatus()(*VppTokenSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncStatus
    }
}
// GetOrganizationName gets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *VppToken) GetOrganizationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationName
    }
}
// GetState gets the state property value. Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM, duplicateLocationId.
func (m *VppToken) GetState()(*VppTokenState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetToken gets the token property value. The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
func (m *VppToken) GetToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.token
    }
}
// GetVppTokenAccountType gets the vppTokenAccountType property value. The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education.
func (m *VppToken) GetVppTokenAccountType()(*VppTokenAccountType) {
    if m == nil {
        return nil
    } else {
        return m.vppTokenAccountType
    }
}
// Serialize serializes information the current object
func (m *VppToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetLastSyncStatus()).String()
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
        cast := (*m.GetState()).String()
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
        cast := (*m.GetVppTokenAccountType()).String()
        err = writer.WriteStringValue("vppTokenAccountType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppleId sets the appleId property value. The apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppToken) SetAppleId(value *string)() {
    if m != nil {
        m.appleId = value
    }
}
// SetAutomaticallyUpdateApps sets the automaticallyUpdateApps property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) SetAutomaticallyUpdateApps(value *bool)() {
    if m != nil {
        m.automaticallyUpdateApps = value
    }
}
// SetCountryOrRegion sets the countryOrRegion property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) SetCountryOrRegion(value *string)() {
    if m != nil {
        m.countryOrRegion = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date time of the Apple Volume Purchase Program Token.
func (m *VppToken) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification date time associated with the Apple Volume Purchase Program Token.
func (m *VppToken) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
func (m *VppToken) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetLastSyncStatus sets the lastSyncStatus property value. Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed.
func (m *VppToken) SetLastSyncStatus(value *VppTokenSyncStatus)() {
    if m != nil {
        m.lastSyncStatus = value
    }
}
// SetOrganizationName sets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *VppToken) SetOrganizationName(value *string)() {
    if m != nil {
        m.organizationName = value
    }
}
// SetState sets the state property value. Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM, duplicateLocationId.
func (m *VppToken) SetState(value *VppTokenState)() {
    if m != nil {
        m.state = value
    }
}
// SetToken sets the token property value. The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
func (m *VppToken) SetToken(value *string)() {
    if m != nil {
        m.token = value
    }
}
// SetVppTokenAccountType sets the vppTokenAccountType property value. The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education.
func (m *VppToken) SetVppTokenAccountType(value *VppTokenAccountType)() {
    if m != nil {
        m.vppTokenAccountType = value
    }
}
