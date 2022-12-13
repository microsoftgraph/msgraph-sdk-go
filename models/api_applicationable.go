package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiApplicationable 
type ApiApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptMappedClaims()(*bool)
    GetKnownClientApplications()([]UUID)
    GetOauth2PermissionScopes()([]PermissionScopeable)
    GetOdataType()(*string)
    GetPreAuthorizedApplications()([]PreAuthorizedApplicationable)
    GetRequestedAccessTokenVersion()(*int32)
    SetAcceptMappedClaims(value *bool)()
    SetKnownClientApplications(value []UUID)()
    SetOauth2PermissionScopes(value []PermissionScopeable)()
    SetOdataType(value *string)()
    SetPreAuthorizedApplications(value []PreAuthorizedApplicationable)()
    SetRequestedAccessTokenVersion(value *int32)()
}
