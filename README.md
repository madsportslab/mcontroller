# mcontroller
Generic update service for Linux

## About

mcontroller is a heavily opinionated service that facilitates common operations for linux appliances like upgrades, reboot, shutdown, etc.  all of these operations could be performed with a terminal, however, the goal is to package these operations for applications via a RESTful API which would allow for simpler access.  imagine being able to reboot an appliance with a mobile app, you wouldn't need SSH libraries or need knowledge of how to operate bash.  mcontroller targets appliances running linux that have some sort of application on top as opposed to generic linux servers, but i assume that mcontroller could be applied to generic linux servers as well.

## Upgrades

mcontroller downloads update packages (debian) from a registered repository and performs installation of that package.  This allows for independent updates of services.

1.  POST /api/update which should include a hash of the update package
1.  update package is then pulled from a registered service
1.  package is installed and service restarted
1.  rollbacks are not allowed for updates, but data is backed up so state can be rebuilt if necessary

## config.json

{
  "repository": "10.241.0.1:8110",
  "mboard": "192.168.1.1:8000"
}
