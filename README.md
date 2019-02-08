# K8slate

K8slate (pronounced like _Kate's late_) is a cli tool for templating Kubernetes resources.

K8slate grabs all `.jinja2` files in a directory and templates them out to an `output/` folder.
These files must be composed of pairs of YAML preambles followed by their corresponding content.
These preambles specify option values used by `k8slate` in the formatting and writing phases.

Files can contain multiple Kubernetes resource descriptions, but each must contain their preamble and contents,
which makes a valid file always have, effectively, a pair number of YAML documents.

Files are named after the resource name, followed by its type.
Like this: `nginx-deployment.yaml`, for a deployment of `metadata.name = nginx` and `kind = Deployment`.

## Preamble spec

Currently, `k8slate` supports one entry in the preamble: `params`.

`params` specifies the values to be injected as variables in the `jinja2` templating phase.
If `params` is a list of dictionaries, the template will be rendered once per dictionary, taking into account its values, which means that, if `metadata.name` does not take into account input from `params`, subsequent rendered paramater dictionaries will replace the prior generated file.
