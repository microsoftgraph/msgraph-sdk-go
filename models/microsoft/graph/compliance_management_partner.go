package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ComplianceManagementPartner 
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
// NewComplianceManagementPartner instantiates a new complianceManagementPartner and sets the default values.
func NewComplianceManagementPartner()(*ComplianceManagementPartner) {
    m := &ComplianceManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// GetAndroidEnrollmentAssignments gets the androidEnrollmentAssignments property value. User groups which enroll Android devices through partner.
func (m *ComplianceManagementPartner) GetAndroidEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.androidEnrollmentAssignments
    }
}
// GetAndroidOnboarded gets the androidOnboarded property value. Partner onboarded for Android devices.
func (m *ComplianceManagementPartner) GetAndroidOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidOnboarded
    }
}
// GetDisplayName gets the displayName property value. Partner display name
func (m *ComplianceManagementPartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIosEnrollmentAssignments gets the iosEnrollmentAssignments property value. User groups which enroll ios devices through partner.
func (m *ComplianceManagementPartner) GetIosEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.iosEnrollmentAssignments
    }
}
// GetIosOnboarded gets the iosOnboarded property value. Partner onboarded for ios devices.
func (m *ComplianceManagementPartner) GetIosOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosOnboarded
    }
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin onboarded to the compliance management partner
func (m *ComplianceManagementPartner) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
// GetMacOsEnrollmentAssignments gets the macOsEnrollmentAssignments property value. User groups which enroll Mac devices through partner.
func (m *ComplianceManagementPartner) GetMacOsEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.macOsEnrollmentAssignments
    }
}
// GetMacOsOnboarded gets the macOsOnboarded property value. Partner onboarded for Mac devices.
func (m *ComplianceManagementPartner) GetMacOsOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macOsOnboarded
    }
}
// GetPartnerState gets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
func (m *ComplianceManagementPartner) GetPartnerState()(*DeviceManagementPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *ComplianceManagementPartner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAndroidEnrollmentAssignments() != nil {
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
    if m.GetIosEnrollmentAssignments() != nil {
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
    if m.GetMacOsEnrollmentAssignments() != nil {
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
// SetAndroidEnrollmentAssignments sets the androidEnrollmentAssignments property value. User groups which enroll Android devices through partner.
func (m *ComplianceManagementPartner) SetAndroidEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    if m != nil {
        m.androidEnrollmentAssignments = value
    }
}
// SetAndroidOnboarded sets the androidOnboarded property value. Partner onboarded for Android devices.
func (m *ComplianceManagementPartner) SetAndroidOnboarded(value *bool)() {
    if m != nil {
        m.androidOnboarded = value
    }
}
// SetDisplayName sets the displayName property value. Partner display name
func (m *ComplianceManagementPartner) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIosEnrollmentAssignments sets the iosEnrollmentAssignments property value. User groups which enroll ios devices through partner.
func (m *ComplianceManagementPartner) SetIosEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    if m != nil {
        m.iosEnrollmentAssignments = value
    }
}
// SetIosOnboarded sets the iosOnboarded property value. Partner onboarded for ios devices.
func (m *ComplianceManagementPartner) SetIosOnboarded(value *bool)() {
    if m != nil {
        m.iosOnboarded = value
    }
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin onboarded to the compliance management partner
func (m *ComplianceManagementPartner) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastHeartbeatDateTime = value
    }
}
// SetMacOsEnrollmentAssignments sets the macOsEnrollmentAssignments property value. User groups which enroll Mac devices through partner.
func (m *ComplianceManagementPartner) SetMacOsEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    if m != nil {
        m.macOsEnrollmentAssignments = value
    }
}
// SetMacOsOnboarded sets the macOsOnboarded property value. Partner onboarded for Mac devices.
func (m *ComplianceManagementPartner) SetMacOsOnboarded(value *bool)() {
    if m != nil {
        m.macOsOnboarded = value
    }
}
// SetPartnerState sets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
func (m *ComplianceManagementPartner) SetPartnerState(value *DeviceManagementPartnerTenantState)() {
    if m != nil {
        m.partnerState = value
    }
}
