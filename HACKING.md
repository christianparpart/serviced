# Requirements

- CRUD deployment environments (aka namespaces)
- CRUD services (aka apps)
  - a service is defined by:
    - globally unique id (required for dependency graphs)
    - name
    - json blob for Marathon to describe the app
    - state:
      - list of configurations (0/1 active, 0+ inactive/pasts)
  - you can create a service without taking it live (instance count = 0)
  - deploy a service only if health checks > 0
  - rolling back
  - ability to create virtual services in order to keep service dependency graph complete
- client authentication via OAuth (must be part of a Github Organization(s))

# HTTP API

```
# App Management
GET /apps                             - list all available apps
GET /apps/:name                       - gets an app with all its available releases
POST /apps/:name                      - create or update an app
POST /apps/:name/:tag                 - create (register) a new app release 
DELETE /apps/:name/:tag               - delete an app release
DELETE /apps/:name                    - delete an app with all its releases & deploys

# Environment Management
GET /environments                     - list all available environments
POST /environments/:name              - create or update an env
DELETE /environments/:name            - delete an environment

# Deploy Management
GET /deployments/:env                 - list all apps in an :env
GET /deployments/:env/:app            - list app with its available releases
POST /deployments/:env/:app/:release  - deploy a given :app :release in :env

GET /environments/:name               - list all services
POST /environments/:name/:app/:tag    - deploy an app:tag
```











# Model
```
Service {
  name: string
  application_protocol: string
  transport_protocol: string
  service_ports: int[]
  health_checks: HealthCheck[]
  requires: Service[]
  uses: Service[]
}

VirtualService < Service {
}

DockerService < Service {
  image: String
  force_pull: Bool
  container_ports: int[]
}

NativeService < Service {
  deploy_cmd: Text
  activate_cmd: Text
  purge_cmd: Text
}

HealthCheck {
  name: String
  grace_period: Duration
  interval: Duration
  failure_threshold: Int
}

HealthCheckHttp {
}

HealthCheckCommand {
}

```
