<h1 align="center">Streaming data pipeline from producer to Clickhouse</h1>
<p align="center">
<a href="https://github.com/nahumsa/streaming-pipeline-clickhouse/actions"><img alt="Actions Status" src="https://github.com/nahumsa/streaming-pipeline-clickhouse/actions/workflows/go.yml/badge.svg"></a>
</p>

The main idea of the project is to create a streaming data pipeline that receives an HTTP request and sends it to
[clickhouse](https://clickhouse.com/). The first step is to add directly to clickhouse and see if it supports the load,
the second step is to add Kafka as a queue to write data to clickhouse, in this way we can handle more load than
writing directly to the server.

The reason for creating the API is two-fold:

1. We have a clear interface between the data producer and our data pipeline, this leads to data quality checks when the data is ingested, thus leading to higher data quality.
2. When the API is created, the data producer will only care about the interface and will not need to care what happens inside the API and what is done to handle the data.

Both points leads to scalability and reliability of data received by data consumers.

## Environment Variables

To run this application, you need to set up the following environment variables:

| Variable Name     | Description                                     |
|-------------------|-------------------------------------------------|
| `CLICKHOUSEHOST` | The hostname or IP address of the Clickhouse server |
| `CLICKHOUSEDB` | The database name          |
| `CLICKHOUSEUSERNAME` | Username for the clickhouse database          |
| `CLICKHOUSEUSERPASS` | Password for the clickhouse database          |

## Running the Application

To run the application, follow these steps:

1. **Clone the repository**:

    ```bash
    git clone https://github.com/nahumsa/streaming-pipeline-clickhouse.git
    cd streaming-pipeline-clickhouse
    # setup the clickhouse database locally
    docker compose up
    ```

2. **Set up environment variables**:
    You can set the environment variables directly in your shell or create a `.env` file in the root directory of your project:

    ```bash
    export CLICKHOUSEHOST="localhost:9000"
    export CLICKHOUSEDB="default"
    export CLICKHOUSEUSERNAME="default"
    export CLICKHOUSEPASS=" "
    ```

    Or create a `.env` file:

    ```dotenv
    CLICKHOUSEHOST="localhost:9000"
    CLICKHOUSEDB="default"
    CLICKHOUSEUSERNAME="default"
    CLICKHOUSEPASS=" "
    ```

    and run:

    ```bash
    export $(cat .env | xargs)
    ```

3. **Run the application**:

    ```bash
    go run main.go
    ```

Now, your application should be running and ready to receive HTTP requests to process data and send it to Clickhouse.
