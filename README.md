# redshift-db-admin

This is a utility to administer redshift databases that are behind a firewall.

The published docker image runs with a lambda entrypoint.
Using a lambda that is on the same VPC as the database, this utility can ensure a database exists with a specific owner.
This utilizes AWS IAM to secure administration instead of using an SSH Tunnel or VPN.
This also limits the actions that a user can take, making it extremely hard to perform malicious commands.

## AWS Lambda setup

The Lambda requires specific configuration to work properly:

- A SecretsManager Secret containing the connection string as a postgres URL.
- `DB_CONN_URL_SECRET_ID` env var containing ARN of the AWS SecretsManager Secret.
- The execution role must have access to the above secret.
- The executing lambda must have network access to the redshift cluster.
