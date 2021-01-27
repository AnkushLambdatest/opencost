package kubecost

import (
	"fmt"
	"strings"
)

// LabelConfig is a port of type AnalyzerConfig. We need to be more thoughtful
// about design at some point, but this is a stop-gap measure, which is required
// because AnalyzerConfig is defined in package main, so it can't be imported.
type LabelConfig struct {
	DepartmentLabel          string `json:"department_label"`
	EnvironmentLabel         string `json:"environment_label"`
	OwnerLabel               string `json:"owner_label"`
	ProductLabel             string `json:"product_label"`
	TeamLabel                string `json:"team_label"`
	ClusterExternalLabel     string `json:"cluster_external_label"`
	NamespaceExternalLabel   string `json:"namespace_external_label"`
	ControllerExternalLabel  string `json:"controller_external_label"`
	DaemonsetExternalLabel   string `json:"daemonset_external_label"`
	DeploymentExternalLabel  string `json:"deployment_external_label"`
	StatefulsetExternalLabel string `json:"statefulset_external_label"`
	ServiceExternalLabel     string `json:"service_external_label"`
	PodExternalLabel         string `json:"pod_external_label"`
	DepartmentExternalLabel  string `json:"department_external_label"`
	EnvironmentExternalLabel string `json:"environment_external_label"`
	OwnerExternalLabel       string `json:"owner_external_label"`
	ProductExternalLabel     string `json:"product_external_label"`
	TeamExternalLabel        string `json:"team_external_label"`
}

// Map returns the config as a basic string map, with default values if not set
func (lc *LabelConfig) Map() map[string]string {
	// Start with default values
	m := map[string]string{
		"department_label":           "department",
		"environment_label":          "env",
		"owner_label":                "owner",
		"product_label":              "app",
		"team_label":                 "team",
		"cluster_external_label":     "kubernetes_cluster",
		"namespace_external_label":   "kubernetes_namespace",
		"controller_external_label":  "kubernetes_controller",
		"daemonset_external_label":   "kubernetes_daemonset",
		"deployment_external_label":  "kubernetes_deployment",
		"statefulset_external_label": "kubernetes_statefulset",
		"service_external_label":     "kubernetes_service",
		"pod_external_label":         "kubernetes_pod",
		"department_external_label":  "kubernetes_label_department",
		"environment_external_label": "kubernetes_label_env",
		"owner_external_label":       "kubernetes_label_owner",
		"product_external_label":     "kubernetes_label_app",
		"team_external_label":        "kubernetes_label_team",
	}

	if lc == nil {
		return m
	}

	if lc.DepartmentLabel != "" {
		m["department_label"] = lc.DepartmentLabel
	}

	if lc.EnvironmentLabel != "" {
		m["environment_label"] = lc.EnvironmentLabel
	}

	if lc.OwnerLabel != "" {
		m["owner_label"] = lc.OwnerLabel
	}

	if lc.ProductLabel != "" {
		m["product_label"] = lc.ProductLabel
	}

	if lc.TeamLabel != "" {
		m["team_label"] = lc.TeamLabel
	}

	if lc.ClusterExternalLabel != "" {
		m["cluster_external_label"] = lc.ClusterExternalLabel
	}

	if lc.NamespaceExternalLabel != "" {
		m["namespace_external_label"] = lc.NamespaceExternalLabel
	}

	if lc.ControllerExternalLabel != "" {
		m["controller_external_label"] = lc.ControllerExternalLabel
	}

	if lc.DaemonsetExternalLabel != "" {
		m["daemonset_external_label"] = lc.DaemonsetExternalLabel
	}

	if lc.DeploymentExternalLabel != "" {
		m["deployment_external_label"] = lc.DeploymentExternalLabel
	}

	if lc.StatefulsetExternalLabel != "" {
		m["statefulset_external_label"] = lc.StatefulsetExternalLabel
	}

	if lc.ServiceExternalLabel != "" {
		m["service_external_label"] = lc.ServiceExternalLabel
	}

	if lc.PodExternalLabel != "" {
		m["pod_external_label"] = lc.PodExternalLabel
	}

	if lc.DepartmentExternalLabel != "" {
		m["department_external_label"] = lc.DepartmentExternalLabel
	} else if lc.DepartmentLabel != "" {
		m["department_external_label"] = "kubernetes_label_" + lc.DepartmentLabel
	}

	if lc.EnvironmentExternalLabel != "" {
		m["environment_external_label"] = lc.EnvironmentExternalLabel
	} else if lc.EnvironmentLabel != "" {
		m["environment_external_label"] = "kubernetes_label_" + lc.EnvironmentLabel
	}

	if lc.OwnerExternalLabel != "" {
		m["owner_external_label"] = lc.OwnerExternalLabel
	} else if lc.OwnerLabel != "" {
		m["owner_external_label"] = "kubernetes_label_" + lc.OwnerLabel
	}

	if lc.ProductExternalLabel != "" {
		m["product_external_label"] = lc.ProductExternalLabel
	} else if lc.ProductLabel != "" {
		m["product_external_label"] = "kubernetes_label_" + lc.ProductLabel
	}

	if lc.TeamExternalLabel != "" {
		m["team_external_label"] = lc.TeamExternalLabel
	} else if lc.TeamLabel != "" {
		m["team_external_label"] = "kubernetes_label_" + lc.TeamLabel
	}

	return m
}

// ExternalQueryLabels returns the config's external labels as a mapping of the
// query column to the label it should set;
// e.g. if the config stores "statefulset_external_label": "kubernetes_sset",
//      then this would return "kubernetes_sset": "statefulset"
func (lc *LabelConfig) ExternalQueryLabels() map[string]string {
	queryLabels := map[string]string{}

	for label, query := range lc.Map() {
		if strings.HasSuffix(label, "external_label") && query != "" {
			queryLabels[query] = label
		}
	}

	return queryLabels
}

// AllocationPropertyLabels returns the config's external resource labels
// as a mapping from k8s resource-to-label name.
// e.g. if the config stores "statefulset_external_label": "kubernetes_sset",
//      then this would return "statefulset": "kubernetes_sset"
// e.g. if the config stores "owner_label": "product_owner",
//      then this would return "label:product_owner": "product_owner"
func (lc *LabelConfig) AllocationPropertyLabels() map[string]string {
	labels := map[string]string{}

	for labelKind, labelName := range lc.Map() {
		if labelName != "" {
			switch labelKind {
			case "namespace_external_label":
				labels["namespace"] = labelName
			case "cluster_external_label":
				labels["cluster"] = labelName
			case "controller_external_label":
				labels["controller"] = labelName
			case "product_external_label":
				labels["product"] = labelName
			case "service_external_label":
				labels["service"] = labelName
			case "deployment_external_label":
				labels["deployment"] = labelName
			case "statefulset_external_label":
				labels["statefulset"] = labelName
			case "daemonset_external_label":
				labels["daemonset"] = labelName
			case "pod_external_label":
				labels["pod"] = labelName
			default:
				labels[fmt.Sprintf("label:%s", labelName)] = labelName
			}
		}
	}

	return labels
}
