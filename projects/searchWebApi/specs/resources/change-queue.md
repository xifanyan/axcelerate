# Change Queue Resource

## Operation

### Wait For Pending Changes

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/changes/queue`
- Raw operationId: `waitForAllCurrentlyScheduledChangeRequests`
- Result schema: `WaitForPendingChangesResult`

## Shared Rules

- This operation waits for changes that are queued or in progress at the time of submission.
- `timeoutMillis` and `onlyHighPriorityChanges` stay wire-exact.
- The result must preserve the backend's distinction between request success and queue-drain completion.
