package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeOff provides operations to manage the drive singleton.
type TimeOff struct {
    ChangeTrackedEntity
    // The draft version of this timeOff that is viewable by managers. Required.
    draftTimeOff TimeOffItemable;
    // The shared version of this timeOff that is viewable by both employees and managers. Required.
    sharedTimeOff TimeOffItemable;
    // ID of the user assigned to the timeOff. Required.
    userId *string;
}
// NewTimeOff instantiates a new timeOff and sets the default values.
func NewTimeOff()(*TimeOff) {
    m := &TimeOff{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// CreateTimeOffFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeOffFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTimeOff(), nil
}
// GetDraftTimeOff gets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
func (m *TimeOff) GetDraftTimeOff()(TimeOffItemable) {
    if m == nil {
        return nil
    } else {
        return m.draftTimeOff
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeOff) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftTimeOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeOffItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftTimeOff(val.(TimeOffItemable))
        }
        return nil
    }
    res["sharedTimeOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeOffItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedTimeOff(val.(TimeOffItemable))
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetSharedTimeOff gets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
func (m *TimeOff) GetSharedTimeOff()(TimeOffItemable) {
    if m == nil {
        return nil
    } else {
        return m.sharedTimeOff
    }
}
// GetUserId gets the userId property value. ID of the user assigned to the timeOff. Required.
func (m *TimeOff) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *TimeOff) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeOff) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftTimeOff", m.GetDraftTimeOff())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedTimeOff", m.GetSharedTimeOff())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftTimeOff sets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
func (m *TimeOff) SetDraftTimeOff(value TimeOffItemable)() {
    if m != nil {
        m.draftTimeOff = value
    }
}
// SetSharedTimeOff sets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
func (m *TimeOff) SetSharedTimeOff(value TimeOffItemable)() {
    if m != nil {
        m.sharedTimeOff = value
    }
}
// SetUserId sets the userId property value. ID of the user assigned to the timeOff. Required.
func (m *TimeOff) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
