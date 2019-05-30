// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.CommonConfig":           schema_pkg_apis_app_v1_CommonConfig(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.Condition":              schema_pkg_apis_app_v1_Condition(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.DatabaseObject":         schema_pkg_apis_app_v1_DatabaseObject(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.Deployments":            schema_pkg_apis_app_v1_Deployments(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.ExternalDatabaseObject": schema_pkg_apis_app_v1_ExternalDatabaseObject(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.GitSource":              schema_pkg_apis_app_v1_GitSource(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieApp":                 schema_pkg_apis_app_v1_KieApp(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppAuthObject":       schema_pkg_apis_app_v1_KieAppAuthObject(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppBuildObject":      schema_pkg_apis_app_v1_KieAppBuildObject(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppObject":           schema_pkg_apis_app_v1_KieAppObject(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppObjects":          schema_pkg_apis_app_v1_KieAppObjects(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppRegistry":         schema_pkg_apis_app_v1_KieAppRegistry(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppSpec":             schema_pkg_apis_app_v1_KieAppSpec(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppStatus":           schema_pkg_apis_app_v1_KieAppStatus(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppUpgrades":         schema_pkg_apis_app_v1_KieAppUpgrades(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieServerSet":           schema_pkg_apis_app_v1_KieServerSet(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.LDAPAuthConfig":         schema_pkg_apis_app_v1_LDAPAuthConfig(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.RoleMapperAuthConfig":   schema_pkg_apis_app_v1_RoleMapperAuthConfig(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthClient":          schema_pkg_apis_app_v1_SSOAuthClient(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthConfig":          schema_pkg_apis_app_v1_SSOAuthConfig(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SecuredKieAppObject":    schema_pkg_apis_app_v1_SecuredKieAppObject(ref),
		"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.WebhookSecret":          schema_pkg_apis_app_v1_WebhookSecret(ref),
	}
}

func schema_pkg_apis_app_v1_CommonConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CommonConfig variables used in the templates",
				Properties: map[string]spec.Schema{
					"applicationName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imageTag": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"keyStorePassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminUser": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"dbPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"amqPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"amqClusterPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"controllerPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"serverPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mavenPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Condition - The condition for the kie-cloud-operator",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_app_v1_DatabaseObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatabaseObject Defines how a KieServer will manage and create a new Database or connect to an existing one",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"size": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"externalConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.ExternalDatabaseObject"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.ExternalDatabaseObject"},
	}
}

func schema_pkg_apis_app_v1_Deployments(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Deployments ...",
				Properties: map[string]spec.Schema{
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments are ready to serve requests",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"starting": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments are starting, may or may not succeed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"stopped": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments are not starting, unclear what next step will be",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"failed": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments failed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_ExternalDatabaseObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExternalDatabaseObject configuration definition of an external database",
				Properties: map[string]spec.Schema{
					"driver": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"dialect": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"host": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"jdbcURL": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"nonXA": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"jndiName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"username": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"password": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"minPoolSize": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"maxPoolSize": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"connectionChecker": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"exceptionSorter": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"backgroundValidation": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"backgroundValidationMillis": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_GitSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GitSource Git coordinates to locate the source code to build",
				Properties: map[string]spec.Schema{
					"uri": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"reference": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"contextDir": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_KieApp(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieApp is the Schema for the kieapps API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppSpec", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1_KieAppAuthObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppAuthObject Authentication specification to be used by the KieApp",
				Properties: map[string]spec.Schema{
					"sso": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthConfig"),
						},
					},
					"ldap": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.LDAPAuthConfig"),
						},
					},
					"roleMapper": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.RoleMapperAuthConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.LDAPAuthConfig", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.RoleMapperAuthConfig", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthConfig"},
	}
}

func schema_pkg_apis_app_v1_KieAppBuildObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppBuildObject Data to define how to build an application from source",
				Properties: map[string]spec.Schema{
					"kieServerContainerDeployment": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"gitSource": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.GitSource"),
						},
					},
					"mavenMirrorURL": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"artifactDir": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"webhooks": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.WebhookSecret"),
									},
								},
							},
						},
					},
					"from": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.GitSource", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.WebhookSecret", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_app_v1_KieAppObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppObject Generic object definition",
				Properties: map[string]spec.Schema{
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"keystoreSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"resources"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_app_v1_KieAppObjects(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppObjects KIE App deployment objects",
				Properties: map[string]spec.Schema{
					"console": {
						SchemaProps: spec.SchemaProps{
							Description: "Business Central container configs",
							Ref:         ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SecuredKieAppObject"),
						},
					},
					"servers": {
						SchemaProps: spec.SchemaProps{
							Description: "KIE Server configuration for individual sets",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieServerSet"),
									},
								},
							},
						},
					},
					"smartRouter": {
						SchemaProps: spec.SchemaProps{
							Description: "SmartRouter container configs",
							Ref:         ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppObject"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppObject", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieServerSet", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SecuredKieAppObject"},
	}
}

func schema_pkg_apis_app_v1_KieAppRegistry(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppRegistry defines the registry that should be used for rhpam images",
				Properties: map[string]spec.Schema{
					"registry": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"insecure": {
						SchemaProps: spec.SchemaProps{
							Description: "Registry to use, can also be set w/ \"REGISTRY\" env variable",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"insecure"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_KieAppSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppSpec defines the desired state of KieApp",
				Properties: map[string]spec.Schema{
					"environment": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file KIE environment type to deploy (prod, authoring, trial, etc)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imageRegistry": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppRegistry"),
						},
					},
					"objects": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppObjects"),
						},
					},
					"commonConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.CommonConfig"),
						},
					},
					"auth": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppAuthObject"),
						},
					},
					"upgrades": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppUpgrades"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.CommonConfig", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppAuthObject", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppObjects", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppRegistry", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppUpgrades"},
	}
}

func schema_pkg_apis_app_v1_KieAppStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppStatus - The status for custom resources managed by the operator-sdk.",
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.Condition"),
									},
								},
							},
						},
					},
					"consoleHost": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"deployments": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.Deployments"),
						},
					},
				},
				Required: []string{"conditions", "deployments"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.Condition", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.Deployments"},
	}
}

func schema_pkg_apis_app_v1_KieAppUpgrades(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieAppUpgrades KIE App product upgrade flags",
				Properties: map[string]spec.Schema{
					"patch": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"minor": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_KieServerSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KieServerSet KIE Server configuration for a single set, or for multiple sets if deployments is set to >1",
				Properties: map[string]spec.Schema{
					"deployments": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of KieServer DeploymentConfigs (defaults to 1)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"id": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"from": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"build": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppBuildObject"),
						},
					},
					"ssoClient": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthClient"),
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"keystoreSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"database": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.DatabaseObject"),
						},
					},
					"jms": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppJmsObject"),
						},
					},
				},
				Required: []string{"resources"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.DatabaseObject", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppBuildObject", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.KieAppJmsObject", "github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthClient", "k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ObjectReference", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_app_v1_LDAPAuthConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LDAPAuthConfig Authentication configuration for LDAP",
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bindDN": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bindCredential": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"jaasSecurityDomain": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"baseCtxDN": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"baseFilter": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"searchScope": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"searchTimeLimit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"distinguishedNameAttribute": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"parseUsername": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"usernameBeginString": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"usernameEndString": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"roleAttributeID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"rolesCtxDN": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"roleFilter": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"roleRecursion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"defaultRole": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"roleNameAttributeID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"parseRoleNameFromDN": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"roleAttributeIsDN": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"referralUserAttributeIDToCheck": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_RoleMapperAuthConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RoleMapperAuthConfig Configuration for RoleMapper Authentication",
				Properties: map[string]spec.Schema{
					"rolesProperties": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"replaceRole": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_SSOAuthClient(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SSOAuthClient Auth client to use for the SSO integration",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"hostnameHTTP": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"hostnameHTTPS": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_SSOAuthConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SSOAuthConfig Authentication configuration for SSO",
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"realm": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminUser": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"disableSSLCertValidation": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"principalAttribute": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1_SecuredKieAppObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecuredKieAppObject Generic object definition",
				Properties: map[string]spec.Schema{
					"ssoClient": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthClient"),
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"keystoreSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"resources"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kie-cloud-operator/pkg/apis/app/v1.SSOAuthClient", "k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_app_v1_WebhookSecret(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WebhookSecret Secret to use for a given webhook",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
