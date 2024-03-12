# Traefik Header Based Proxy Plugin

## Overview
This plugin extends the functionality of Traefik, a modern reverse proxy and load balancer for microservices. 
Traefik provides automatic discovery of services, dynamic configuration, and flexible routing capabilities. 

The Traefik Header Based Proxy Plugin enhances these features by offering a redirection option based on header values.

## Features
- **Filter based on reqquest headers**: Make a proxy redirect based on the value of specified header keys.

## Usage

You have to install the plugin in the traefik definition like:

**Ex: Using Docker Labels**:
 - 'traefik.http.middlewares.my-proxy.plugin.proxy.Header=X-Header-Custom'
 - 'traefik.http.middlewares.my-proxy.plugin.proxy.Mapping.value1=https://service1'
 - 'traefik.http.middlewares.my-proxy.plugin.proxy.Mapping.value2=https://service2'
 - "traefik.http.routers.my-router.middlewares=my-proxy"

## Contributing
We welcome contributions from the community to enhance the functionality and usability of the Traefik Public Plugin. If you'd like to contribute, please follow these guidelines:

- **Fork the Repository**: Create your own fork of the repository to make changes.
- **Create a Branch**: Create a new branch for your changes and work on them independently.
- **Submit a Pull Request**: Once your changes are ready, submit a pull request to merge them into the main repository.
- **Follow Coding Standards**: Adhere to coding standards, conventions, and best practices to maintain code quality.

## License
The Traefik Header Based Proxy Plugin is released under the Apache 2 Licence.