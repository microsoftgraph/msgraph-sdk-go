package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse()(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse{
        ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse(), nil
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
