package preview

import (
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// PreviewResponseable 
type PreviewResponseable interface {
    AdditionalDataHolder
    Parsable
    GetOnenotePagePreview()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePagePreviewable)
    SetOnenotePagePreview(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePagePreviewable)()
}
