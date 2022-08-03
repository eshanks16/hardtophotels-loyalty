# Loyalty App

This application serves as a demo app for Technical Marketing. 
The application helps fictional customers sign up for a loyalty program at the fictional company Hard Top Hotels.

## Purpose

The purpose of this application is to have a sample app to use for demonstrations.
During demos/blog posts/webinars, etc this app can be deployed on virtual machines, or through containers to be used for migration purposes, backup/restore demonstrations, or app deconstruction.

## Architecture

This application has two parts associated with it.

1. The golang web server
2. A mongo database to store "loyalist" information.

# Deployment

## <a name="k8s-deploy"></a>Deployment - Kubernetes

There are manifests in the `kubernetes` directory to deploy a prebuilt image.

> NOTE: In order to use the existing manifests, you'll need a Kubernetes cluster with a cloud provider capable of creating load balancers. If you need a bare metal load balancer there are manifests in the [metal-lb](kubernetes/metal-lb) folder.

You can deploy all of the kubernetes components by deploying each of the yaml manifests in the [kubernetes](kubernetes) folder. There is also an option to deploy with letsencrypt certificates if you can use cert-manager with letsencrypt in your environment.

1. Configure your kubectl to use a KUBECONFIG file with access to your Kubernetes cluster.
2. Update the credentials in the deployment YAML to match your database.
3. Deploy the manifests within the [Kubernetes](kubernetes) folder in sequential order.
4. Set a DNS entry to match your loyalty service endpoint. This should be the entry in the virtualhost config in the [loyalty app manifest](kubernetes/4-ht-loyalty_dep-svc-ingress.yml).

## <a name="documentDB"></a>Deployment - DocumentDB

The Loyalty App can be deployed while using the Amazon DocumentDB service.

Kubernetes Deployment

If you're deloying Loyalty App through Kubernetes, the [Loyalty Manifest](kubernetes/4-ht-loyalty_dep-svc-ingress.yml) must have the `MONGO_TLS` value set to a non `nil` value. When the application starts up with a non-nil value, it appends the AWS Certificate details to the connection string. The certificate is included in the image build.
