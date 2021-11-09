package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedDeviceOverview struct {
    Entity
    // Distribution of Exchange Access State in Intune
    deviceExchangeAccessStateSummary *DeviceExchangeAccessStateSummary;
    // Device operating system summary.
    deviceOperatingSystemSummary *DeviceOperatingSystemSummary;
    // The number of devices enrolled in both MDM and EAS
    dualEnrolledDeviceCount *int32;
    // Total enrolled device count. Does not include PC devices managed via Intune PC Agent
    enrolledDeviceCount *int32;
    // The number of devices enrolled in MDM
    mdmEnrolledCount *int32;
}
// Instantiates a new managedDeviceOverview and sets the default values.
func NewManagedDeviceOverview()(*ManagedDeviceOverview) {
    m := &ManagedDeviceOverview{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceExchangeAccessStateSummary property value. Distribution of Exchange Access State in Intune
func (m *ManagedDeviceOverview) GetDeviceExchangeAccessStateSummary()(*DeviceExchangeAccessStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceExchangeAccessStateSummary
    }
}
// Gets the deviceOperatingSystemSummary property value. Device operating system summary.
func (m *ManagedDeviceOverview) GetDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceOperatingSystemSummary
    }
}
// Gets the dualEnrolledDeviceCount property value. The number of devices enrolled in both MDM and EAS
func (m *ManagedDeviceOverview) GetDualEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dualEnrolledDeviceCount
    }
}
// Gets the enrolledDeviceCount property value. Total enrolled device count. Does not include PC devices managed via Intune PC Agent
func (m *ManagedDeviceOverview) GetEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDeviceCount
    }
}
// Gets the mdmEnrolledCount property value. The number of devices enrolled in MDM
func (m *ManagedDeviceOverview) GetMdmEnrolledCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mdmEnrolledCount
    }
}
// The deserialization information for the current model
func (m *ManagedDeviceOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceExchangeAccessStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceExchangeAccessStateSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceExchangeAccessStateSummary(val.(*DeviceExchangeAccessStateSummary))
        }
        return nil
    }
    res["deviceOperatingSystemSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceOperatingSystemSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOperatingSystemSummary(val.(*DeviceOperatingSystemSummary))
        }
        return nil
    }
    res["dualEnrolledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDualEnrolledDeviceCount(val)
        }
        return nil
    }
    res["enrolledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDeviceCount(val)
        }
        return nil
    }
    res["mdmEnrolledCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmEnrolledCount(val)
        }
        return nil
    }
    return res
}
func (m *ManagedDeviceOverview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedDeviceOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceExchangeAccessStateSummary", m.GetDeviceExchangeAccessStateSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceOperatingSystemSummary", m.GetDeviceOperatingSystemSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("dualEnrolledDeviceCount", m.GetDualEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enrolledDeviceCount", m.GetEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mdmEnrolledCount", m.GetMdmEnrolledCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceExchangeAccessStateSummary property value. Distribution of Exchange Access State in Intune
// Parameters:
//  - value : Value to set for the deviceExchangeAccessStateSummary property.
func (m *ManagedDeviceOverview) SetDeviceExchangeAccessStateSummary(value *DeviceExchangeAccessStateSummary)() {
    m.deviceExchangeAccessStateSummary = value
}
// Sets the deviceOperatingSystemSummary property value. Device operating system summary.
// Parameters:
//  - value : Value to set for the deviceOperatingSystemSummary property.
func (m *ManagedDeviceOverview) SetDeviceOperatingSystemSummary(value *DeviceOperatingSystemSummary)() {
    m.deviceOperatingSystemSummary = value
}
// Sets the dualEnrolledDeviceCount property value. The number of devices enrolled in both MDM and EAS
// Parameters:
//  - value : Value to set for the dualEnrolledDeviceCount property.
func (m *ManagedDeviceOverview) SetDualEnrolledDeviceCount(value *int32)() {
    m.dualEnrolledDeviceCount = value
}
// Sets the enrolledDeviceCount property value. Total enrolled device count. Does not include PC devices managed via Intune PC Agent
// Parameters:
//  - value : Value to set for the enrolledDeviceCount property.
func (m *ManagedDeviceOverview) SetEnrolledDeviceCount(value *int32)() {
    m.enrolledDeviceCount = value
}
// Sets the mdmEnrolledCount property value. The number of devices enrolled in MDM
// Parameters:
//  - value : Value to set for the mdmEnrolledCount property.
func (m *ManagedDeviceOverview) SetMdmEnrolledCount(value *int32)() {
    m.mdmEnrolledCount = value
}
