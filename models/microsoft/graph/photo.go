package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Photo struct {
    additionalData map[string]interface{};
    cameraMake *string;
    cameraModel *string;
    exposureDenominator *float64;
    exposureNumerator *float64;
    fNumber *float64;
    focalLength *float64;
    iso *int32;
    orientation *int32;
    takenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewPhoto()(*Photo) {
    m := &Photo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Photo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Photo) GetCameraMake()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cameraMake
    }
}
func (m *Photo) GetCameraModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cameraModel
    }
}
func (m *Photo) GetExposureDenominator()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.exposureDenominator
    }
}
func (m *Photo) GetExposureNumerator()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.exposureNumerator
    }
}
func (m *Photo) GetFNumber()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.fNumber
    }
}
func (m *Photo) GetFocalLength()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.focalLength
    }
}
func (m *Photo) GetIso()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iso
    }
}
func (m *Photo) GetOrientation()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
func (m *Photo) GetTakenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.takenDateTime
    }
}
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
func (m *Photo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Photo) SetCameraMake(value *string)() {
    m.cameraMake = value
}
func (m *Photo) SetCameraModel(value *string)() {
    m.cameraModel = value
}
func (m *Photo) SetExposureDenominator(value *float64)() {
    m.exposureDenominator = value
}
func (m *Photo) SetExposureNumerator(value *float64)() {
    m.exposureNumerator = value
}
func (m *Photo) SetFNumber(value *float64)() {
    m.fNumber = value
}
func (m *Photo) SetFocalLength(value *float64)() {
    m.focalLength = value
}
func (m *Photo) SetIso(value *int32)() {
    m.iso = value
}
func (m *Photo) SetOrientation(value *int32)() {
    m.orientation = value
}
func (m *Photo) SetTakenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.takenDateTime = value
}
