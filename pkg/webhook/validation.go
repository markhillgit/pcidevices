package webhook

import (
	"net/http"

	"github.com/harvester/harvester/pkg/webhook/types"
	"github.com/rancher/wrangler/pkg/webhook"
)

func Validation(clients *Clients) (http.Handler, []types.Resource, error) {
	var resources []types.Resource
	validators := []types.Validator{
		NewSriovNetworkDeviceValidator(clients.PCIFactory.Devices().V1beta1().PCIDeviceClaim().Cache()),
		NewPCIDeviceClaimValidator(clients.PCIFactory.Devices().V1beta1().PCIDevice().Cache(), clients.KubevirtFactory.Kubevirt().V1().VirtualMachine().Cache()),
	}

	router := webhook.NewRouter()
	for _, v := range validators {
		addHandler(router, types.AdmissionTypeValidation, types.NewValidatorAdapter(v))
		resources = append(resources, v.Resource())
	}

	return router, resources, nil
}
