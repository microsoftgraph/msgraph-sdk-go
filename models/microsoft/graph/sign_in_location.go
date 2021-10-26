package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SignInLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
    city *string;
    // Provides the country code info (2 letter code) where the sign-in originated.  This is calculated using latitude/longitude information from the sign-in activity.
    countryOrRegion *string;
    // Provides the latitude, longitude and altitude where the sign-in originated.
    geoCoordinates *GeoCoordinates;
    // Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
    state *string;
}
// Instantiates a new signInLocation and sets the default values.
func NewSignInLocation()(*SignInLocation) {
    m := &SignInLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SignInLocation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the city property value. Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the countryOrRegion property value. Provides the country code info (2 letter code) where the sign-in originated.  This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// Gets the geoCoordinates property value. Provides the latitude, longitude and altitude where the sign-in originated.
func (m *SignInLocation) GetGeoCoordinates()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.geoCoordinates
    }
}
// Gets the state property value. Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *SignInLocation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["geoCoordinates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeoCoordinates() })
        if err != nil {
            return err
        }
        m.SetGeoCoordinates(val.(*GeoCoordinates))
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    return res
}
func (m *SignInLocation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SignInLocation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("geoCoordinates", m.GetGeoCoordinates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *SignInLocation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the city property value. Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
// Parameters:
//  - value : Value to set for the city property.
func (m *SignInLocation) SetCity(value *string)() {
    m.city = value
}
// Sets the countryOrRegion property value. Provides the country code info (2 letter code) where the sign-in originated.  This is calculated using latitude/longitude information from the sign-in activity.
// Parameters:
//  - value : Value to set for the countryOrRegion property.
func (m *SignInLocation) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// Sets the geoCoordinates property value. Provides the latitude, longitude and altitude where the sign-in originated.
// Parameters:
//  - value : Value to set for the geoCoordinates property.
func (m *SignInLocation) SetGeoCoordinates(value *GeoCoordinates)() {
    m.geoCoordinates = value
}
// Sets the state property value. Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
// Parameters:
//  - value : Value to set for the state property.
func (m *SignInLocation) SetState(value *string)() {
    m.state = value
}
