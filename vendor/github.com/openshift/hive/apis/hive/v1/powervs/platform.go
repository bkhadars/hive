package powervs

import (
	corev1 "k8s.io/api/core/v1"
)

// Platform stores all the global configuration that all machinesets use.
type Platform struct {
	// CredentialsSecretRef refers to a secret that contains the IBM account access
	// credentials.
	CredentialsSecretRef corev1.LocalObjectReference `json:"credentialsSecretRef"`

	// PowerVSResourceGroup is the resource group in which Power VS resources will be created.
	PowerVSResourceGroup string `json:"powervsResourceGroup"`

	// Region specifies the IBM Cloud colo region where the cluster will be created.
	Region string `json:"region"`

	// UserID is the login for the user's IBM Cloud account.
	UserID string `json:"userID"`

	// Zone specifies the IBM Cloud colo region where the cluster will be created.
	// At this time, only single-zone clusters are supported.
	Zone string `json:"zone"`
}