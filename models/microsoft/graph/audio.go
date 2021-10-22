package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Audio struct {
    additionalData map[string]interface{};
    album *string;
    albumArtist *string;
    artist *string;
    bitrate *int64;
    composers *string;
    copyright *string;
    disc *int32;
    discCount *int32;
    duration *int64;
    genre *string;
    hasDrm *bool;
    isVariableBitrate *bool;
    title *string;
    track *int32;
    trackCount *int32;
    year *int32;
}
func NewAudio()(*Audio) {
    m := &Audio{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Audio) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Audio) GetAlbum()(*string) {
    if m == nil {
        return nil
    } else {
        return m.album
    }
}
func (m *Audio) GetAlbumArtist()(*string) {
    if m == nil {
        return nil
    } else {
        return m.albumArtist
    }
}
func (m *Audio) GetArtist()(*string) {
    if m == nil {
        return nil
    } else {
        return m.artist
    }
}
func (m *Audio) GetBitrate()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.bitrate
    }
}
func (m *Audio) GetComposers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.composers
    }
}
func (m *Audio) GetCopyright()(*string) {
    if m == nil {
        return nil
    } else {
        return m.copyright
    }
}
func (m *Audio) GetDisc()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.disc
    }
}
func (m *Audio) GetDiscCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.discCount
    }
}
func (m *Audio) GetDuration()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
func (m *Audio) GetGenre()(*string) {
    if m == nil {
        return nil
    } else {
        return m.genre
    }
}
func (m *Audio) GetHasDrm()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasDrm
    }
}
func (m *Audio) GetIsVariableBitrate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVariableBitrate
    }
}
func (m *Audio) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *Audio) GetTrack()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.track
    }
}
func (m *Audio) GetTrackCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.trackCount
    }
}
func (m *Audio) GetYear()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.year
    }
}
func (m *Audio) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["album"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlbum(val)
        return nil
    }
    res["albumArtist"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlbumArtist(val)
        return nil
    }
    res["artist"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetArtist(val)
        return nil
    }
    res["bitrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetBitrate(val)
        return nil
    }
    res["composers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetComposers(val)
        return nil
    }
    res["copyright"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCopyright(val)
        return nil
    }
    res["disc"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDisc(val)
        return nil
    }
    res["discCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDiscCount(val)
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
    res["genre"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGenre(val)
        return nil
    }
    res["hasDrm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasDrm(val)
        return nil
    }
    res["isVariableBitrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsVariableBitrate(val)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    res["track"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTrack(val)
        return nil
    }
    res["trackCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTrackCount(val)
        return nil
    }
    res["year"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetYear(val)
        return nil
    }
    return res
}
func (m *Audio) IsNil()(bool) {
    return m == nil
}
func (m *Audio) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("album", m.GetAlbum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("albumArtist", m.GetAlbumArtist())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("artist", m.GetArtist())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("bitrate", m.GetBitrate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("composers", m.GetComposers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("copyright", m.GetCopyright())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("disc", m.GetDisc())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("discCount", m.GetDiscCount())
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
        err := writer.WriteStringValue("genre", m.GetGenre())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasDrm", m.GetHasDrm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVariableBitrate", m.GetIsVariableBitrate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("track", m.GetTrack())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("trackCount", m.GetTrackCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("year", m.GetYear())
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
func (m *Audio) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Audio) SetAlbum(value *string)() {
    m.album = value
}
func (m *Audio) SetAlbumArtist(value *string)() {
    m.albumArtist = value
}
func (m *Audio) SetArtist(value *string)() {
    m.artist = value
}
func (m *Audio) SetBitrate(value *int64)() {
    m.bitrate = value
}
func (m *Audio) SetComposers(value *string)() {
    m.composers = value
}
func (m *Audio) SetCopyright(value *string)() {
    m.copyright = value
}
func (m *Audio) SetDisc(value *int32)() {
    m.disc = value
}
func (m *Audio) SetDiscCount(value *int32)() {
    m.discCount = value
}
func (m *Audio) SetDuration(value *int64)() {
    m.duration = value
}
func (m *Audio) SetGenre(value *string)() {
    m.genre = value
}
func (m *Audio) SetHasDrm(value *bool)() {
    m.hasDrm = value
}
func (m *Audio) SetIsVariableBitrate(value *bool)() {
    m.isVariableBitrate = value
}
func (m *Audio) SetTitle(value *string)() {
    m.title = value
}
func (m *Audio) SetTrack(value *int32)() {
    m.track = value
}
func (m *Audio) SetTrackCount(value *int32)() {
    m.trackCount = value
}
func (m *Audio) SetYear(value *int32)() {
    m.year = value
}
