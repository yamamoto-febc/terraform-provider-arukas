package arukas

import (
	"fmt"
	API "github.com/arukasio/cli"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
	"time"
)

func resourceArukasContainer() *schema.Resource {
	return &schema.Resource{
		Create: resourceArukasContainerCreate,
		Read:   resourceArukasContainerRead,
		Update: resourceArukasContainerUpdate,
		Delete: resourceArukasContainerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instances": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validateIntegerInRange(1, 10),
			},
			"memory": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      256,
				ValidateFunc: validateIntInWord([]string{"256", "512"}),
			},
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 20,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "tcp",
							ValidateFunc: validateStringInWord([]string{"tcp", "udp"}),
						},
						"number": &schema.Schema{
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      "80",
							ValidateFunc: validateIntegerInRange(1, 65535),
						},
					},
				},
			},
			"environments": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 20,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"cmd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"port_mappings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"container_port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"service_port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"endpoint_full_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoint_full_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceArukasContainerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*API.Client)

	var appSet API.AppSet

	// create an app
	newApp := API.App{Name: d.Get("name").(string)}

	var parsedEnvs API.Envs
	var parsedPorts API.Ports

	if rawEnvs, ok := d.GetOk("environments"); ok {
		parsedEnvs = expandEnvs(rawEnvs)
	}
	if rawPorts, ok := d.GetOk("ports"); ok {
		parsedPorts = expandPorts(rawPorts)
	}

	newContainer := API.Container{
		Envs:      parsedEnvs,
		Ports:     parsedPorts,
		ImageName: d.Get("image").(string),
		Mem:       d.Get("memory").(int),
		Instances: d.Get("instances").(int),
		Cmd:       d.Get("cmd").(string),
		Name:      d.Get("endpoint").(string),
	}
	newAppSet := API.AppSet{
		App:       newApp,
		Container: newContainer,
	}

	// create
	if err := client.Post(&appSet, "/app-sets", newAppSet); err != nil {
		return err
	}

	// start container
	if err := client.Post(nil, fmt.Sprintf("/containers/%s/power", appSet.Container.ID), nil); err != nil {
		return err
	}

	//wait for boot
	timeout := time.Duration(d.Get("timeout").(int)) * time.Second
	if err := sleepUntilUp(client, appSet.Container.ID, timeout); err != nil {
		return err
	}

	d.SetId(appSet.Container.ID)
	return resourceArukasContainerRead(d, meta)
}

func resourceArukasContainerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*API.Client)

	var container API.Container
	var app API.App

	if err := client.Get(&container, fmt.Sprintf("/containers/%s", d.Id())); err != nil {
		return err
	}
	if err := client.Get(&app, fmt.Sprintf("/apps/%s", container.AppID)); err != nil {
		return err
	}

	d.Set("app_id", container.AppID)
	d.Set("name", app.Name)
	d.Set("image", container.ImageName)
	d.Set("instances", container.Instances)
	d.Set("memory", container.Mem)
	endpoint := container.Endpoint
	if strings.HasSuffix(endpoint, ".arukascloud.io") {
		endpoint = strings.Replace(endpoint, ".arukascloud.io", "", -1)
	}

	d.Set("endpoint", endpoint)
	d.Set("endpoint_full_hostname", container.Endpoint)
	d.Set("endpoint_full_url", fmt.Sprintf("https://%s", container.Endpoint))

	d.Set("cmd", container.Cmd)

	//ports
	d.Set("ports", flattenPorts(container.Ports))

	//port mappings
	d.Set("port_mappings", flattenPortMappings(container.PortMappings))

	//envs
	d.Set("environments", flattenEnvs(container.Envs))

	return nil
}

func resourceArukasContainerUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*API.Client)
	var container API.Container

	if err := client.Get(&container, fmt.Sprintf("/containers/%s", d.Id())); err != nil {
		return err
	}

	var parsedEnvs API.Envs
	var parsedPorts API.Ports

	if rawEnvs, ok := d.GetOk("environments"); ok {
		parsedEnvs = expandEnvs(rawEnvs)
	}
	if rawPorts, ok := d.GetOk("ports"); ok {
		parsedPorts = expandPorts(rawPorts)
	}

	newContainer := API.Container{
		Envs:      parsedEnvs,
		Ports:     parsedPorts,
		ImageName: d.Get("image").(string),
		Mem:       d.Get("memory").(int),
		Instances: d.Get("instances").(int),
		Cmd:       d.Get("cmd").(string),
		Name:      d.Get("endpoint").(string),
	}

	// update
	if err := client.Patch(nil, fmt.Sprintf("/containers/%s", d.Id()), newContainer); err != nil {
		return err
	}

	return resourceArukasContainerRead(d, meta)

}

func resourceArukasContainerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*API.Client)
	var container API.Container

	if err := client.Get(&container, fmt.Sprintf("/containers/%s", d.Id())); err != nil {
		return err
	}

	if err := client.Delete(fmt.Sprintf("/apps/%s", container.AppID)); err != nil {
		return err
	}

	return nil
}

func sleepUntilUp(client *API.Client, containerID string, timeout time.Duration) error {
	current := 0 * time.Second
	interval := 5 * time.Second
	for {
		var container API.Container
		if err := client.Get(&container, fmt.Sprintf("/containers/%s", containerID)); err != nil {
			return err
		}

		if container.IsRunning {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: sleepUntilUp")
		}
	}
}
