package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Videoable 
type Videoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAudioBitsPerSample()(*int32)
    GetAudioChannels()(*int32)
    GetAudioFormat()(*string)
    GetAudioSamplesPerSecond()(*int32)
    GetBitrate()(*int32)
    GetDuration()(*int32)
    GetFourCC()(*string)
    GetFrameRate()(*float64)
    GetHeight()(*int32)
    GetWidth()(*int32)
    SetAudioBitsPerSample(value *int32)()
    SetAudioChannels(value *int32)()
    SetAudioFormat(value *string)()
    SetAudioSamplesPerSecond(value *int32)()
    SetBitrate(value *int32)()
    SetDuration(value *int32)()
    SetFourCC(value *string)()
    SetFrameRate(value *float64)()
    SetHeight(value *int32)()
    SetWidth(value *int32)()
}
