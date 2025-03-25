package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemTablesCountgetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemTablesCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemTablesCountgetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemTablesCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesCountResponse()(*ItemItemsItemWorkbookWorksheetsItemTablesCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesCountResponse{
        ItemItemsItemWorkbookWorksheetsItemTablesCountgetResponse: *NewItemItemsItemWorkbookWorksheetsItemTablesCountgetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTablesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemTablesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemTablesCountgetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemTablesCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemTablesCountgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
