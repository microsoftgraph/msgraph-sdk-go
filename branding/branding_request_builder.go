package branding

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i11606dfdbbf1d1bc678546ed98e21fe3520002d2ebc49f2e965d358c6ec48612 "github.com/microsoftgraph/msgraph-sdk-go/branding/localizations"
    i1bd33413feee997e32357cccfaf363eb0ea1df44f2f048b98d4bf0d8ba264765 "github.com/microsoftgraph/msgraph-sdk-go/branding/bannerlogo"
    i66d3bf3612261a255368fda89fdc148f565c3afcfadd3d6aa59f39b9ca851cd3 "github.com/microsoftgraph/msgraph-sdk-go/branding/backgroundimage"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    icc91487517040d6fd90bfc70b7fdff5086c556c304c7d37df61cf4fa1927b390 "github.com/microsoftgraph/msgraph-sdk-go/branding/squarelogo"
    i91122bc73f66085eb8e21b0f9aeadfa4ab90c1d5f4782e2d5b68f2f03fbf9409 "github.com/microsoftgraph/msgraph-sdk-go/branding/localizations/item"
)

// BrandingRequestBuilder provides operations to manage the organizationalBranding singleton.
type BrandingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BrandingRequestBuilderGetOptions options for Get
type BrandingRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *BrandingRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// BrandingRequestBuilderGetQueryParameters get branding
type BrandingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BrandingRequestBuilderPatchOptions options for Patch
type BrandingRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// BackgroundImage the backgroundImage property
func (m *BrandingRequestBuilder) BackgroundImage()(*i66d3bf3612261a255368fda89fdc148f565c3afcfadd3d6aa59f39b9ca851cd3.BackgroundImageRequestBuilder) {
    return i66d3bf3612261a255368fda89fdc148f565c3afcfadd3d6aa59f39b9ca851cd3.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo the bannerLogo property
func (m *BrandingRequestBuilder) BannerLogo()(*i1bd33413feee997e32357cccfaf363eb0ea1df44f2f048b98d4bf0d8ba264765.BannerLogoRequestBuilder) {
    return i1bd33413feee997e32357cccfaf363eb0ea1df44f2f048b98d4bf0d8ba264765.NewBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBrandingRequestBuilderInternal instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BrandingRequestBuilder) {
    m := &BrandingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/branding{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBrandingRequestBuilder instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BrandingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBrandingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get branding
func (m *BrandingRequestBuilder) CreateGetRequestInformation(options *BrandingRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update branding
func (m *BrandingRequestBuilder) CreatePatchRequestInformation(options *BrandingRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get branding
func (m *BrandingRequestBuilder) Get(options *BrandingRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationalBrandingFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable), nil
}
// Localizations the localizations property
func (m *BrandingRequestBuilder) Localizations()(*i11606dfdbbf1d1bc678546ed98e21fe3520002d2ebc49f2e965d358c6ec48612.LocalizationsRequestBuilder) {
    return i11606dfdbbf1d1bc678546ed98e21fe3520002d2ebc49f2e965d358c6ec48612.NewLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.branding.localizations.item collection
func (m *BrandingRequestBuilder) LocalizationsById(id string)(*i91122bc73f66085eb8e21b0f9aeadfa4ab90c1d5f4782e2d5b68f2f03fbf9409.OrganizationalBrandingLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organizationalBrandingLocalization_id"] = id
    }
    return i91122bc73f66085eb8e21b0f9aeadfa4ab90c1d5f4782e2d5b68f2f03fbf9409.NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update branding
func (m *BrandingRequestBuilder) Patch(options *BrandingRequestBuilderPatchOptions)(error) {
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
// SquareLogo the squareLogo property
func (m *BrandingRequestBuilder) SquareLogo()(*icc91487517040d6fd90bfc70b7fdff5086c556c304c7d37df61cf4fa1927b390.SquareLogoRequestBuilder) {
    return icc91487517040d6fd90bfc70b7fdff5086c556c304c7d37df61cf4fa1927b390.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
