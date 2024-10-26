# Daily Progress for langchain-ai/langchain

## Issues
- fix the grammar and markdown component (URL: https://github.com/langchain-ai/langchain/pull/27657)
- Multi-agent supervisor example not working with ChatOllama, breaks when attempting to use two system messages. Must be an issue with using two system prompts in ChatOllama. (URL: https://github.com/langchain-ai/langchain/issues/27656)
- community: Update Replicate LLM and fix tests (URL: https://github.com/langchain-ai/langchain/pull/27655)
- community: Add Writer integration (URL: https://github.com/langchain-ai/langchain/pull/27646)
- Add ChatModels wrapper for Cloudflare Workers AI (URL: https://github.com/langchain-ai/langchain/pull/27645)
- community: add ObjectBox as a vector-store (URL: https://github.com/langchain-ai/langchain/pull/27644)
- AzureSearch error (URL: https://github.com/langchain-ai/langchain/issues/27643)
- How do I use langchain for vllm serve tool calls？ (URL: https://github.com/langchain-ai/langchain/issues/27642)
- Fix typo (missing letter) in elasticsearch_retriever.ipynb (URL: https://github.com/langchain-ai/langchain/pull/27639)
- DOC: data parsing error handling (if retry and data fixing does not work) (URL: https://github.com/langchain-ai/langchain/issues/27635)
- [AIMessage]tool_calls.0.args is not always returned as a valid dictionary (URL: https://github.com/langchain-ai/langchain/issues/27632)
- community: failing unit test with sqlalchemy 2.0.36 (URL: https://github.com/langchain-ai/langchain/issues/27627)
- Error during FAISS save_local due to __pydantic_private__ attribute (URL: https://github.com/langchain-ai/langchain/issues/27625)
- Cant import any of the HuggingFaceEmbeddings because 'openssl' has no attribute 'ciphers' (URL: https://github.com/langchain-ai/langchain/issues/27624)
- Confluence Loader: Fix CQL loading (URL: https://github.com/langchain-ai/langchain/pull/27620)
- core: Modified RunnableWithMessageHistory get_input_schema to make use of the underlying runnable input keys. (URL: https://github.com/langchain-ai/langchain/pull/27619)
- docs: run how-to guides in CI (URL: https://github.com/langchain-ai/langchain/pull/27615)
- check broken links (URL: https://github.com/langchain-ai/langchain/pull/27614)
- Various warnings due to Pydantic protected namespaces, such as UserWarning: Field "model_name" in JinaEmbeddings has conflict with protected namespace "model_". (URL: https://github.com/langchain-ai/langchain/issues/27609)
- community: Add `@mozilla/readability` document transformer (URL: https://github.com/langchain-ai/langchain/pull/27604)
- community: handle chatdeepinfra jsondecode error (URL: https://github.com/langchain-ai/langchain/pull/27603)
- ChatDeepInfra JSONDecode error when tool call args is empty string (URL: https://github.com/langchain-ai/langchain/issues/27602)
- DallEAPIWrapper not support base_url setting which supported by OpenAI sdk. (URL: https://github.com/langchain-ai/langchain/issues/27597)
- rfc: trace standard tool schema directly (URL: https://github.com/langchain-ai/langchain/pull/27596)
- langchain_openai: AZURE_OPENAI_DEPLOYMENT environment variable (URL: https://github.com/langchain-ai/langchain/pull/27595)
- langchain-openai: add check for none values when summing token usage (URL: https://github.com/langchain-ai/langchain/pull/27585)
- core: update load chat prompt template multi mesages (URL: https://github.com/langchain-ai/langchain/pull/27584)
- PowerBIDataSet credential not working: field "credential" not yet prepared so type is still a ForwardRef, you might need to call PowerBIDataset.update_forward_refs() (URL: https://github.com/langchain-ai/langchain/issues/27583)
- neo4j: initial package (URL: https://github.com/langchain-ai/langchain/pull/27582)
- [community] Pebblo: Policy-based semantic context in PebbloRetrievalQA (URL: https://github.com/langchain-ai/langchain/pull/27581)

## Pull Requests
- fix the grammar and markdown component (URL: https://github.com/langchain-ai/langchain/pull/27657)
- community: Update Replicate LLM and fix tests (URL: https://github.com/langchain-ai/langchain/pull/27655)
- community: Add Writer integration (URL: https://github.com/langchain-ai/langchain/pull/27646)
- Add ChatModels wrapper for Cloudflare Workers AI (URL: https://github.com/langchain-ai/langchain/pull/27645)
- community: add ObjectBox as a vector-store (URL: https://github.com/langchain-ai/langchain/pull/27644)
- Fix typo (missing letter) in elasticsearch_retriever.ipynb (URL: https://github.com/langchain-ai/langchain/pull/27639)
- Confluence Loader: Fix CQL loading (URL: https://github.com/langchain-ai/langchain/pull/27620)
- core: Modified RunnableWithMessageHistory get_input_schema to make use of the underlying runnable input keys. (URL: https://github.com/langchain-ai/langchain/pull/27619)
- docs: run how-to guides in CI (URL: https://github.com/langchain-ai/langchain/pull/27615)
- check broken links (URL: https://github.com/langchain-ai/langchain/pull/27614)
- community: Add `@mozilla/readability` document transformer (URL: https://github.com/langchain-ai/langchain/pull/27604)
- community: handle chatdeepinfra jsondecode error (URL: https://github.com/langchain-ai/langchain/pull/27603)
- rfc: trace standard tool schema directly (URL: https://github.com/langchain-ai/langchain/pull/27596)
- langchain_openai: AZURE_OPENAI_DEPLOYMENT environment variable (URL: https://github.com/langchain-ai/langchain/pull/27595)
- langchain-openai: add check for none values when summing token usage (URL: https://github.com/langchain-ai/langchain/pull/27585)
- core: update load chat prompt template multi mesages (URL: https://github.com/langchain-ai/langchain/pull/27584)
- neo4j: initial package (URL: https://github.com/langchain-ai/langchain/pull/27582)
- [community] Pebblo: Policy-based semantic context in PebbloRetrievalQA (URL: https://github.com/langchain-ai/langchain/pull/27581)
- docs: anchor redirect (URL: https://github.com/langchain-ai/langchain/pull/27556)
- community: Update Polygon.io API (URL: https://github.com/langchain-ai/langchain/pull/27552)
- core[patch]: support handle_tool_error=(Exception, ...) tuple (URL: https://github.com/langchain-ai/langchain/pull/27547)
- community: add router chain based on the classify API from Jina AI (URL: https://github.com/langchain-ai/langchain/pull/27540)
- Update mongodb_atlas docs (URL: https://github.com/langchain-ai/langchain/pull/27530)
- community: sambanovacloud llm integration (URL: https://github.com/langchain-ai/langchain/pull/27526)
- partners/huggingface: fix HuggingFacePipeline model_id parameter (URL: https://github.com/langchain-ai/langchain/pull/27514)
- core: Add support for creating structured tools from partial functions (URL: https://github.com/langchain-ai/langchain/pull/27501)
- community: [fix] add missing tool_calls kwargs of delta message in openai adapter (URL: https://github.com/langchain-ai/langchain/pull/27492)
- community: add `AzureAIFaceAnalysisTool` (URL: https://github.com/langchain-ai/langchain/pull/27481)
- community: add `AzureAIContentSafetyChain` (URL: https://github.com/langchain-ai/langchain/pull/27480)
- Adding Firecrawl webscraper tool to langchain-community (URL: https://github.com/langchain-ai/langchain/pull/27476)

## Latest Release
- Name: langchain-core==0.3.13
- Tag: langchain-core==0.3.13
- Published At: 2024-10-25 19:36:18 +0000 UTC
- URL: https://github.com/langchain-ai/langchain/releases/tag/langchain-core%3D%3D0.3.13