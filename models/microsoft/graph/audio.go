package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new audio and sets the default values.
func NewAudio()(*Audio) {
    m := &Audio{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Audio) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the album property value. The title of the album for this audio file.
func (m *Audio) GetAlbum()(*string) {
    if m == nil {
        return nil
    } else {
        return m.album
    }
}
// Gets the albumArtist property value. The artist named on the album for the audio file.
func (m *Audio) GetAlbumArtist()(*string) {
    if m == nil {
        return nil
    } else {
        return m.albumArtist
    }
}
// Gets the artist property value. The performing artist for the audio file.
func (m *Audio) GetArtist()(*string) {
    if m == nil {
        return nil
    } else {
        return m.artist
    }
}
// Gets the bitrate property value. Bitrate expressed in kbps.
func (m *Audio) GetBitrate()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.bitrate
    }
}
// Gets the composers property value. The name of the composer of the audio file.
func (m *Audio) GetComposers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.composers
    }
}
// Gets the copyright property value. Copyright information for the audio file.
func (m *Audio) GetCopyright()(*string) {
    if m == nil {
        return nil
    } else {
        return m.copyright
    }
}
// Gets the disc property value. The number of the disc this audio file came from.
func (m *Audio) GetDisc()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.disc
    }
}
// Gets the discCount property value. The total number of discs in this album.
func (m *Audio) GetDiscCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.discCount
    }
}
// Gets the duration property value. Duration of the audio file, expressed in milliseconds
func (m *Audio) GetDuration()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// Gets the genre property value. The genre of this audio file.
func (m *Audio) GetGenre()(*string) {
    if m == nil {
        return nil
    } else {
        return m.genre
    }
}
// Gets the hasDrm property value. Indicates if the file is protected with digital rights management.
func (m *Audio) GetHasDrm()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasDrm
    }
}
// Gets the isVariableBitrate property value. Indicates if the file is encoded with a variable bitrate.
func (m *Audio) GetIsVariableBitrate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVariableBitrate
    }
}
// Gets the title property value. The title of the audio file.
func (m *Audio) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the track property value. The number of the track on the original disc for this audio file.
func (m *Audio) GetTrack()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.track
    }
}
// Gets the trackCount property value. The total number of tracks on the original disc for this audio file.
func (m *Audio) GetTrackCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.trackCount
    }
}
// Gets the year property value. The year the audio file was recorded.
func (m *Audio) GetYear()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.year
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Audio) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the album property value. The title of the album for this audio file.
// Parameters:
//  - value : Value to set for the album property.
func (m *Audio) SetAlbum(value *string)() {
    m.album = value
}
// Sets the albumArtist property value. The artist named on the album for the audio file.
// Parameters:
//  - value : Value to set for the albumArtist property.
func (m *Audio) SetAlbumArtist(value *string)() {
    m.albumArtist = value
}
// Sets the artist property value. The performing artist for the audio file.
// Parameters:
//  - value : Value to set for the artist property.
func (m *Audio) SetArtist(value *string)() {
    m.artist = value
}
// Sets the bitrate property value. Bitrate expressed in kbps.
// Parameters:
//  - value : Value to set for the bitrate property.
func (m *Audio) SetBitrate(value *int64)() {
    m.bitrate = value
}
// Sets the composers property value. The name of the composer of the audio file.
// Parameters:
//  - value : Value to set for the composers property.
func (m *Audio) SetComposers(value *string)() {
    m.composers = value
}
// Sets the copyright property value. Copyright information for the audio file.
// Parameters:
//  - value : Value to set for the copyright property.
func (m *Audio) SetCopyright(value *string)() {
    m.copyright = value
}
// Sets the disc property value. The number of the disc this audio file came from.
// Parameters:
//  - value : Value to set for the disc property.
func (m *Audio) SetDisc(value *int32)() {
    m.disc = value
}
// Sets the discCount property value. The total number of discs in this album.
// Parameters:
//  - value : Value to set for the discCount property.
func (m *Audio) SetDiscCount(value *int32)() {
    m.discCount = value
}
// Sets the duration property value. Duration of the audio file, expressed in milliseconds
// Parameters:
//  - value : Value to set for the duration property.
func (m *Audio) SetDuration(value *int64)() {
    m.duration = value
}
// Sets the genre property value. The genre of this audio file.
// Parameters:
//  - value : Value to set for the genre property.
func (m *Audio) SetGenre(value *string)() {
    m.genre = value
}
// Sets the hasDrm property value. Indicates if the file is protected with digital rights management.
// Parameters:
//  - value : Value to set for the hasDrm property.
func (m *Audio) SetHasDrm(value *bool)() {
    m.hasDrm = value
}
// Sets the isVariableBitrate property value. Indicates if the file is encoded with a variable bitrate.
// Parameters:
//  - value : Value to set for the isVariableBitrate property.
func (m *Audio) SetIsVariableBitrate(value *bool)() {
    m.isVariableBitrate = value
}
// Sets the title property value. The title of the audio file.
// Parameters:
//  - value : Value to set for the title property.
func (m *Audio) SetTitle(value *string)() {
    m.title = value
}
// Sets the track property value. The number of the track on the original disc for this audio file.
// Parameters:
//  - value : Value to set for the track property.
func (m *Audio) SetTrack(value *int32)() {
    m.track = value
}
// Sets the trackCount property value. The total number of tracks on the original disc for this audio file.
// Parameters:
//  - value : Value to set for the trackCount property.
func (m *Audio) SetTrackCount(value *int32)() {
    m.trackCount = value
}
// Sets the year property value. The year the audio file was recorded.
// Parameters:
//  - value : Value to set for the year property.
func (m *Audio) SetYear(value *int32)() {
    m.year = value
}
