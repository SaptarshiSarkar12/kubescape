package getter

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/attacktrack/v1alpha1"
)

type (
	// IPolicyGetter knows how to retrieve policies, i.e. frameworks and their controls.
	IPolicyGetter interface {
		GetFramework(name string) (*reporthandling.Framework, error)
		GetFrameworks() ([]reporthandling.Framework, error)
		GetControl(ID string) (*reporthandling.Control, error)

		ListFrameworks() ([]string, error)
		ListControls() ([]string, error)
	}

	// IExceptionsGetter knows how to retrieve exceptions.
	IExceptionsGetter interface {
		GetExceptions(clusterName string) ([]armotypes.PostureExceptionPolicy, error)
	}

	// IControlsInputsGetter knows how to retrieve controls inputs.
	IControlsInputsGetter interface {
		GetControlsInputs(clusterName string) (map[string][]string, error)
	}

	// IAttackTracksGetter knows how to retrieve attack tracks.
	IAttackTracksGetter interface {
		GetAttackTracks() ([]v1alpha1.AttackTrack, error)
	}

	// IBackend knows how to configure a KS Cloud client
	IBackend interface {
		GetAccountID() string
		GetClientID() string
		GetSecretKey() string
		GetCloudReportURL() string
		GetCloudAPIURL() string
		GetCloudUIURL() string
		GetCloudAuthURL() string

		SetAccountID(accountID string)
		SetClientID(clientID string)
		SetSecretKey(secretKey string)
		SetCloudReportURL(cloudReportURL string)
		SetCloudAPIURL(cloudAPIURL string)
		SetCloudUIURL(cloudUIURL string)
		SetCloudAuthURL(cloudAuthURL string)

		GetTenant() (*TenantResponse, error)
	}
)
