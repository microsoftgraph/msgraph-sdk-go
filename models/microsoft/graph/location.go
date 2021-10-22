package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Location struct {
    additionalData map[string]interface{};
    address *PhysicalAddress;
    coordinates *OutlookGeoCoordinates;
    displayName *string;
    locationEmailAddress *string;
    locationType *LocationType;
    locationUri *string;
    uniqueId *string;
    uniqueIdType *LocationUniqueIdType;
}
func NewLocation()(*Location) {
    m := &Location{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Location) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Location) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *Location) GetCoordinates()(*OutlookGeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.coordinates
    }
}
func (m *Location) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Location) GetLocationEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locationEmailAddress
    }
}
func (m *Location) GetLocationType()(*LocationType) {
    if m == nil {
        return nil
    } else {
        return m.locationType
    }
}
func (m *Location) GetLocationUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locationUri
    }
}
func (m *Location) GetUniqueId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueId
    }
}
func (m *Location) GetUniqueIdType()(*LocationUniqueIdType) {
    if m == nil {
        return nil
    } else {
        return m.uniqueIdType
    }
}
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
func (m *Location) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Location) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
func (m *Location) SetCoordinates(value *OutlookGeoCoordinates)() {
    m.coordinates = value
}
func (m *Location) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Location) SetLocationEmailAddress(value *string)() {
    m.locationEmailAddress = value
}
func (m *Location) SetLocationType(value *LocationType)() {
    m.locationType = value
}
func (m *Location) SetLocationUri(value *string)() {
    m.locationUri = value
}
func (m *Location) SetUniqueId(value *string)() {
    m.uniqueId = value
}
func (m *Location) SetUniqueIdType(value *LocationUniqueIdType)() {
    m.uniqueIdType = value
}
