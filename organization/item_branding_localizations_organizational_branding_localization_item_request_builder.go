package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
type ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters read the properties and relationships of an organizationalBrandingLocalization object. To retrieve a localization branding object, specify the value of **id** in the URL.
type ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters
}
// ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackgroundImage provides operations to manage the media for the organization entity.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) BackgroundImage()(*ItemBrandingLocalizationsItemBackgroundImageRequestBuilder) {
    return NewItemBrandingLocalizationsItemBackgroundImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BannerLogo provides operations to manage the media for the organization entity.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) BannerLogo()(*ItemBrandingLocalizationsItemBannerLogoRequestBuilder) {
    return NewItemBrandingLocalizationsItemBannerLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) {
    m := &ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/branding/localizations/{organizationalBrandingLocalization%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomCSS provides operations to manage the media for the organization entity.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) CustomCSS()(*ItemBrandingLocalizationsItemCustomCSSRequestBuilder) {
    return NewItemBrandingLocalizationsItemCustomCSSRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a localized branding object. To delete the organizationalBrandingLocalization object, all images (Stream types) must first be removed from the object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/organizationalbrandinglocalization-delete?view=graph-rest-1.0
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Favicon provides operations to manage the media for the organization entity.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Favicon()(*ItemBrandingLocalizationsItemFaviconRequestBuilder) {
    return NewItemBrandingLocalizationsItemFaviconRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of an organizationalBrandingLocalization object. To retrieve a localization branding object, specify the value of **id** in the URL.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/organizationalbrandinglocalization-get?view=graph-rest-1.0
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingLocalizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationalBrandingLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingLocalizationable), nil
}
// HeaderLogo provides operations to manage the media for the organization entity.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) HeaderLogo()(*ItemBrandingLocalizationsItemHeaderLogoRequestBuilder) {
    return NewItemBrandingLocalizationsItemHeaderLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an organizationalBrandingLocalization object for a specific localization.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/organizationalbrandinglocalization-update?view=graph-rest-1.0
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingLocalizationable, requestConfiguration *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingLocalizationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationalBrandingLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingLocalizationable), nil
}
// SquareLogo provides operations to manage the media for the organization entity.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) SquareLogo()(*ItemBrandingLocalizationsItemSquareLogoRequestBuilder) {
    return NewItemBrandingLocalizationsItemSquareLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SquareLogoDark provides operations to manage the media for the organization entity.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) SquareLogoDark()(*ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) {
    return NewItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a localized branding object. To delete the organizationalBrandingLocalization object, all images (Stream types) must first be removed from the object.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an organizationalBrandingLocalization object. To retrieve a localization branding object, specify the value of **id** in the URL.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of an organizationalBrandingLocalization object for a specific localization.
func (m *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingLocalizationable, requestConfiguration *ItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
