package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomer represents a customer of the business.
type BookingCustomer struct {
    BookingCustomerBase
}
// NewBookingCustomer instantiates a new bookingCustomer and sets the default values.
func NewBookingCustomer()(*BookingCustomer) {
    m := &BookingCustomer{
        BookingCustomerBase: *NewBookingCustomerBase(),
    }
    odataTypeValue := "#microsoft.graph.bookingCustomer"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBookingCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingCustomer(), nil
}
// GetAddresses gets the addresses property value. Addresses associated with the customer. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingCustomer) GetAddresses()([]PhysicalAddressable) {
    val, err := m.GetBackingStore().Get("addresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PhysicalAddressable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the customer.
func (m *BookingCustomer) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailAddress gets the emailAddress property value. The SMTP address of the customer.
func (m *BookingCustomer) GetEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("emailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingCustomerBase.GetFieldDeserializers()
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhysicalAddressable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PhysicalAddressable)
                }
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phoneable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Phoneable)
                }
            }
            m.SetPhones(res)
        }
        return nil
    }
    return res
}
// GetPhones gets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
func (m *BookingCustomer) GetPhones()([]Phoneable) {
    val, err := m.GetBackingStore().Get("phones")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Phoneable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BookingCustomer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingCustomerBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
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
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddresses sets the addresses property value. Addresses associated with the customer. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingCustomer) SetAddresses(value []PhysicalAddressable)() {
    err := m.GetBackingStore().Set("addresses", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the customer.
func (m *BookingCustomer) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddress sets the emailAddress property value. The SMTP address of the customer.
func (m *BookingCustomer) SetEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("emailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetPhones sets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
func (m *BookingCustomer) SetPhones(value []Phoneable)() {
    err := m.GetBackingStore().Set("phones", value)
    if err != nil {
        panic(err)
    }
}
// BookingCustomerable 
type BookingCustomerable interface {
    BookingCustomerBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddresses()([]PhysicalAddressable)
    GetDisplayName()(*string)
    GetEmailAddress()(*string)
    GetPhones()([]Phoneable)
    SetAddresses(value []PhysicalAddressable)()
    SetDisplayName(value *string)()
    SetEmailAddress(value *string)()
    SetPhones(value []Phoneable)()
}
