
The private/public key pairs used by Tessera can be [stored](../Keys) in and [retrieved](../../../Configuration/Keys) from a key vault, preventing the need to store the keys locally.

This page details how to set up and configure an _AWS Secrets Manager_ for use with Tessera.

The _AWS Secrets Manager_ documentation provides much of the information needed to get started.  The information in this section has been taken from the following pages of the _AWS_ documentation:

* [AWS Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) 
* [AWS SDK for Java Developer Guide](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/welcome.html)

## Creating the AWS Secrets Manager

Once you have set up your AWS profile, you will be able to use AWS Secrets Manager.

## Enabling Tessera to use the _AWS Secrets Manager_
### Environment Variables

If using an _AWS Secrets Manager_, configuration credentials can be provided in many ways as outlined in the [AWS docs - Supplying and Retrieving AWS Credentials](https://docs.aws.amazon.com/sdk-for-java/v2/developer-guide/credentials.html).  

To use environment variables set the following:

1. `AWS_REGION`: region_to_connect_to (i.e. us-west-2)
1. `AWS_ACCESS_KEY_ID`: your_access_key_id
1. `AWS_SECRET_ACCESS_KEY`: your_secret_access_key

The `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` for a particular user can be retrieved from the [AWS IAM Management Console](https://console.aws.amazon.com/iam).

### Dependencies
The AWS Secrets Manager dependencies are included in the `tessera-app-<version>-app.jar`.  If using the `tessera-simple-<version>-app.jar` then `aws-key-vault-<version>-all.jar` must be added to the classpath.
