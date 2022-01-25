# Microsoft Graph SDK for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/microsoftgraph/msgraph-sdk-go/)](https://pkg.go.dev/github.com/microsoftgraph/msgraph-sdk-go/)

Get started with the Microsoft Graph SDK for Go by integrating the [Microsoft Graph API](https://docs.microsoft.com/graph/overview) into your Go application!

> **Note:** this SDK allows you to build applications using the [v1.0](https://docs.microsoft.com/graph/use-the-api#version) of Microsoft Graph. If you want to try the latest Microsoft Graph APIs under beta, use our [beta SDK](https://github.com/microsoftgraph/msgraph-beta-sdk-go) instead.
>
> **Note:** the Microsoft Graph Go SDK is currently in Community Preview. During this period we're expecting breaking changes to happen to the SDK based on community's feedback. Checkout the [known limitations](https://github.com/microsoftgraph/msgraph-sdk-go-core/issues/1).

## 1. Installation

```Shell
go get -u github.com/microsoftgraph/msgraph-sdk-go
go get -u github.com/Azure/azure-sdk-for-go/sdk/azidentity
go get -u github.com/microsoft/kiota/authentication/go/azure
```

## 2. Getting started

### 2.1 Register your application

Register your application by following the steps at [Register your app with the Microsoft Identity Platform](https://docs.microsoft.com/graph/auth-register-app-v2).

### 2.2 Create an AuthenticationProvider object

An instance of the **GraphRequestAdapter** class handles building client. To create a new instance of this class, you need to provide an instance of **AuthenticationProvider**, which can authenticate requests to Microsoft Graph.

For an example of how to get an authentication provider, see [choose a Microsoft Graph authentication provider](https://docs.microsoft.com/graph/sdks/choose-authentication-providers?tabs=Go).

> Note: we are working to add the getting started information for Go to our public documentation, in the meantime the following sample should help you getting started.

```Golang
import (
    azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    a          "github.com/microsoft/kiota/authentication/go/azure"
    "context"
)

cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
    TenantID: "<the tenant id from your app registration>",
    ClientID: "<the client id from your app registration>",
    UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
        fmt.Println(message.Message)
        return nil
    },
})

if err != nil {
    fmt.Printf("Error creating credentials: %v\n", err)
}

auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"Files.Read"})
if err != nil {
    fmt.Printf("Error authentication provider: %v\n", err)
    return
}
```

### 2.3 Get a Graph Service Client Adapter object

You must get a **GraphRequestAdapter** object to make requests against the service.

```Golang
import msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
if err != nil {
    fmt.Printf("Error creating adapter: %v\n", err)
    return
}
client := msgraphsdk.NewGraphServiceClient(adapter)
```

## 3. Make requests against the service

After you have a **GraphServiceClient** that is authenticated, you can begin making calls against the service. The requests against the service look like our [REST API](https://docs.microsoft.com/graph/api/overview?view=graph-rest-1.0).

### 3.1 Get the user's drive

To retrieve the user's drive:

```Golang
result, err := client.Me().Drive().Get(nil)
if err != nil {
    fmt.Printf("Error getting the drive: %v\n", err)
}
fmt.Printf("Found Drive : %v\n", result.GetId())
```

## 4. Getting results that span across multiple pages

Items in a collection response can span across multiple pages. To get the complete set of items in the collection, your application must make additional calls to get the subsequent pages until no more next link is provided in the response.

### 4.1 Get all the users in an environment

To retrieve the users:

```Golang
result, err := client.Users().Get(nil)
if err != nil {
    fmt.Printf("Error getting users: %v\n", err)
    return err
}

for {
    if result.GetNextLink() == nil {
        break
    }

    result, err = users.NewUsersRequestBuilder(*result.GetNextLink(), adapter).Get(nil)
    if err != nil {
        fmt.Printf("Error getting users: %v\n", err)
        break
    }
}
```

## 5. Documentation

For more detailed documentation, see:

* [Overview](https://docs.microsoft.com/graph/overview)
* [Collections](https://docs.microsoft.com/graph/sdks/paging)
* [Making requests](https://docs.microsoft.com/graph/sdks/create-requests)
* [Known issues](https://github.com/MicrosoftGraph/msgraph-sdk-go/issues)
* [Contributions](https://github.com/microsoftgraph/msgraph-sdk-go/blob/main/CONTRIBUTING.md)

## 6. Issues

For known issues, see [issues](https://github.com/MicrosoftGraph/msgraph-sdk-go/issues).

## 7. Contributions

The Microsoft Graph SDK is open for contribution. To contribute to this project, see [Contributing](https://github.com/microsoftgraph/msgraph-sdk-go/blob/main/CONTRIBUTING.md).

## 8. License

Copyright (c) Microsoft Corporation. All Rights Reserved. Licensed under the [MIT license](LICENSE).

## 9. Third-party notices

[Third-party notices](THIRD%20PARTY%20NOTICES)
