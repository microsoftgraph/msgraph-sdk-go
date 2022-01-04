package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingCurrency 
type BookingCurrency struct {
    Entity
    // The currency symbol. For example, the currency symbol for the US dollar and for the Australian dollar is $.
    symbol *string;
}
// NewBookingCurrency instantiates a new bookingCurrency and sets the default values.
func NewBookingCurrency()(*BookingCurrency) {
    m := &BookingCurrency{
        Entity: *NewEntity(),
    }
    return m
}
// GetSymbol gets the symbol property value. The currency symbol. For example, the currency symbol for the US dollar and for the Australian dollar is $.
func (m *BookingCurrency) GetSymbol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.symbol
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCurrency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["symbol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSymbol(val)
        }
        return nil
    }
    return res
}
func (m *BookingCurrency) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingCurrency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("symbol", m.GetSymbol())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSymbol sets the symbol property value. The currency symbol. For example, the currency symbol for the US dollar and for the Australian dollar is $.
func (m *BookingCurrency) SetSymbol(value *string)() {
    if m != nil {
        m.symbol = value
    }
}
