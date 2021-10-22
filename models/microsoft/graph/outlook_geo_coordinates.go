package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OutlookGeoCoordinates struct {
    accuracy *float64;
    additionalData map[string]interface{};
    altitude *float64;
    altitudeAccuracy *float64;
    latitude *float64;
    longitude *float64;
}
func NewOutlookGeoCoordinates()(*OutlookGeoCoordinates) {
    m := &OutlookGeoCoordinates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OutlookGeoCoordinates) GetAccuracy()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.accuracy
    }
}
func (m *OutlookGeoCoordinates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OutlookGeoCoordinates) GetAltitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.altitude
    }
}
func (m *OutlookGeoCoordinates) GetAltitudeAccuracy()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.altitudeAccuracy
    }
}
func (m *OutlookGeoCoordinates) GetLatitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.latitude
    }
}
func (m *OutlookGeoCoordinates) GetLongitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.longitude
    }
}
func (m *OutlookGeoCoordinates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accuracy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAccuracy(val)
        return nil
    }
    res["altitude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAltitude(val)
        return nil
    }
    res["altitudeAccuracy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAltitudeAccuracy(val)
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
    return res
}
func (m *OutlookGeoCoordinates) IsNil()(bool) {
    return m == nil
}
func (m *OutlookGeoCoordinates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("accuracy", m.GetAccuracy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("altitude", m.GetAltitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("altitudeAccuracy", m.GetAltitudeAccuracy())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OutlookGeoCoordinates) SetAccuracy(value *float64)() {
    m.accuracy = value
}
func (m *OutlookGeoCoordinates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OutlookGeoCoordinates) SetAltitude(value *float64)() {
    m.altitude = value
}
func (m *OutlookGeoCoordinates) SetAltitudeAccuracy(value *float64)() {
    m.altitudeAccuracy = value
}
func (m *OutlookGeoCoordinates) SetLatitude(value *float64)() {
    m.latitude = value
}
func (m *OutlookGeoCoordinates) SetLongitude(value *float64)() {
    m.longitude = value
}
