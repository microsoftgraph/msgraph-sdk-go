package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookTablesCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookTablesCountResponse struct {
    ItemItemsItemWorkbookTablesCountGetResponse
}
// NewItemItemsItemWorkbookTablesCountResponse instantiates a new ItemItemsItemWorkbookTablesCountResponse and sets the default values.
func NewItemItemsItemWorkbookTablesCountResponse()(*ItemItemsItemWorkbookTablesCountResponse) {
    m := &ItemItemsItemWorkbookTablesCountResponse{
        ItemItemsItemWorkbookTablesCountGetResponse: *NewItemItemsItemWorkbookTablesCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookTablesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookTablesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesCountResponse(), nil
}
// ItemItemsItemWorkbookTablesCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookTablesCountResponseable interface {
    ItemItemsItemWorkbookTablesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
