package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kentik/community_sdk_golang/apiv6/examples"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/cloudexport"
)

func main() {
	client := examples.NewClient()
	var err error

	if err = runCRUD(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = runGetAll(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runCRUD(client *kentikapi.Client) error {
	export := cloudexport.NewV202101beta1CloudExport()

	fmt.Println("### CREATE")
	gce := *cloudexport.NewV202101beta1GceProperties()
	gce.SetProject("test gce project")
	gce.SetSubscription("test gce subscription")
	export.SetGce(gce)
	export.SetCloudProvider("gce")
	export.SetName("test_gce_export")
	export.SetPlanId("11467")
	export.SetType(cloudexport.V202101BETA1CLOUDEXPORTTYPE_KENTIK_MANAGED)
	createReqPayload := *cloudexport.NewV202101beta1CreateCloudExportRequest()
	createReqPayload.Export = export

	createResp, httpResp, err := client.CloudExportAdminServiceApi.
		ExportCreate(context.Background()).
		Body(createReqPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(createResp)
	fmt.Println()

	created := createResp.Export

	fmt.Println("### UPDATE")
	created.SetDescription("Updated description")
	updateReqPayload := *cloudexport.NewV202101beta1UpdateCloudExportRequest()
	updateReqPayload.Export = created

	updateResp, httpResp, err := client.CloudExportAdminServiceApi.
		ExportUpdate(context.Background(), *created.Id).
		Body(updateReqPayload).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(updateResp)
	fmt.Println()

	fmt.Println("### GET")
	getResp, httpResp, err := client.CloudExportAdminServiceApi.
		ExportGet(context.Background(), *created.Id).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	examples.PrettyPrint(getResp)
	fmt.Println()

	fmt.Println("### DELETE")
	deleteResp, httpResp, err := client.CloudExportAdminServiceApi.
		ExportDelete(context.Background(), *created.Id).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	fmt.Println("Success")
	examples.PrettyPrint(deleteResp)
	fmt.Println()

	return nil
}

func runGetAll(client *kentikapi.Client) error {
	fmt.Println("### GET ALL")
	getAllResp, httpResp, err := client.CloudExportAdminServiceApi.
		ExportList(context.Background()).
		Execute()
	if err != nil {
		return fmt.Errorf("%v %v", err, httpResp)
	}
	exports := *getAllResp.Exports
	fmt.Println("Num exports:", len(exports))
	examples.PrettyPrint(exports)
	fmt.Println()

	return nil
}
