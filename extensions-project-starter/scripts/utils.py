import os
import requests


class Dynatrace:
    SCHEMAS_API = "api/v2/extensions/schemas"
    CREDENTIALS_API = "api/config/v1/credentials"
    EXTENSIONS_API = "api/v2/extensions"


    def __init__(self, tenant_url: str, token: str):
        self.url = tenant_url
        self.token = parse_token(token)
        self.auth_header = {'Authorization': f'Api-Token {self.token}'}

    def make_request(self, path: str, method: str="GET", json: dict=None) -> dict:
        url = f"{self.url}/{path}"

        resp = requests.request(
            url=url,
            method=method, 
            headers=self.auth_header,
            json=json
        )
        if resp.status_code not in [200, 201, 204]:
            print("Could not complete request")
            print(resp.text)
            raise Exception(resp.text)

        return resp


def parse_token(raw_token: str):
    if raw_token.startswith(".Env."):
        return os.environ.get(raw_token[5:], "")
    return raw_token