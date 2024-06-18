import (
	"context"
	"fmt"
	"io"

	appengine "cloud.google.com/go/appengine/apiv1"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

// updateService updates the service to the given version.
func updateService(w io.Writer, projectID, serviceID, versionID string) error {
	// projectID := "my-project-id"
	// serviceID := "my-service"
	// versionID := "v1"
	ctx := context.Background()
	client, err := appengine.NewServicesClient(ctx)
	if err != nil {
		return fmt.Errorf("appengine.NewServicesClient: %v", err)
	}
	defer client.Close()

	req := &appenginepb.UpdateServiceRequest{
		Name:     fmt.Sprintf("apps/%s/services/%s", projectID, serviceID),
		Service: &appenginepb.Service{
			Id: serviceID,
			TrafficSplit: map[string]float64{
				versionID: 1.0,
			},
		},
		UpdateMask: &appenginepb.UpdateServiceRequest_UpdateMask{
			Paths: []string{"traffic_split"},
		},
	}

	op, err := client.UpdateService(ctx, req)
	if err != nil {
		return fmt.Errorf("client.UpdateService: %v", err)
	}
	resp, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("op.Wait: %v", err)
	}
	fmt.Fprintf(w, "Updated service: %v\n", resp.GetName())
	return nil
}
  
