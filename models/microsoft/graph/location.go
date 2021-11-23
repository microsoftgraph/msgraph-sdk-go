package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// location 
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
// NewLocation instantiates a new location and sets the default values.
func NewLocation()(*Location) {
    m := &Location{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Location) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddress gets the address property value. The street address of the location.
func (m *Location) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetCoordinates gets the coordinates property value. The geographic coordinates and elevation of the location.
func (m *Location) GetCoordinates()(*OutlookGeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.coordinates
    }
}
// GetDisplayName gets the displayName property value. The name associated with the location.
func (m *Location) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLocationEmailAddress gets the locationEmailAddress property value. Optional email address of the location.
func (m *Location) GetLocationEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locationEmailAddress
    }
}
// GetLocationType gets the locationType property value. The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
func (m *Location) GetLocationType()(*LocationType) {
    if m == nil {
        return nil
    } else {
        return m.locationType
    }
}
// GetLocationUri gets the locationUri property value. Optional URI representing the location.
func (m *Location) GetLocationUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locationUri
    }
}
// GetUniqueId gets the uniqueId property value. For internal use only.
func (m *Location) GetUniqueId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueId
    }
}
// GetUniqueIdType gets the uniqueIdType property value. For internal use only.
func (m *Location) GetUniqueIdType()(*LocationUniqueIdType) {
    if m == nil {
        return nil
    } else {
        return m.uniqueIdType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Location) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["coordinates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookGeoCoordinates() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoordinates(val.(*OutlookGeoCoordinates))
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
    res["locationEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationEmailAddress(val)
        }
        return nil
    }
    res["locationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LocationType)
            m.SetLocationType(&cast)
        }
        return nil
    }
    res["locationUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationUri(val)
        }
        return nil
    }
    res["uniqueId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueId(val)
        }
        return nil
    }
    res["uniqueIdType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationUniqueIdType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LocationUniqueIdType)
            m.SetUniqueIdType(&cast)
        }
        return nil
    }
    return res
}
func (m *Location) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Location) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddress sets the address property value. The street address of the location.
func (m *Location) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
// SetCoordinates sets the coordinates property value. The geographic coordinates and elevation of the location.
func (m *Location) SetCoordinates(value *OutlookGeoCoordinates)() {
    m.coordinates = value
}
// SetDisplayName sets the displayName property value. The name associated with the location.
func (m *Location) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLocationEmailAddress sets the locationEmailAddress property value. Optional email address of the location.
func (m *Location) SetLocationEmailAddress(value *string)() {
    m.locationEmailAddress = value
}
// SetLocationType sets the locationType property value. The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
func (m *Location) SetLocationType(value *LocationType)() {
    m.locationType = value
}
// SetLocationUri sets the locationUri property value. Optional URI representing the location.
func (m *Location) SetLocationUri(value *string)() {
    m.locationUri = value
}
// SetUniqueId sets the uniqueId property value. For internal use only.
func (m *Location) SetUniqueId(value *string)() {
    m.uniqueId = value
}
// SetUniqueIdType sets the uniqueIdType property value. For internal use only.
func (m *Location) SetUniqueIdType(value *LocationUniqueIdType)() {
    m.uniqueIdType = value
}
