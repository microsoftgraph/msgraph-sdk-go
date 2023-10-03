package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse(), nil
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseable 
// Deprecated: This class is obsolete. Use countGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
