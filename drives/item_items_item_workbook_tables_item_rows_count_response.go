package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookTablesItemRowsCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookTablesItemRowsCountResponse struct {
    ItemItemsItemWorkbookTablesItemRowsCountGetResponse
}
// NewItemItemsItemWorkbookTablesItemRowsCountResponse instantiates a new ItemItemsItemWorkbookTablesItemRowsCountResponse and sets the default values.
func NewItemItemsItemWorkbookTablesItemRowsCountResponse()(*ItemItemsItemWorkbookTablesItemRowsCountResponse) {
    m := &ItemItemsItemWorkbookTablesItemRowsCountResponse{
        ItemItemsItemWorkbookTablesItemRowsCountGetResponse: *NewItemItemsItemWorkbookTablesItemRowsCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookTablesItemRowsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookTablesItemRowsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesItemRowsCountResponse(), nil
}
// ItemItemsItemWorkbookTablesItemRowsCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookTablesItemRowsCountResponseable interface {
    ItemItemsItemWorkbookTablesItemRowsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
