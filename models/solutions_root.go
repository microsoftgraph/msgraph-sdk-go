package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SolutionsRoot 
type SolutionsRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The bookingBusinesses property
    bookingBusinesses []BookingBusinessable
    // The bookingCurrencies property
    bookingCurrencies []BookingCurrencyable
}
// NewSolutionsRoot instantiates a new SolutionsRoot and sets the default values.
func NewSolutionsRoot()(*SolutionsRoot) {
    m := &SolutionsRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSolutionsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSolutionsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSolutionsRoot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsRoot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBookingBusinesses gets the bookingBusinesses property value. The bookingBusinesses property
func (m *SolutionsRoot) GetBookingBusinesses()([]BookingBusinessable) {
    if m == nil {
        return nil
    } else {
        return m.bookingBusinesses
    }
}
// GetBookingCurrencies gets the bookingCurrencies property value. The bookingCurrencies property
func (m *SolutionsRoot) GetBookingCurrencies()([]BookingCurrencyable) {
    if m == nil {
        return nil
    } else {
        return m.bookingCurrencies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SolutionsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bookingBusinesses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingBusinessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingBusinessable, len(val))
            for i, v := range val {
                res[i] = v.(BookingBusinessable)
            }
            m.SetBookingBusinesses(res)
        }
        return nil
    }
    res["bookingCurrencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCurrencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCurrencyable, len(val))
            for i, v := range val {
                res[i] = v.(BookingCurrencyable)
            }
            m.SetBookingCurrencies(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SolutionsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBookingBusinesses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBookingBusinesses()))
        for i, v := range m.GetBookingBusinesses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("bookingBusinesses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBookingCurrencies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBookingCurrencies()))
        for i, v := range m.GetBookingCurrencies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("bookingCurrencies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsRoot) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBookingBusinesses sets the bookingBusinesses property value. The bookingBusinesses property
func (m *SolutionsRoot) SetBookingBusinesses(value []BookingBusinessable)() {
    if m != nil {
        m.bookingBusinesses = value
    }
}
// SetBookingCurrencies sets the bookingCurrencies property value. The bookingCurrencies property
func (m *SolutionsRoot) SetBookingCurrencies(value []BookingCurrencyable)() {
    if m != nil {
        m.bookingCurrencies = value
    }
}
