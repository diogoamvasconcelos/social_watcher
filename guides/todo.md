- cloudwatch event trigger for lambda
- twitterLambda(keyword) -> searches twitter (last hour/day) and adds to ddb
  - v1: searches last hour
  - v2: also searches for replies on main entries 7days
- ddbstream consumer lambda -> consumes and posts discord
  - v1: consumes and logs
  - v2: posts to discord
  - v3: translates if not english?


- CI/Dev
  - Run lambda locally
  - Unit Tests
  - Integration Tests
  - Environment Tests
  - Test coverage
  - Vulnerability scans
  - Behaviour test (cucumber?)