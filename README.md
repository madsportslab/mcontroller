# Updater
Generic update service for Linux

## About

Updater is a service that facilitates the downloading of an update package from a registered repository and performs installation of that package on Linux.  This allows for independent updates of services. 

1.  POST /api/update which should include a hash of the update package
1.  update package is then pulled from a registered service
1.  package is installed and service restarted
1.  rollbacks are not allowed for updates, but data is backed up so that state can be rebuilt
