package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the capture device driver used by the media endpoint.
    captureDeviceDriver *string;
    // Name of the capture device used by the media endpoint.
    captureDeviceName *string;
    // Fraction of the call that the media endpoint detected the capture device was not working properly.
    captureNotFunctioningEventRatio *float32;
    // Fraction of the call that the media endpoint detected the CPU resources available were insufficient and caused poor quality of the audio sent and received.
    cpuInsufficentEventRatio *float32;
    // Fraction of the call that the media endpoint detected clipping in the captured audio that caused poor quality of the audio being sent.
    deviceClippingEventRatio *float32;
    // Fraction of the call that the media endpoint detected glitches or gaps in the audio played or captured that caused poor quality of the audio being sent or received.
    deviceGlitchEventRatio *float32;
    // Number of times during the call that the media endpoint detected howling or screeching audio.
    howlingEventCount *int32;
    // The root mean square (RMS) of the incoming signal of up to the first 30 seconds of the call.
    initialSignalLevelRootMeanSquare *float32;
    // Fraction of the call that the media endpoint detected low speech level that caused poor quality of the audio being sent.
    lowSpeechLevelEventRatio *float32;
    // Fraction of the call that the media endpoint detected low speech to noise level that caused poor quality of the audio being sent.
    lowSpeechToNoiseEventRatio *float32;
    // Glitches per 5 minute interval for the media endpoint's microphone.
    micGlitchRate *float32;
    // Average energy level of received audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
    receivedNoiseLevel *int32;
    // Average energy level of received audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
    receivedSignalLevel *int32;
    // Name of the render device driver used by the media endpoint.
    renderDeviceDriver *string;
    // Name of the render device used by the media endpoint.
    renderDeviceName *string;
    // Fraction of the call that media endpoint detected device render is muted.
    renderMuteEventRatio *float32;
    // Fraction of the call that the media endpoint detected the render device was not working properly.
    renderNotFunctioningEventRatio *float32;
    // Fraction of the call that media endpoint detected device render volume is set to 0.
    renderZeroVolumeEventRatio *float32;
    // Average energy level of sent audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
    sentNoiseLevel *int32;
    // Average energy level of sent audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
    sentSignalLevel *int32;
    // Glitches per 5 minute internal for the media endpoint's loudspeaker.
    speakerGlitchRate *float32;
}
// Instantiates a new deviceInfo and sets the default values.
func NewDeviceInfo()(*DeviceInfo) {
    m := &DeviceInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the captureDeviceDriver property value. Name of the capture device driver used by the media endpoint.
func (m *DeviceInfo) GetCaptureDeviceDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.captureDeviceDriver
    }
}
// Gets the captureDeviceName property value. Name of the capture device used by the media endpoint.
func (m *DeviceInfo) GetCaptureDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.captureDeviceName
    }
}
// Gets the captureNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the capture device was not working properly.
func (m *DeviceInfo) GetCaptureNotFunctioningEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.captureNotFunctioningEventRatio
    }
}
// Gets the cpuInsufficentEventRatio property value. Fraction of the call that the media endpoint detected the CPU resources available were insufficient and caused poor quality of the audio sent and received.
func (m *DeviceInfo) GetCpuInsufficentEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.cpuInsufficentEventRatio
    }
}
// Gets the deviceClippingEventRatio property value. Fraction of the call that the media endpoint detected clipping in the captured audio that caused poor quality of the audio being sent.
func (m *DeviceInfo) GetDeviceClippingEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.deviceClippingEventRatio
    }
}
// Gets the deviceGlitchEventRatio property value. Fraction of the call that the media endpoint detected glitches or gaps in the audio played or captured that caused poor quality of the audio being sent or received.
func (m *DeviceInfo) GetDeviceGlitchEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.deviceGlitchEventRatio
    }
}
// Gets the howlingEventCount property value. Number of times during the call that the media endpoint detected howling or screeching audio.
func (m *DeviceInfo) GetHowlingEventCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.howlingEventCount
    }
}
// Gets the initialSignalLevelRootMeanSquare property value. The root mean square (RMS) of the incoming signal of up to the first 30 seconds of the call.
func (m *DeviceInfo) GetInitialSignalLevelRootMeanSquare()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.initialSignalLevelRootMeanSquare
    }
}
// Gets the lowSpeechLevelEventRatio property value. Fraction of the call that the media endpoint detected low speech level that caused poor quality of the audio being sent.
func (m *DeviceInfo) GetLowSpeechLevelEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.lowSpeechLevelEventRatio
    }
}
// Gets the lowSpeechToNoiseEventRatio property value. Fraction of the call that the media endpoint detected low speech to noise level that caused poor quality of the audio being sent.
func (m *DeviceInfo) GetLowSpeechToNoiseEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.lowSpeechToNoiseEventRatio
    }
}
// Gets the micGlitchRate property value. Glitches per 5 minute interval for the media endpoint's microphone.
func (m *DeviceInfo) GetMicGlitchRate()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.micGlitchRate
    }
}
// Gets the receivedNoiseLevel property value. Average energy level of received audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
func (m *DeviceInfo) GetReceivedNoiseLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.receivedNoiseLevel
    }
}
// Gets the receivedSignalLevel property value. Average energy level of received audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
func (m *DeviceInfo) GetReceivedSignalLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.receivedSignalLevel
    }
}
// Gets the renderDeviceDriver property value. Name of the render device driver used by the media endpoint.
func (m *DeviceInfo) GetRenderDeviceDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.renderDeviceDriver
    }
}
// Gets the renderDeviceName property value. Name of the render device used by the media endpoint.
func (m *DeviceInfo) GetRenderDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.renderDeviceName
    }
}
// Gets the renderMuteEventRatio property value. Fraction of the call that media endpoint detected device render is muted.
func (m *DeviceInfo) GetRenderMuteEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.renderMuteEventRatio
    }
}
// Gets the renderNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the render device was not working properly.
func (m *DeviceInfo) GetRenderNotFunctioningEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.renderNotFunctioningEventRatio
    }
}
// Gets the renderZeroVolumeEventRatio property value. Fraction of the call that media endpoint detected device render volume is set to 0.
func (m *DeviceInfo) GetRenderZeroVolumeEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.renderZeroVolumeEventRatio
    }
}
// Gets the sentNoiseLevel property value. Average energy level of sent audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
func (m *DeviceInfo) GetSentNoiseLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sentNoiseLevel
    }
}
// Gets the sentSignalLevel property value. Average energy level of sent audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
func (m *DeviceInfo) GetSentSignalLevel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sentSignalLevel
    }
}
// Gets the speakerGlitchRate property value. Glitches per 5 minute internal for the media endpoint's loudspeaker.
func (m *DeviceInfo) GetSpeakerGlitchRate()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.speakerGlitchRate
    }
}
// The deserialization information for the current model
func (m *DeviceInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["captureDeviceDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptureDeviceDriver(val)
        }
        return nil
    }
    res["captureDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptureDeviceName(val)
        }
        return nil
    }
    res["captureNotFunctioningEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptureNotFunctioningEventRatio(val)
        }
        return nil
    }
    res["cpuInsufficentEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuInsufficentEventRatio(val)
        }
        return nil
    }
    res["deviceClippingEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceClippingEventRatio(val)
        }
        return nil
    }
    res["deviceGlitchEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGlitchEventRatio(val)
        }
        return nil
    }
    res["howlingEventCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHowlingEventCount(val)
        }
        return nil
    }
    res["initialSignalLevelRootMeanSquare"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitialSignalLevelRootMeanSquare(val)
        }
        return nil
    }
    res["lowSpeechLevelEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowSpeechLevelEventRatio(val)
        }
        return nil
    }
    res["lowSpeechToNoiseEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowSpeechToNoiseEventRatio(val)
        }
        return nil
    }
    res["micGlitchRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicGlitchRate(val)
        }
        return nil
    }
    res["receivedNoiseLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedNoiseLevel(val)
        }
        return nil
    }
    res["receivedSignalLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedSignalLevel(val)
        }
        return nil
    }
    res["renderDeviceDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderDeviceDriver(val)
        }
        return nil
    }
    res["renderDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderDeviceName(val)
        }
        return nil
    }
    res["renderMuteEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderMuteEventRatio(val)
        }
        return nil
    }
    res["renderNotFunctioningEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderNotFunctioningEventRatio(val)
        }
        return nil
    }
    res["renderZeroVolumeEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenderZeroVolumeEventRatio(val)
        }
        return nil
    }
    res["sentNoiseLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentNoiseLevel(val)
        }
        return nil
    }
    res["sentSignalLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentSignalLevel(val)
        }
        return nil
    }
    res["speakerGlitchRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerGlitchRate(val)
        }
        return nil
    }
    return res
}
func (m *DeviceInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the captureDeviceDriver property value. Name of the capture device driver used by the media endpoint.
// Parameters:
//  - value : Value to set for the captureDeviceDriver property.
func (m *DeviceInfo) SetCaptureDeviceDriver(value *string)() {
    m.captureDeviceDriver = value
}
// Sets the captureDeviceName property value. Name of the capture device used by the media endpoint.
// Parameters:
//  - value : Value to set for the captureDeviceName property.
func (m *DeviceInfo) SetCaptureDeviceName(value *string)() {
    m.captureDeviceName = value
}
// Sets the captureNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the capture device was not working properly.
// Parameters:
//  - value : Value to set for the captureNotFunctioningEventRatio property.
func (m *DeviceInfo) SetCaptureNotFunctioningEventRatio(value *float32)() {
    m.captureNotFunctioningEventRatio = value
}
// Sets the cpuInsufficentEventRatio property value. Fraction of the call that the media endpoint detected the CPU resources available were insufficient and caused poor quality of the audio sent and received.
// Parameters:
//  - value : Value to set for the cpuInsufficentEventRatio property.
func (m *DeviceInfo) SetCpuInsufficentEventRatio(value *float32)() {
    m.cpuInsufficentEventRatio = value
}
// Sets the deviceClippingEventRatio property value. Fraction of the call that the media endpoint detected clipping in the captured audio that caused poor quality of the audio being sent.
// Parameters:
//  - value : Value to set for the deviceClippingEventRatio property.
func (m *DeviceInfo) SetDeviceClippingEventRatio(value *float32)() {
    m.deviceClippingEventRatio = value
}
// Sets the deviceGlitchEventRatio property value. Fraction of the call that the media endpoint detected glitches or gaps in the audio played or captured that caused poor quality of the audio being sent or received.
// Parameters:
//  - value : Value to set for the deviceGlitchEventRatio property.
func (m *DeviceInfo) SetDeviceGlitchEventRatio(value *float32)() {
    m.deviceGlitchEventRatio = value
}
// Sets the howlingEventCount property value. Number of times during the call that the media endpoint detected howling or screeching audio.
// Parameters:
//  - value : Value to set for the howlingEventCount property.
func (m *DeviceInfo) SetHowlingEventCount(value *int32)() {
    m.howlingEventCount = value
}
// Sets the initialSignalLevelRootMeanSquare property value. The root mean square (RMS) of the incoming signal of up to the first 30 seconds of the call.
// Parameters:
//  - value : Value to set for the initialSignalLevelRootMeanSquare property.
func (m *DeviceInfo) SetInitialSignalLevelRootMeanSquare(value *float32)() {
    m.initialSignalLevelRootMeanSquare = value
}
// Sets the lowSpeechLevelEventRatio property value. Fraction of the call that the media endpoint detected low speech level that caused poor quality of the audio being sent.
// Parameters:
//  - value : Value to set for the lowSpeechLevelEventRatio property.
func (m *DeviceInfo) SetLowSpeechLevelEventRatio(value *float32)() {
    m.lowSpeechLevelEventRatio = value
}
// Sets the lowSpeechToNoiseEventRatio property value. Fraction of the call that the media endpoint detected low speech to noise level that caused poor quality of the audio being sent.
// Parameters:
//  - value : Value to set for the lowSpeechToNoiseEventRatio property.
func (m *DeviceInfo) SetLowSpeechToNoiseEventRatio(value *float32)() {
    m.lowSpeechToNoiseEventRatio = value
}
// Sets the micGlitchRate property value. Glitches per 5 minute interval for the media endpoint's microphone.
// Parameters:
//  - value : Value to set for the micGlitchRate property.
func (m *DeviceInfo) SetMicGlitchRate(value *float32)() {
    m.micGlitchRate = value
}
// Sets the receivedNoiseLevel property value. Average energy level of received audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
// Parameters:
//  - value : Value to set for the receivedNoiseLevel property.
func (m *DeviceInfo) SetReceivedNoiseLevel(value *int32)() {
    m.receivedNoiseLevel = value
}
// Sets the receivedSignalLevel property value. Average energy level of received audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
// Parameters:
//  - value : Value to set for the receivedSignalLevel property.
func (m *DeviceInfo) SetReceivedSignalLevel(value *int32)() {
    m.receivedSignalLevel = value
}
// Sets the renderDeviceDriver property value. Name of the render device driver used by the media endpoint.
// Parameters:
//  - value : Value to set for the renderDeviceDriver property.
func (m *DeviceInfo) SetRenderDeviceDriver(value *string)() {
    m.renderDeviceDriver = value
}
// Sets the renderDeviceName property value. Name of the render device used by the media endpoint.
// Parameters:
//  - value : Value to set for the renderDeviceName property.
func (m *DeviceInfo) SetRenderDeviceName(value *string)() {
    m.renderDeviceName = value
}
// Sets the renderMuteEventRatio property value. Fraction of the call that media endpoint detected device render is muted.
// Parameters:
//  - value : Value to set for the renderMuteEventRatio property.
func (m *DeviceInfo) SetRenderMuteEventRatio(value *float32)() {
    m.renderMuteEventRatio = value
}
// Sets the renderNotFunctioningEventRatio property value. Fraction of the call that the media endpoint detected the render device was not working properly.
// Parameters:
//  - value : Value to set for the renderNotFunctioningEventRatio property.
func (m *DeviceInfo) SetRenderNotFunctioningEventRatio(value *float32)() {
    m.renderNotFunctioningEventRatio = value
}
// Sets the renderZeroVolumeEventRatio property value. Fraction of the call that media endpoint detected device render volume is set to 0.
// Parameters:
//  - value : Value to set for the renderZeroVolumeEventRatio property.
func (m *DeviceInfo) SetRenderZeroVolumeEventRatio(value *float32)() {
    m.renderZeroVolumeEventRatio = value
}
// Sets the sentNoiseLevel property value. Average energy level of sent audio for audio classified as mono noise or left channel of stereo noise by the media endpoint.
// Parameters:
//  - value : Value to set for the sentNoiseLevel property.
func (m *DeviceInfo) SetSentNoiseLevel(value *int32)() {
    m.sentNoiseLevel = value
}
// Sets the sentSignalLevel property value. Average energy level of sent audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint.
// Parameters:
//  - value : Value to set for the sentSignalLevel property.
func (m *DeviceInfo) SetSentSignalLevel(value *int32)() {
    m.sentSignalLevel = value
}
// Sets the speakerGlitchRate property value. Glitches per 5 minute internal for the media endpoint's loudspeaker.
// Parameters:
//  - value : Value to set for the speakerGlitchRate property.
func (m *DeviceInfo) SetSpeakerGlitchRate(value *float32)() {
    m.speakerGlitchRate = value
}
