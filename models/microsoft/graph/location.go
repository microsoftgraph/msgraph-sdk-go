package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Location struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The street address of the location.
    address *PhysicalAddress;
    // The geographic coordinates and elevation of the location.
    coordinates *OutlookGeoCoordinates;
    // The name associated with the location.
    displayName *string;
    // Optional email address of the location.
    locationEmailAddress *string;
    // The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
    locationType *LocationType;
    // Optional URI representing the location.
    locationUri *string;
    // For internal use only.
    uniqueId *string;
    // For internal use only.
    uniqueIdType *LocationUniqueIdType;
}
// Instantiates a new location and sets the default values.
func NewLocation()(*Location) {
    m := &Location{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Location) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the address property value. The street address of the location.
func (m *Location) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the coordinates property value. The geographic coordinates and elevation of the location.
func (m *Location) GetCoordinates()(*OutlookGeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.coordinates
    }
}
// Gets the displayName property value. The name associated with the location.
func (m *Location) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the locationEmailAddress property value. Optional email address of the location.
func (m *Location) GetLocationEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locationEmailAddress
    }
}
// Gets the locationType property value. The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
func (m *Location) GetLocationType()(*LocationType) {
    if m == nil {
        return nil
    } else {
        return m.locationType
    }
}
// Gets the locationUri property value. Optional URI representing the location.
func (m *Location) GetLocationUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locationUri
    }
}
// Gets the uniqueId property value. For internal use only.
func (m *Location) GetUniqueId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueId
    }
}
// Gets the uniqueIdType property value. For internal use only.
func (m *Location) GetUniqueIdType()(*LocationUniqueIdType) {
    if m == nil {
        return nil
    } else {
        return m.uniqueIdType
    }
}
// The deserialization information for the current model
func (m *Location) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetAddress(val.(*PhysicalAddress))
        return nil
    }
    res["coordinates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookGeoCoordinates() })
        if err != nil {
            return err
        }
        m.SetCoordinates(val.(*OutlookGeoCoordinates))
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
    res["locationEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocationEmailAddress(val)
        return nil
    }
    res["locationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationType)
        if err != nil {
            return err
        }
        cast := val.(LocationType)
        m.SetLocationType(&cast)
        return nil
    }
    res["locationUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocationUri(val)
        return nil
    }
    res["uniqueId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUniqueId(val)
        return nil
    }
    res["uniqueIdType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationUniqueIdType)
        if err != nil {
            return err
        }
        cast := val.(LocationUniqueIdType)
        m.SetUniqueIdType(&cast)
        return nil
    }
    return res
}
func (m *Location) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Location) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("coordinates", m.GetCoordinates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locationEmailAddress", m.GetLocationEmailAddress())
        if err != nil {
            return err
        }
    }
    if m.GetLocationType() != nil {
        cast := m.GetLocationType().String()
        err := writer.WriteStringValue("locationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locationUri", m.GetLocationUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uniqueId", m.GetUniqueId())
        if err != nil {
            return err
        }
    }
    if m.GetUniqueIdType() != nil {
        cast := m.GetUniqueIdType().String()
        err := writer.WriteStringValue("uniqueIdType", &cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Location) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the address property value. The street address of the location.
// Parameters:
//  - value : Value to set for the address property.
func (m *Location) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
// Sets the coordinates property value. The geographic coordinates and elevation of the location.
// Parameters:
//  - value : Value to set for the coordinates property.
func (m *Location) SetCoordinates(value *OutlookGeoCoordinates)() {
    m.coordinates = value
}
// Sets the displayName property value. The name associated with the location.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Location) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the locationEmailAddress property value. Optional email address of the location.
// Parameters:
//  - value : Value to set for the locationEmailAddress property.
func (m *Location) SetLocationEmailAddress(value *string)() {
    m.locationEmailAddress = value
}
// Sets the locationType property value. The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
// Parameters:
//  - value : Value to set for the locationType property.
func (m *Location) SetLocationType(value *LocationType)() {
    m.locationType = value
}
// Sets the locationUri property value. Optional URI representing the location.
// Parameters:
//  - value : Value to set for the locationUri property.
func (m *Location) SetLocationUri(value *string)() {
    m.locationUri = value
}
// Sets the uniqueId property value. For internal use only.
// Parameters:
//  - value : Value to set for the uniqueId property.
func (m *Location) SetUniqueId(value *string)() {
    m.uniqueId = value
}
// Sets the uniqueIdType property value. For internal use only.
// Parameters:
//  - value : Value to set for the uniqueIdType property.
func (m *Location) SetUniqueIdType(value *LocationUniqueIdType)() {
    m.uniqueIdType = value
}
