package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SolutionsRoot 
type SolutionsRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    bookingBusinesses []BookingBusiness;
    // 
    bookingCurrencies []BookingCurrency;
}
// NewSolutionsRoot instantiates a new SolutionsRoot and sets the default values.
func NewSolutionsRoot()(*SolutionsRoot) {
    m := &SolutionsRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsRoot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBookingBusinesses gets the bookingBusinesses property value. 
func (m *SolutionsRoot) GetBookingBusinesses()([]BookingBusiness) {
    if m == nil {
        return nil
    } else {
        return m.bookingBusinesses
    }
}
// GetBookingCurrencies gets the bookingCurrencies property value. 
func (m *SolutionsRoot) GetBookingCurrencies()([]BookingCurrency) {
    if m == nil {
        return nil
    } else {
        return m.bookingCurrencies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SolutionsRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bookingBusinesses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingBusiness() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingBusiness, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingBusiness))
            }
            m.SetBookingBusinesses(res)
        }
        return nil
    }
    res["bookingCurrencies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingCurrency() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCurrency, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingCurrency))
            }
            m.SetBookingCurrencies(res)
        }
        return nil
    }
    return res
}
func (m *SolutionsRoot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SolutionsRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetBookingBusinesses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBookingBusinesses()))
        for i, v := range m.GetBookingBusinesses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("bookingBusinesses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBookingCurrencies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBookingCurrencies()))
        for i, v := range m.GetBookingCurrencies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// SetBookingBusinesses sets the bookingBusinesses property value. 
func (m *SolutionsRoot) SetBookingBusinesses(value []BookingBusiness)() {
    if m != nil {
        m.bookingBusinesses = value
    }
}
// SetBookingCurrencies sets the bookingCurrencies property value. 
func (m *SolutionsRoot) SetBookingCurrencies(value []BookingCurrency)() {
    if m != nil {
        m.bookingCurrencies = value
    }
}
