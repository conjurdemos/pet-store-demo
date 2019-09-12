# k8s-conjur-deploy
The `./start` script at this folder deploys k8s Conjur environment by calling `./start` of project `kubernetes-conjur-deploy`. 
Afterward it loads relevant policies and intialize conjur CA. On completion it is possible to run `pet-store-demo` 
from root-project folder with `./bin/build` and `./bin/deploy`.

### Configuration
Configure file `bootstrap.env`.

### Running
Take care that docker repo has image of `cyberark-secrets-provider-for-k8s`. Otherwise build project `cyberark-secrets-provider-for-k8s`. 

Run `./start` to raise Conjur environment for k8s, load relevant policies and initialize Conjur CA.
