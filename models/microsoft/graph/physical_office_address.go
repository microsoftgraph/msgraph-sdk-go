package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PhysicalOfficeAddress 
type PhysicalOfficeAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The city.
    city *string;
    // The country or region. It's a free-format string value, for example, 'United States'.
    countryOrRegion *string;
    // Office location such as building and office number for an organizational contact.
    officeLocation *string;
    // The postal code.
    postalCode *string;
    // The state.
    state *string;
    // The street.
    street *string;
}
// NewPhysicalOfficeAddress instantiates a new physicalOfficeAddress and sets the default values.
func NewPhysicalOfficeAddress()(*PhysicalOfficeAddress) {
    m := &PhysicalOfficeAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PhysicalOfficeAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCity gets the city property value. The city.
func (m *PhysicalOfficeAddress) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountryOrRegion gets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalOfficeAddress) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// GetOfficeLocation gets the officeLocation property value. Office location such as building and office number for an organizational contact.
func (m *PhysicalOfficeAddress) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetPostalCode gets the postalCode property value. The postal code.
func (m *PhysicalOfficeAddress) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// GetState gets the state property value. The state.
func (m *PhysicalOfficeAddress) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStreet gets the street property value. The street.
func (m *PhysicalOfficeAddress) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PhysicalOfficeAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryOrRegion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegion(val)
        }
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *PhysicalOfficeAddress) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PhysicalOfficeAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
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
func (m *PhysicalOfficeAddress) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCity sets the city property value. The city.
func (m *PhysicalOfficeAddress) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCountryOrRegion sets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalOfficeAddress) SetCountryOrRegion(value *string)() {
    if m != nil {
        m.countryOrRegion = value
    }
}
// SetOfficeLocation sets the officeLocation property value. Office location such as building and office number for an organizational contact.
func (m *PhysicalOfficeAddress) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetPostalCode sets the postalCode property value. The postal code.
func (m *PhysicalOfficeAddress) SetPostalCode(value *string)() {
    if m != nil {
        m.postalCode = value
    }
}
// SetState sets the state property value. The state.
func (m *PhysicalOfficeAddress) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetStreet sets the street property value. The street.
func (m *PhysicalOfficeAddress) SetStreet(value *string)() {
    if m != nil {
        m.street = value
    }
}
