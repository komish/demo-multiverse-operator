# Webhook Conversion Demo - "Multiverse Operator"

## Summary

This operator manages a sample custom resource `MyCustomKind` for demonstrating
webhook conversion. The reconciliation logic is not implemented as it is out of
scope.

## Walkthrough

At the time of publication, the controller-manager container exists on
[Quay.io](https://quay.io). This means that you can just run the following
command from the repo root to install the controller-manager in your cluster.
This does require [Cert-Manager](https://cert-manager.io/docs/) be installed
already to automatically manager certificates needed by your webhook.

```shell
make deploy
```

From there, you can install either sample in the `config/samples/` directory.

```shell
oc apply -f config/samples/demo_v1beta1_mycustomkind.yaml
```

With the instance installed, you should then be able to see the conversion in
play:

```shell
oc get mycustomkind.v1alpha1.demo.example.com mycustomkind-v1beta1-sample -o yaml
```

In this particular example, you would see that the `bar` key with value `baz` is
represented in the annotations because it does not exists in the spec of the
`v1alpha1` representation of `MyCustomKind`

```yaml
apiVersion: demo.example.com/v1alpha1
kind: MyCustomKind
metadata:
  annotations:
    conversion.mycustomkind.demo.example.com/bar: baz
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"demo.example.com/v1beta1","kind":"MyCustomKind","metadata":{"annotations":{},"name":"mycustomkind-v1beta1-sample","namespace":"multiverse-operator-system"},"spec":{"bar":"baz","newFoo":"teehee"}}
  creationTimestamp: "2021-04-15T22:22:09Z"
  generation: 1
  managedFields:
    # omitted
  name: mycustomkind-v1beta1-sample
  namespace: multiverse-operator-system
  resourceVersion: "107033543"
  selfLink: /apis/demo.example.com/v1alpha1/namespaces/multiverse-operator-system/mycustomkinds/mycustomkind-v1beta1-sample
  uid: e19b3d2e-5d95-474d-ab7b-36a8c282134c
spec:
  foo: teehee
status: {}
```

Conversely, the `v1alpha1` representation of `MyCustomKind` accepts a key `zap`
with an `int32` value. You can edit the existing resource and add `.spec.zap=3`,
after which you should see the value persisted to annotations in the `v1beta1`
representation of the instance.
