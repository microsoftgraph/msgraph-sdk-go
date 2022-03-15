package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Audioable 
type Audioable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAlbum()(*string)
    GetAlbumArtist()(*string)
    GetArtist()(*string)
    GetBitrate()(*int32)
    GetComposers()(*string)
    GetCopyright()(*string)
    GetDisc()(*int32)
    GetDiscCount()(*int32)
    GetDuration()(*int32)
    GetGenre()(*string)
    GetHasDrm()(*bool)
    GetIsVariableBitrate()(*bool)
    GetTitle()(*string)
    GetTrack()(*int32)
    GetTrackCount()(*int32)
    GetYear()(*int32)
    SetAlbum(value *string)()
    SetAlbumArtist(value *string)()
    SetArtist(value *string)()
    SetBitrate(value *int32)()
    SetComposers(value *string)()
    SetCopyright(value *string)()
    SetDisc(value *int32)()
    SetDiscCount(value *int32)()
    SetDuration(value *int32)()
    SetGenre(value *string)()
    SetHasDrm(value *bool)()
    SetIsVariableBitrate(value *bool)()
    SetTitle(value *string)()
    SetTrack(value *int32)()
    SetTrackCount(value *int32)()
    SetYear(value *int32)()
}
