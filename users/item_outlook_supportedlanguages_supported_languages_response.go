package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemOutlookSupportedlanguagesSupportedLanguagesGetResponseable instead.
type ItemOutlookSupportedlanguagesSupportedLanguagesResponse struct {
    ItemOutlookSupportedlanguagesSupportedLanguagesGetResponse
}
// NewItemOutlookSupportedlanguagesSupportedLanguagesResponse instantiates a new ItemOutlookSupportedlanguagesSupportedLanguagesResponse and sets the default values.
func NewItemOutlookSupportedlanguagesSupportedLanguagesResponse()(*ItemOutlookSupportedlanguagesSupportedLanguagesResponse) {
    m := &ItemOutlookSupportedlanguagesSupportedLanguagesResponse{
        ItemOutlookSupportedlanguagesSupportedLanguagesGetResponse: *NewItemOutlookSupportedlanguagesSupportedLanguagesGetResponse(),
    }
    return m
}
// CreateItemOutlookSupportedlanguagesSupportedLanguagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOutlookSupportedlanguagesSupportedLanguagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookSupportedlanguagesSupportedLanguagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemOutlookSupportedlanguagesSupportedLanguagesGetResponseable instead.
type ItemOutlookSupportedlanguagesSupportedLanguagesResponseable interface {
    ItemOutlookSupportedlanguagesSupportedLanguagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
