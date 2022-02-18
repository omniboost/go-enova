package enova_test

import (
	"encoding/json"
	"fmt"
	"testing"

	enova "github.com/omniboost/go-enova"
)

func TestInvokeServiceMethod(t *testing.T) {
	req := client.NewInvokeServiceMethodRequest()
	req.RequestBody().InvokerParams.DatabaseHandle = enova.DatabaseHandle(client.DBName())
	req.RequestBody().InvokerParams.MethodName = enova.MethodName("Update")
	req.RequestBody().InvokerParams.Operator = enova.Operator(client.DBUsername())
	req.RequestBody().InvokerParams.Password = enova.Password(client.DBPassword())
	req.RequestBody().InvokerParams.MethodArgs = enova.MethodArgs{
		"TableName":      client.TableName(),
		"SchemaName":     client.SchemaName(),
		"ExternalSystem": client.ExternalSystemGUID(),
		"Data": enova.UpdateParamsData{
			Rows: enova.Rows{enova.Row{XML: "<Dokument />"}},
		},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
