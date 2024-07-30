package deprovision

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/openshift/hive/contrib/pkg/utils"
	powervsutils "github.com/openshift/hive/contrib/pkg/utils/powervs"
	"github.com/openshift/hive/pkg/powervsclient"
	"github.com/openshift/installer/pkg/destroy/powervs"
	"github.com/openshift/installer/pkg/types"
	typespowervs "github.com/openshift/installer/pkg/types/powervs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// powerVSDeprovisionOptions is the set of options to deprovision an IBM Cloud cluster
type powerVSDeprovisionOptions struct {
	baseDomain           string
	cisInstanceCRN       string
	clusterName          string
	infraID              string
	logLevel             string
	powerVSResourceGroup string
	zone                 string
}

// NewDeprovisionPowerVSCommand is the entrypoint to create the IBM Cloud deprovision subcommand
func NewDeprovisionPowerVSCommand() *cobra.Command {
	opt := &powerVSDeprovisionOptions{}
	cmd := &cobra.Command{
		Use:   "powervs <INFRAID> --base-domain=<BASE_DOMAIN> --cluster-name=<CLUSTERNAME> --powervs-resource-group=<RESOURCEGROUP_NAME> --zone=<ZONE>",
		Short: "Deprovision PowerVS Cloud assets (as created by openshift-installer)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := opt.Complete(cmd, args); err != nil {
				log.WithError(err).Fatal("failed to complete options")
			}
			if err := opt.Validate(cmd); err != nil {
				log.WithError(err).Fatal("validation failed")
			}
			if err := opt.Run(); err != nil {
				log.WithError(err).Fatal("Runtime error")
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opt.logLevel, "loglevel", "info", "log level, one of: debug, info, warn, error, fatal, panic")

	// Required flags
	flags.StringVar(&opt.baseDomain, "base-domain", "", "cluster's base domain")
	flags.StringVar(&opt.clusterName, "cluster-name", "", "cluster's name")
	flags.StringVar(&opt.powerVSResourceGroup, "powervs-resource-group", "", "resource group where the cluster is installed.")
	flags.StringVar(&opt.zone, "zone", "", "zone in which to deprovision cluster")

	return cmd
}

// Complete finishes parsing arguments for the command
func (o *powerVSDeprovisionOptions) Complete(cmd *cobra.Command, args []string) error {
	o.infraID = os.Args[0]

	client, err := utils.GetClient()
	if err != nil {
		return errors.Wrap(err, "failed to get client")
	}
	powervsutils.ConfigureCreds(client)

	// Create PowerVS Client
	content, err := os.ReadFile(os.Getenv("POWERVS_AUTH_FILEPATH"))
	if err != nil {
		return err
	}
	var ss powervsutils.SessionStore
	err = json.Unmarshal(content, &ss)
	if ss.APIKey == "" {
		return fmt.Errorf("failed to read the apikey")
	}
	powerVSClient, err := powervsclient.NewClient(ss.APIKey)
	if err != nil {
		return errors.Wrap(err, "Unable to create PowerVS Cloud client")
	}

	// Retrieve CISInstanceCRN
	cisInstanceCRN, err := powervsclient.GetCISInstanceCRN(powerVSClient, context.TODO(), o.baseDomain)
	if err != nil {
		return err
	}
	o.cisInstanceCRN = cisInstanceCRN

	return nil
}

// Validate ensures that option values make sense
func (o *powerVSDeprovisionOptions) Validate(cmd *cobra.Command) error {
	if o.baseDomain == "" {
		cmd.Usage()
		return fmt.Errorf("no --base-domain provided, cannot proceed")
	}
	if o.clusterName == "" {
		cmd.Usage()
		return fmt.Errorf("no --cluster-name provided, cannot proceed")
	}
	if o.powerVSResourceGroup == "" {
		cmd.Usage()
		return fmt.Errorf("no --powervs-resource-group provided, cannot proceed")
	}
	if o.zone == "" {
		cmd.Usage()
		return fmt.Errorf("no --zone provided, cannot proceed")
	}
	return nil
}

// Run executes the command
func (o *powerVSDeprovisionOptions) Run() error {
	logger, err := utils.NewLogger(o.logLevel)
	if err != nil {
		return err
	}

	metadata := &types.ClusterMetadata{
		ClusterName: o.clusterName,
		InfraID:     o.infraID,
		ClusterPlatformMetadata: types.ClusterPlatformMetadata{
			PowerVS: &typespowervs.Metadata{
				BaseDomain:           o.baseDomain,
				CISInstanceCRN:       o.cisInstanceCRN,
				Region:               powervsutils.GetRegionFromZone(o.zone),
				Zone:                 o.zone,
				PowerVSResourceGroup: o.powerVSResourceGroup,
			},
		},
	}

	destroyer, err := powervs.New(logger, metadata)
	if err != nil {
		return err
	}

	// ClusterQuota stomped in return
	_, err = destroyer.Run()
	return err
}
