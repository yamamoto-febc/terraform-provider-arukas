# terraform-provider-arukas

Terraform provider for Arukas. - `Terraform for Arukas`

## Installation

1. Download the plugin from the [releases page](https://github.com/yamamoto-febc/terraform-provider-arukas/releases/latest)
2. Put it in the same directory as the terraform binary. ex:`$GOPATH/bin/`.


## Usage

  - [Provider Configuration](#provider-configuration)
  - Resource Configuration
    - [arukas_container](#resource-configuration-arukas_container)
  - [Samples](#samples)

## Provider Configuration

### Example

```
provider "arukas" {
    token = "your API token"
    secret = "your API secret"
}
```
    
### Argument Reference

The following arguments are supported:

* `token` - (Required) This is the Arukas API token. This can also be specified
  with the `ARUKAS_JSON_API_TOKEN` shell environment variable.

* `secret` - (Required) This is the Arukas API secret. This can also be specified
  with the `ARUKAS_JSON_API_SECRET` shell environment variable.

* `api_url` - (Optional) This is the ArukasAPI root URL. This can also be specified
  with the `ARUKAS_JSON_API_URL` shell environment variable.

* `trace` - (Optional) Flag of trace mode. This can also be specified
  with the `ARUKAS_DEBUG` shell environment variable.

## Resource Configuration `arukas_container`

Provides a Arukas container resource. This can be used to create, modify,
and delete container.

### Example Usage

```
resource "arukas_container" "foobar" {
    name = "terraform_for_arukas_test_foobar"
    image = "nginx:latest"
    instances = 2
    memory = 512
    endpoint = "terraform-for-arukas-test-endpoint-upd"
    ports = {
        protocol = "tcp"
        number = "80"
    }
    ports = {
        protocol = "tcp"
        number = "443"
    }
    environments {
        key = "key"
        value = "value"
    }
    environments {
        key = "key2"
        value = "value2"
    }
    cmd = "/foo/bar.sh"
}
```

### Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the App.
* `image` - (Required) The name of the DockerImage.Specify an image that exists in `https://hub.docker.com/`.
* `instances` - (Optional) The number of the container instance. Must be between `1` and `10`. Default is `1`.
* `memory` - (Optional) The size of the memory. Must be in [`256` , `512`]. Default is `256`.
* `endpoint` - (Optional) The endpoint of the app.
* `timeout` - (Optional) The number of seconds for waiting boot.Default is `300`
* `ports` - (Required) The ports of the app.
  * `protocol` - The protocol of the port.Must be in [`tcp` , `udp`]
  * `number` - The number of the port. Must be between `1` and `65535`.
* `environments` - (Optional) The description of the switch.
  * `key` - The key of environment variable.
  * `value` - The value of environment variable.
* `cmd` - (Optional) The description of the switch.

### Attributes Reference

The following attributes are exported:

* `id` - The ID of container.
* `app_id` - The ID of app.
* `name` - The name of the App.
* `image` - The name of the DockerImage.Specify an image that exists in `https://hub.docker.com/`.
* `instances` - The number of the container instance. Must be between `1` and `10`. Default is `1`.
* `memory` - The size of the memory. Must be in [`256` , `512`]. Default is `256`.
* `endpoint` - The endpoint of the app.
* `ports` - The ports of the app.
  * `protocol` - The protocol of the port.Must be in [`tcp` , `udp`]
  * `number` - The number of the port. Must be between `1` and `65535`.
* `environments` - The description of the switch.
  * `key` - The key of environment variable.
  * `value` - The value of environment variable.
* `cmd` - The description of the switch.

* `endpoint_full_hostname` - The full name of endpoint host.
* `endpoint_full_url` - The URL of endpoint.
* `port_mappings` - The mappings of service port.
  * `host` -  The hostname of thr service host.
  * `container_port` - The number of port was container exposed.
  * `service_port` - The number of service port.

## License

  This project is published under [Apache 2.0 License](LICENSE).

## Author

  * Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
