package powervs

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/form3tech-oss/jwt-go"
	"github.com/openshift/hive/contrib/pkg/utils"
	"github.com/openshift/hive/pkg/constants"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Region describes resources associated with a region in Power VS.
// We're using a few items from the IBM Cloud VPC offering. The region names
// for VPC are different so another function of this is to correlate those.
type Region struct {
	Description string
	VPCRegion   string
	COSRegion   string
	Zones       []string
	SysTypes    []string
	VPCZones    []string
}

type SessionStore struct {
	ID                   string `json:"id,omitempty"`
	APIKey               string `json:"apikey,omitempty"`
	DefaultRegion        string `json:"region,omitempty"`
	DefaultZone          string `json:"zone,omitempty"`
	PowerVSResourceGroup string `json:"resourcegroup,omitempty"`
}

const serviceIBMCloud = "IBMCLOUD"

// Regions holds the regions for IBM Power VS, and descriptions used during the survey.
var Regions = map[string]Region{
	"dal": {
		Description: "Dallas, USA",
		VPCRegion:   "us-south",
		COSRegion:   "us-south",
		Zones:       []string{"dal10", "dal12"},
		SysTypes:    []string{"s922", "e980"},
		VPCZones:    []string{"us-south-1", "us-south-2", "us-south-3"},
	},
	"eu-de": {
		Description: "Frankfurt, Germany",
		VPCRegion:   "eu-de",
		COSRegion:   "eu-de",
		Zones:       []string{"eu-de-1", "eu-de-2"},
		SysTypes:    []string{"s922", "e980"},
		VPCZones:    []string{"eu-de-2", "eu-de-3"},
	},
	"lon": {
		Description: "London, UK",
		VPCRegion:   "eu-gb",
		COSRegion:   "eu-gb",
		Zones:       []string{"lon06"},
		SysTypes:    []string{"s922", "e980"},
		VPCZones:    []string{"eu-gb-1", "eu-gb-2", "eu-gb-3"},
	},
	"mad": {
		Description: "Madrid, Spain",
		VPCRegion:   "eu-es",
		COSRegion:   "eu-de", // @HACK - PowerVS says COS not supported in this region
		Zones:       []string{"mad02", "mad04"},
		SysTypes:    []string{"s1022"},
		VPCZones:    []string{"eu-es-1", "eu-es-2"},
	},
	"osa": {
		Description: "Osaka, Japan",
		VPCRegion:   "jp-osa",
		COSRegion:   "jp-osa",
		Zones:       []string{"osa21"},
		SysTypes:    []string{"s922", "e980"},
		VPCZones:    []string{"jp-osa-1", "jp-osa-2", "jp-osa-3"},
	},
	"sao": {
		Description: "São Paulo, Brazil",
		VPCRegion:   "br-sao",
		COSRegion:   "br-sao",
		Zones:       []string{"sao01", "sao04"},
		SysTypes:    []string{"s922", "e980"},
		VPCZones:    []string{"br-sao-1", "br-sao-2", "br-sao-3"},
	},
	"syd": {
		Description: "Sydney, Australia",
		VPCRegion:   "au-syd",
		COSRegion:   "au-syd",
		Zones:       []string{"syd04"},
		SysTypes:    []string{"s922", "e980"},
		VPCZones:    []string{"au-syd-1", "au-syd-2", "au-syd-3"},
	},
	"wdc": {
		Description: "Washington DC, USA",
		VPCRegion:   "us-east",
		COSRegion:   "us-east",
		Zones:       []string{"wdc06", "wdc07"},
		SysTypes:    []string{"s922", "e980"},
		VPCZones:    []string{"us-east-1", "us-east-2", "us-east-3"},
	},
}

// RegionFromZone returns the region name for a given zone name.
func GetRegionFromZone(zone string) string {
	for r := range Regions {
		for z := range Regions[r].Zones {
			if zone == Regions[r].Zones[z] {
				return r
			}
		}
	}
	return ""
}

// This expects the credential file in the following search order:
// 1) ${IBM_CREDENTIALS_FILE}
// 2) <user-home-dir>/ibm-credentials.env
// 3) <current-working-directory>/ibm-credentials.env
//
// and the format is:
// $ cat ibm-credentials.env
// IBMCLOUD_AUTH_TYPE=iam
// IBMCLOUD_APIKEY=xxxxxxxxxxxxx
// IBMCLOUD_AUTH_URL=https://iam.cloud.ibm.com

// GetAuthenticator will get the authenticator for ibmcloud.
func GetAuthenticatorFromEnvironment() (core.Authenticator, error) {
	auth, err := core.GetAuthenticatorFromEnvironment(serviceIBMCloud)
	if err != nil {
		return nil, err
	}
	if auth == nil {
		return nil, fmt.Errorf("authenticator can't be nil, please set proper authentication")
	}
	return auth, nil
}

// GetProperties returns a map containing configuration properties for the specified service that are retrieved from external configuration sources.
func GetProperties() (map[string]string, error) {
	properties, err := core.GetServiceProperties(serviceIBMCloud)
	if err != nil {
		return nil, fmt.Errorf("error while fetching service properties")
	}
	return properties, nil
}

// GetUserID is function parses the user id.
func GetUserID(auth core.Authenticator) (string, error) {
	// fake request to get a barer token from the request header
	ctx := context.TODO()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com", http.NoBody)
	if err != nil {
		return "", err
	}
	err = auth.Authenticate(req)
	if err != nil {
		return "", err
	}
	bearerToken := req.Header.Get("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer") {
		bearerToken = bearerToken[7:]
	}
	token, err := jwt.Parse(bearerToken, func(_ *jwt.Token) (interface{}, error) {
		return "", nil
	})
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return "", err
	}

	return token.Claims.(jwt.MapClaims)["id"].(string), nil
}

// ConfigureCreds loads secrets designated by the environment variables CLUSTERDEPLOYMENT_NAMESPACE
// and CREDS_SECRET_NAME and configures IBMPower VS credential environment variables accordingly.
func ConfigureCreds(c client.Client) {
	credsSecret := utils.LoadSecretOrDie(c, "CREDS_SECRET_NAME")
	if credsSecret == nil {
		return
	}

	utils.ProjectToDir(credsSecret, constants.PowerVSCredentialsDir, nil)
	os.Setenv("POWERVS_AUTH_FILEPATH", constants.PowerVSCredentialsDir+"/"+constants.PowervsCredentialsName)

	// Install cluster proxy trusted CA bundle
	utils.InstallCerts(constants.TrustedCABundleDir)
}
