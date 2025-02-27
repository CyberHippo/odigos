package resources

import (
	"github.com/keyval-dev/odigos/cli/pkg/containers"

	"github.com/keyval-dev/odigos/cli/pkg/labels"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	keyvalProxyServiceName            = "odigos-cloud-k8s"
	keyvalProxyImage                  = "keyval/odigos-proxy-k8s"
	keyvalProxyAppName                = "odigos-cloud-proxy"
	keyvalProxyDeploymentName         = "odigos-cloud-proxy"
	keyvalProxyServiceAccountName     = "odigos-cloud-proxy"
	keyvalProxyRoleName               = "odigos-cloud-proxy"
	keyvalProxyRoleBindingName        = "odigos-cloud-proxy"
	keyvalProxyClusterRoleName        = "odigos-cloud-proxy"
	keyvalProxyClusterRoleBindingName = "odigos-cloud-proxy"
)

func NewKeyvalProxyServiceAccount() *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   keyvalProxyServiceAccountName,
			Labels: labels.OdigosSystem,
		},
	}
}

func NewKeyvalProxyRole(ns string) *rbacv1.Role {
	return &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Role",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyvalProxyRoleName,
			Namespace: ns,
			Labels:    labels.OdigosSystem,
		},
		Rules: []rbacv1.PolicyRule{
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"update",
					"watch",
				},
				APIGroups: []string{""},
				Resources: []string{
					"secrets",
				},
			},
		},
	}
}

func NewKeyvalProxyRoleBinding(ns string) *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyvalProxyRoleBindingName,
			Namespace: ns,
			Labels:    labels.OdigosSystem,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      keyvalProxyServiceAccountName,
				Namespace: ns,
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "Role",
			Name:     keyvalProxyRoleName,
			APIGroup: "rbac.authorization.k8s.io",
		},
	}
}

func NewKeyvalProxyClusterRole() *rbacv1.ClusterRole {
	return &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   keyvalProxyClusterRoleName,
			Labels: labels.OdigosSystem,
		},
		Rules: []rbacv1.PolicyRule{
			{
				Verbs: []string{
					"list",
					"watch",
					"get",
					"update",
				},
				APIGroups: []string{""},
				Resources: []string{
					"namespaces",
				},
			},
			{
				Verbs: []string{
					"list",
					"watch",
					"get",
					"update",
				},
				APIGroups: []string{"apps"},
				Resources: []string{
					"deployments",
				},
			},
			{
				Verbs: []string{
					"list",
					"watch",
					"get",
					"update",
				},
				APIGroups: []string{"apps"},
				Resources: []string{
					"daemonsets",
				},
			},
			{
				Verbs: []string{
					"list",
					"watch",
					"get",
					"update",
				},
				APIGroups: []string{"apps"},
				Resources: []string{
					"statefulsets",
				},
			},
			{
				Verbs: []string{
					"list",
					"watch",
					"get",
					"update",
				},
				APIGroups: []string{"apps"},
				Resources: []string{
					"statefulsets",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"odigos.io",
				},
				Resources: []string{
					"destinations",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"odigos.io",
				},
				Resources: []string{
					"instrumentedapplications",
				},
			},
		},
	}
}

func NewKeyvalProxyClusterRoleBinding(ns string) *rbacv1.ClusterRoleBinding {
	return &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   keyvalProxyClusterRoleBindingName,
			Labels: labels.OdigosSystem,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      keyvalProxyServiceAccountName,
				Namespace: ns,
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "ClusterRole",
			Name:     keyvalProxyClusterRoleName,
			APIGroup: "rbac.authorization.k8s.io",
		},
	}
}

func NewKeyvalProxyDeployment(version string, ns string) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: keyvalProxyDeploymentName,
			Labels: map[string]string{
				"app":                       keyvalProxyAppName,
				labels.OdigosSystemLabelKey: labels.OdigosSystemLabelValue,
			},
			Annotations: map[string]string{
				"odigos.io/skip": "true",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptrint32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": keyvalProxyAppName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": keyvalProxyAppName,
					},
					Annotations: map[string]string{
						"kubectl.kubernetes.io/default-container": keyvalProxyAppName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  keyvalProxyAppName,
							Image: containers.GetImageName(keyvalProxyImage, version),
							Args: []string{
								"--health-probe-bind-address=:8081",
								"--metrics-bind-address=127.0.0.1:8080",
								// "--leader-elect",
							},
							Env: []corev1.EnvVar{
								{
									Name: "CURRENT_NS",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
								{
									Name:  "OTEL_SERVICE_NAME",
									Value: keyvalProxyServiceName,
								},
								{
									Name: odigosCloudTokenEnvName,
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: odigosCloudSecretName,
											},
											Key: odigosCloudApiKeySecretKey,
										},
									},
								},
							},
							EnvFrom: []corev1.EnvFromSource{
								{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: ownTelemetryOtelConfig,
										},
									},
								},
								{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: OdigosDeploymentConfigMapName,
										},
									},
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									"cpu":    resource.MustParse("500m"),
									"memory": *resource.NewQuantity(134217728, resource.BinarySI),
								},
								Requests: corev1.ResourceList{
									"cpu":    resource.MustParse("10m"),
									"memory": *resource.NewQuantity(67108864, resource.BinarySI),
								},
							},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{
											Type:   intstr.Type(0),
											IntVal: 8081,
										},
									},
								},
								InitialDelaySeconds: 15,
								TimeoutSeconds:      0,
								PeriodSeconds:       20,
								SuccessThreshold:    0,
								FailureThreshold:    0,
							},
							SecurityContext: &corev1.SecurityContext{},
						},
					},
					TerminationGracePeriodSeconds: ptrint64(10),
					ServiceAccountName:            keyvalProxyServiceAccountName,
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: ptrbool(true),
					},
				},
			},
			Strategy:        appsv1.DeploymentStrategy{},
			MinReadySeconds: 0,
		},
	}
}
