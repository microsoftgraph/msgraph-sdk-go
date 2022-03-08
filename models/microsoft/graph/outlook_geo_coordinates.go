package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookGeoCoordinates provides operations to manage the drive singleton.
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
// NewOutlookGeoCoordinates instantiates a new outlookGeoCoordinates and sets the default values.
func NewOutlookGeoCoordinates()(*OutlookGeoCoordinates) {
    m := &OutlookGeoCoordinates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOutlookGeoCoordinatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookGeoCoordinatesFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOutlookGeoCoordinates(), nil
}
// GetAccuracy gets the accuracy property value. The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
func (m *OutlookGeoCoordinates) GetAccuracy()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.accuracy
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutlookGeoCoordinates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAltitude gets the altitude property value. The altitude of the location.
func (m *OutlookGeoCoordinates) GetAltitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.altitude
    }
}
// GetAltitudeAccuracy gets the altitudeAccuracy property value. The accuracy of the altitude.
func (m *OutlookGeoCoordinates) GetAltitudeAccuracy()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.altitudeAccuracy
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookGeoCoordinates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accuracy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccuracy(val)
        }
        return nil
    }
    res["altitude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitude(val)
        }
        return nil
    }
    res["altitudeAccuracy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitudeAccuracy(val)
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
    return res
}
// GetLatitude gets the latitude property value. The latitude of the location.
func (m *OutlookGeoCoordinates) GetLatitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.latitude
    }
}
// GetLongitude gets the longitude property value. The longitude of the location.
func (m *OutlookGeoCoordinates) GetLongitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.longitude
    }
}
func (m *OutlookGeoCoordinates) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAccuracy sets the accuracy property value. The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
func (m *OutlookGeoCoordinates) SetAccuracy(value *float64)() {
    if m != nil {
        m.accuracy = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutlookGeoCoordinates) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAltitude sets the altitude property value. The altitude of the location.
func (m *OutlookGeoCoordinates) SetAltitude(value *float64)() {
    if m != nil {
        m.altitude = value
    }
}
// SetAltitudeAccuracy sets the altitudeAccuracy property value. The accuracy of the altitude.
func (m *OutlookGeoCoordinates) SetAltitudeAccuracy(value *float64)() {
    if m != nil {
        m.altitudeAccuracy = value
    }
}
// SetLatitude sets the latitude property value. The latitude of the location.
func (m *OutlookGeoCoordinates) SetLatitude(value *float64)() {
    if m != nil {
        m.latitude = value
    }
}
// SetLongitude sets the longitude property value. The longitude of the location.
func (m *OutlookGeoCoordinates) SetLongitude(value *float64)() {
    if m != nil {
        m.longitude = value
    }
}
