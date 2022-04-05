import json
import yaml
from utils import Dynatrace


def get_target_version():
    versions = dt.make_request(dt.SCHEMAS_API).json().get("versions", [])

    if target_version == "latest":
        return versions[-1]
    
    matches = [v for v in versions if v.startswith(target_version)]
    if matches:
        return matches[0]
    
    print(f"Target version {target_version} does not exist.")
    raise SystemExit


if __name__ == "__main__":
    # Read config
    with open(file="config.yaml", mode="r") as f:
        config = yaml.safe_load(f)

    # Set parameters
    target_version = str(config.get("schema_version", "latest"))
    download_dir = config.get("download_folder", "../schemas")
    tenant_url = config["tenant_url"]
    api_token = config["api_token"]

    # Get a DT client
    dt = Dynatrace(tenant_url, api_token)

    version = get_target_version()        
    print(f"Downloading schemas for version {version}")

    # Download all the schemas
    files = dt.make_request(f"{dt.SCHEMAS_API}/{version}").json().get("files", [])
    for file in files:
        schema = dt.make_request(f"{dt.SCHEMAS_API}/{version}/{file}").json()
        with open(file=f"{download_dir}/{file}", mode="w") as f:
            json.dump(schema, f, indent=2)

    print("Finished.")