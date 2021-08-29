# Deploy CI

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/arthurkay/deploy-ci)
![GitHub](https://img.shields.io/github/license/arthurkay/deploy-ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkay/deploy-ci)](https://goreportcard.com/report/github.com/arthurkay/deploy-ci)
[![Go](https://github.com/arthurkay/env/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/arthurkay/env/actions/workflows/go.yml)
![GitHub issues](https://img.shields.io/github/issues-raw/arthurkay/deploy-ci)

CI and CD are two acronyms frequently used in modern development practices and DevOps. CI stands for continuous integration, a fundamental DevOps best practice where developers frequently merge code changes into a central repository where automated builds and tests run.

## What is Deploy CI

Deploy CI is a web service that listens for http calls, and then executes an ansible playbook.


## What is Ansible

Ansible is an open-source software provisioning, configuration management, and application-deployment tool enabling infrastructure as code.

## What is a Playbook

An organized unit of scripts that defines work for a server configuration managed by the automation tool Ansible.

## How Deploy CI Works

Deploy CI creates an API end point that gets hit and responds by executing and ansible play, the playbook can perform any set of tasks beyond just running tests and deployments. Possibilities are endless.

Setting It Up

The project can be cloned to a resource directory, e.g `/usr/local/src`.
From there the user can edit the [Evironment File](environment.conf) to addjust parameters like the port, the playbook to run and the user to run the play as.

The config file looks something like the following:

```conf

Environment=DEPLOY_PORT=5000
Environment=DEPLOY_ANSIBLE_FILE=/home/arthur/Documents/Dev/misc/deploy-ci/environment.conf
Environment=DEPLOY_USER=arthur

```

Alternatively the config parameters can be set during the creation of the systemd services file, e.g:

```bash

export DEPLOY_PORT=7000
export DEPLOY_ANSIBLE_FILE=/home/arthur/deployment.yml
export DEPLOY_USER=arthur

make setup

```

The above sequence of comands sets the values to be used as paramters when the make `setup` target runs.

***NOTE***: Environment variables set as export flags in the terminal tak precedence over the ones in the [Evironment File](environment.conf).

After setting up the tool, it can be managed with systemd.