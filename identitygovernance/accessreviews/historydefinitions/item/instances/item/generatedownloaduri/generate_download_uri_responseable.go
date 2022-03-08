package generatedownloaduri

import (
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// GenerateDownloadUriResponseable 
type GenerateDownloadUriResponseable interface {
    AdditionalDataHolder
    Parsable
    GetAccessReviewHistoryInstance()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable)
    SetAccessReviewHistoryInstance(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable)()
}
