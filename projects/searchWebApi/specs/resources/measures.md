# Measures Resource

## Operation

### Get Measure Cube

- Raw operation: `POST /projects/{projectId}/collections/{collectionId}/measures`
- Raw operationId: `getMeasureCube`
- Request schema: array of `DimensionRequest`
- Result schema: `MeasureCube`

## Shared Rules

- Measure requests may contain zero, one, or two dimension requests.
- Measure behavior depends on `MeasureTypeParameter` and must be documented without collapsing count and aggregate cases into one simplified description.
