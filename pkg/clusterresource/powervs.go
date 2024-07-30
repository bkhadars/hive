package clusterresource

import (
	"encoding/json"
	"fmt"

	machinev1 "github.com/openshift/api/machine/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"

	installertypes "github.com/openshift/installer/pkg/types"
	installerpowervs "github.com/openshift/installer/pkg/types/powervs"

	hivev1 "github.com/openshift/hive/apis/hive/v1"
	hivev1powervs "github.com/openshift/hive/apis/hive/v1/powervs"
	"github.com/openshift/hive/pkg/constants"
)

var _ CloudBuilder = (*PowerVSBuilder)(nil)

// PowerVSBuilder encapsulates cluster artifact generation logic specific to PowerVS.
type PowerVSBuilder struct {
	// APIKey is the IBM account api key
	APIKey string `json:"apikey"`

	// PowerVSResourceGroup is the resource group in which IBMPower VS resources will be created.
	PowerVSResourceGroup string `json:"resourcegroup"`

	// Region specifies the PowerVS region where the cluster will be created
	Region string `json:"region"`

	// UserID is the login for the user's IBM Cloud account.
	UserID string `json:"id"`

	// Zone specifies the PowerVS zone where the cluster will be created
	Zone string `json:"zone"`
}

func getPowerVSSessionVars(session *PowerVSBuilder) []byte {
	ssVars, _ := json.Marshal(session)
	return ssVars
}

func (p *PowerVSBuilder) GenerateCredentialsSecret(o *Builder) *corev1.Secret {
	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.CredsSecretName(o),
			Namespace: o.Namespace,
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			// This API KEY will be passed to the installer as constants.PowerVSAPIKeyEnvVar
			//constants.PowerVSAPIKeySecretKey: p.APIKey,
			constants.PowervsCredentialsName: getPowerVSSessionVars(p),
		},
	}
}

func (p *PowerVSBuilder) GenerateCloudObjects(o *Builder) []runtime.Object {
	return nil
}

func (p *PowerVSBuilder) GetCloudPlatform(o *Builder) hivev1.Platform {
	return hivev1.Platform{
		PowerVS: &hivev1powervs.Platform{
			CredentialsSecretRef: corev1.LocalObjectReference{
				Name: p.CredsSecretName(o),
			},
			PowerVSResourceGroup: p.PowerVSResourceGroup,
			Region:               p.Region,
			UserID:               p.UserID,
			Zone:                 p.Zone,
		},
	}
}

func (p *PowerVSBuilder) addMachinePoolPlatform(o *Builder, mp *hivev1.MachinePool) {
	mp.Spec.Platform.PowerVS = &hivev1powervs.MachinePool{
		MemoryGiB:  32,
		Processors: intstr.FromString("0.5"),
		ProcType:   machinev1.PowerVSProcessorTypeShared,
		SysType:    "s922",
	}
}

func (p *PowerVSBuilder) addInstallConfigPlatform(o *Builder, ic *installertypes.InstallConfig) {
	//change default arch
	ic.Compute[0].Architecture = installertypes.ArchitecturePPC64LE
	ic.ControlPlane.Architecture = installertypes.ArchitecturePPC64LE

	ic.Platform = installertypes.Platform{
		PowerVS: &installerpowervs.Platform{
			PowerVSResourceGroup: p.PowerVSResourceGroup,
			Region:               p.Region,
			UserID:               p.UserID,
			Zone:                 p.Zone,
		},
	}
}

func (p *PowerVSBuilder) CredsSecretName(o *Builder) string {
	return fmt.Sprintf("%s-powervs-creds", o.Name)
}
