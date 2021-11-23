package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0afd60ce3dbcbcfa575a272a3f01f1816238e962adc50d27383cf38add45546d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/versions"
    i0b91669df33041122a3239128150d9881d1c5148ce1110db82dc471ac8f03380 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/deltawithtoken"
    i19c8a435e35f3a91440869da89a429da57fcd7561f3e38373a0760184ece7ee5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/listitem"
    i1b79ba93e8f9ce8e91cbb68bbe93059197f68509d9f1ec1e4951be971507e670 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/validatepermission"
    i252fd7c35689b2a02e3f3f6a369c7b92d60eac212982bcc0ca288ad8c6180f70 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/restore"
    i30c63e24788689962791595cb53052ff09bcee5b947bdcad4bdf76549580e375 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/copy"
    i336cb2040266afc7f34389ad32e9cde75866748d18dc0af981052b23597b2f18 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/unfollow"
    i4a1b5cb316219ccd2b0d5274ee44a684951623b97be4ac059d58ab79180da176 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/invite"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4aee501c094d82ef0b97fddf8ca4acda070ff6458057d83027e62f87832e701d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/searchwithq"
    i52b5cab9e6301cea985fb5afcd313f4e023c24e7354864c186b0932dc71b3c10 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/checkout"
    i52d55698f37207ae8cfc4d0581661790447acf0fa06972a6c6079678d1716c3e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/createlink"
    i614fa74ffac53752a7c403ccba0bf7595d1c3d753c9758a84c28c3708801511a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook"
    i63af7806e0ae5f4175097e470d929fb4adcc391a4308468068a7e4d6e3198ddd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/analytics"
    i727bc3da6e9c7bcaefc7e93c6537625263c40b0040ccfb4ea7fb98ee33606468 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/checkin"
    i8322282da8aba85dea7f464f9817c042833a1e61796c1743fab901ceb2b9e0bf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/createuploadsession"
    ia597af73a896059b6e28f8af56e903ffd69180dae0086fb241c206ac9158e73f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/children"
    ia889c3d885eb9a9e1be826c747ce32ac0dc738051dfafa2515bcb74296576e50 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/thumbnails"
    iadfb7f5925f25f6d21cbe4775107ce213058b65eca0dc1737beb1b2a1ea01947 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/delta"
    ib121a556cf3e26404fc08358679728faa63ad120d407307b65b8fe2f5ae147c4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ibf7cf0b36ab7fa681127712ba38052a6b0b76bd276af6398c8fcfe3b646a574d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/preview"
    ic6fb998d47a797063386fd6d0a833ead6f93b1706e3af2a61f08cf17f3b12b9c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/subscriptions"
    icf1df6087ee055e0d82133b727f46f8e65b392162d6927ac19d0404cfca0cc06 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/content"
    id47f3cd96ac62060cd4706158e84ca804380b00ffa0deb0c987cbd9861fdc941 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/follow"
    id98738fe99e7cee62a1c2f7c738811ddbbda251bc9a20b42fbba36e7cd90b957 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/permissions"
    ifbb48734282372c41bb9bbc447d9984c550ba21aaa310f937681aa66356ddde6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/getactivitiesbyinterval"
    i1045b039fef62d5ef72566ac302564d7a39c35320f7d3ed85a713473dd79ecc3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/children/item"
    i1e600c52f111ed567ffcb9cd01fe8720b9c93c7af81196af4ce1b46a2dfac925 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/thumbnails/item"
    i273f6427d4e6aa2bb741a8d68e377c20d8b7a98bf0a4c581769ace831a969301 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/versions/item"
    ia90ca1da21635979709e703dc27c0bd00834adebaad1df486f890c40213d08d4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/subscriptions/item"
    ie61cc009266b855785fa76d2bad46bc9dcf5abfdd3228403f2d8507b2f2e9e83 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/permissions/item"
)

// DriveItemRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}
type DriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemRequestBuilderDeleteOptions options for Delete
type DriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetOptions options for Get
type DriveItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetQueryParameters get entity from workbooks by key
type DriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// DriveItemRequestBuilderPatchOptions options for Patch
type DriveItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveItemRequestBuilder) Analytics()(*i63af7806e0ae5f4175097e470d929fb4adcc391a4308468068a7e4d6e3198ddd.AnalyticsRequestBuilder) {
    return i63af7806e0ae5f4175097e470d929fb4adcc391a4308468068a7e4d6e3198ddd.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Checkin()(*i727bc3da6e9c7bcaefc7e93c6537625263c40b0040ccfb4ea7fb98ee33606468.CheckinRequestBuilder) {
    return i727bc3da6e9c7bcaefc7e93c6537625263c40b0040ccfb4ea7fb98ee33606468.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Checkout()(*i52b5cab9e6301cea985fb5afcd313f4e023c24e7354864c186b0932dc71b3c10.CheckoutRequestBuilder) {
    return i52b5cab9e6301cea985fb5afcd313f4e023c24e7354864c186b0932dc71b3c10.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Children()(*ia597af73a896059b6e28f8af56e903ffd69180dae0086fb241c206ac9158e73f.ChildrenRequestBuilder) {
    return ia597af73a896059b6e28f8af56e903ffd69180dae0086fb241c206ac9158e73f.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.children.item collection
func (m *DriveItemRequestBuilder) ChildrenById(id string)(*i1045b039fef62d5ef72566ac302564d7a39c35320f7d3ed85a713473dd79ecc3.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i1045b039fef62d5ef72566ac302564d7a39c35320f7d3ed85a713473dd79ecc3.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DriveItemRequestBuilder) Content()(*icf1df6087ee055e0d82133b727f46f8e65b392162d6927ac19d0404cfca0cc06.ContentRequestBuilder) {
    return icf1df6087ee055e0d82133b727f46f8e65b392162d6927ac19d0404cfca0cc06.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Copy()(*i30c63e24788689962791595cb53052ff09bcee5b947bdcad4bdf76549580e375.CopyRequestBuilder) {
    return i30c63e24788689962791595cb53052ff09bcee5b947bdcad4bdf76549580e375.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from workbooks
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from workbooks by key
func (m *DriveItemRequestBuilder) CreateGetRequestInformation(options *DriveItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveItemRequestBuilder) CreateLink()(*i52d55698f37207ae8cfc4d0581661790447acf0fa06972a6c6079678d1716c3e.CreateLinkRequestBuilder) {
    return i52d55698f37207ae8cfc4d0581661790447acf0fa06972a6c6079678d1716c3e.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update entity in workbooks
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveItemRequestBuilder) CreateUploadSession()(*i8322282da8aba85dea7f464f9817c042833a1e61796c1743fab901ceb2b9e0bf.CreateUploadSessionRequestBuilder) {
    return i8322282da8aba85dea7f464f9817c042833a1e61796c1743fab901ceb2b9e0bf.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete entity from workbooks
func (m *DriveItemRequestBuilder) Delete(options *DriveItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Delta builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.delta()
func (m *DriveItemRequestBuilder) Delta()(*iadfb7f5925f25f6d21cbe4775107ce213058b65eca0dc1737beb1b2a1ea01947.DeltaRequestBuilder) {
    return iadfb7f5925f25f6d21cbe4775107ce213058b65eca0dc1737beb1b2a1ea01947.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.delta(token='{token}')
func (m *DriveItemRequestBuilder) DeltaWithToken(token *string)(*i0b91669df33041122a3239128150d9881d1c5148ce1110db82dc471ac8f03380.DeltaWithTokenRequestBuilder) {
    return i0b91669df33041122a3239128150d9881d1c5148ce1110db82dc471ac8f03380.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
func (m *DriveItemRequestBuilder) Follow()(*id47f3cd96ac62060cd4706158e84ca804380b00ffa0deb0c987cbd9861fdc941.FollowRequestBuilder) {
    return id47f3cd96ac62060cd4706158e84ca804380b00ffa0deb0c987cbd9861fdc941.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from workbooks by key
func (m *DriveItemRequestBuilder) Get(options *DriveItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDriveItem() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItem), nil
}
// GetActivitiesByInterval builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.getActivitiesByInterval()
func (m *DriveItemRequestBuilder) GetActivitiesByInterval()(*ifbb48734282372c41bb9bbc447d9984c550ba21aaa310f937681aa66356ddde6.GetActivitiesByIntervalRequestBuilder) {
    return ifbb48734282372c41bb9bbc447d9984c550ba21aaa310f937681aa66356ddde6.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
func (m *DriveItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*ib121a556cf3e26404fc08358679728faa63ad120d407307b65b8fe2f5ae147c4.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ib121a556cf3e26404fc08358679728faa63ad120d407307b65b8fe2f5ae147c4.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
func (m *DriveItemRequestBuilder) Invite()(*i4a1b5cb316219ccd2b0d5274ee44a684951623b97be4ac059d58ab79180da176.InviteRequestBuilder) {
    return i4a1b5cb316219ccd2b0d5274ee44a684951623b97be4ac059d58ab79180da176.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) ListItem()(*i19c8a435e35f3a91440869da89a429da57fcd7561f3e38373a0760184ece7ee5.ListItemRequestBuilder) {
    return i19c8a435e35f3a91440869da89a429da57fcd7561f3e38373a0760184ece7ee5.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in workbooks
func (m *DriveItemRequestBuilder) Patch(options *DriveItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveItemRequestBuilder) Permissions()(*id98738fe99e7cee62a1c2f7c738811ddbbda251bc9a20b42fbba36e7cd90b957.PermissionsRequestBuilder) {
    return id98738fe99e7cee62a1c2f7c738811ddbbda251bc9a20b42fbba36e7cd90b957.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.permissions.item collection
func (m *DriveItemRequestBuilder) PermissionsById(id string)(*ie61cc009266b855785fa76d2bad46bc9dcf5abfdd3228403f2d8507b2f2e9e83.PermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ie61cc009266b855785fa76d2bad46bc9dcf5abfdd3228403f2d8507b2f2e9e83.NewPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Preview()(*ibf7cf0b36ab7fa681127712ba38052a6b0b76bd276af6398c8fcfe3b646a574d.PreviewRequestBuilder) {
    return ibf7cf0b36ab7fa681127712ba38052a6b0b76bd276af6398c8fcfe3b646a574d.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Restore()(*i252fd7c35689b2a02e3f3f6a369c7b92d60eac212982bcc0ca288ad8c6180f70.RestoreRequestBuilder) {
    return i252fd7c35689b2a02e3f3f6a369c7b92d60eac212982bcc0ca288ad8c6180f70.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.search(q='{q}')
func (m *DriveItemRequestBuilder) SearchWithQ(q *string)(*i4aee501c094d82ef0b97fddf8ca4acda070ff6458057d83027e62f87832e701d.SearchWithQRequestBuilder) {
    return i4aee501c094d82ef0b97fddf8ca4acda070ff6458057d83027e62f87832e701d.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
func (m *DriveItemRequestBuilder) Subscriptions()(*ic6fb998d47a797063386fd6d0a833ead6f93b1706e3af2a61f08cf17f3b12b9c.SubscriptionsRequestBuilder) {
    return ic6fb998d47a797063386fd6d0a833ead6f93b1706e3af2a61f08cf17f3b12b9c.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.subscriptions.item collection
func (m *DriveItemRequestBuilder) SubscriptionsById(id string)(*ia90ca1da21635979709e703dc27c0bd00834adebaad1df486f890c40213d08d4.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ia90ca1da21635979709e703dc27c0bd00834adebaad1df486f890c40213d08d4.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Thumbnails()(*ia889c3d885eb9a9e1be826c747ce32ac0dc738051dfafa2515bcb74296576e50.ThumbnailsRequestBuilder) {
    return ia889c3d885eb9a9e1be826c747ce32ac0dc738051dfafa2515bcb74296576e50.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.thumbnails.item collection
func (m *DriveItemRequestBuilder) ThumbnailsById(id string)(*i1e600c52f111ed567ffcb9cd01fe8720b9c93c7af81196af4ce1b46a2dfac925.ThumbnailSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i1e600c52f111ed567ffcb9cd01fe8720b9c93c7af81196af4ce1b46a2dfac925.NewThumbnailSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Unfollow()(*i336cb2040266afc7f34389ad32e9cde75866748d18dc0af981052b23597b2f18.UnfollowRequestBuilder) {
    return i336cb2040266afc7f34389ad32e9cde75866748d18dc0af981052b23597b2f18.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) ValidatePermission()(*i1b79ba93e8f9ce8e91cbb68bbe93059197f68509d9f1ec1e4951be971507e670.ValidatePermissionRequestBuilder) {
    return i1b79ba93e8f9ce8e91cbb68bbe93059197f68509d9f1ec1e4951be971507e670.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Versions()(*i0afd60ce3dbcbcfa575a272a3f01f1816238e962adc50d27383cf38add45546d.VersionsRequestBuilder) {
    return i0afd60ce3dbcbcfa575a272a3f01f1816238e962adc50d27383cf38add45546d.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.versions.item collection
func (m *DriveItemRequestBuilder) VersionsById(id string)(*i273f6427d4e6aa2bb741a8d68e377c20d8b7a98bf0a4c581769ace831a969301.DriveItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i273f6427d4e6aa2bb741a8d68e377c20d8b7a98bf0a4c581769ace831a969301.NewDriveItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Workbook()(*i614fa74ffac53752a7c403ccba0bf7595d1c3d753c9758a84c28c3708801511a.WorkbookRequestBuilder) {
    return i614fa74ffac53752a7c403ccba0bf7595d1c3d753c9758a84c28c3708801511a.NewWorkbookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
