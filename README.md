# Streamlining CI/CD Workflow: A GitOps Pipeline with Jenkins, GitHub Actions, and ArgoCD (CI/CD)
This repo contains a dockerfile to build a echoserver docker image using Jenkinsfile and then creates a new PR in [echoserver app kubernetes](https://github.com/yogeshraj-au/echoserver_app_kubernetes.git) repo. Upon PR creation, the Github Action will trigger a workflow which will merge the PR. ArgoCD will monitor the main branch for new commits. When it detects a new commit, ArgoCD Sync the change and deploy the latest changes in Kubernetes Cluster.

![Architecture](GitOps_Pipeline_with_Argocd_and_Github_Actions.png.png)

[Blog post](https://medium.com/@yogeshraj-au/streamlining-your-development-workflow-a-gitops-pipeline-with-jenkins-github-actions-and-argocd-411d5c1adca7)

Tech Stack:

- Jenkins
- Github
- Docker
- ArgoCD
- Kubernetes
- Go

# Jenkins

The Jenkins folder contains a jenkinsfile `Jenkins/Jenkinsfile` which will copy the application code and package into a Docker image. It publishes to Docker registry.

# ArgoCD Pipeline

[Argocd repo](https://github.com/yogeshraj-au/echoserver_app_kubernetes.git)

## Github Actions

The `.github/workflows/mergepr.yml` folder contians config to merge the PR when a PR is created. 

## ArgoCD Config

The `environment/stage/apps` folder has config which will create application and it deploys resources in Kubernetes.