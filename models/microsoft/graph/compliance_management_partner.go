package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ComplianceManagementPartner struct {
    Entity
    androidEnrollmentAssignments []ComplianceManagementPartnerAssignment;
    androidOnboarded *bool;
    displayName *string;
    iosEnrollmentAssignments []ComplianceManagementPartnerAssignment;
    iosOnboarded *bool;
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    macOsEnrollmentAssignments []ComplianceManagementPartnerAssignment;
    macOsOnboarded *bool;
    partnerState *DeviceManagementPartnerTenantState;
}
func NewComplianceManagementPartner()(*ComplianceManagementPartner) {
    m := &ComplianceManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ComplianceManagementPartner) GetAndroidEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.androidEnrollmentAssignments
    }
}
func (m *ComplianceManagementPartner) GetAndroidOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidOnboarded
    }
}
func (m *ComplianceManagementPartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ComplianceManagementPartner) GetIosEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.iosEnrollmentAssignments
    }
}
func (m *ComplianceManagementPartner) GetIosOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosOnboarded
    }
}
func (m *ComplianceManagementPartner) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
func (m *ComplianceManagementPartner) GetMacOsEnrollmentAssignments()([]ComplianceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.macOsEnrollmentAssignments
    }
}
func (m *ComplianceManagementPartner) GetMacOsOnboarded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macOsOnboarded
    }
}
func (m *ComplianceManagementPartner) GetPartnerState()(*DeviceManagementPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
func (m *ComplianceManagementPartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidEnrollmentAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartnerAssignment() })
        if err != nil {
            return err
        }
        res := make([]ComplianceManagementPartnerAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*ComplianceManagementPartnerAssignment))
        }
        m.SetAndroidEnrollmentAssignments(res)
        return nil
    }
    res["androidOnboarded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAndroidOnboarded(val)
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
    res["iosEnrollmentAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartnerAssignment() })
        if err != nil {
            return err
        }
        res := make([]ComplianceManagementPartnerAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*ComplianceManagementPartnerAssignment))
        }
        m.SetIosEnrollmentAssignments(res)
        return nil
    }
    res["iosOnboarded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIosOnboarded(val)
        return nil
    }
    res["lastHeartbeatDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastHeartbeatDateTime(val)
        return nil
    }
    res["macOsEnrollmentAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartnerAssignment() })
        if err != nil {
            return err
        }
        res := make([]ComplianceManagementPartnerAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*ComplianceManagementPartnerAssignment))
        }
        m.SetMacOsEnrollmentAssignments(res)
        return nil
    }
    res["macOsOnboarded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMacOsOnboarded(val)
        return nil
    }
    res["partnerState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerTenantState)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementPartnerTenantState)
        m.SetPartnerState(&cast)
        return nil
    }
    return res
}
func (m *ComplianceManagementPartner) IsNil()(bool) {
    return m == nil
}
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
func (m *ComplianceManagementPartner) SetAndroidEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    m.androidEnrollmentAssignments = value
}
func (m *ComplianceManagementPartner) SetAndroidOnboarded(value *bool)() {
    m.androidOnboarded = value
}
func (m *ComplianceManagementPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ComplianceManagementPartner) SetIosEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    m.iosEnrollmentAssignments = value
}
func (m *ComplianceManagementPartner) SetIosOnboarded(value *bool)() {
    m.iosOnboarded = value
}
func (m *ComplianceManagementPartner) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
func (m *ComplianceManagementPartner) SetMacOsEnrollmentAssignments(value []ComplianceManagementPartnerAssignment)() {
    m.macOsEnrollmentAssignments = value
}
func (m *ComplianceManagementPartner) SetMacOsOnboarded(value *bool)() {
    m.macOsOnboarded = value
}
func (m *ComplianceManagementPartner) SetPartnerState(value *DeviceManagementPartnerTenantState)() {
    m.partnerState = value
}
