package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceInfoable 
type DeviceInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCaptureDeviceDriver()(*string)
    GetCaptureDeviceName()(*string)
    GetCaptureNotFunctioningEventRatio()(*float32)
    GetCpuInsufficentEventRatio()(*float32)
    GetDeviceClippingEventRatio()(*float32)
    GetDeviceGlitchEventRatio()(*float32)
    GetHowlingEventCount()(*int32)
    GetInitialSignalLevelRootMeanSquare()(*float32)
    GetLowSpeechLevelEventRatio()(*float32)
    GetLowSpeechToNoiseEventRatio()(*float32)
    GetMicGlitchRate()(*float32)
    GetReceivedNoiseLevel()(*int32)
    GetReceivedSignalLevel()(*int32)
    GetRenderDeviceDriver()(*string)
    GetRenderDeviceName()(*string)
    GetRenderMuteEventRatio()(*float32)
    GetRenderNotFunctioningEventRatio()(*float32)
    GetRenderZeroVolumeEventRatio()(*float32)
    GetSentNoiseLevel()(*int32)
    GetSentSignalLevel()(*int32)
    GetSpeakerGlitchRate()(*float32)
    SetCaptureDeviceDriver(value *string)()
    SetCaptureDeviceName(value *string)()
    SetCaptureNotFunctioningEventRatio(value *float32)()
    SetCpuInsufficentEventRatio(value *float32)()
    SetDeviceClippingEventRatio(value *float32)()
    SetDeviceGlitchEventRatio(value *float32)()
    SetHowlingEventCount(value *int32)()
    SetInitialSignalLevelRootMeanSquare(value *float32)()
    SetLowSpeechLevelEventRatio(value *float32)()
    SetLowSpeechToNoiseEventRatio(value *float32)()
    SetMicGlitchRate(value *float32)()
    SetReceivedNoiseLevel(value *int32)()
    SetReceivedSignalLevel(value *int32)()
    SetRenderDeviceDriver(value *string)()
    SetRenderDeviceName(value *string)()
    SetRenderMuteEventRatio(value *float32)()
    SetRenderNotFunctioningEventRatio(value *float32)()
    SetRenderZeroVolumeEventRatio(value *float32)()
    SetSentNoiseLevel(value *int32)()
    SetSentSignalLevel(value *int32)()
    SetSpeakerGlitchRate(value *float32)()
}
