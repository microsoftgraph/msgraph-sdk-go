package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterLocation 
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
// NewPrinterLocation instantiates a new printerLocation and sets the default values.
func NewPrinterLocation()(*PrinterLocation) {
    m := &PrinterLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterLocation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAltitudeInMeters gets the altitudeInMeters property value. The altitude, in meters, that the printer is located at.
func (m *PrinterLocation) GetAltitudeInMeters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.altitudeInMeters
    }
}
// GetBuilding gets the building property value. The building that the printer is located in.
func (m *PrinterLocation) GetBuilding()(*string) {
    if m == nil {
        return nil
    } else {
        return m.building
    }
}
// GetCity gets the city property value. The city that the printer is located in.
func (m *PrinterLocation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountryOrRegion gets the countryOrRegion property value. The country or region that the printer is located in.
func (m *PrinterLocation) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// GetFloor gets the floor property value. The floor that the printer is located on. Only numerical values are supported right now.
func (m *PrinterLocation) GetFloor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.floor
    }
}
// GetFloorDescription gets the floorDescription property value. The description of the floor that the printer is located on.
func (m *PrinterLocation) GetFloorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.floorDescription
    }
}
// GetLatitude gets the latitude property value. The latitude that the printer is located at.
func (m *PrinterLocation) GetLatitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.latitude
    }
}
// GetLongitude gets the longitude property value. The longitude that the printer is located at.
func (m *PrinterLocation) GetLongitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.longitude
    }
}
// GetOrganization gets the organization property value. The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
func (m *PrinterLocation) GetOrganization()([]string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
// GetPostalCode gets the postalCode property value. The postal code that the printer is located in.
func (m *PrinterLocation) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// GetRoomDescription gets the roomDescription property value. The description of the room that the printer is located in.
func (m *PrinterLocation) GetRoomDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roomDescription
    }
}
// GetRoomName gets the roomName property value. The room that the printer is located in. Only numerical values are supported right now.
func (m *PrinterLocation) GetRoomName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roomName
    }
}
// GetSite gets the site property value. The site that the printer is located in.
func (m *PrinterLocation) GetSite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.site
    }
}
// GetStateOrProvince gets the stateOrProvince property value. The state or province that the printer is located in.
func (m *PrinterLocation) GetStateOrProvince()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateOrProvince
    }
}
// GetStreetAddress gets the streetAddress property value. The street address where the printer is located.
func (m *PrinterLocation) GetStreetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.streetAddress
    }
}
// GetSubdivision gets the subdivision property value. The subdivision that the printer is located in. The elements should be in hierarchical order.
func (m *PrinterLocation) GetSubdivision()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subdivision
    }
}
// GetSubunit gets the subunit property value. 
func (m *PrinterLocation) GetSubunit()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subunit
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterLocation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["altitudeInMeters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitudeInMeters(val)
        }
        return nil
    }
    res["building"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuilding(val)
        }
        return nil
    }
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
    res["floor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloor(val)
        }
        return nil
    }
    res["floorDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorDescription(val)
        }
        return nil
    }
    res["latitude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val)
        }
        return nil
    }
    res["longitude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val)
        }
        return nil
    }
    res["organization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrganization(res)
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
    res["roomDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomDescription(val)
        }
        return nil
    }
    res["roomName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomName(val)
        }
        return nil
    }
    res["site"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val)
        }
        return nil
    }
    res["stateOrProvince"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateOrProvince(val)
        }
        return nil
    }
    res["streetAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreetAddress(val)
        }
        return nil
    }
    res["subdivision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSubdivision(res)
        }
        return nil
    }
    res["subunit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSubunit(res)
        }
        return nil
    }
    return res
}
func (m *PrinterLocation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterLocation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAltitudeInMeters sets the altitudeInMeters property value. The altitude, in meters, that the printer is located at.
func (m *PrinterLocation) SetAltitudeInMeters(value *int32)() {
    if m != nil {
        m.altitudeInMeters = value
    }
}
// SetBuilding sets the building property value. The building that the printer is located in.
func (m *PrinterLocation) SetBuilding(value *string)() {
    if m != nil {
        m.building = value
    }
}
// SetCity sets the city property value. The city that the printer is located in.
func (m *PrinterLocation) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCountryOrRegion sets the countryOrRegion property value. The country or region that the printer is located in.
func (m *PrinterLocation) SetCountryOrRegion(value *string)() {
    if m != nil {
        m.countryOrRegion = value
    }
}
// SetFloor sets the floor property value. The floor that the printer is located on. Only numerical values are supported right now.
func (m *PrinterLocation) SetFloor(value *string)() {
    if m != nil {
        m.floor = value
    }
}
// SetFloorDescription sets the floorDescription property value. The description of the floor that the printer is located on.
func (m *PrinterLocation) SetFloorDescription(value *string)() {
    if m != nil {
        m.floorDescription = value
    }
}
// SetLatitude sets the latitude property value. The latitude that the printer is located at.
func (m *PrinterLocation) SetLatitude(value *float64)() {
    if m != nil {
        m.latitude = value
    }
}
// SetLongitude sets the longitude property value. The longitude that the printer is located at.
func (m *PrinterLocation) SetLongitude(value *float64)() {
    if m != nil {
        m.longitude = value
    }
}
// SetOrganization sets the organization property value. The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
func (m *PrinterLocation) SetOrganization(value []string)() {
    if m != nil {
        m.organization = value
    }
}
// SetPostalCode sets the postalCode property value. The postal code that the printer is located in.
func (m *PrinterLocation) SetPostalCode(value *string)() {
    if m != nil {
        m.postalCode = value
    }
}
// SetRoomDescription sets the roomDescription property value. The description of the room that the printer is located in.
func (m *PrinterLocation) SetRoomDescription(value *string)() {
    if m != nil {
        m.roomDescription = value
    }
}
// SetRoomName sets the roomName property value. The room that the printer is located in. Only numerical values are supported right now.
func (m *PrinterLocation) SetRoomName(value *string)() {
    if m != nil {
        m.roomName = value
    }
}
// SetSite sets the site property value. The site that the printer is located in.
func (m *PrinterLocation) SetSite(value *string)() {
    if m != nil {
        m.site = value
    }
}
// SetStateOrProvince sets the stateOrProvince property value. The state or province that the printer is located in.
func (m *PrinterLocation) SetStateOrProvince(value *string)() {
    if m != nil {
        m.stateOrProvince = value
    }
}
// SetStreetAddress sets the streetAddress property value. The street address where the printer is located.
func (m *PrinterLocation) SetStreetAddress(value *string)() {
    if m != nil {
        m.streetAddress = value
    }
}
// SetSubdivision sets the subdivision property value. The subdivision that the printer is located in. The elements should be in hierarchical order.
func (m *PrinterLocation) SetSubdivision(value []string)() {
    if m != nil {
        m.subdivision = value
    }
}
// SetSubunit sets the subunit property value. 
func (m *PrinterLocation) SetSubunit(value []string)() {
    if m != nil {
        m.subunit = value
    }
}
