package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Place struct {
    Entity
    // The street address of the place.
    address *PhysicalAddress;
    // The name associated with the place.
    displayName *string;
    // Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
    geoCoordinates *OutlookGeoCoordinates;
    // The phone number of the place.
    phone *string;
}
// Instantiates a new place and sets the default values.
func NewPlace()(*Place) {
    m := &Place{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the address property value. The street address of the place.
func (m *Place) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the displayName property value. The name associated with the place.
func (m *Place) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the geoCoordinates property value. Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
func (m *Place) GetGeoCoordinates()(*OutlookGeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.geoCoordinates
    }
}
// Gets the phone property value. The phone number of the place.
func (m *Place) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
// The deserialization information for the current model
func (m *Place) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetAddress(val.(*PhysicalAddress))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["geoCoordinates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookGeoCoordinates() })
        if err != nil {
            return err
        }
        m.SetGeoCoordinates(val.(*OutlookGeoCoordinates))
        return nil
    }
    res["phone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhone(val)
        return nil
    }
    return res
}
func (m *Place) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Place) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
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
        err = writer.WriteObjectValue("geoCoordinates", m.GetGeoCoordinates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the address property value. The street address of the place.
// Parameters:
//  - value : Value to set for the address property.
func (m *Place) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
// Sets the displayName property value. The name associated with the place.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Place) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the geoCoordinates property value. Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
// Parameters:
//  - value : Value to set for the geoCoordinates property.
func (m *Place) SetGeoCoordinates(value *OutlookGeoCoordinates)() {
    m.geoCoordinates = value
}
// Sets the phone property value. The phone number of the place.
// Parameters:
//  - value : Value to set for the phone property.
func (m *Place) SetPhone(value *string)() {
    m.phone = value
}
