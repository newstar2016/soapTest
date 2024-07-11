package main

import (
	"encoding/xml"
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	"soapTest/rich"
)

func main() {
	client := soap.NewClient("https://api-test.rich-healthcare.com/zj/RichHealthThridInterfaceForCard.svc/RichHealthService")
	richClient := rich.NewRichHealthThridInterfaceForCard(client)
	requestj := "{\"IdentityId\":\"ZhongKang\",\"Password\":\"123456\",\"MacAddress\":\"00:16:3E:06:58:BB\"}"
	resp, err := richClient.CreateToken(&rich.CreateToken{
		XMLName:     xml.Name{},
		RequestJson: &requestj,
	})
	fmt.Println(*resp.CreateTokenResult)
	fmt.Println(err)
}
