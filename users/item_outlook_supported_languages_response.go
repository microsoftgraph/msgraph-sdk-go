package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOutlookSupportedLanguagesResponse 
// Deprecated: This class is obsolete. Use supportedLanguagesGetResponse instead.
type ItemOutlookSupportedLanguagesResponse struct {
    ItemOutlookSupportedLanguagesGetResponse
}
// NewItemOutlookSupportedLanguagesResponse instantiates a new ItemOutlookSupportedLanguagesResponse and sets the default values.
func NewItemOutlookSupportedLanguagesResponse()(*ItemOutlookSupportedLanguagesResponse) {
    m := &ItemOutlookSupportedLanguagesResponse{
        ItemOutlookSupportedLanguagesGetResponse: *NewItemOutlookSupportedLanguagesGetResponse(),
    }
    return m
}
// CreateItemOutlookSupportedLanguagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOutlookSupportedLanguagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutlookSupportedLanguagesResponse(), nil
}
// ItemOutlookSupportedLanguagesResponseable 
// Deprecated: This class is obsolete. Use supportedLanguagesGetResponse instead.
type ItemOutlookSupportedLanguagesResponseable interface {
    ItemOutlookSupportedLanguagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
