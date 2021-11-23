package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// timeOffReason 
type TimeOffReason struct {
    ChangeTrackedEntity
    // The name of the timeOffReason. Required.
    displayName *string;
    // Supported icon types: none; car; calendar; running; plane; firstAid; doctor; notWorking; clock; juryDuty; globe; cup; phone; weather; umbrella; piggyBank; dog; cake; trafficCone; pin; sunny. Required.
    iconType *TimeOffReasonIconType;
    // Indicates whether the timeOffReason can be used when creating new entities or updating existing ones. Required.
    isActive *bool;
}
// NewTimeOffReason instantiates a new timeOffReason and sets the default values.
func NewTimeOffReason()(*TimeOffReason) {
    m := &TimeOffReason{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The name of the timeOffReason. Required.
func (m *TimeOffReason) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIconType gets the iconType property value. Supported icon types: none; car; calendar; running; plane; firstAid; doctor; notWorking; clock; juryDuty; globe; cup; phone; weather; umbrella; piggyBank; dog; cake; trafficCone; pin; sunny. Required.
func (m *TimeOffReason) GetIconType()(*TimeOffReasonIconType) {
    if m == nil {
        return nil
    } else {
        return m.iconType
    }
}
// GetIsActive gets the isActive property value. Indicates whether the timeOffReason can be used when creating new entities or updating existing ones. Required.
func (m *TimeOffReason) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeOffReason) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
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
    res["iconType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTimeOffReasonIconType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TimeOffReasonIconType)
            m.SetIconType(&cast)
        }
        return nil
    }
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    return res
}
func (m *TimeOffReason) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeOffReason) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetIconType() != nil {
        cast := m.GetIconType().String()
        err = writer.WriteStringValue("iconType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the timeOffReason. Required.
func (m *TimeOffReason) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIconType sets the iconType property value. Supported icon types: none; car; calendar; running; plane; firstAid; doctor; notWorking; clock; juryDuty; globe; cup; phone; weather; umbrella; piggyBank; dog; cake; trafficCone; pin; sunny. Required.
func (m *TimeOffReason) SetIconType(value *TimeOffReasonIconType)() {
    m.iconType = value
}
// SetIsActive sets the isActive property value. Indicates whether the timeOffReason can be used when creating new entities or updating existing ones. Required.
func (m *TimeOffReason) SetIsActive(value *bool)() {
    m.isActive = value
}
