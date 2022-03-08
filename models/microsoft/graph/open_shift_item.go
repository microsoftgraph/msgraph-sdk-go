package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OpenShiftItem provides operations to manage the drive singleton.
type OpenShiftItem struct {
    ShiftItem
    // Count of the number of slots for the given open shift.
    openSlotCount *int32;
}
// NewOpenShiftItem instantiates a new openShiftItem and sets the default values.
func NewOpenShiftItem()(*OpenShiftItem) {
    m := &OpenShiftItem{
        ShiftItem: *NewShiftItem(),
    }
    return m
}
// CreateOpenShiftItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenShiftItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOpenShiftItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenShiftItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ShiftItem.GetFieldDeserializers()
    res["openSlotCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenSlotCount(val)
        }
        return nil
    }
    return res
}
// GetOpenSlotCount gets the openSlotCount property value. Count of the number of slots for the given open shift.
func (m *OpenShiftItem) GetOpenSlotCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.openSlotCount
    }
}
func (m *OpenShiftItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OpenShiftItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ShiftItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("openSlotCount", m.GetOpenSlotCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOpenSlotCount sets the openSlotCount property value. Count of the number of slots for the given open shift.
func (m *OpenShiftItem) SetOpenSlotCount(value *int32)() {
    if m != nil {
        m.openSlotCount = value
    }
}
