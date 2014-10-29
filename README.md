aws-env
=======

About
-----

aws-env will read the profiles from your ~/.aws/config file (or another similarly structured file of your choosing) and print out a list of the export directives to set your Environment variables for use with commandline aws tools.

Installation
------------

Download the latest release, or build from source.

Usage
-----

```bash
$ aws-env --profile specialized-profile
export AWS_ACCESS_KEY_ID=A1B2C3D4E5F6G7H8I9J0K
export AWS_SECRET_ACCESS_KEY=aHR0cHM6Ly9tb2R1c2NyZWF0ZS5jb20vY2FyZWVycw==
export OUTPUT=json
export REGION=us-east-1
```

You could do something like the following, to automatically export the environment variables.

```bash
$ $(aws-env --profile not-default)
```

Enjoy!