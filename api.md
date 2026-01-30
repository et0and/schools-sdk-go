# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Root

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#RootGetResponse">RootGetResponse</a>

Methods:

- <code title="get /">client.Root.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#RootService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#RootGetResponse">RootGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Schools

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolGetResponse">SchoolGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolListResponse">SchoolListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolSearchResponse">SchoolSearchResponse</a>

Methods:

- <code title="get /v1/schools/id/{schoolId}">client.Schools.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, schoolID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolGetResponse">SchoolGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/schools">client.Schools.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolListParams">SchoolListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolListResponse">SchoolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/schools/authority/{authority}">client.Schools.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolService.ByAuthority">ByAuthority</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, authority <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolByAuthorityParams">SchoolByAuthorityParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/schools/city/{city}">client.Schools.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolService.ByCity">ByCity</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, city <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolByCityParams">SchoolByCityParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/schools/status/{status}">client.Schools.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolService.ByStatus">ByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, status <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolByStatusParams">SchoolByStatusParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/schools/suburb/{suburb}">client.Schools.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolService.BySuburb">BySuburb</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, suburb <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolBySuburbParams">SchoolBySuburbParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/schools/search">client.Schools.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolSearchParams">SchoolSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SchoolSearchResponse">SchoolSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sync

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SyncGetStatusResponse">SyncGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SyncTriggerResponse">SyncTriggerResponse</a>

Methods:

- <code title="get /v1/sync/status">client.Sync.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SyncService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SyncGetStatusResponse">SyncGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sync">client.Sync.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SyncService.Trigger">Trigger</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go">schools</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/schools-go#SyncTriggerResponse">SyncTriggerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
