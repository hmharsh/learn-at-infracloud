## Start
- install HELM from instructions listed in Doc

## Commands
- helm search hub  // list available helm hub
- helm search hub wordpress // List search results from publicly available helm hubs
- helm repo ls // list extra Added repo 
- helm search repo // list all charts available in all added repo

- helm search repo argo // List all charts available to install in specific repo

ex:  **helm install harshit-mariadb stable/mariadb** // will just start the installation, not wait till it finish
stable: repo added
mariadb: chart name to install  
harshit-mariadb: release name (a unique name for installation)
// use --generate-name flag to let helm generate the unique name 

- helm list // list the deployed/installed charts 
- helm status harshit-mariadb // initial line of output show status of overall release (if it is RUNNING or not)
- helm show values stable/mariadb // check helm chart default values (configuration parameters) before to install 
- echo '{mariadbUser: user0, mariadbDatabase: user0db}' > config.yaml && helm install **-f config.yaml** stable/mariadb --generate-name // to override configuration Values parameters at time of installation 
**Two ways to override configuration 1. --values (or -f), 2. --set**
```
If both are used, --set values are merged into --values with higher precedence. 
Overrides specified with --set are persisted in a ConfigMap. 
Values that have been --set can be viewed for a given release with helm get values <release-name>. 
Values that have been --set can be cleared by running helm upgrade with --reset-values specified.
```
- helm upgrade // When a new version of a chart is released, **or** when you want to change the configuration of your release
- helm upgrade -f new_values.yaml harshit-mariadb stable/mariadb // to update the configuration values
- helm get // command is a useful tool for looking at a release in the cluster.
- helm rollback harshit-mariadb 1 // helm rollback [RELEASE] [REVISION], get this revision value from helm history
- helm history harshit-mariadb

