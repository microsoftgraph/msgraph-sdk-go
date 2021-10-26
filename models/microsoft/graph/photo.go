package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Photo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Camera manufacturer. Read-only.
    cameraMake *string;
    // Camera model. Read-only.
    cameraModel *string;
    // The denominator for the exposure time fraction from the camera. Read-only.
    exposureDenominator *float64;
    // The numerator for the exposure time fraction from the camera. Read-only.
    exposureNumerator *float64;
    // The F-stop value from the camera. Read-only.
    fNumber *float64;
    // The focal length from the camera. Read-only.
    focalLength *float64;
    // The ISO value from the camera. Read-only.
    iso *int32;
    // The orientation value from the camera. Writable on OneDrive Personal.
    orientation *int32;
    // Represents the date and time the photo was taken. Read-only.
    takenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new photo and sets the default values.
func NewPhoto()(*Photo) {
    m := &Photo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Photo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the cameraMake property value. Camera manufacturer. Read-only.
func (m *Photo) GetCameraMake()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cameraMake
    }
}
// Gets the cameraModel property value. Camera model. Read-only.
func (m *Photo) GetCameraModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cameraModel
    }
}
// Gets the exposureDenominator property value. The denominator for the exposure time fraction from the camera. Read-only.
func (m *Photo) GetExposureDenominator()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.exposureDenominator
    }
}
// Gets the exposureNumerator property value. The numerator for the exposure time fraction from the camera. Read-only.
func (m *Photo) GetExposureNumerator()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.exposureNumerator
    }
}
// Gets the fNumber property value. The F-stop value from the camera. Read-only.
func (m *Photo) GetFNumber()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.fNumber
    }
}
// Gets the focalLength property value. The focal length from the camera. Read-only.
func (m *Photo) GetFocalLength()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.focalLength
    }
}
// Gets the iso property value. The ISO value from the camera. Read-only.
func (m *Photo) GetIso()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iso
    }
}
// Gets the orientation property value. The orientation value from the camera. Writable on OneDrive Personal.
func (m *Photo) GetOrientation()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
// Gets the takenDateTime property value. Represents the date and time the photo was taken. Read-only.
func (m *Photo) GetTakenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.takenDateTime
    }
}
// The deserialization information for the current model
func (m *Photo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cameraMake"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCameraMake(val)
        return nil
    }
    res["cameraModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCameraModel(val)
        return nil
    }
    res["exposureDenominator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetExposureDenominator(val)
        return nil
    }
    res["exposureNumerator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetExposureNumerator(val)
        return nil
    }
    res["fNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetFNumber(val)
        return nil
    }
    res["focalLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetFocalLength(val)
        return nil
    }
    res["iso"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIso(val)
        return nil
    }
    res["orientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOrientation(val)
        return nil
    }
    res["takenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetTakenDateTime(val)
        return nil
    }
    return res
}
func (m *Photo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Photo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cameraMake", m.GetCameraMake())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cameraModel", m.GetCameraModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("exposureDenominator", m.GetExposureDenominator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("exposureNumerator", m.GetExposureNumerator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("fNumber", m.GetFNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("focalLength", m.GetFocalLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("iso", m.GetIso())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("orientation", m.GetOrientation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("takenDateTime", m.GetTakenDateTime())
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
func (m *Photo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the cameraMake property value. Camera manufacturer. Read-only.
// Parameters:
//  - value : Value to set for the cameraMake property.
func (m *Photo) SetCameraMake(value *string)() {
    m.cameraMake = value
}
// Sets the cameraModel property value. Camera model. Read-only.
// Parameters:
//  - value : Value to set for the cameraModel property.
func (m *Photo) SetCameraModel(value *string)() {
    m.cameraModel = value
}
// Sets the exposureDenominator property value. The denominator for the exposure time fraction from the camera. Read-only.
// Parameters:
//  - value : Value to set for the exposureDenominator property.
func (m *Photo) SetExposureDenominator(value *float64)() {
    m.exposureDenominator = value
}
// Sets the exposureNumerator property value. The numerator for the exposure time fraction from the camera. Read-only.
// Parameters:
//  - value : Value to set for the exposureNumerator property.
func (m *Photo) SetExposureNumerator(value *float64)() {
    m.exposureNumerator = value
}
// Sets the fNumber property value. The F-stop value from the camera. Read-only.
// Parameters:
//  - value : Value to set for the fNumber property.
func (m *Photo) SetFNumber(value *float64)() {
    m.fNumber = value
}
// Sets the focalLength property value. The focal length from the camera. Read-only.
// Parameters:
//  - value : Value to set for the focalLength property.
func (m *Photo) SetFocalLength(value *float64)() {
    m.focalLength = value
}
// Sets the iso property value. The ISO value from the camera. Read-only.
// Parameters:
//  - value : Value to set for the iso property.
func (m *Photo) SetIso(value *int32)() {
    m.iso = value
}
// Sets the orientation property value. The orientation value from the camera. Writable on OneDrive Personal.
// Parameters:
//  - value : Value to set for the orientation property.
func (m *Photo) SetOrientation(value *int32)() {
    m.orientation = value
}
// Sets the takenDateTime property value. Represents the date and time the photo was taken. Read-only.
// Parameters:
//  - value : Value to set for the takenDateTime property.
func (m *Photo) SetTakenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.takenDateTime = value
}
