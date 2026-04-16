# Projects Resource

## Purpose

This resource area covers project discovery and project-level resource discovery.

## Client Shape

- Root client exposes project discovery.
- Project-scoped access starts with a `projectId` returned by the discovery endpoints.

## Operations

### List Projects

- Raw operation: `GET /projects`
- Raw operationId: `getProjects`
- Result schema: `ProjectsResult`
- Purpose: discover all available projects in the installation.

### Get Project Resources

- Raw operation: `GET /projects/{projectId}`
- Raw operationId: `getProjectResources`
- Result schema: `ProjectResourcesResult`
- Purpose: discover project-level resources for a selected project.

## Shared Rules

- `projectId` is the wire-level identifier and must not be renamed in request construction.
- Project discovery is read-oriented and follows the normal session and authentication rules.
