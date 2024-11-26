package main

import (
	"fmt"
	"strconv"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiClient "github.com/secret-server/client-api/client"
	authClient "github.com/secret-server/client-api/client/authentication"
	apiRoles "github.com/secret-server/client-api/client/roles"
)

func main() {
	fmt.Println("Hello, Mock Server!")

	//================================================================================================
	transport := httptransport.New("127.0.0.1:9000", "/SecretServer", []string{"http"});
	aClient := apiClient.New(transport, strfmt.Default);

	//================================================================================================
	var clientParams = authClient.NewOAuth2ServiceAuthorizeParams();
	clientParams.GrantType = "password";
	username := "KevinKelche"
	clientParams.Username = &username
	password := "bleach.out.34"
	clientParams.Password = &password;
	authorized, err := aClient.Authentication.OAuth2ServiceAuthorize(clientParams);
    if err != nil {
		fmt.Println("Error authenticating user", clientParams.Username);
		return ;
    }
	fmt.Println("authorized="+authorized.String());

	var accessToken = pointerToString(authorized.Payload.AccessToken);
	fmt.Println("authorized.Payload.AccessToken="+accessToken);

	bearerTokenAuth := httptransport.BearerToken(accessToken)

	//================================================================================================
	var roleParams = apiRoles.NewRolesServiceGetParams();
	roleParams.WithID(2);
	role, err := aClient.Roles.RolesServiceGet(roleParams, bearerTokenAuth);
    if err != nil {
		fmt.Println("Error retrieving role =", roleParams.ID);
		return ;
    }
	fmt.Println("role="+role.String());

	// var accessToken = pointerToString(&role.Payload.Name);
	fmt.Println("role.Payload.Name="+role.Payload.Name);
	fmt.Println("role.Payload.Created="+role.Payload.Created.String());
	fmt.Println("role.Payload.Enabled="+strconv.FormatBool(role.Payload.Enabled));
	fmt.Println("role.Payload.ID="+strconv.Itoa(int(role.Payload.ID)));
}

func pointerToString(value *string) string {
	if( value==nil ) {
		return "";
	}
	return *value;
}

