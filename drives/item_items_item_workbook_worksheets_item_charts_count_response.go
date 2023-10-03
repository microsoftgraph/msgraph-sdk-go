package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemChartsCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsCountResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsCountResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsCountResponse(), nil
}
// ItemItemsItemWorkbookWorksheetsItemChartsCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
