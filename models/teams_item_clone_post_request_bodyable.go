package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsItemClonePostRequestBodyable 
type TeamsItemClonePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassification()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetMailNickname()(*string)
    GetPartsToClone()(*ClonableTeamParts)
    GetVisibility()(*TeamVisibilityType)
    SetClassification(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetMailNickname(value *string)()
    SetPartsToClone(value *ClonableTeamParts)()
    SetVisibility(value *TeamVisibilityType)()
}
