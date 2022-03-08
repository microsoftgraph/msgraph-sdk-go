package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LocationConstraintItem provides operations to call the findMeetingTimes method.
type LocationConstraintItem struct {
    Location
    // If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
    resolveAvailability *bool;
}
// NewLocationConstraintItem instantiates a new locationConstraintItem and sets the default values.
func NewLocationConstraintItem()(*LocationConstraintItem) {
    m := &LocationConstraintItem{
        Location: *NewLocation(),
    }
    return m
}
// CreateLocationConstraintItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocationConstraintItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewLocationConstraintItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocationConstraintItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Location.GetFieldDeserializers()
    res["resolveAvailability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolveAvailability(val)
        }
        return nil
    }
    return res
}
// GetResolveAvailability gets the resolveAvailability property value. If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
func (m *LocationConstraintItem) GetResolveAvailability()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.resolveAvailability
    }
}
func (m *LocationConstraintItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LocationConstraintItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Location.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("resolveAvailability", m.GetResolveAvailability())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResolveAvailability sets the resolveAvailability property value. If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
func (m *LocationConstraintItem) SetResolveAvailability(value *bool)() {
    if m != nil {
        m.resolveAvailability = value
    }
}
