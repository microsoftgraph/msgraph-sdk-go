package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0afd60ce3dbcbcfa575a272a3f01f1816238e962adc50d27383cf38add45546d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/versions"
    i0b91669df33041122a3239128150d9881d1c5148ce1110db82dc471ac8f03380 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/deltawithtoken"
    i19c8a435e35f3a91440869da89a429da57fcd7561f3e38373a0760184ece7ee5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/listitem"
    i1b79ba93e8f9ce8e91cbb68bbe93059197f68509d9f1ec1e4951be971507e670 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/validatepermission"
    i252fd7c35689b2a02e3f3f6a369c7b92d60eac212982bcc0ca288ad8c6180f70 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/restore"
    i30c63e24788689962791595cb53052ff09bcee5b947bdcad4bdf76549580e375 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/copy"
    i336cb2040266afc7f34389ad32e9cde75866748d18dc0af981052b23597b2f18 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/unfollow"
    i4a1b5cb316219ccd2b0d5274ee44a684951623b97be4ac059d58ab79180da176 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/invite"
    i4aee501c094d82ef0b97fddf8ca4acda070ff6458057d83027e62f87832e701d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/searchwithq"
    i52b5cab9e6301cea985fb5afcd313f4e023c24e7354864c186b0932dc71b3c10 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/checkout"
    i52d55698f37207ae8cfc4d0581661790447acf0fa06972a6c6079678d1716c3e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/createlink"
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

// DriveItemItemRequestBuilder provides operations to manage the collection of driveItem entities.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemItemRequestBuilderDeleteOptions options for Delete
type DriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DriveItemItemRequestBuilderGetOptions options for Get
type DriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DriveItemItemRequestBuilderGetQueryParameters get entity from workbooks by key
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemItemRequestBuilderPatchOptions options for Patch
type DriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*i63af7806e0ae5f4175097e470d929fb4adcc391a4308468068a7e4d6e3198ddd.AnalyticsRequestBuilder) {
    return i63af7806e0ae5f4175097e470d929fb4adcc391a4308468068a7e4d6e3198ddd.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin the checkin property
func (m *DriveItemItemRequestBuilder) Checkin()(*i727bc3da6e9c7bcaefc7e93c6537625263c40b0040ccfb4ea7fb98ee33606468.CheckinRequestBuilder) {
    return i727bc3da6e9c7bcaefc7e93c6537625263c40b0040ccfb4ea7fb98ee33606468.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout the checkout property
func (m *DriveItemItemRequestBuilder) Checkout()(*i52b5cab9e6301cea985fb5afcd313f4e023c24e7354864c186b0932dc71b3c10.CheckoutRequestBuilder) {
    return i52b5cab9e6301cea985fb5afcd313f4e023c24e7354864c186b0932dc71b3c10.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*ia597af73a896059b6e28f8af56e903ffd69180dae0086fb241c206ac9158e73f.ChildrenRequestBuilder) {
    return ia597af73a896059b6e28f8af56e903ffd69180dae0086fb241c206ac9158e73f.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i1045b039fef62d5ef72566ac302564d7a39c35320f7d3ed85a713473dd79ecc3.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i1045b039fef62d5ef72566ac302564d7a39c35320f7d3ed85a713473dd79ecc3.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *DriveItemItemRequestBuilder) Content()(*icf1df6087ee055e0d82133b727f46f8e65b392162d6927ac19d0404cfca0cc06.ContentRequestBuilder) {
    return icf1df6087ee055e0d82133b727f46f8e65b392162d6927ac19d0404cfca0cc06.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *DriveItemItemRequestBuilder) Copy()(*i30c63e24788689962791595cb53052ff09bcee5b947bdcad4bdf76549580e375.CopyRequestBuilder) {
    return i30c63e24788689962791595cb53052ff09bcee5b947bdcad4bdf76549580e375.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from workbooks
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from workbooks by key
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation(options *DriveItemItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateLink the createLink property
func (m *DriveItemItemRequestBuilder) CreateLink()(*i52d55698f37207ae8cfc4d0581661790447acf0fa06972a6c6079678d1716c3e.CreateLinkRequestBuilder) {
    return i52d55698f37207ae8cfc4d0581661790447acf0fa06972a6c6079678d1716c3e.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update entity in workbooks
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateUploadSession the createUploadSession property
func (m *DriveItemItemRequestBuilder) CreateUploadSession()(*i8322282da8aba85dea7f464f9817c042833a1e61796c1743fab901ceb2b9e0bf.CreateUploadSessionRequestBuilder) {
    return i8322282da8aba85dea7f464f9817c042833a1e61796c1743fab901ceb2b9e0bf.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete entity from workbooks
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Delta provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) Delta()(*iadfb7f5925f25f6d21cbe4775107ce213058b65eca0dc1737beb1b2a1ea01947.DeltaRequestBuilder) {
    return iadfb7f5925f25f6d21cbe4775107ce213058b65eca0dc1737beb1b2a1ea01947.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DriveItemItemRequestBuilder) DeltaWithToken(token *string)(*i0b91669df33041122a3239128150d9881d1c5148ce1110db82dc471ac8f03380.DeltaWithTokenRequestBuilder) {
    return i0b91669df33041122a3239128150d9881d1c5148ce1110db82dc471ac8f03380.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// Follow the follow property
func (m *DriveItemItemRequestBuilder) Follow()(*id47f3cd96ac62060cd4706158e84ca804380b00ffa0deb0c987cbd9861fdc941.FollowRequestBuilder) {
    return id47f3cd96ac62060cd4706158e84ca804380b00ffa0deb0c987cbd9861fdc941.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from workbooks by key
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// GetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
func (m *DriveItemItemRequestBuilder) GetActivitiesByInterval()(*ifbb48734282372c41bb9bbc447d9984c550ba21aaa310f937681aa66356ddde6.GetActivitiesByIntervalRequestBuilder) {
    return ifbb48734282372c41bb9bbc447d9984c550ba21aaa310f937681aa66356ddde6.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *DriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ib121a556cf3e26404fc08358679728faa63ad120d407307b65b8fe2f5ae147c4.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ib121a556cf3e26404fc08358679728faa63ad120d407307b65b8fe2f5ae147c4.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite the invite property
func (m *DriveItemItemRequestBuilder) Invite()(*i4a1b5cb316219ccd2b0d5274ee44a684951623b97be4ac059d58ab79180da176.InviteRequestBuilder) {
    return i4a1b5cb316219ccd2b0d5274ee44a684951623b97be4ac059d58ab79180da176.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *DriveItemItemRequestBuilder) ListItem()(*i19c8a435e35f3a91440869da89a429da57fcd7561f3e38373a0760184ece7ee5.ListItemRequestBuilder) {
    return i19c8a435e35f3a91440869da89a429da57fcd7561f3e38373a0760184ece7ee5.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in workbooks
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Permissions the permissions property
func (m *DriveItemItemRequestBuilder) Permissions()(*id98738fe99e7cee62a1c2f7c738811ddbbda251bc9a20b42fbba36e7cd90b957.PermissionsRequestBuilder) {
    return id98738fe99e7cee62a1c2f7c738811ddbbda251bc9a20b42fbba36e7cd90b957.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ie61cc009266b855785fa76d2bad46bc9dcf5abfdd3228403f2d8507b2f2e9e83.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ie61cc009266b855785fa76d2bad46bc9dcf5abfdd3228403f2d8507b2f2e9e83.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview the preview property
func (m *DriveItemItemRequestBuilder) Preview()(*ibf7cf0b36ab7fa681127712ba38052a6b0b76bd276af6398c8fcfe3b646a574d.PreviewRequestBuilder) {
    return ibf7cf0b36ab7fa681127712ba38052a6b0b76bd276af6398c8fcfe3b646a574d.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *DriveItemItemRequestBuilder) Restore()(*i252fd7c35689b2a02e3f3f6a369c7b92d60eac212982bcc0ca288ad8c6180f70.RestoreRequestBuilder) {
    return i252fd7c35689b2a02e3f3f6a369c7b92d60eac212982bcc0ca288ad8c6180f70.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemItemRequestBuilder) SearchWithQ(q *string)(*i4aee501c094d82ef0b97fddf8ca4acda070ff6458057d83027e62f87832e701d.SearchWithQRequestBuilder) {
    return i4aee501c094d82ef0b97fddf8ca4acda070ff6458057d83027e62f87832e701d.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*ic6fb998d47a797063386fd6d0a833ead6f93b1706e3af2a61f08cf17f3b12b9c.SubscriptionsRequestBuilder) {
    return ic6fb998d47a797063386fd6d0a833ead6f93b1706e3af2a61f08cf17f3b12b9c.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*ia90ca1da21635979709e703dc27c0bd00834adebaad1df486f890c40213d08d4.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ia90ca1da21635979709e703dc27c0bd00834adebaad1df486f890c40213d08d4.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ia889c3d885eb9a9e1be826c747ce32ac0dc738051dfafa2515bcb74296576e50.ThumbnailsRequestBuilder) {
    return ia889c3d885eb9a9e1be826c747ce32ac0dc738051dfafa2515bcb74296576e50.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i1e600c52f111ed567ffcb9cd01fe8720b9c93c7af81196af4ce1b46a2dfac925.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i1e600c52f111ed567ffcb9cd01fe8720b9c93c7af81196af4ce1b46a2dfac925.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow the unfollow property
func (m *DriveItemItemRequestBuilder) Unfollow()(*i336cb2040266afc7f34389ad32e9cde75866748d18dc0af981052b23597b2f18.UnfollowRequestBuilder) {
    return i336cb2040266afc7f34389ad32e9cde75866748d18dc0af981052b23597b2f18.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission the validatePermission property
func (m *DriveItemItemRequestBuilder) ValidatePermission()(*i1b79ba93e8f9ce8e91cbb68bbe93059197f68509d9f1ec1e4951be971507e670.ValidatePermissionRequestBuilder) {
    return i1b79ba93e8f9ce8e91cbb68bbe93059197f68509d9f1ec1e4951be971507e670.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i0afd60ce3dbcbcfa575a272a3f01f1816238e962adc50d27383cf38add45546d.VersionsRequestBuilder) {
    return i0afd60ce3dbcbcfa575a272a3f01f1816238e962adc50d27383cf38add45546d.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i273f6427d4e6aa2bb741a8d68e377c20d8b7a98bf0a4c581769ace831a969301.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i273f6427d4e6aa2bb741a8d68e377c20d8b7a98bf0a4c581769ace831a969301.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
