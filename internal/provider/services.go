package provider

import (
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/aadb2c"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/advisor"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/analysisservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/apimanagement"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/appconfiguration"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/applicationinsights"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/appservice"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/attestation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/authorization"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/automation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/azurestackhci"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/batch"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/billing"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/blueprints"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/bot"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cdn"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cognitive"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/communication"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/compute"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/confidentialledger"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/connections"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/consumption"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/containers"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cosmos"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/costmanagement"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/customproviders"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/databasemigration"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/databoxedge"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/databricks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/datafactory"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/dataprotection"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/datashare"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/desktopvirtualization"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/devtestlabs"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/digitaltwins"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/disks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/dns"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/domainservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/elastic"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/eventgrid"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/eventhub"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/firewall"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/fluidrelay"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/frontdoor"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/hdinsight"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/healthcare"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/hpccache"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/hsm"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iotcentral"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iothub"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iottimeseriesinsights"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/keyvault"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/kusto"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/legacy"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/lighthouse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/loadbalancer"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/loadtest"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/loganalytics"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/logic"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/logz"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/machinelearning"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/maintenance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/managedapplications"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/managementgroup"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/maps"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mariadb"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/media"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mixedreality"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/monitor"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/msi"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mssql"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mysql"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/netapp"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/network"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/notificationhub"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/policy"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/portal"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/postgres"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/powerbi"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/privatedns"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/purview"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/recoveryservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/redis"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/redisenterprise"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/relay"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/resource"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/search"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/securitycenter"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/sentinel"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicebus"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabric"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/signalr"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/springcloud"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/sql"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/storage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/streamanalytics"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/subscription"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/synapse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/trafficmanager"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/videoanalyzer"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/vmware"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/web"
)

//go:generate go run ../tools/generator-services/main.go -path=../../

func SupportedTypedServices() []sdk.TypedServiceRegistration {
	return []sdk.TypedServiceRegistration{
		aadb2c.Registration{},
		apimanagement.Registration{},
		appconfiguration.Registration{},
		applicationinsights.Registration{},
		appservice.Registration{},
		batch.Registration{},
		bot.Registration{},
		compute.Registration{},
		consumption.Registration{},
		containers.Registration{},
		costmanagement.Registration{},
		disks.Registration{},
		domainservices.Registration{},
		eventhub.Registration{},
		fluidrelay.Registration{},
		keyvault.Registration{},
		loadbalancer.Registration{},
		loadtest.Registration{},
		mssql.Registration{},
		policy.Registration{},
		resource.Registration{},
		sentinel.Registration{},
		servicefabricmanaged.Registration{},
		streamanalytics.Registration{},
		web.Registration{},
	}
}

func SupportedUntypedServices() []sdk.UntypedServiceRegistration {
	return func() []sdk.UntypedServiceRegistration {
		out := []sdk.UntypedServiceRegistration{
			advisor.Registration{},
			analysisservices.Registration{},
			apimanagement.Registration{},
			appconfiguration.Registration{},
			springcloud.Registration{},
			applicationinsights.Registration{},
			attestation.Registration{},
			authorization.Registration{},
			automation.Registration{},
			azurestackhci.Registration{},
			batch.Registration{},
			billing.Registration{},
			blueprints.Registration{},
			bot.Registration{},
			cdn.Registration{},
			cognitive.Registration{},
			communication.Registration{},
			compute.Registration{},
			confidentialledger.Registration{},
			connections.Registration{},
			containers.Registration{},
			consumption.Registration{},
			cosmos.Registration{},
			customproviders.Registration{},
			databricks.Registration{},
			datafactory.Registration{},
			databasemigration.Registration{},
			databoxedge.Registration{},
			dataprotection.Registration{},
			datashare.Registration{},
			desktopvirtualization.Registration{},
			devtestlabs.Registration{},
			digitaltwins.Registration{},
			dns.Registration{},
			domainservices.Registration{},
			elastic.Registration{},
			eventgrid.Registration{},
			eventhub.Registration{},
			firewall.Registration{},
			frontdoor.Registration{},
			hpccache.Registration{},
			hsm.Registration{},
			hdinsight.Registration{},
			healthcare.Registration{},
			iothub.Registration{},
			iotcentral.Registration{},
			keyvault.Registration{},
			kusto.Registration{},
			legacy.Registration{},
			loadbalancer.Registration{},
			loganalytics.Registration{},
			logic.Registration{},
			logz.Registration{},
			machinelearning.Registration{},
			maintenance.Registration{},
			managedapplications.Registration{},
			lighthouse.Registration{},
			managementgroup.Registration{},
			maps.Registration{},
			mariadb.Registration{},
			media.Registration{},
			mixedreality.Registration{},
			monitor.Registration{},
			msi.Registration{},
			mssql.Registration{},
			mysql.Registration{},
			netapp.Registration{},
			network.Registration{},
			notificationhub.Registration{},
			policy.Registration{},
			portal.Registration{},
			postgres.Registration{},
			powerbi.Registration{},
			privatedns.Registration{},
			purview.Registration{},
			recoveryservices.Registration{},
			redis.Registration{},
			redisenterprise.Registration{},
			relay.Registration{},
			resource.Registration{},
			search.Registration{},
			securitycenter.Registration{},
			sentinel.Registration{},
			servicebus.Registration{},
			servicefabric.Registration{},
			signalr.Registration{},
			sql.Registration{},
			storage.Registration{},
			streamanalytics.Registration{},
			subscription.Registration{},
			synapse.Registration{},
			iottimeseriesinsights.Registration{},
			trafficmanager.Registration{},
			videoanalyzer.Registration{},
			vmware.Registration{},
			web.Registration{},
		}
		return out
	}()
}
