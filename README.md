RAG Architecture:

1. We can generate vector embedding for the companies/organisation documents using openAI embedding API or other free embedding API’s in the market.
2. From my research we can store the vector embeddings in database like Milvus DB or Postgres PG vector which supports vector similarity search. 
3. As the user prompts questions in the chatbot , we can turn the queries into vector embedding again using openAI embedding API. 
4. Then we can run the similarity vector search to find the similar context based on the user query that is stored in the database as vectors.
5. Having these similar documents in the context of the Open AI chat completion API we get the answer for the user query and return the response back to the user in UI. 

Referred links: https://platform.openai.com/docs/guides/embeddings
https://github.com/openai/openai-go
https://www.nvidia.com/en-in/glossary/retrieval-augmented-generation/
