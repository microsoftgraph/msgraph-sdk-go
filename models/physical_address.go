package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PhysicalAddress 
type PhysicalAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The city.
    city *string;
    // The country or region. It's a free-format string value, for example, 'United States'.
    countryOrRegion *string;
    // The postal code.
    postalCode *string;
    // The state.
    state *string;
    // The street.
    street *string;
}
// NewPhysicalAddress instantiates a new physicalAddress and sets the default values.
func NewPhysicalAddress()(*PhysicalAddress) {
    m := &PhysicalAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePhysicalAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePhysicalAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPhysicalAddress(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PhysicalAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCity gets the city property value. The city.
func (m *PhysicalAddress) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountryOrRegion gets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalAddress) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PhysicalAddress) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["city"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryOrRegion"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegion(val)
        }
        return nil
    }
    res["postalCode"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreet(val)
        }
        return nil
    }
    return res
}
// GetPostalCode gets the postalCode property value. The postal code.
func (m *PhysicalAddress) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// GetState gets the state property value. The state.
func (m *PhysicalAddress) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStreet gets the street property value. The street.
func (m *PhysicalAddress) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
// Serialize serializes information the current object
func (m *PhysicalAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("street", m.GetStreet())
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
func (m *PhysicalAddress) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCity sets the city property value. The city.
func (m *PhysicalAddress) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCountryOrRegion sets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalAddress) SetCountryOrRegion(value *string)() {
    if m != nil {
        m.countryOrRegion = value
    }
}
// SetPostalCode sets the postalCode property value. The postal code.
func (m *PhysicalAddress) SetPostalCode(value *string)() {
    if m != nil {
        m.postalCode = value
    }
}
// SetState sets the state property value. The state.
func (m *PhysicalAddress) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetStreet sets the street property value. The street.
func (m *PhysicalAddress) SetStreet(value *string)() {
    if m != nil {
        m.street = value
    }
}
