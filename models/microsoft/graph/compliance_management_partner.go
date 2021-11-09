package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ComplianceManagementPartner struct {
    Entity
    // User groups which enroll Android devices through partner.
    androidEnrollmentAssignments []ComplianceManagementPartnerAssignment;
    // Partner onboarded for Android devices.
    androidOnboarded *bool;
    // Partner display name
    displayName *string;
    // User groups which enroll ios devices through partner.
    iosEnrollmentAssignments []ComplianceManagementPartnerAssignment;
    // Partner onboarded for ios devices.
    iosOnboarded *bool;
    // Timestamp of last heartbeat after admin onboarded to the compliance management partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // User groups which enroll Mac devices through partner.
    macOsEnrollmentAssignments []ComplianceManagementPartnerAssignment;
    // Partner onboarded for Mac devices.
    macOsOnboarded *bool;
    // Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
    partnerState *DeviceManagementPartnerTenantState;
}
// Instantiates a new complianceManagementPartner and sets the default values.
func NewComplianceManagementPartner()(*ComplianceManagementPartner) {
    m := &ComplianceManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the androidEnrollmentAssignments property value. User groups which enroll Android devices through partner.
func (m *ComplianceManagementPartner) GetAndroidEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.androidEnrollmentAssignments
    }
}
// Gets the androidOnboarded property value. Partner onboarded for Android devices.
func (m *ComplianceManagementPartner) GetAndroidOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidOnboarded
    }
}
// Gets the displayName property value. Partner display name
func (m *ComplianceManagementPartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the iosEnrollmentAssignments property value. User groups which enroll ios devices through partner.
func (m *ComplianceManagementPartner) GetIosEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.iosEnrollmentAssignments
    }
}
// Gets the iosOnboarded property value. Partner onboarded for ios devices.
func (m *ComplianceManagementPartner) GetIosOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosOnboarded
    }
}
// Gets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin onboarded to the compliance management partner
func (m *ComplianceManagementPartner) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
// Gets the macOsEnrollmentAssignments property value. User groups which enroll Mac devices through partner.
func (m *ComplianceManagementPartner) GetMacOsEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.macOsEnrollmentAssignments
    }
}
// Gets the macOsOnboarded property value. Partner onboarded for Mac devices.
func (m *ComplianceManagementPartner) GetMacOsOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macOsOnboarded
    }
}
// Gets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
func (m *ComplianceManagementPartner) GetPartnerState()(*DeviceManagementPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
// The deserialization information for the current model
func (m *ComplianceManagementPartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidEnrollmentAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartnerAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*ComplianceManagementPartnerAssignment))
            }
            m.SetAndroidEnrollmentAssignments(res)
        }
        return nil
    }
    res["androidOnboarded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidOnboarded(val)
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
    res["iosEnrollmentAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartnerAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*ComplianceManagementPartnerAssignment))
            }
            m.SetIosEnrollmentAssignments(res)
        }
        return nil
    }
    res["iosOnboarded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosOnboarded(val)
        }
        return nil
    }
    res["lastHeartbeatDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHeartbeatDateTime(val)
        }
        return nil
    }
    res["macOsEnrollmentAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartnerAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*ComplianceManagementPartnerAssignment))
            }
            m.SetMacOsEnrollmentAssignments(res)
        }
        return nil
    }
    res["macOsOnboarded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOsOnboarded(val)
        }
        return nil
    }
    res["partnerState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementPartnerTenantState)
            m.SetPartnerState(&cast)
        }
        return nil
    }
    return res
}
func (m *ComplianceManagementPartner) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ComplianceManagementPartner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAndroidEnrollmentAssignments()))
        for i, v := range m.GetAndroidEnrollmentAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("androidEnrollmentAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidOnboarded", m.GetAndroidOnboarded())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIosEnrollmentAssignments()))
        for i, v := range m.GetIosEnrollmentAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("iosEnrollmentAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosOnboarded", m.GetIosOnboarded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastHeartbeatDateTime", m.GetLastHeartbeatDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMacOsEnrollmentAssignments()))
        for i, v := range m.GetMacOsEnrollmentAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("macOsEnrollmentAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("macOsOnboarded", m.GetMacOsOnboarded())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := m.GetPartnerState().String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the androidEnrollmentAssignments property value. User groups which enroll Android devices through partner.
// Parameters:
//  - value : Value to set for the androidEnrollmentAssignments property.
func (m *ComplianceManagementPartner) SetAndroidEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    m.androidEnrollmentAssignments = value
}
// Sets the androidOnboarded property value. Partner onboarded for Android devices.
// Parameters:
//  - value : Value to set for the androidOnboarded property.
func (m *ComplianceManagementPartner) SetAndroidOnboarded(value *bool)() {
    m.androidOnboarded = value
}
// Sets the displayName property value. Partner display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ComplianceManagementPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the iosEnrollmentAssignments property value. User groups which enroll ios devices through partner.
// Parameters:
//  - value : Value to set for the iosEnrollmentAssignments property.
func (m *ComplianceManagementPartner) SetIosEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    m.iosEnrollmentAssignments = value
}
// Sets the iosOnboarded property value. Partner onboarded for ios devices.
// Parameters:
//  - value : Value to set for the iosOnboarded property.
func (m *ComplianceManagementPartner) SetIosOnboarded(value *bool)() {
    m.iosOnboarded = value
}
// Sets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin onboarded to the compliance management partner
// Parameters:
//  - value : Value to set for the lastHeartbeatDateTime property.
func (m *ComplianceManagementPartner) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
// Sets the macOsEnrollmentAssignments property value. User groups which enroll Mac devices through partner.
// Parameters:
//  - value : Value to set for the macOsEnrollmentAssignments property.
func (m *ComplianceManagementPartner) SetMacOsEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    m.macOsEnrollmentAssignments = value
}
// Sets the macOsOnboarded property value. Partner onboarded for Mac devices.
// Parameters:
//  - value : Value to set for the macOsOnboarded property.
func (m *ComplianceManagementPartner) SetMacOsOnboarded(value *bool)() {
    m.macOsOnboarded = value
}
// Sets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
// Parameters:
//  - value : Value to set for the partnerState property.
func (m *ComplianceManagementPartner) SetPartnerState(value *DeviceManagementPartnerTenantState)() {
    m.partnerState = value
}
