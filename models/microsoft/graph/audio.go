package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Audio 
type Audio struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The title of the album for this audio file.
    album *string;
    // The artist named on the album for the audio file.
    albumArtist *string;
    // The performing artist for the audio file.
    artist *string;
    // Bitrate expressed in kbps.
    bitrate *int64;
    // The name of the composer of the audio file.
    composers *string;
    // Copyright information for the audio file.
    copyright *string;
    // The number of the disc this audio file came from.
    disc *int32;
    // The total number of discs in this album.
    discCount *int32;
    // Duration of the audio file, expressed in milliseconds
    duration *int64;
    // The genre of this audio file.
    genre *string;
    // Indicates if the file is protected with digital rights management.
    hasDrm *bool;
    // Indicates if the file is encoded with a variable bitrate.
    isVariableBitrate *bool;
    // The title of the audio file.
    title *string;
    // The number of the track on the original disc for this audio file.
    track *int32;
    // The total number of tracks on the original disc for this audio file.
    trackCount *int32;
    // The year the audio file was recorded.
    year *int32;
}
// NewAudio instantiates a new audio and sets the default values.
func NewAudio()(*Audio) {
    m := &Audio{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Audio) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlbum gets the album property value. The title of the album for this audio file.
func (m *Audio) GetAlbum()(*string) {
    if m == nil {
        return nil
    } else {
        return m.album
    }
}
// GetAlbumArtist gets the albumArtist property value. The artist named on the album for the audio file.
func (m *Audio) GetAlbumArtist()(*string) {
    if m == nil {
        return nil
    } else {
        return m.albumArtist
    }
}
// GetArtist gets the artist property value. The performing artist for the audio file.
func (m *Audio) GetArtist()(*string) {
    if m == nil {
        return nil
    } else {
        return m.artist
    }
}
// GetBitrate gets the bitrate property value. Bitrate expressed in kbps.
func (m *Audio) GetBitrate()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.bitrate
    }
}
// GetComposers gets the composers property value. The name of the composer of the audio file.
func (m *Audio) GetComposers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.composers
    }
}
// GetCopyright gets the copyright property value. Copyright information for the audio file.
func (m *Audio) GetCopyright()(*string) {
    if m == nil {
        return nil
    } else {
        return m.copyright
    }
}
// GetDisc gets the disc property value. The number of the disc this audio file came from.
func (m *Audio) GetDisc()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.disc
    }
}
// GetDiscCount gets the discCount property value. The total number of discs in this album.
func (m *Audio) GetDiscCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.discCount
    }
}
// GetDuration gets the duration property value. Duration of the audio file, expressed in milliseconds
func (m *Audio) GetDuration()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetGenre gets the genre property value. The genre of this audio file.
func (m *Audio) GetGenre()(*string) {
    if m == nil {
        return nil
    } else {
        return m.genre
    }
}
// GetHasDrm gets the hasDrm property value. Indicates if the file is protected with digital rights management.
func (m *Audio) GetHasDrm()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasDrm
    }
}
// GetIsVariableBitrate gets the isVariableBitrate property value. Indicates if the file is encoded with a variable bitrate.
func (m *Audio) GetIsVariableBitrate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVariableBitrate
    }
}
// GetTitle gets the title property value. The title of the audio file.
func (m *Audio) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetTrack gets the track property value. The number of the track on the original disc for this audio file.
func (m *Audio) GetTrack()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.track
    }
}
// GetTrackCount gets the trackCount property value. The total number of tracks on the original disc for this audio file.
func (m *Audio) GetTrackCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.trackCount
    }
}
// GetYear gets the year property value. The year the audio file was recorded.
func (m *Audio) GetYear()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.year
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Audio) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["album"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlbum(val)
        }
        return nil
    }
    res["albumArtist"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlbumArtist(val)
        }
        return nil
    }
    res["artist"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtist(val)
        }
        return nil
    }
    res["bitrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate(val)
        }
        return nil
    }
    res["composers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComposers(val)
        }
        return nil
    }
    res["copyright"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyright(val)
        }
        return nil
    }
    res["disc"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisc(val)
        }
        return nil
    }
    res["discCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscCount(val)
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
    res["genre"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGenre(val)
        }
        return nil
    }
    res["hasDrm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDrm(val)
        }
        return nil
    }
    res["isVariableBitrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVariableBitrate(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["track"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrack(val)
        }
        return nil
    }
    res["trackCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackCount(val)
        }
        return nil
    }
    res["year"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYear(val)
        }
        return nil
    }
    return res
}
func (m *Audio) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Audio) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlbum sets the album property value. The title of the album for this audio file.
func (m *Audio) SetAlbum(value *string)() {
    if m != nil {
        m.album = value
    }
}
// SetAlbumArtist sets the albumArtist property value. The artist named on the album for the audio file.
func (m *Audio) SetAlbumArtist(value *string)() {
    if m != nil {
        m.albumArtist = value
    }
}
// SetArtist sets the artist property value. The performing artist for the audio file.
func (m *Audio) SetArtist(value *string)() {
    if m != nil {
        m.artist = value
    }
}
// SetBitrate sets the bitrate property value. Bitrate expressed in kbps.
func (m *Audio) SetBitrate(value *int64)() {
    if m != nil {
        m.bitrate = value
    }
}
// SetComposers sets the composers property value. The name of the composer of the audio file.
func (m *Audio) SetComposers(value *string)() {
    if m != nil {
        m.composers = value
    }
}
// SetCopyright sets the copyright property value. Copyright information for the audio file.
func (m *Audio) SetCopyright(value *string)() {
    if m != nil {
        m.copyright = value
    }
}
// SetDisc sets the disc property value. The number of the disc this audio file came from.
func (m *Audio) SetDisc(value *int32)() {
    if m != nil {
        m.disc = value
    }
}
// SetDiscCount sets the discCount property value. The total number of discs in this album.
func (m *Audio) SetDiscCount(value *int32)() {
    if m != nil {
        m.discCount = value
    }
}
// SetDuration sets the duration property value. Duration of the audio file, expressed in milliseconds
func (m *Audio) SetDuration(value *int64)() {
    if m != nil {
        m.duration = value
    }
}
// SetGenre sets the genre property value. The genre of this audio file.
func (m *Audio) SetGenre(value *string)() {
    if m != nil {
        m.genre = value
    }
}
// SetHasDrm sets the hasDrm property value. Indicates if the file is protected with digital rights management.
func (m *Audio) SetHasDrm(value *bool)() {
    if m != nil {
        m.hasDrm = value
    }
}
// SetIsVariableBitrate sets the isVariableBitrate property value. Indicates if the file is encoded with a variable bitrate.
func (m *Audio) SetIsVariableBitrate(value *bool)() {
    if m != nil {
        m.isVariableBitrate = value
    }
}
// SetTitle sets the title property value. The title of the audio file.
func (m *Audio) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
// SetTrack sets the track property value. The number of the track on the original disc for this audio file.
func (m *Audio) SetTrack(value *int32)() {
    if m != nil {
        m.track = value
    }
}
// SetTrackCount sets the trackCount property value. The total number of tracks on the original disc for this audio file.
func (m *Audio) SetTrackCount(value *int32)() {
    if m != nil {
        m.trackCount = value
    }
}
// SetYear sets the year property value. The year the audio file was recorded.
func (m *Audio) SetYear(value *int32)() {
    if m != nil {
        m.year = value
    }
}
