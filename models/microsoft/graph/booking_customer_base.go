package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingCustomerBase provides operations to manage the solutionsRoot singleton.
type BookingCustomerBase struct {
    Entity
}
// NewBookingCustomerBase instantiates a new bookingCustomerBase and sets the default values.
func NewBookingCustomerBase()(*BookingCustomerBase) {
    m := &BookingCustomerBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingCustomerBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomerBaseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingCustomerBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomerBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *BookingCustomerBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingCustomerBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
