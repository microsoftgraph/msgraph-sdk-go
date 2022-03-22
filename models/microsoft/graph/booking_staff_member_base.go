package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingStaffMemberBase 
type BookingStaffMemberBase struct {
    Entity
}
// NewBookingStaffMemberBase instantiates a new bookingStaffMemberBase and sets the default values.
func NewBookingStaffMemberBase()(*BookingStaffMemberBase) {
    m := &BookingStaffMemberBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingStaffMemberBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingStaffMemberBaseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingStaffMemberBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingStaffMemberBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *BookingStaffMemberBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
