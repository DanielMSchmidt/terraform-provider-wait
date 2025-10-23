# Terraform Provider Wait

This Terrform Provider allows you to wait for certain events. Right now it is focussed around the new Terraform Actions.
This Provider will never wait for an amount of time since this enables brittle infrastructure code, we will just wait for things to happen, e.g. ports opening.

## Using the provider


See the [documentation](http://registry.terraform.io/providers/DanielMSchmidt/wait/latest) for details on how to use the provider and its actions.


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
