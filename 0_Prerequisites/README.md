# Prerequisites

To successfully develop Extensions 2.0 and be able to complete this tutorial you will need to sort out some pre-requisites:
* Admin access to a Dynatrace SaaS or Managed environment (min. version 1.227)
* A Windows host (virtual machine)
* OneAgent deployed on the said Windows host
* A developer certificate and key (to be used for signing your extensions)
* Either [OpenSSL](https://www.openssl.org/source/) or [dt-cli](https://github.com/dynatrace-oss/dt-cli) on your machine
* Your root certificate uploaded to Dynatrace and on the OneAgent host

## Generate a developer certificate and key

### Option 1: via dt-cli
1. Install [Python 3.9+](https://www.python.org/downloads/) on your machine
2. Install `dt-cli` python module
  ```shell
  pip install dt-cli
  ```
3. Download [dt.py script](https://raw.githubusercontent.com/dynatrace-oss/dt-cli/main/dtcli/scripts/dt.py)

4. Generate all the required files
  ```shell
  python dt.py extension gencerts --days-valid 10000
  ```

   Files are created as:
   * `developer.pem` - your developer certificate
   * `developer.key` - your developer key
   * `ca.pem` - your root certificate
   * `ca.key` - your root key

### Option 2: via OpenSSL
1. Download OpenSSL and install it on your machine
   * Windows: download [OpenSSL v1.1L](https://slproweb.com/download/Win64OpenSSL-1_1_1L.exe) and run the installer
      Or: `choco install openssl` (if you have chocolatey)
   * Mac: `brew install openssl@1.1`
2. Create the root key and certificate
   ```shell
   openssl genrsa -out root.key 2048
   openssl req -days 10000 -new -x509 -key root.key -out root.pem
   openssl rsa -in root.key -pubout -out root.pub.key
   ```
3. Create a certificate signing request
   ```shell
   openssl genrsa -out developer.key 2048
   openssl rsa -in developer.key -pubout -out developer.pub.key
   openssl req -new -key developer.key -out developer.csr
   ```
4. Issue your developer certificate and key
   ```shell
   openssl x509 -req -days 10000 -in developer.csr -CA root.pem -Cakey root.key -Cacreateserial -out developer.pem
   ```
   Files are created as:
   * `developer.pem` - your developer certificate
   * `developer.key` - your developer key
   * `root.pem` - your root certificate
   * `root.key` - your root key

## Distribute the root certificate to Dynatrace components

### 1. Upload to the Dynatrace Credential Vault
1. Go to `Settings`>`Web and mobile monitoring`>`Credential Vault`>`Add credential`
2. Set the credential type to `Public Certificate`
3. Upload the `ca.pem` or `root.pem` generated earlier

### 2. Upload to OneAgent host that runs the extension
1. Go to `C:\ProgramData\dynatrace\oneagent\agent\config`
2. Go to `certificates` folder or create it it doesn't exist
3. Upload your `ca.pem` or `root.pem` generated earlier


###[Click here to start the tutorial ▶](/1_Basic-Extension)