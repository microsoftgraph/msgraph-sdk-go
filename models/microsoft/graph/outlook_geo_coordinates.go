package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OutlookGeoCoordinates struct {
    // The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
    accuracy *float64;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The altitude of the location.
    altitude *float64;
    // The accuracy of the altitude.
    altitudeAccuracy *float64;
    // The latitude of the location.
    latitude *float64;
    // The longitude of the location.
    longitude *float64;
}
// Instantiates a new outlookGeoCoordinates and sets the default values.
func NewOutlookGeoCoordinates()(*OutlookGeoCoordinates) {
    m := &OutlookGeoCoordinates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the accuracy property value. The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
func (m *OutlookGeoCoordinates) GetAccuracy()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.accuracy
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutlookGeoCoordinates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the altitude property value. The altitude of the location.
func (m *OutlookGeoCoordinates) GetAltitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.altitude
    }
}
// Gets the altitudeAccuracy property value. The accuracy of the altitude.
func (m *OutlookGeoCoordinates) GetAltitudeAccuracy()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.altitudeAccuracy
    }
}
// Gets the latitude property value. The latitude of the location.
func (m *OutlookGeoCoordinates) GetLatitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.latitude
    }
}
// Gets the longitude property value. The longitude of the location.
func (m *OutlookGeoCoordinates) GetLongitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.longitude
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the accuracy property value. The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
// Parameters:
//  - value : Value to set for the accuracy property.
func (m *OutlookGeoCoordinates) SetAccuracy(value *float64)() {
    m.accuracy = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *OutlookGeoCoordinates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the altitude property value. The altitude of the location.
// Parameters:
//  - value : Value to set for the altitude property.
func (m *OutlookGeoCoordinates) SetAltitude(value *float64)() {
    m.altitude = value
}
// Sets the altitudeAccuracy property value. The accuracy of the altitude.
// Parameters:
//  - value : Value to set for the altitudeAccuracy property.
func (m *OutlookGeoCoordinates) SetAltitudeAccuracy(value *float64)() {
    m.altitudeAccuracy = value
}
// Sets the latitude property value. The latitude of the location.
// Parameters:
//  - value : Value to set for the latitude property.
func (m *OutlookGeoCoordinates) SetLatitude(value *float64)() {
    m.latitude = value
}
// Sets the longitude property value. The longitude of the location.
// Parameters:
//  - value : Value to set for the longitude property.
func (m *OutlookGeoCoordinates) SetLongitude(value *float64)() {
    m.longitude = value
}
