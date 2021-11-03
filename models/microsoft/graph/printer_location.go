package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrinterLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The altitude, in meters, that the printer is located at.
    altitudeInMeters *int32;
    // The building that the printer is located in.
    building *string;
    // The city that the printer is located in.
    city *string;
    // The country or region that the printer is located in.
    countryOrRegion *string;
    // The floor that the printer is located on. Only numerical values are supported right now.
    floor *string;
    // The description of the floor that the printer is located on.
    floorDescription *string;
    // The latitude that the printer is located at.
    latitude *float64;
    // The longitude that the printer is located at.
    longitude *float64;
    // The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
    organization []string;
    // The postal code that the printer is located in.
    postalCode *string;
    // The description of the room that the printer is located in.
    roomDescription *string;
    // The room that the printer is located in. Only numerical values are supported right now.
    roomName *string;
    // The site that the printer is located in.
    site *string;
    // The state or province that the printer is located in.
    stateOrProvince *string;
    // The street address where the printer is located.
    streetAddress *string;
    // The subdivision that the printer is located in. The elements should be in hierarchical order.
    subdivision []string;
    // 
    subunit []string;
}
// Instantiates a new printerLocation and sets the default values.
func NewPrinterLocation()(*PrinterLocation) {
    m := &PrinterLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterLocation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the altitudeInMeters property value. The altitude, in meters, that the printer is located at.
func (m *PrinterLocation) GetAltitudeInMeters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.altitudeInMeters
    }
}
// Gets the building property value. The building that the printer is located in.
func (m *PrinterLocation) GetBuilding()(*string) {
    if m == nil {
        return nil
    } else {
        return m.building
    }
}
// Gets the city property value. The city that the printer is located in.
func (m *PrinterLocation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the countryOrRegion property value. The country or region that the printer is located in.
func (m *PrinterLocation) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// Gets the floor property value. The floor that the printer is located on. Only numerical values are supported right now.
func (m *PrinterLocation) GetFloor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.floor
    }
}
// Gets the floorDescription property value. The description of the floor that the printer is located on.
func (m *PrinterLocation) GetFloorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.floorDescription
    }
}
// Gets the latitude property value. The latitude that the printer is located at.
func (m *PrinterLocation) GetLatitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.latitude
    }
}
// Gets the longitude property value. The longitude that the printer is located at.
func (m *PrinterLocation) GetLongitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.longitude
    }
}
// Gets the organization property value. The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
func (m *PrinterLocation) GetOrganization()([]string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
// Gets the postalCode property value. The postal code that the printer is located in.
func (m *PrinterLocation) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// Gets the roomDescription property value. The description of the room that the printer is located in.
func (m *PrinterLocation) GetRoomDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roomDescription
    }
}
// Gets the roomName property value. The room that the printer is located in. Only numerical values are supported right now.
func (m *PrinterLocation) GetRoomName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roomName
    }
}
// Gets the site property value. The site that the printer is located in.
func (m *PrinterLocation) GetSite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.site
    }
}
// Gets the stateOrProvince property value. The state or province that the printer is located in.
func (m *PrinterLocation) GetStateOrProvince()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateOrProvince
    }
}
// Gets the streetAddress property value. The street address where the printer is located.
func (m *PrinterLocation) GetStreetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.streetAddress
    }
}
// Gets the subdivision property value. The subdivision that the printer is located in. The elements should be in hierarchical order.
func (m *PrinterLocation) GetSubdivision()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subdivision
    }
}
// Gets the subunit property value. 
func (m *PrinterLocation) GetSubunit()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subunit
    }
}
// The deserialization information for the current model
func (m *PrinterLocation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["altitudeInMeters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAltitudeInMeters(val)
        return nil
    }
    res["building"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBuilding(val)
        return nil
    }
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCity(val)
        return nil
    }
    res["countryOrRegion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryOrRegion(val)
        return nil
    }
    res["floor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFloor(val)
        return nil
    }
    res["floorDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFloorDescription(val)
        return nil
    }
    res["latitude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetLatitude(val)
        return nil
    }
    res["longitude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetLongitude(val)
        return nil
    }
    res["organization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetOrganization(res)
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostalCode(val)
        return nil
    }
    res["roomDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoomDescription(val)
        return nil
    }
    res["roomName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoomName(val)
        return nil
    }
    res["site"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSite(val)
        return nil
    }
    res["stateOrProvince"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStateOrProvince(val)
        return nil
    }
    res["streetAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStreetAddress(val)
        return nil
    }
    res["subdivision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSubdivision(res)
        return nil
    }
    res["subunit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSubunit(res)
        return nil
    }
    return res
}
func (m *PrinterLocation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrinterLocation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("altitudeInMeters", m.GetAltitudeInMeters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("building", m.GetBuilding())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("floor", m.GetFloor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("floorDescription", m.GetFloorDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("longitude", m.GetLongitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("organization", m.GetOrganization())
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
        err := writer.WriteStringValue("roomDescription", m.GetRoomDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roomName", m.GetRoomName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stateOrProvince", m.GetStateOrProvince())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("streetAddress", m.GetStreetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("subdivision", m.GetSubdivision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("subunit", m.GetSubunit())
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
func (m *PrinterLocation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the altitudeInMeters property value. The altitude, in meters, that the printer is located at.
// Parameters:
//  - value : Value to set for the altitudeInMeters property.
func (m *PrinterLocation) SetAltitudeInMeters(value *int32)() {
    m.altitudeInMeters = value
}
// Sets the building property value. The building that the printer is located in.
// Parameters:
//  - value : Value to set for the building property.
func (m *PrinterLocation) SetBuilding(value *string)() {
    m.building = value
}
// Sets the city property value. The city that the printer is located in.
// Parameters:
//  - value : Value to set for the city property.
func (m *PrinterLocation) SetCity(value *string)() {
    m.city = value
}
// Sets the countryOrRegion property value. The country or region that the printer is located in.
// Parameters:
//  - value : Value to set for the countryOrRegion property.
func (m *PrinterLocation) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// Sets the floor property value. The floor that the printer is located on. Only numerical values are supported right now.
// Parameters:
//  - value : Value to set for the floor property.
func (m *PrinterLocation) SetFloor(value *string)() {
    m.floor = value
}
// Sets the floorDescription property value. The description of the floor that the printer is located on.
// Parameters:
//  - value : Value to set for the floorDescription property.
func (m *PrinterLocation) SetFloorDescription(value *string)() {
    m.floorDescription = value
}
// Sets the latitude property value. The latitude that the printer is located at.
// Parameters:
//  - value : Value to set for the latitude property.
func (m *PrinterLocation) SetLatitude(value *float64)() {
    m.latitude = value
}
// Sets the longitude property value. The longitude that the printer is located at.
// Parameters:
//  - value : Value to set for the longitude property.
func (m *PrinterLocation) SetLongitude(value *float64)() {
    m.longitude = value
}
// Sets the organization property value. The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
// Parameters:
//  - value : Value to set for the organization property.
func (m *PrinterLocation) SetOrganization(value []string)() {
    m.organization = value
}
// Sets the postalCode property value. The postal code that the printer is located in.
// Parameters:
//  - value : Value to set for the postalCode property.
func (m *PrinterLocation) SetPostalCode(value *string)() {
    m.postalCode = value
}
// Sets the roomDescription property value. The description of the room that the printer is located in.
// Parameters:
//  - value : Value to set for the roomDescription property.
func (m *PrinterLocation) SetRoomDescription(value *string)() {
    m.roomDescription = value
}
// Sets the roomName property value. The room that the printer is located in. Only numerical values are supported right now.
// Parameters:
//  - value : Value to set for the roomName property.
func (m *PrinterLocation) SetRoomName(value *string)() {
    m.roomName = value
}
// Sets the site property value. The site that the printer is located in.
// Parameters:
//  - value : Value to set for the site property.
func (m *PrinterLocation) SetSite(value *string)() {
    m.site = value
}
// Sets the stateOrProvince property value. The state or province that the printer is located in.
// Parameters:
//  - value : Value to set for the stateOrProvince property.
func (m *PrinterLocation) SetStateOrProvince(value *string)() {
    m.stateOrProvince = value
}
// Sets the streetAddress property value. The street address where the printer is located.
// Parameters:
//  - value : Value to set for the streetAddress property.
func (m *PrinterLocation) SetStreetAddress(value *string)() {
    m.streetAddress = value
}
// Sets the subdivision property value. The subdivision that the printer is located in. The elements should be in hierarchical order.
// Parameters:
//  - value : Value to set for the subdivision property.
func (m *PrinterLocation) SetSubdivision(value []string)() {
    m.subdivision = value
}
// Sets the subunit property value. 
// Parameters:
//  - value : Value to set for the subunit property.
func (m *PrinterLocation) SetSubunit(value []string)() {
    m.subunit = value
}
