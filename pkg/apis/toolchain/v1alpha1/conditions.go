package v1alpha1

import (
	toolchainv1alpha1 "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
)

const (
	// status condition type

	CheReady    toolchainv1alpha1.ConditionType = "CheReady"
	TektonReady toolchainv1alpha1.ConditionType = "TektonReady"
	Ready       toolchainv1alpha1.ConditionType = "Ready"

	// Status condition reasons

	InstallingReason      = "Installing"
	FailedToInstallReason = "FailedToInstall"
	InstalledReason       = "Installed"
)
