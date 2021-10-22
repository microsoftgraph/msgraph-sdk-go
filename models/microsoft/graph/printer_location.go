package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrinterLocation struct {
    additionalData map[string]interface{};
    altitudeInMeters *int32;
    building *string;
    city *string;
    countryOrRegion *string;
    floor *string;
    floorDescription *string;
    latitude *float64;
    longitude *float64;
    organization []string;
    postalCode *string;
    roomDescription *string;
    roomName *string;
    site *string;
    stateOrProvince *string;
    streetAddress *string;
    subdivision []string;
    subunit []string;
}
func NewPrinterLocation()(*PrinterLocation) {
    m := &PrinterLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrinterLocation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrinterLocation) GetAltitudeInMeters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.altitudeInMeters
    }
}
func (m *PrinterLocation) GetBuilding()(*string) {
    if m == nil {
        return nil
    } else {
        return m.building
    }
}
func (m *PrinterLocation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
func (m *PrinterLocation) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
func (m *PrinterLocation) GetFloor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.floor
    }
}
func (m *PrinterLocation) GetFloorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.floorDescription
    }
}
func (m *PrinterLocation) GetLatitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.latitude
    }
}
func (m *PrinterLocation) GetLongitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.longitude
    }
}
func (m *PrinterLocation) GetOrganization()([]string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
func (m *PrinterLocation) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
func (m *PrinterLocation) GetRoomDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roomDescription
    }
}
func (m *PrinterLocation) GetRoomName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roomName
    }
}
func (m *PrinterLocation) GetSite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.site
    }
}
func (m *PrinterLocation) GetStateOrProvince()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateOrProvince
    }
}
func (m *PrinterLocation) GetStreetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.streetAddress
    }
}
func (m *PrinterLocation) GetSubdivision()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subdivision
    }
}
func (m *PrinterLocation) GetSubunit()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subunit
    }
}
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
        }
        m.SetSubunit(res)
        return nil
    }
    return res
}
func (m *PrinterLocation) IsNil()(bool) {
    return m == nil
}
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
func (m *PrinterLocation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrinterLocation) SetAltitudeInMeters(value *int32)() {
    m.altitudeInMeters = value
}
func (m *PrinterLocation) SetBuilding(value *string)() {
    m.building = value
}
func (m *PrinterLocation) SetCity(value *string)() {
    m.city = value
}
func (m *PrinterLocation) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
func (m *PrinterLocation) SetFloor(value *string)() {
    m.floor = value
}
func (m *PrinterLocation) SetFloorDescription(value *string)() {
    m.floorDescription = value
}
func (m *PrinterLocation) SetLatitude(value *float64)() {
    m.latitude = value
}
func (m *PrinterLocation) SetLongitude(value *float64)() {
    m.longitude = value
}
func (m *PrinterLocation) SetOrganization(value []string)() {
    m.organization = value
}
func (m *PrinterLocation) SetPostalCode(value *string)() {
    m.postalCode = value
}
func (m *PrinterLocation) SetRoomDescription(value *string)() {
    m.roomDescription = value
}
func (m *PrinterLocation) SetRoomName(value *string)() {
    m.roomName = value
}
func (m *PrinterLocation) SetSite(value *string)() {
    m.site = value
}
func (m *PrinterLocation) SetStateOrProvince(value *string)() {
    m.stateOrProvince = value
}
func (m *PrinterLocation) SetStreetAddress(value *string)() {
    m.streetAddress = value
}
func (m *PrinterLocation) SetSubdivision(value []string)() {
    m.subdivision = value
}
func (m *PrinterLocation) SetSubunit(value []string)() {
    m.subunit = value
}
