package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// place 
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
// NewPlace instantiates a new place and sets the default values.
func NewPlace()(*Place) {
    m := &Place{
        Entity: *NewEntity(),
    }
    return m
}
// GetAddress gets the address property value. The street address of the place.
func (m *Place) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetDisplayName gets the displayName property value. The name associated with the place.
func (m *Place) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGeoCoordinates gets the geoCoordinates property value. Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
func (m *Place) GetGeoCoordinates()(*OutlookGeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.geoCoordinates
    }
}
// GetPhone gets the phone property value. The phone number of the place.
func (m *Place) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Place) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(*PhysicalAddress))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["geoCoordinates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookGeoCoordinates() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeoCoordinates(val.(*OutlookGeoCoordinates))
        }
        return nil
    }
    res["phone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    return res
}
func (m *Place) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAddress sets the address property value. The street address of the place.
func (m *Place) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
// SetDisplayName sets the displayName property value. The name associated with the place.
func (m *Place) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeoCoordinates sets the geoCoordinates property value. Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
func (m *Place) SetGeoCoordinates(value *OutlookGeoCoordinates)() {
    m.geoCoordinates = value
}
// SetPhone sets the phone property value. The phone number of the place.
func (m *Place) SetPhone(value *string)() {
    m.phone = value
}
