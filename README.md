<h1 align="center">Streaming data pipeline from producer to Clickhouse</h1>
<p align="center">
<a href="https://github.com/nahumsa/streaming-pipeline-clickhouse/actions"><img alt="Actions Status" src="https://github.com/nahumsa/streaming-pipeline-clickhouse/actions/workflows/go.yml/badge.svg"></a>
</p>

The main idea of the project is to create a streaming data pipeline that receives an HTTP request and sends it to
[clickhouse](https://clickhouse.com/). The first step is to add directly to clickhouse and see if it supports the load,
the second step is to add Kafka as a queue to write data to clickhouse, in this way we can handle more load than
writing directly to the server.

There are two main points in creating an API for receiving the data:

1. We have a clear interface between the data producer and our data pipeline, this leads to data quality checks when the data is ingested, thus leading to higher data quality.
2. When the API is created, the data producer will only care about the interface and will not need to care what happens inside the API and what is done to handle the data.

Both points leads to scalability and reliability of data received by data producers.

