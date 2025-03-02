// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDataFusionInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataFusionInstanceCreate,
		Read:   resourceDataFusionInstanceRead,
		Update: resourceDataFusionInstanceUpdate,
		Delete: resourceDataFusionInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataFusionInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(25 * time.Minute),
			Delete: schema.DefaultTimeout(50 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the instance or a fully qualified identifier for the instance.`,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"BASIC", "ENTERPRISE", "DEVELOPER"}),
				Description: `Represents the type of Data Fusion instance. Each type is configured with
the default settings for processing and memory.
- BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
using point and click UI. However, there are certain limitations, such as fewer number
of concurrent pipelines, no support for streaming pipelines, etc.
- ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
- DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration 
pipelines at low cost. Possible values: ["BASIC", "ENTERPRISE", "DEVELOPER"]`,
			},
			"dataproc_service_account": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of the instance.`,
			},
			"enable_stackdriver_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Option to enable Stackdriver Logging.`,
			},
			"enable_stackdriver_monitoring": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Option to enable Stackdriver Monitoring.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The resource labels for instance to use to annotate any related underlying resources,
such as Compute Engine VMs.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"network_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Network configuration options. These are required when a private Data Fusion instance is to be created.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_allocation": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The IP range in CIDR notation to use for the managed Data Fusion instance
nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.`,
						},
						"network": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `Name of the network in the project with which the tenant project
will be peered for executing pipelines. In case of shared VPC where the network resides in another host
project the network should specified in the form of projects/{host-project-id}/global/networks/{network}`,
						},
					},
				},
			},
			"options": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Map of additional options used to configure the behavior of Data Fusion instance.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"private_instance": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Specifies whether the Data Fusion instance should be private. If set to
true, all Data Fusion nodes will have private IP addresses and will not be
able to access the public internet.`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the Data Fusion instance.`,
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Current version of the Data Fusion.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.`,
			},
			"service_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Service account which will be used to access resources in the customer project.`,
			},
			"service_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Endpoint on which the Data Fusion UI and REST APIs are accessible.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of this Data Fusion instance.
- CREATING: Instance is being created
- RUNNING: Instance is running and ready for requests
- FAILED: Instance creation failed
- DELETING: Instance is being deleted
- UPGRADING: Instance is being upgraded
- RESTARTING: Instance is being restarted`,
			},
			"state_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Additional information about the current state of this Data Fusion instance if available.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDataFusionInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandDataFusionInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandDataFusionInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	typeProp, err := expandDataFusionInstanceType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	enableStackdriverLoggingProp, err := expandDataFusionInstanceEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !isEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableStackdriverMonitoringProp, err := expandDataFusionInstanceEnableStackdriverMonitoring(d.Get("enable_stackdriver_monitoring"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_monitoring"); !isEmptyValue(reflect.ValueOf(enableStackdriverMonitoringProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverMonitoringProp)) {
		obj["enableStackdriverMonitoring"] = enableStackdriverMonitoringProp
	}
	labelsProp, err := expandDataFusionInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	optionsProp, err := expandDataFusionInstanceOptions(d.Get("options"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("options"); !isEmptyValue(reflect.ValueOf(optionsProp)) && (ok || !reflect.DeepEqual(v, optionsProp)) {
		obj["options"] = optionsProp
	}
	versionProp, err := expandDataFusionInstanceVersion(d.Get("version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version"); !isEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	privateInstanceProp, err := expandDataFusionInstancePrivateInstance(d.Get("private_instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_instance"); !isEmptyValue(reflect.ValueOf(privateInstanceProp)) && (ok || !reflect.DeepEqual(v, privateInstanceProp)) {
		obj["privateInstance"] = privateInstanceProp
	}
	dataprocServiceAccountProp, err := expandDataFusionInstanceDataprocServiceAccount(d.Get("dataproc_service_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dataproc_service_account"); !isEmptyValue(reflect.ValueOf(dataprocServiceAccountProp)) && (ok || !reflect.DeepEqual(v, dataprocServiceAccountProp)) {
		obj["dataprocServiceAccount"] = dataprocServiceAccountProp
	}
	networkConfigProp, err := expandDataFusionInstanceNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_config"); !isEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}

	url, err := replaceVars(d, config, "{{DataFusionBasePath}}projects/{{project}}/locations/{{region}}/instances?instanceId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = dataFusionOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	if err := d.Set("name", flattenDataFusionInstanceName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceDataFusionInstanceRead(d, meta)
}

func resourceDataFusionInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DataFusionBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DataFusionInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("name", flattenDataFusionInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("description", flattenDataFusionInstanceDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("type", flattenDataFusionInstanceType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("enable_stackdriver_logging", flattenDataFusionInstanceEnableStackdriverLogging(res["enableStackdriverLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("enable_stackdriver_monitoring", flattenDataFusionInstanceEnableStackdriverMonitoring(res["enableStackdriverMonitoring"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenDataFusionInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("options", flattenDataFusionInstanceOptions(res["options"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenDataFusionInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("update_time", flattenDataFusionInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenDataFusionInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state_message", flattenDataFusionInstanceStateMessage(res["stateMessage"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("service_endpoint", flattenDataFusionInstanceServiceEndpoint(res["serviceEndpoint"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("version", flattenDataFusionInstanceVersion(res["version"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("service_account", flattenDataFusionInstanceServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("private_instance", flattenDataFusionInstancePrivateInstance(res["privateInstance"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("dataproc_service_account", flattenDataFusionInstanceDataprocServiceAccount(res["dataprocServiceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("network_config", flattenDataFusionInstanceNetworkConfig(res["networkConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceDataFusionInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	enableStackdriverLoggingProp, err := expandDataFusionInstanceEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableStackdriverMonitoringProp, err := expandDataFusionInstanceEnableStackdriverMonitoring(d.Get("enable_stackdriver_monitoring"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_monitoring"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableStackdriverMonitoringProp)) {
		obj["enableStackdriverMonitoring"] = enableStackdriverMonitoringProp
	}
	labelsProp, err := expandDataFusionInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "{{DataFusionBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = dataFusionOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceDataFusionInstanceRead(d, meta)
}

func resourceDataFusionInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{DataFusionBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	err = dataFusionOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceDataFusionInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/instances/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataFusionInstanceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenDataFusionInstanceDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceEnableStackdriverLogging(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceEnableStackdriverMonitoring(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceOptions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceStateMessage(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceServiceEndpoint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceServiceAccount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstancePrivateInstance(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceDataprocServiceAccount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceNetworkConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ip_allocation"] =
		flattenDataFusionInstanceNetworkConfigIpAllocation(original["ipAllocation"], d, config)
	transformed["network"] =
		flattenDataFusionInstanceNetworkConfigNetwork(original["network"], d, config)
	return []interface{}{transformed}
}
func flattenDataFusionInstanceNetworkConfigIpAllocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataFusionInstanceNetworkConfigNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandDataFusionInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
}

func expandDataFusionInstanceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableStackdriverLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableStackdriverMonitoring(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataFusionInstanceOptions(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataFusionInstanceVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstancePrivateInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceDataprocServiceAccount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAllocation, err := expandDataFusionInstanceNetworkConfigIpAllocation(original["ip_allocation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAllocation); val.IsValid() && !isEmptyValue(val) {
		transformed["ipAllocation"] = transformedIpAllocation
	}

	transformedNetwork, err := expandDataFusionInstanceNetworkConfigNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !isEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	return transformed, nil
}

func expandDataFusionInstanceNetworkConfigIpAllocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfigNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
