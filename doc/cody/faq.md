<style>

  .markdown-body .cards {
  display: flex;
  align-items: stretch;
}

.markdown-body .cards .card {
  flex: 1;
  margin: 0.5em;
  color: var(--text-color);
  border-radius: 4px;
  border: 1px solid var(--sidebar-nav-active-bg);
  padding: 1.5rem;
  padding-top: 1.25rem;
}

.markdown-body .cards .card:hover {
  color: var(--link-color);
}

.markdown-body .cards .card span {
  color: var(--link-color);
  font-weight: bold;
}

.limg {
  list-style: none;
  margin: 3rem 0 !important;
  padding: 0 !important;
}
.limg li {
  margin-bottom: 1rem;
  padding: 0 !important;
}

.limg li:last {
  margin-bottom: 0;
}

.limg a {
    display: flex;
    flex-direction: column;
    transition-property: all;
   transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
     transition-duration: 350ms;
     border-radius: 0.75rem;
  padding-top: 1rem;
  padding-bottom: 1rem;

}

.limg a {
  padding-left: 1rem;
  padding-right: 1rem;
  background: rgb(113 220 232 / 19%);
}

.limg p {
  margin: 0rem;
}
.limg a img {
  width: 1rem;
}

.limg h3 {
  display:flex;
  gap: 0.6rem;
  margin-top: 0;
  margin-bottom: .25rem

}

</style>

# Cody FAQ

<p class="subtitle">Find answers to the most common questions about Cody.</p>

## General

### Does Cody train on my code?

No, Cody does not train on your code. Our third-party Language Model (LLM) providers also do not train on your specific codebase. Cody operates by following a specific process to generate answers to your queries:

- **User query**: A user asks a question
- **Code retrieval**: Sourcegraph, our underlying code intelligence platform, performs a search and code intelligence operation to retrieve code snippets relevant to the user's question. During this process, strict permissions are enforced to ensure that only code that the user has read permission for is retrieved
- **Prompt to Language Model**: Sourcegraph sends a prompt, and the code snippets are retrieved to a Language Model (LLM). This prompt provides the context for the LLM to generate a meaningful response
- **Response to user**: The response generated by the LLM is then sent back to Cody and presented to the user

This process ensures that Cody can provide helpful answers to your questions while respecting data privacy and security by not training on or retaining your specific code.

### Does Cody work with self-hosted Sourcegraph?

Yes, Cody is compatible with self-hosted Sourcegraph instances. However, there are a few considerations:

- Cody operates by sending code snippets (up to 28 KB per request) to a third-party cloud service. By default, this service is Anthropic but can also be OpenAI
- For certain repositories, Cody may utilize embeddings, which involves sending repository data to another third-party service like OpenAI
- To use Cody effectively, your self-hosted Sourcegraph instance must have internet access for these interactions with external services

### Is there a public facing Cody API?

Currently, there is no public-facing Cody API available.

### Does Cody require Sourcegraph to function?

Yes, Cody relies on Sourcegraph for two essential functions:

- It is used to retrieve context relevant to user queries
- Sourcegraph acts as a proxy for the LLM provider to facilitate the interaction between Cody and the LLM

### What programming languages Cody supports?

Cody supports a wide range of programming languages, including:

JavaScript, TypeScript, PHP, Python, Java, C/C++, C#, Ruby, Go, SQL, Swift, Objective-C, Perl, Rust, Kotlin, Scala, Groovy, R, MATLAB, Dart, Lua, Julia, Cobol and Shell scripting languages (like Bash, PowerShell).

### Can Cody answer non-programming questions?

Cody is an expert in answering a wide range of coding-related questions on topics including questions about your codebase, general programming concepts, test cases, debugging, and more. Cody Chat is not designed to answer non-coding questions or provide general information on topics outside of coding or your codebase.

### What happened to the Cody App?

We’ve deprecated the Cody App to streamline the experience for our Cody Free and Cody Pro users. Now, anyone with a Sourcegraph.com account can generate local embeddings for their personal projects within the VS Code extension without downloading and connecting the Cody App. Local embeddings are only supported for VS Code, but we’re working on adding the same functionality to JetBrains IDEs.

## Embeddings

### What are embeddings for?

Embeddings help Sourcegraph retrieve relevant code to feed the Large Language Model as context. Embeddings, often associated with vector search, complement other strategies in the code retrieval process.

While embeddings excel in semantic matching — determining "what is this code about" and "what does it do" — they may not capture syntax and other specific matching details as effectively. Sourcegraph's approach involves getting the best results from various sources to deliver the most accurate and comprehensive answers possible.

### Do embeddings enforce permissions? Does Cody receive code that users don't have access to?

When using embeddings, permissions are enforced to ensure Cody does not receive code the user cannot access. Currently, Sourcegraph uses embeddings search for a single repository, with a prior check to confirm user access.

In the future, the process will involve the following steps:

- Determine which repositories the user has access to
- Query embeddings for each of these repositories
- Select the most relevant results and provide them to the user

This approach safeguards data privacy and ensures that Cody's responses are based on code accessible to the user.

### Why isn't my scheduled embedding job listed?

There can be several reasons why your scheduled one-off embedding job isn't appearing in the job list:

- The repository is already in the queue or currently being processed
- The system has successfully completed a job for the same repository and revision
- Another job for the same repository is in the queue, scheduled within the [`embeddings.MinimumInterval`](./core-concepts/embeddings.md#minimum-time-interval-between-automatically-scheduled-embeddings) time window

### How do I stop a running embeddings job?

A running embeddings job with the state `QUEUED` or `PROCESSING` can be stopped by admins from the **Cody > Embeddings Jobs** page. To do so:

- Click on the "Cancel" button associated with the job you wish to terminate
- The job will then be tagged for cancellation. Please note that the time required for the job to be fully canceled may vary depending on its current state, ranging from a few seconds to a few minutes

### Why are files skipped?

Files may be skipped for the following reasons:

- The file size exceeds 1 MB
- The file path matches an [exclusion pattern](./core-concepts/embeddings/manage-embeddings.md#filter-files-from-embeddings)
- The repository has already reached the maximum limit for generated embeddings, as specified by [`embeddings.maxCodeEmbeddingsPerRepo`](./core-concepts/embeddings/usage-and-limits.md#limit-the-number-of-embeddings-that-can-be-generated) or [`embeddings.maxTextEmbeddingsPerRepo`](./core-concepts/embeddings/usage-and-limits.md#limit-the-number-of-embeddings-that-can-be-generated)

## Third party dependencies

### What is the default `sourcegraph` provider for completions and embeddings?

The default provider for completions and embeddings, specified as `"provider": "sourcegraph"` refers to the [Sourcegraph Cody Gateway](./core-concepts/cody-gateway.md). The Cody Gateway facilitates access to completions and embeddings for Sourcegraph enterprise instances by leveraging third-party services such as Anthropic and OpenAI.

### What third-party cloud services does Cody depend on?

Cody relies on one primary third-party dependency, i.e., Anthropic's Claude API. Users can use this with the OpenAI API configuration.

Additionally, Cody can optionally use OpenAI for generating embeddings, enhancing the quality of its context snippets, although this is not mandatory.

It's worth noting that these dependencies remain consistent when utilizing the [default `sourcegraph` provider, Cody Gateway](./core-concepts/cody-gateway.md), which uses the same third-party providers.

### What is the retention policy for Anthropic and OpenAI?

Please refer to this [terms and conditions](https://sourcegraph.com/terms/cody-notice) for details regarding the retention policy for data managed by Anthropic and OpenAI.

### Can I use my own API keys?

Yes! you can use your own API keys.

### Can I use with my Cloud IDE?

Yes, Cody supports the following cloud development environments:

- vscode.dev and GitHub Codespaces (install from the VS Code extension marketplace)
- Any editor supporting the [Open VSX Registry](https://open-vsx.org/extension/sourcegraph/cody-ai), including [Gitpod](https://www.gitpod.io/blog/boosting-developer-productivity-unleashing-the-power-of-sourcegraph-cody-in-gitpod), Coder, and `code-server` (install from the [Open VSX Registry](https://open-vsx.org/extension/sourcegraph/cody-ai))

## More resources

For more information on what to do next, we recommend the following resources:

<div class="cards">
  <a class="card text-left" href="./quickstart"><b>Cody Quickstart</b><p>This guide recommends how to use Cody once you have installed the extension in your VS Code editor.</p></a>
  <a class="card text-left" href="troubleshooting"><b>Troubleshooting with Cody</b><p>Learn how you can troubleshoot and debug common development issues with Cody.</p></a>
</div>
