package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GeoCoordinates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
    altitude *float64;
    // Optional. The latitude, in decimal, for the item. Read-only.
    latitude *float64;
    // Optional. The longitude, in decimal, for the item. Read-only.
    longitude *float64;
}
// Instantiates a new geoCoordinates and sets the default values.
func NewGeoCoordinates()(*GeoCoordinates) {
    m := &GeoCoordinates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GeoCoordinates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the altitude property value. Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
func (m *GeoCoordinates) GetAltitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.altitude
    }
}
// Gets the latitude property value. Optional. The latitude, in decimal, for the item. Read-only.
func (m *GeoCoordinates) GetLatitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.latitude
    }
}
// Gets the longitude property value. Optional. The longitude, in decimal, for the item. Read-only.
func (m *GeoCoordinates) GetLongitude()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.longitude
    }
}
// The deserialization information for the current model
func (m *GeoCoordinates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["altitude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAltitude(val)
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
func (m *GeoCoordinates) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GeoCoordinates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("altitude", m.GetAltitude())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *GeoCoordinates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the altitude property value. Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
// Parameters:
//  - value : Value to set for the altitude property.
func (m *GeoCoordinates) SetAltitude(value *float64)() {
    m.altitude = value
}
// Sets the latitude property value. Optional. The latitude, in decimal, for the item. Read-only.
// Parameters:
//  - value : Value to set for the latitude property.
func (m *GeoCoordinates) SetLatitude(value *float64)() {
    m.latitude = value
}
// Sets the longitude property value. Optional. The longitude, in decimal, for the item. Read-only.
// Parameters:
//  - value : Value to set for the longitude property.
func (m *GeoCoordinates) SetLongitude(value *float64)() {
    m.longitude = value
}
