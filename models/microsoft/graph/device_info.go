package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceInfo struct {
    additionalData map[string]interface{};
    captureDeviceDriver *string;
    captureDeviceName *string;
    captureNotFunctioningEventRatio *float32;
    cpuInsufficentEventRatio *float32;
    deviceClippingEventRatio *float32;
    deviceGlitchEventRatio *float32;
    howlingEventCount *int32;
    initialSignalLevelRootMeanSquare *float32;
    lowSpeechLevelEventRatio *float32;
    lowSpeechToNoiseEventRatio *float32;
    micGlitchRate *float32;
    receivedNoiseLevel *int32;
    receivedSignalLevel *int32;
    renderDeviceDriver *string;
    renderDeviceName *string;
    renderMuteEventRatio *float32;
    renderNotFunctioningEventRatio *float32;
    renderZeroVolumeEventRatio *float32;
    sentNoiseLevel *int32;
    sentSignalLevel *int32;
    speakerGlitchRate *float32;
}
func NewDeviceInfo()(*DeviceInfo) {
    m := &DeviceInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceInfo) GetCaptureDeviceDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.captureDeviceDriver
    }
}
func (m *DeviceInfo) GetCaptureDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.captureDeviceName
    }
}
func (m *DeviceInfo) GetCaptureNotFunctioningEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.captureNotFunctioningEventRatio
    }
}
func (m *DeviceInfo) GetCpuInsufficentEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.cpuInsufficentEventRatio
    }
}
func (m *DeviceInfo) GetDeviceClippingEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.deviceClippingEventRatio
    }
}
func (m *DeviceInfo) GetDeviceGlitchEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.deviceGlitchEventRatio
    }
}
func (m *DeviceInfo) GetHowlingEventCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.howlingEventCount
    }
}
func (m *DeviceInfo) GetInitialSignalLevelRootMeanSquare()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.initialSignalLevelRootMeanSquare
    }
}
func (m *DeviceInfo) GetLowSpeechLevelEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.lowSpeechLevelEventRatio
    }
}
func (m *DeviceInfo) GetLowSpeechToNoiseEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.lowSpeechToNoiseEventRatio
    }
}
func (m *DeviceInfo) GetMicGlitchRate()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.micGlitchRate
    }
}
func (m *DeviceInfo) GetReceivedNoiseLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.receivedNoiseLevel
    }
}
func (m *DeviceInfo) GetReceivedSignalLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.receivedSignalLevel
    }
}
func (m *DeviceInfo) GetRenderDeviceDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.renderDeviceDriver
    }
}
func (m *DeviceInfo) GetRenderDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.renderDeviceName
    }
}
func (m *DeviceInfo) GetRenderMuteEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.renderMuteEventRatio
    }
}
func (m *DeviceInfo) GetRenderNotFunctioningEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.renderNotFunctioningEventRatio
    }
}
func (m *DeviceInfo) GetRenderZeroVolumeEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.renderZeroVolumeEventRatio
    }
}
func (m *DeviceInfo) GetSentNoiseLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sentNoiseLevel
    }
}
func (m *DeviceInfo) GetSentSignalLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sentSignalLevel
    }
}
func (m *DeviceInfo) GetSpeakerGlitchRate()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.speakerGlitchRate
    }
}
func (m *DeviceInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["captureDeviceDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCaptureDeviceDriver(val)
        return nil
    }
    res["captureDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCaptureDeviceName(val)
        return nil
    }
    res["captureNotFunctioningEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetCaptureNotFunctioningEventRatio(val)
        return nil
    }
    res["cpuInsufficentEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetCpuInsufficentEventRatio(val)
        return nil
    }
    res["deviceClippingEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetDeviceClippingEventRatio(val)
        return nil
    }
    res["deviceGlitchEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetDeviceGlitchEventRatio(val)
        return nil
    }
    res["howlingEventCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetHowlingEventCount(val)
        return nil
    }
    res["initialSignalLevelRootMeanSquare"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetInitialSignalLevelRootMeanSquare(val)
        return nil
    }
    res["lowSpeechLevelEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetLowSpeechLevelEventRatio(val)
        return nil
    }
    res["lowSpeechToNoiseEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetLowSpeechToNoiseEventRatio(val)
        return nil
    }
    res["micGlitchRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetMicGlitchRate(val)
        return nil
    }
    res["receivedNoiseLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetReceivedNoiseLevel(val)
        return nil
    }
    res["receivedSignalLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetReceivedSignalLevel(val)
        return nil
    }
    res["renderDeviceDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRenderDeviceDriver(val)
        return nil
    }
    res["renderDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRenderDeviceName(val)
        return nil
    }
    res["renderMuteEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetRenderMuteEventRatio(val)
        return nil
    }
    res["renderNotFunctioningEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetRenderNotFunctioningEventRatio(val)
        return nil
    }
    res["renderZeroVolumeEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetRenderZeroVolumeEventRatio(val)
        return nil
    }
    res["sentNoiseLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSentNoiseLevel(val)
        return nil
    }
    res["sentSignalLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSentSignalLevel(val)
        return nil
    }
    res["speakerGlitchRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetSpeakerGlitchRate(val)
        return nil
    }
    return res
}
func (m *DeviceInfo) IsNil()(bool) {
    return m == nil
}
func (m *DeviceInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("captureDeviceDriver", m.GetCaptureDeviceDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("captureDeviceName", m.GetCaptureDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("captureNotFunctioningEventRatio", m.GetCaptureNotFunctioningEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("cpuInsufficentEventRatio", m.GetCpuInsufficentEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("deviceClippingEventRatio", m.GetDeviceClippingEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("deviceGlitchEventRatio", m.GetDeviceGlitchEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("howlingEventCount", m.GetHowlingEventCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("initialSignalLevelRootMeanSquare", m.GetInitialSignalLevelRootMeanSquare())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("lowSpeechLevelEventRatio", m.GetLowSpeechLevelEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("lowSpeechToNoiseEventRatio", m.GetLowSpeechToNoiseEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("micGlitchRate", m.GetMicGlitchRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("receivedNoiseLevel", m.GetReceivedNoiseLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("receivedSignalLevel", m.GetReceivedSignalLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("renderDeviceDriver", m.GetRenderDeviceDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("renderDeviceName", m.GetRenderDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("renderMuteEventRatio", m.GetRenderMuteEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("renderNotFunctioningEventRatio", m.GetRenderNotFunctioningEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("renderZeroVolumeEventRatio", m.GetRenderZeroVolumeEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sentNoiseLevel", m.GetSentNoiseLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sentSignalLevel", m.GetSentSignalLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("speakerGlitchRate", m.GetSpeakerGlitchRate())
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
func (m *DeviceInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceInfo) SetCaptureDeviceDriver(value *string)() {
    m.captureDeviceDriver = value
}
func (m *DeviceInfo) SetCaptureDeviceName(value *string)() {
    m.captureDeviceName = value
}
func (m *DeviceInfo) SetCaptureNotFunctioningEventRatio(value *float32)() {
    m.captureNotFunctioningEventRatio = value
}
func (m *DeviceInfo) SetCpuInsufficentEventRatio(value *float32)() {
    m.cpuInsufficentEventRatio = value
}
func (m *DeviceInfo) SetDeviceClippingEventRatio(value *float32)() {
    m.deviceClippingEventRatio = value
}
func (m *DeviceInfo) SetDeviceGlitchEventRatio(value *float32)() {
    m.deviceGlitchEventRatio = value
}
func (m *DeviceInfo) SetHowlingEventCount(value *int32)() {
    m.howlingEventCount = value
}
func (m *DeviceInfo) SetInitialSignalLevelRootMeanSquare(value *float32)() {
    m.initialSignalLevelRootMeanSquare = value
}
func (m *DeviceInfo) SetLowSpeechLevelEventRatio(value *float32)() {
    m.lowSpeechLevelEventRatio = value
}
func (m *DeviceInfo) SetLowSpeechToNoiseEventRatio(value *float32)() {
    m.lowSpeechToNoiseEventRatio = value
}
func (m *DeviceInfo) SetMicGlitchRate(value *float32)() {
    m.micGlitchRate = value
}
func (m *DeviceInfo) SetReceivedNoiseLevel(value *int32)() {
    m.receivedNoiseLevel = value
}
func (m *DeviceInfo) SetReceivedSignalLevel(value *int32)() {
    m.receivedSignalLevel = value
}
func (m *DeviceInfo) SetRenderDeviceDriver(value *string)() {
    m.renderDeviceDriver = value
}
func (m *DeviceInfo) SetRenderDeviceName(value *string)() {
    m.renderDeviceName = value
}
func (m *DeviceInfo) SetRenderMuteEventRatio(value *float32)() {
    m.renderMuteEventRatio = value
}
func (m *DeviceInfo) SetRenderNotFunctioningEventRatio(value *float32)() {
    m.renderNotFunctioningEventRatio = value
}
func (m *DeviceInfo) SetRenderZeroVolumeEventRatio(value *float32)() {
    m.renderZeroVolumeEventRatio = value
}
func (m *DeviceInfo) SetSentNoiseLevel(value *int32)() {
    m.sentNoiseLevel = value
}
func (m *DeviceInfo) SetSentSignalLevel(value *int32)() {
    m.sentSignalLevel = value
}
func (m *DeviceInfo) SetSpeakerGlitchRate(value *float32)() {
    m.speakerGlitchRate = value
}
