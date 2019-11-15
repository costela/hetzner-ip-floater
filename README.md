![Docker Automated build](https://img.shields.io/docker/automated/costela/hetzner-ip-floater.svg)
![Docker Build Status](https://img.shields.io/docker/build/costela/hetzner-ip-floater.svg)
![Image Info](https://images.microbadger.com/badges/image/costela/hetzner-ip-floater.svg)

# ⚠ Deprecated in favor of [hcloud-ip-floater](/costela/hcloud-ip-floater)

# hetzner-ip-floater

Minimalistic floating IP setter for container clusters (currently tested on docker swarm) running on [Hetzner Cloud](https://www.hetzner.com/cloud).

## Usage

This project should be used as a container. It will run once and update a given floating IP to point to the node currently running it. This can be used to ensure the floating IP will be reassigned upon node failure, by relying on the underlying cluster to redeploy this service on a healthy node.

This is an example deployment for `docker stack deploy`:
```yaml
services:
  app:
    image: costela/hetzner-ip-floater
    secrets:
      - hetzner_api_key_for_floating_ip  # set via `docker secret create`
    environment:
      API_KEY_FILE: /run/secrets/hetzner_api_key_for_floating_ip
      TARGET_HOST: '{{ .Node.Hostname }}'  # uses docker swarm's templating to get node name
      FLOATING_IP_ID: 12345  # taken from Hetzner cloud console
    deploy:
      replicas: 1
```

This assumes the node's hostnames are the same as their API names, which is the case unless the hostname has been changed after provisioning.

## Other considerations

The deployment of `hetzner-ip-floater` should be limited to those nodes where the floating IP is locally configured, otherwise incoming trafic will be dropped.

The same nodes should also be configured as ingress nodes. When using the [default mesh networking](https://docs.docker.com/engine/swarm/ingress/) on docker swarm, this is already the case for all worker nodes.
