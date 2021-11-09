package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BitlockerRecoveryKey struct {
    Entity
    // The date and time when the key was originally backed up to Azure Active Directory. Not nullable.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identifier of the device the BitLocker key is originally backed up from. Supports $filter (eq).
    deviceId *string;
    // The BitLocker recovery key. Returned only on $select. Not nullable.
    key *string;
    // Indicates the type of volume the BitLocker key is associated with. The possible values are: 1 (for operatingSystemVolume), 2 (for fixedDataVolume), 3 (for removableDataVolume), and 4 (for unknownFutureValue).
    volumeType *VolumeType;
}
// Instantiates a new bitlockerRecoveryKey and sets the default values.
func NewBitlockerRecoveryKey()(*BitlockerRecoveryKey) {
    m := &BitlockerRecoveryKey{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. The date and time when the key was originally backed up to Azure Active Directory. Not nullable.
func (m *BitlockerRecoveryKey) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deviceId property value. Identifier of the device the BitLocker key is originally backed up from. Supports $filter (eq).
func (m *BitlockerRecoveryKey) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the key property value. The BitLocker recovery key. Returned only on $select. Not nullable.
func (m *BitlockerRecoveryKey) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// Gets the volumeType property value. Indicates the type of volume the BitLocker key is associated with. The possible values are: 1 (for operatingSystemVolume), 2 (for fixedDataVolume), 3 (for removableDataVolume), and 4 (for unknownFutureValue).
func (m *BitlockerRecoveryKey) GetVolumeType()(*VolumeType) {
    if m == nil {
        return nil
    } else {
        return m.volumeType
    }
}
// The deserialization information for the current model
func (m *BitlockerRecoveryKey) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["volumeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseVolumeType)
        if err != nil {
            return err
        }
        cast := val.(VolumeType)
        m.SetVolumeType(&cast)
        return nil
    }
    return res
}
func (m *BitlockerRecoveryKey) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BitlockerRecoveryKey) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    if m.GetVolumeType() != nil {
        cast := m.GetVolumeType().String()
        err = writer.WriteStringValue("volumeType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. The date and time when the key was originally backed up to Azure Active Directory. Not nullable.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *BitlockerRecoveryKey) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deviceId property value. Identifier of the device the BitLocker key is originally backed up from. Supports $filter (eq).
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *BitlockerRecoveryKey) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the key property value. The BitLocker recovery key. Returned only on $select. Not nullable.
// Parameters:
//  - value : Value to set for the key property.
func (m *BitlockerRecoveryKey) SetKey(value *string)() {
    m.key = value
}
// Sets the volumeType property value. Indicates the type of volume the BitLocker key is associated with. The possible values are: 1 (for operatingSystemVolume), 2 (for fixedDataVolume), 3 (for removableDataVolume), and 4 (for unknownFutureValue).
// Parameters:
//  - value : Value to set for the volumeType property.
func (m *BitlockerRecoveryKey) SetVolumeType(value *VolumeType)() {
    m.volumeType = value
}
