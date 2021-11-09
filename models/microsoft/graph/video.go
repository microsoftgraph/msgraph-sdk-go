package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Video struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of audio bits per sample.
    audioBitsPerSample *int32;
    // Number of audio channels.
    audioChannels *int32;
    // Name of the audio format (AAC, MP3, etc.).
    audioFormat *string;
    // Number of audio samples per second.
    audioSamplesPerSecond *int32;
    // Bit rate of the video in bits per second.
    bitrate *int32;
    // Duration of the file in milliseconds.
    duration *int64;
    // 'Four character code' name of the video format.
    fourCC *string;
    // Frame rate of the video.
    frameRate *float64;
    // Height of the video, in pixels.
    height *int32;
    // Width of the video, in pixels.
    width *int32;
}
// Instantiates a new video and sets the default values.
func NewVideo()(*Video) {
    m := &Video{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Video) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the audioBitsPerSample property value. Number of audio bits per sample.
func (m *Video) GetAudioBitsPerSample()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioBitsPerSample
    }
}
// Gets the audioChannels property value. Number of audio channels.
func (m *Video) GetAudioChannels()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioChannels
    }
}
// Gets the audioFormat property value. Name of the audio format (AAC, MP3, etc.).
func (m *Video) GetAudioFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.audioFormat
    }
}
// Gets the audioSamplesPerSecond property value. Number of audio samples per second.
func (m *Video) GetAudioSamplesPerSecond()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioSamplesPerSecond
    }
}
// Gets the bitrate property value. Bit rate of the video in bits per second.
func (m *Video) GetBitrate()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bitrate
    }
}
// Gets the duration property value. Duration of the file in milliseconds.
func (m *Video) GetDuration()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// Gets the fourCC property value. 'Four character code' name of the video format.
func (m *Video) GetFourCC()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fourCC
    }
}
// Gets the frameRate property value. Frame rate of the video.
func (m *Video) GetFrameRate()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.frameRate
    }
}
// Gets the height property value. Height of the video, in pixels.
func (m *Video) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// Gets the width property value. Width of the video, in pixels.
func (m *Video) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
// The deserialization information for the current model
func (m *Video) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["audioBitsPerSample"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioBitsPerSample(val)
        }
        return nil
    }
    res["audioChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioChannels(val)
        }
        return nil
    }
    res["audioFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioFormat(val)
        }
        return nil
    }
    res["audioSamplesPerSecond"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioSamplesPerSecond(val)
        }
        return nil
    }
    res["bitrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate(val)
        }
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["fourCC"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFourCC(val)
        }
        return nil
    }
    res["frameRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrameRate(val)
        }
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
func (m *Video) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Video) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("audioBitsPerSample", m.GetAudioBitsPerSample())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("audioChannels", m.GetAudioChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("audioFormat", m.GetAudioFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("audioSamplesPerSecond", m.GetAudioSamplesPerSecond())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("bitrate", m.GetBitrate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fourCC", m.GetFourCC())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("frameRate", m.GetFrameRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("width", m.GetWidth())
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
func (m *Video) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the audioBitsPerSample property value. Number of audio bits per sample.
// Parameters:
//  - value : Value to set for the audioBitsPerSample property.
func (m *Video) SetAudioBitsPerSample(value *int32)() {
    m.audioBitsPerSample = value
}
// Sets the audioChannels property value. Number of audio channels.
// Parameters:
//  - value : Value to set for the audioChannels property.
func (m *Video) SetAudioChannels(value *int32)() {
    m.audioChannels = value
}
// Sets the audioFormat property value. Name of the audio format (AAC, MP3, etc.).
// Parameters:
//  - value : Value to set for the audioFormat property.
func (m *Video) SetAudioFormat(value *string)() {
    m.audioFormat = value
}
// Sets the audioSamplesPerSecond property value. Number of audio samples per second.
// Parameters:
//  - value : Value to set for the audioSamplesPerSecond property.
func (m *Video) SetAudioSamplesPerSecond(value *int32)() {
    m.audioSamplesPerSecond = value
}
// Sets the bitrate property value. Bit rate of the video in bits per second.
// Parameters:
//  - value : Value to set for the bitrate property.
func (m *Video) SetBitrate(value *int32)() {
    m.bitrate = value
}
// Sets the duration property value. Duration of the file in milliseconds.
// Parameters:
//  - value : Value to set for the duration property.
func (m *Video) SetDuration(value *int64)() {
    m.duration = value
}
// Sets the fourCC property value. 'Four character code' name of the video format.
// Parameters:
//  - value : Value to set for the fourCC property.
func (m *Video) SetFourCC(value *string)() {
    m.fourCC = value
}
// Sets the frameRate property value. Frame rate of the video.
// Parameters:
//  - value : Value to set for the frameRate property.
func (m *Video) SetFrameRate(value *float64)() {
    m.frameRate = value
}
// Sets the height property value. Height of the video, in pixels.
// Parameters:
//  - value : Value to set for the height property.
func (m *Video) SetHeight(value *int32)() {
    m.height = value
}
// Sets the width property value. Width of the video, in pixels.
// Parameters:
//  - value : Value to set for the width property.
func (m *Video) SetWidth(value *int32)() {
    m.width = value
}
