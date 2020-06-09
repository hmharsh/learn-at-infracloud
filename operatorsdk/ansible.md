# Installation 
- install ansible with apt or yum 2.9+
- install ansible-runnerv1.1.0+ (with pip or dnf)
- install poerator sdk (https://sdk.operatorframework.io/docs/install-operator-sdk/)

# Generate all files
- operator-sdk new harshit-operator --api-version=harshit.example.com/v1alpha1 --kind=Harshit --type=ansible

# Update the operator code
-  Edir CRD (for all input) and modify role accordingly with same variable input 
-  create more CR with the same operator: operator-sdk add api --api-version=mahajan.example.com/v1alpha1 --kind=Mahajan

# Run the operator
-   kubectl create -f deploy/crds/harshit.example.com_harshit_crd.yaml        (create crd, to gnerate API endpoint)      
-   operator-sdk build hmharsh3/myoperator:1                                  (create docker image for the operator)
-   docker push hmharsh3/myoperator:1                                         (push the operator image)
-   sed -i 's|REPLACE_IMAGE|hmharsh3/myoperator:1|g' deploy/operator.yaml     (update the image name operator descriptor to create background deployment crrosponding to the operator)
-   kubectl apply -f deploy/crds/harshit.example.com_v1alpha1_harshit_cr.yaml (now use the custom CR with custom variables, which will supplied to ansible as extra vars)

# Test operator 
- already pushed to docker hub
   `https://hub.docker.com/repository/docker/hmharsh3/operatorharshit`


- reference: https://sdk.operatorframework.io/docs/ansible/quickstart/
