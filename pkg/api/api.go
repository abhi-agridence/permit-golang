package api

import (
	"github.com/permitio/permit-golang/openapi"
	permit "github.com/permitio/permit-golang/pkg/permit"
	"go.uber.org/zap"
)

type PermitBaseApi struct {
	client *openapi.APIClient
	config *permit.PermitConfig
	logger *zap.Logger
}

type PermitApiClient struct {
	config             *permit.PermitConfig
	logger             *zap.Logger
	client             *openapi.APIClient
	tenants            *Tenants
	environments       *Environments
	projects           *Projects
	resourceActions    *ResourceActions
	resourceAttributes *ResourceAttributes
	resources          *Resources
	roles              *Roles
	users              *Users
	Elements           *Elements
}

func NewPermitApiClient(config *permit.PermitConfig) *PermitApiClient {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	return &PermitApiClient{
		config:             config,
		logger:             config.logger,
		client:             client,
		tenants:            NewTenantsApi(client, config),
		environments:       NewEnvironmentsApi(client, config),
		projects:           NewProjectsApi(client, config),
		resourceActions:    NewResourceActionsApi(client, config),
		resourceAttributes: NewResourceAttributesApi(client, config),
		resources:          NewResourcesApi(client, config),
		roles:              NewRolesApi(client, config),
		users:              NewUsersApi(client, config),
		Elements:           NewElementsApi(client, config),
	}
}
