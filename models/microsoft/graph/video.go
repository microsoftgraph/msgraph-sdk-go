package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Video provides operations to manage the drive singleton.
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
// NewVideo instantiates a new video and sets the default values.
func NewVideo()(*Video) {
    m := &Video{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVideoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVideoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewVideo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Video) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAudioBitsPerSample gets the audioBitsPerSample property value. Number of audio bits per sample.
func (m *Video) GetAudioBitsPerSample()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioBitsPerSample
    }
}
// GetAudioChannels gets the audioChannels property value. Number of audio channels.
func (m *Video) GetAudioChannels()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioChannels
    }
}
// GetAudioFormat gets the audioFormat property value. Name of the audio format (AAC, MP3, etc.).
func (m *Video) GetAudioFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.audioFormat
    }
}
// GetAudioSamplesPerSecond gets the audioSamplesPerSecond property value. Number of audio samples per second.
func (m *Video) GetAudioSamplesPerSecond()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioSamplesPerSecond
    }
}
// GetBitrate gets the bitrate property value. Bit rate of the video in bits per second.
func (m *Video) GetBitrate()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bitrate
    }
}
// GetDuration gets the duration property value. Duration of the file in milliseconds.
func (m *Video) GetDuration()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetFourCC gets the fourCC property value. 'Four character code' name of the video format.
func (m *Video) GetFourCC()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fourCC
    }
}
// GetFrameRate gets the frameRate property value. Frame rate of the video.
func (m *Video) GetFrameRate()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.frameRate
    }
}
// GetHeight gets the height property value. Height of the video, in pixels.
func (m *Video) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// GetWidth gets the width property value. Width of the video, in pixels.
func (m *Video) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
func (m *Video) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Video) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAudioBitsPerSample sets the audioBitsPerSample property value. Number of audio bits per sample.
func (m *Video) SetAudioBitsPerSample(value *int32)() {
    if m != nil {
        m.audioBitsPerSample = value
    }
}
// SetAudioChannels sets the audioChannels property value. Number of audio channels.
func (m *Video) SetAudioChannels(value *int32)() {
    if m != nil {
        m.audioChannels = value
    }
}
// SetAudioFormat sets the audioFormat property value. Name of the audio format (AAC, MP3, etc.).
func (m *Video) SetAudioFormat(value *string)() {
    if m != nil {
        m.audioFormat = value
    }
}
// SetAudioSamplesPerSecond sets the audioSamplesPerSecond property value. Number of audio samples per second.
func (m *Video) SetAudioSamplesPerSecond(value *int32)() {
    if m != nil {
        m.audioSamplesPerSecond = value
    }
}
// SetBitrate sets the bitrate property value. Bit rate of the video in bits per second.
func (m *Video) SetBitrate(value *int32)() {
    if m != nil {
        m.bitrate = value
    }
}
// SetDuration sets the duration property value. Duration of the file in milliseconds.
func (m *Video) SetDuration(value *int64)() {
    if m != nil {
        m.duration = value
    }
}
// SetFourCC sets the fourCC property value. 'Four character code' name of the video format.
func (m *Video) SetFourCC(value *string)() {
    if m != nil {
        m.fourCC = value
    }
}
// SetFrameRate sets the frameRate property value. Frame rate of the video.
func (m *Video) SetFrameRate(value *float64)() {
    if m != nil {
        m.frameRate = value
    }
}
// SetHeight sets the height property value. Height of the video, in pixels.
func (m *Video) SetHeight(value *int32)() {
    if m != nil {
        m.height = value
    }
}
// SetWidth sets the width property value. Width of the video, in pixels.
func (m *Video) SetWidth(value *int32)() {
    if m != nil {
        m.width = value
    }
}
