# API Response Verification Checklist

Before finalizing any spec, verify against actual API:

## Field Names
- [ ] All response fields use camelCase (e.g., `executionId`, not `ExecutionID`)
- [ ] Test with real API call using debug logging

## Field Types
- [ ] `loggingEnabled` - is it boolean or string?
- [ ] `executionPersistent` - is it boolean or string?
- [ ] `progressPercentage` - is it integer or float?
- [ ] `executionMetaData` - is it required or optional?

## Error Cases
- [ ] Test failure scenarios
- [ ] Document `errorMessage` field presence
- [ ] Document `executionMetaData` behavior on failure (null?)

## Testing
Always test with:
```bash
# Enable debug to see raw JSON
./client --debug command --args
```

## Common Pitfalls

1. **Field Casing**: API uses camelCase, not PascalCase
2. **Boolean Fields**: Many fields that look like booleans are actually strings ("true"/"false")
3. **Optional Fields**: `executionMetaData` may be null on failure
4. **Error Messages**: Always check for `errorMessage` field when status is "failed"
