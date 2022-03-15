package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Photo provides operations to manage the collection of drive entities.
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
// NewPhoto instantiates a new photo and sets the default values.
func NewPhoto()(*Photo) {
    m := &Photo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePhotoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePhotoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPhoto(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Photo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCameraMake gets the cameraMake property value. Camera manufacturer. Read-only.
func (m *Photo) GetCameraMake()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cameraMake
    }
}
// GetCameraModel gets the cameraModel property value. Camera model. Read-only.
func (m *Photo) GetCameraModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cameraModel
    }
}
// GetExposureDenominator gets the exposureDenominator property value. The denominator for the exposure time fraction from the camera. Read-only.
func (m *Photo) GetExposureDenominator()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.exposureDenominator
    }
}
// GetExposureNumerator gets the exposureNumerator property value. The numerator for the exposure time fraction from the camera. Read-only.
func (m *Photo) GetExposureNumerator()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.exposureNumerator
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Photo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cameraMake"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraMake(val)
        }
        return nil
    }
    res["cameraModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraModel(val)
        }
        return nil
    }
    res["exposureDenominator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExposureDenominator(val)
        }
        return nil
    }
    res["exposureNumerator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExposureNumerator(val)
        }
        return nil
    }
    res["fNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFNumber(val)
        }
        return nil
    }
    res["focalLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFocalLength(val)
        }
        return nil
    }
    res["iso"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIso(val)
        }
        return nil
    }
    res["orientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrientation(val)
        }
        return nil
    }
    res["takenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTakenDateTime(val)
        }
        return nil
    }
    return res
}
// GetFNumber gets the fNumber property value. The F-stop value from the camera. Read-only.
func (m *Photo) GetFNumber()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.fNumber
    }
}
// GetFocalLength gets the focalLength property value. The focal length from the camera. Read-only.
func (m *Photo) GetFocalLength()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.focalLength
    }
}
// GetIso gets the iso property value. The ISO value from the camera. Read-only.
func (m *Photo) GetIso()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iso
    }
}
// GetOrientation gets the orientation property value. The orientation value from the camera. Writable on OneDrive Personal.
func (m *Photo) GetOrientation()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
// GetTakenDateTime gets the takenDateTime property value. Represents the date and time the photo was taken. Read-only.
func (m *Photo) GetTakenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.takenDateTime
    }
}
func (m *Photo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Photo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCameraMake sets the cameraMake property value. Camera manufacturer. Read-only.
func (m *Photo) SetCameraMake(value *string)() {
    if m != nil {
        m.cameraMake = value
    }
}
// SetCameraModel sets the cameraModel property value. Camera model. Read-only.
func (m *Photo) SetCameraModel(value *string)() {
    if m != nil {
        m.cameraModel = value
    }
}
// SetExposureDenominator sets the exposureDenominator property value. The denominator for the exposure time fraction from the camera. Read-only.
func (m *Photo) SetExposureDenominator(value *float64)() {
    if m != nil {
        m.exposureDenominator = value
    }
}
// SetExposureNumerator sets the exposureNumerator property value. The numerator for the exposure time fraction from the camera. Read-only.
func (m *Photo) SetExposureNumerator(value *float64)() {
    if m != nil {
        m.exposureNumerator = value
    }
}
// SetFNumber sets the fNumber property value. The F-stop value from the camera. Read-only.
func (m *Photo) SetFNumber(value *float64)() {
    if m != nil {
        m.fNumber = value
    }
}
// SetFocalLength sets the focalLength property value. The focal length from the camera. Read-only.
func (m *Photo) SetFocalLength(value *float64)() {
    if m != nil {
        m.focalLength = value
    }
}
// SetIso sets the iso property value. The ISO value from the camera. Read-only.
func (m *Photo) SetIso(value *int32)() {
    if m != nil {
        m.iso = value
    }
}
// SetOrientation sets the orientation property value. The orientation value from the camera. Writable on OneDrive Personal.
func (m *Photo) SetOrientation(value *int32)() {
    if m != nil {
        m.orientation = value
    }
}
// SetTakenDateTime sets the takenDateTime property value. Represents the date and time the photo was taken. Read-only.
func (m *Photo) SetTakenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.takenDateTime = value
    }
}
