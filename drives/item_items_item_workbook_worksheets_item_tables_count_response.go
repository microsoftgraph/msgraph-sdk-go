package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemTablesCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemTablesCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemTablesCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemTablesCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesCountResponse()(*ItemItemsItemWorkbookWorksheetsItemTablesCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesCountResponse{
        ItemItemsItemWorkbookWorksheetsItemTablesCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemTablesCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTablesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemTablesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesCountResponse(), nil
}
// ItemItemsItemWorkbookWorksheetsItemTablesCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemTablesCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemTablesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
