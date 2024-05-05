# Organizations

Methods:

- <code title="post /organizations/create">client.Organizations.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go">farquest</a>.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#OrganizationNewParams">OrganizationNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Session

Methods:

- <code title="post /session/{correlatedId}">client.Session.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#SessionService.NewCorrelated">NewCorrelated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go">farquest</a>.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#SessionNewCorrelatedParams">SessionNewCorrelatedParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Auth

Methods:

- <code title="post /auth/callback">client.Auth.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#AuthService.Callback">Callback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go">farquest</a>.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#AuthCallbackParams">AuthCallbackParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /auth/{state}">client.Auth.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#AuthService.GetState">GetState</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, state <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Quest

Methods:

- <code title="post /quest/create">client.Quest.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go">farquest</a>.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestNewParams">QuestNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /quest/{id}">client.Quest.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /quest/list/{filter}">client.Quest.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filter <a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go">farquest</a>.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestListParamsFilter">QuestListParamsFilter</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Types

Methods:

- <code title="get /quest/types">client.Quest.Types.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Completions

Methods:

- <code title="get /quest/completions/count/{id}">client.Quest.Completions.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestCompletionService.Count">Count</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Validation

Methods:

- <code title="get /quest/validation/{id}">client.Quest.Validation.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestValidationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Quests

Methods:

- <code title="post /quest/complete">client.Quests.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go">farquest</a>.<a href="https://pkg.go.dev/github.com/FarquestSocial/farquest-go#QuestCompleteParams">QuestCompleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
