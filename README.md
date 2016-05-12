# sample-app-pipeline

This is a sample of a Concourse pipeline for a Go application that performs the
following tasks:
1. builds (compiles) the binary artifact
1. unit tests it
1. deploys it to a Cloud Foundry staging space
1. load tests it
1. deploys it to a Cloud Foundry production space

Each step is configured to perform automatically only if the previous step has been successfully executed.

The pre-reqs for the pipeline deployment instructions below are as follows:
1. An instance of Concourse is installed either local as a vagrant machine or in a remote server. Please refer to the documentation on [how to install Concourse](http://concourse.ci/installing.html).
1. The Concourse Fly command line interface is installed on the local VM. The fly cli can be downloaded directly from the link provided on the Concourse web interface. Please refer to the [Fly cli documentation](http://concourse.ci/fly-cli.html) for details.

How to setup this sample pipeline on your Concourse server:

1. Clone this git repository on your local machine
...* ```git clone https://github.com/lsilvapvt/sample-app-pipeline.git```
...* ```cd sample-app-pipeline```
1. Setup the pipeline credentials file
...* ```cp ci/credentials.yml.sample ci/credentials.yml```
...* Edit _ci/credentials.yml_ and fill out all the required credentials.
....* _deploy-username:_ the userID to deploy apps on the Cloud Foundry deployment
....* _deploy-password:_ the corresponding password to deploy apps on the Cloud Foundry deployment
....* _pws-organization:_ the name of your organization in Cloud Foundry
....* _pws-staging-space:_ the name of the staging/development space to deploy the sample app to in CF
....* _pws-production-space:_ the name of the production space to deploy the sample app to in CF
....* _pws-api:_ the url of the CF API. (e.g. https://api.run.pivotal.io)
1. Configure the sample pipeline in Concourse
..* ```fly -t local login <concourse-url>``` . Example: ```fly -t local login http://192.168.100.4:8080```
..* ```fly -t local set-pipeline -c ci/pipeline.yml -p sample-app-pipeline -l ci/credentials.yml```
1. Access to the Concourse web interface (e.g. http://192.168.100.4:8080 ), click on the list of pipelines and un-pause the _sample-app-pipeline_ and click on it to visualize the pipeline diagram.

You will see the pipeline's tasks get execution, one at a time if the previous pipeline is executed successfully.
Notice that the pipeline is organized in two groups: _delivery_ and _deployment_, with corresponding links located at the top of the pipeline's diagram. The _delivery_ group contains the jobs associated with a typical build and test pipeline for development organizations. The _deployment_ group displays the job associated with the typical task of promoting a successful build from development into production.

Edit file _ci/pipeline.yml_ to inspect how this sample Concourse pipeline was defined and structured.

# Credits
This project was forked from the original work of [J.Calabrese's sample concourse app](https://github.com/xchapter7x/concourse-demo-app).
