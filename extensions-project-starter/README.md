# Extensions 2.0 Project Starter

## About

This is a template project structure and some convenience scripts to speed up through some of the steps of developing Dynatrace Extensions 2.0. It takes away some of the manual effort involved in setting everything up for first time use.

The scripts are primarily based on the Dynatrace [dt-cli](https://github.com/dynatrace-oss/dt-cli) utility which remains the main developer's toolbox when it comes to Extensions 2.0

## Pre-requisites

1. **VisualStudio Code** - this repo has been tailored for VisualStudio Code users and a configuration has been added for one of its extensions.
    * VisualStudio Extension for YAML language support
2. **Python** - the convenience scripts are written in Python
    * [dt-cli](https://github.com/dynatrace-oss/dt-cli) module is required
3. **Dynatrace** - a Dynatrace SaaS or Managed environment where you're planning to develop Extensions 2.0
4. **Dynatrace API token** - a Dynatrace API token with appropriate permissions depending on which scripts you're planning to use

## Project structure

* **.vscode**
    * This repo has been tailored for VisualStudio Code users
    * Workspace settings have been configured to recognise JSON schemas from the `/schemas` folder and automatically apply them to the `extension/extension.yaml` file
* **build**
    * This is where the convenience scripts will build the extensions that you are developing.
* **certs**
    * This is where the convenience scripts will generate the certificates required for signing Dynatrace Extensions 2.0
* **extension**
    * The `extension` folder where the `extension.yaml` file should reside. Both of these must have these exact names otherwise the archived extension package will not comply with the Extension 2.0 requirements.
* **schemas**
    * This is where the convenience scripts will download all the schemas by default.
    * This location has been configured for VS Code to automatically apply validation to your `extension.yaml` file
* **scripts**
    * Convenience scripts to help you automate boring tasks so you can build more and worry less!
    * Scripts can be customized with parameters taken from `config.yaml`

## Getting started

Install the python dependencies

```bash
pip install -r requirements.txt
```

Change into the scripts directory

```bash
cd scripts
```

Modify `config.yaml` and add your Dynatrace Environment URL and API Token.

Run the first scripts
```bash
python initialize.py
python download_schemas.py
```

## Convenience scripts

The scripts are meant to run from within the scripts folder.

### [initialize.py](scripts/initialize.py)

_API permissions needed: `credentialVault.read` and `credentialVault.write`_

This script should only be needed when first starting your work with Extensions 2.0. It does the following:
* Generate the certificates needed to work with Extensions 2.0
* Upload the CA certificate to the Dynatrace Credential Vault

### [build_and_upload.py](scripts/build_and_upload.py)

_API permissions needed: `extensions.read` and `extensions.write`_

This script should be used every time an update should be built and published to Dynatrace. It does the following:
* Package and sign your Extension, producing a .zip archive in the build folder
* Check whether the limit for extension versions has been reached, and remove the oldest one
* Upload the latest version of the extension to Dynatrace
* Activate the latest extension version

### [download_schemas.py](scripts/download_schemas.py)

_API permissions needed: `extensions.read`_

This script should only be needed once, whenever schema files are missing or you want to target a different version than what you already have. It does the following:
* Downloads all the extension schema files of a specific version
* Schemas are downloaded to `schemas` folder and `.vscode/settings.json` is configured to recognise this
