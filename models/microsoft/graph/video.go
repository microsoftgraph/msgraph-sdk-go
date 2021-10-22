package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Video struct {
    additionalData map[string]interface{};
    audioBitsPerSample *int32;
    audioChannels *int32;
    audioFormat *string;
    audioSamplesPerSecond *int32;
    bitrate *int32;
    duration *int64;
    fourCC *string;
    frameRate *float64;
    height *int32;
    width *int32;
}
func NewVideo()(*Video) {
    m := &Video{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Video) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Video) GetAudioBitsPerSample()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioBitsPerSample
    }
}
func (m *Video) GetAudioChannels()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioChannels
    }
}
func (m *Video) GetAudioFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.audioFormat
    }
}
func (m *Video) GetAudioSamplesPerSecond()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.audioSamplesPerSecond
    }
}
func (m *Video) GetBitrate()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bitrate
    }
}
func (m *Video) GetDuration()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
func (m *Video) GetFourCC()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fourCC
    }
}
func (m *Video) GetFrameRate()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.frameRate
    }
}
func (m *Video) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
func (m *Video) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
func (m *Video) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["audioBitsPerSample"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAudioBitsPerSample(val)
        return nil
    }
    res["audioChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAudioChannels(val)
        return nil
    }
    res["audioFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAudioFormat(val)
        return nil
    }
    res["audioSamplesPerSecond"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAudioSamplesPerSecond(val)
        return nil
    }
    res["bitrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBitrate(val)
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDuration(val)
        return nil
    }
    res["fourCC"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFourCC(val)
        return nil
    }
    res["frameRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetFrameRate(val)
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetHeight(val)
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWidth(val)
        return nil
    }
    return res
}
func (m *Video) IsNil()(bool) {
    return m == nil
}
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
func (m *Video) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Video) SetAudioBitsPerSample(value *int32)() {
    m.audioBitsPerSample = value
}
func (m *Video) SetAudioChannels(value *int32)() {
    m.audioChannels = value
}
func (m *Video) SetAudioFormat(value *string)() {
    m.audioFormat = value
}
func (m *Video) SetAudioSamplesPerSecond(value *int32)() {
    m.audioSamplesPerSecond = value
}
func (m *Video) SetBitrate(value *int32)() {
    m.bitrate = value
}
func (m *Video) SetDuration(value *int64)() {
    m.duration = value
}
func (m *Video) SetFourCC(value *string)() {
    m.fourCC = value
}
func (m *Video) SetFrameRate(value *float64)() {
    m.frameRate = value
}
func (m *Video) SetHeight(value *int32)() {
    m.height = value
}
func (m *Video) SetWidth(value *int32)() {
    m.width = value
}
