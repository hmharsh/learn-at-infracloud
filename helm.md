## To do
- helm lint
- helm charemeusium lib
- helm dependencies

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

# Commands
```
 
-- Repo Management -- 
helm repo add stable https://kubernetes-charts.storage.googleapis.com/
helm repo list or helm repo ls
helm repo update
helm repo remove stable

-- searching chart --
// search is from substring
helm search repo mysql  (search for string stable in available repos)
helm search repo stable  (stable is there in all rows, it will show all charts under stable repo)
helm search hub etcd (search on helm hub for specified string)


-- Checking the chart --
helm show values stable/mariadb // check the chart values available before installing, just like a doc or readme
helm pull chartrepo/chartname // Download a chart and check the files without installing it


-- Chart management --
helm install stable/mysql --generate-name // with random release name
helm install happy-panda stable/mariadb // Install with release name
helm install -f config.yaml stable/mariadb --generate-name // pass values by file
helm install --set key=value stable/mariadb --generate-name


helm ls or helm list
helm list --uninstalled
helm list --all

helm status happy-panda //To check the installation notes later

helm get values happy-panda // shoow the values supplied at runrime

helm uninstall smiling-penguin //will completly delete
helm uninstall --keep-history

helm package deis-workflow // package chart inro compressed formate for sharing
helm dep update ./somechart // installing the dependencies as listed in requirements.yaml


-- Upgrade, Rollback --
- helm upgrade -f panda.yaml happy-panda stable/mariadb // to update chart to new version, or to update values, supply same release name the one which we want to update, It will only update things that have changed since the last release.

- helm rollback happy-panda 1 // 1 is revision, easily found in helm ls or helm status happy-panda, or history command as below
- helm history happy-panda // check all older versions of specific repo
```

# NOTES
  - '--set' values merged into --values
  - Values that have been --set can be viewed for a given release with helm get values <release-name>
   Values that have been --set can be cleared by running helm upgrade with --reset-values specified.

