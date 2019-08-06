# kubeval-playground

## Output

```
GO111MODULE=on go run main.go manifests/invalid-deployment.yaml
```

```
[]kubeval.ValidationResult{
  kubeval.ValidationResult{
    FileName:               "stdin",
    Kind:                   "Deployment",
    APIVersion:             "apps/v1",
    ValidatedAgainstSchema: true,
    Errors:                 []gojsonschema.ResultError{
      &gojsonschema.RequiredError{
        ResultErrorFields: gojsonschema.ResultErrorFields{
          errorType: "required",
          context:   &gojsonschema.JsonContext{
            head: "spec",
            tail: &gojsonschema.JsonContext{
              head: "(root)",
              tail: (*gojsonschema.JsonContext)(nil),
            },
          },
          description:       "selector is required",
          descriptionFormat: "{{.property}} is required",
          value:             map[string]interface {}{
            "replicas": "many",
            "template": map[string]interface {}{
              "spec": map[string]interface {}{
                "containers": []interface {}{
                  map[string]interface {}{
                    "image": "ngix",
                    "name":  "ngix",
                    "ports": []interface {}{
                      map[string]interface {}{
                        "protocol":               "TCP",
                        "invalidAdditionalField": "invalid",
                        "name":                   "grpc",
                      },
                    },
                  },
                },
              },
            },
          },
          details: gojsonschema.ErrorDetails{
            "property": "selector",
            "field":    "selector",
            "context":  "(root).spec",
          },
        },
      },
      &gojsonschema.RequiredError{
        ResultErrorFields: gojsonschema.ResultErrorFields{
          errorType: "required",
          context:   &gojsonschema.JsonContext{
            head: "0",
            tail: &gojsonschema.JsonContext{
              head: "ports",
              tail: &gojsonschema.JsonContext{
                head: "0",
                tail: &gojsonschema.JsonContext{
                  head: "containers",
                  tail: &gojsonschema.JsonContext{
                    head: "spec",
                    tail: &gojsonschema.JsonContext{
                      head: "template",
                      tail: &gojsonschema.JsonContext{...},
                    },
                  },
                },
              },
            },
          },
          description:       "containerPort is required",
          descriptionFormat: "{{.property}} is required",
          value:             map[string]interface {}{...},
          details:           gojsonschema.ErrorDetails{
            "field":    "containerPort",
            "context":  "(root).spec.template.spec.containers.0.ports.0",
            "property": "containerPort",
          },
        },
      },
      &gojsonschema.AdditionalPropertyNotAllowedError{
        ResultErrorFields: gojsonschema.ResultErrorFields{
          errorType:         "additional_property_not_allowed",
          context:           &gojsonschema.JsonContext{...},
          description:       "Additional property invalidAdditionalField is not allowed",
          descriptionFormat: "Additional property {{.property}} is not allowed",
          value:             "invalid",
          details:           gojsonschema.ErrorDetails{
            "property": "invalidAdditionalField",
            "field":    "invalidAdditionalField",
            "context":  "(root).spec.template.spec.containers.0.ports.0",
          },
        },
      },
      &gojsonschema.InvalidTypeError{
        ResultErrorFields: gojsonschema.ResultErrorFields{
          errorType: "invalid_type",
          context:   &gojsonschema.JsonContext{
            head: "replicas",
            tail: &gojsonschema.JsonContext{...},
          },
          description:       "Invalid type. Expected: [integer,null], given: string",
          descriptionFormat: "Invalid type. Expected: {{.expected}}, given: {{.given}}",
          value:             "many",
          details:           gojsonschema.ErrorDetails{
            "expected": "[integer,null]",
            "given":    "string",
            "field":    "spec.replicas",
            "context":  "(root).spec.replicas",
          },
        },
      },
    },
  },
}
```

## Reference

- [Kubeval](https://kubeval.instrumenta.dev/)
- [instrumenta/kubeval](https://github.com/instrumenta/kubeval)
