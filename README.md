# Application CI pipeline targeting multiple CF spaces

This is an example of a Concourse pipeline that performs the following tasks for a sample _Golang_ application:

1. builds (compiles) the binary artifact

1. unit tests it

1. deploys it to a Cloud Foundry staging space

1. load tests it

1. deploys it to a Cloud Foundry production space

![Delivery pipeline][pipeline01]

Each step is configured to perform automatically only if the previous step has been successfully executed.

## Pre-requisites

The requirements for this pipeline's setup are as follows:

1. An instance of Concourse is installed either as a local vagrant machine or as a remote server.

   Please refer to the documentation on [how to install Concourse](http://concourse.ci/installing.html).

1. The Concourse Fly command line interface is installed on the local VM.

   The Fly cli can be downloaded directly from the link provided on the Concourse web interface.

   Please refer to the [Fly cli documentation](http://concourse.ci/fly-cli.html) for details.


## Pipeline setup and execution

How to setup this sample pipeline on your Concourse server:

1. Clone this git repository on your local machine  
  __git clone https://github.com/lsilvapvt/sample-app-pipeline.git __  
  __cd sample-app-pipeline__

1. Setup the pipeline credentials file  
  __cp ci/credentials.yml.sample ci/credentials.yml__  

  Edit _ci/credentials.yml_ and fill out all the required credentials:  
  _deploy-username:_ the userID to deploy apps on the Cloud Foundry deployment  
  _deploy-password:_ the corresponding password to deploy apps on the Cloud Foundry deployment  
  _pws-organization:_ the name of your organization in Cloud Foundry  
  _pws-staging-space:_ the name of the staging/development space to deploy the sample app to in CF  
  _pws-production-space:_ the name of the production space to deploy the sample app to in CF  
  _pws-api:_ the url of the CF API. (e.g. https://api.run.pivotal.io)  

1. Configure the sample pipeline in Concourse with the following commands:  
   __fly -t local login <concourse-url>__  

   Example:  
   __fly -t local login http://192.168.100.4:8080  __  
   __fly -t local set-pipeline -c ci/pipeline.yml -p sample-app-pipeline -l ci/credentials.yml__  

1. Access to the Concourse web interface (e.g. http://192.168.100.4:8080 ), click on the list of pipelines, un-pause the _sample-app-pipeline_ and then click on its link to visualize its pipeline diagram.

You will then notice the pipeline's jobs getting executed within a few seconds, one at a time, if the previous job in the pipeline is executed successfully.


## Notes

Notice that the pipeline is organized in two groups: _delivery_ and _deployment_, with corresponding links located at the top of the pipeline's diagram.

The _delivery_ group contains the jobs associated with a typical build and test pipeline for development organizations and/or a staging environment. See the pipeline image above.

The _deployment_ group displays the job associated with the typical task of promoting a successful build from development/staging into production.

![Deployment pipeline][pipeline02]

Edit file _ci/pipeline.yml_ to inspect how this sample Concourse pipeline was defined and structured.

## Read more

- [Blue-Green application deployment pipeline with Concourse](https://github.com/lsilvapvt/concourse-pipeline-samples/tree/master/blue-green-app-deployment)

- [Deploying Concourse on Bosh-lite](https://github.com/lsilvapvt/concourse-pipeline-samples/tree/master/concourse-on-bosh-lite)

- [Deploying Concourse on a Bosh 1.0 Director](https://github.com/lsilvapvt/concourse-pipeline-samples/tree/master/concourse-on-bosh-1.0)

- [Concourse pipelines with a local Docker Registry](https://github.com/lsilvapvt/concourse-pipeline-samples/tree/master/private-docker-registry)


## Credits

This project was forked from the original work of [J.Calabrese's sample concourse app](https://github.com/xchapter7x/concourse-demo-app).

[pipeline01]: https://raw.githubusercontent.com/lsilvapvt/sample-app-pipeline/master/images/pipeline01.png "Delivery pipeline"
[pipeline02]: https://raw.githubusercontent.com/lsilvapvt/sample-app-pipeline/master/images/pipeline02.png "Deployment pipeline"
