package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppEBookable 
type IosVppEBookable interface {
    ManagedEBookable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleId()(*string)
    GetGenres()([]string)
    GetLanguage()(*string)
    GetSeller()(*string)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    GetVppOrganizationName()(*string)
    GetVppTokenId()(*UUID)
    SetAppleId(value *string)()
    SetGenres(value []string)()
    SetLanguage(value *string)()
    SetSeller(value *string)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
    SetVppOrganizationName(value *string)()
    SetVppTokenId(value *UUID)()
}
