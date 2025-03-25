package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookTablesItemRowsCountgetResponseable instead.
type ItemItemsItemWorkbookTablesItemRowsCountResponse struct {
    ItemItemsItemWorkbookTablesItemRowsCountgetResponse
}
// NewItemItemsItemWorkbookTablesItemRowsCountResponse instantiates a new ItemItemsItemWorkbookTablesItemRowsCountResponse and sets the default values.
func NewItemItemsItemWorkbookTablesItemRowsCountResponse()(*ItemItemsItemWorkbookTablesItemRowsCountResponse) {
    m := &ItemItemsItemWorkbookTablesItemRowsCountResponse{
        ItemItemsItemWorkbookTablesItemRowsCountgetResponse: *NewItemItemsItemWorkbookTablesItemRowsCountgetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookTablesItemRowsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookTablesItemRowsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesItemRowsCountResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookTablesItemRowsCountgetResponseable instead.
type ItemItemsItemWorkbookTablesItemRowsCountResponseable interface {
    ItemItemsItemWorkbookTablesItemRowsCountgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
