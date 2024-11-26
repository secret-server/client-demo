package main

import (
	"flag"
	"fmt"
	"strconv"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiClient "github.com/secret-server/client-api/client"
	apiRoles "github.com/secret-server/client-api/client/roles"
)

// var demoFlags = struct {
// 	AccessToken string `short:"a" long:"AccessToken" description:"Secret Server Session access token"`
// }{}


func main() {
	fmt.Println("Hello, Mock Server!")
	accessToken := flag.String("AccessToken", "", "Secret Server Session access token") 
	flag.Parse();

	var token = pointerToString(accessToken);
	if(token=="") {
		fmt.Println("AccessToken is required, missing AccessToken");
		return ;
	}

	//================================================================================================
	transport := httptransport.New("127.0.0.1:9000", "/SecretServer", []string{"http"});
	aClient := apiClient.New(transport, strfmt.Default);
	bearerTokenAuth := httptransport.BearerToken(token)

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

