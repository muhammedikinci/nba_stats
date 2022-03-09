<br />
<p align="center">
  <h3 align="center">NBA Simulator</h3>
  <p align="center">
    NBA Simulator Project With Go
    <br />
    <br />
  </p>
</p>

<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#installation-without-docker">Installation Without Docker</a></li>
      </ul>
    </li>
    <li><a href="#api-reference">API Reference</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## Getting Started

### Prerequisites

* Golang
* Docker

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/muhammedikinci/nba_stats
   cd ./nba_stats
   ```
2. Build docker compose file
   ```sh
   docker compose build --no-cache
   ```
3. Run
   ```sh
   docker compose up
   ```

- API can be restarted many times until MySQL is ready to accept the connection
  ```yaml
  restart: on-failure
  ```

### Installation Without Docker

This action needs the installation of PostgreSQL manually.

1. Clone the repo
   ```sh
   git clone https://github.com/muhammedikinci/nba_stats
   cd ./nba_stats
   ```
2. Install Go dependencies
   ```sh
   go mod download
   ```

3. Start MYSQL Service and Create `nba` database

4. Change database root password from `repository`

4. Start API without build
   ```sh
   go run .
   ```

## API Reference

### GET /simulation

Getting current simulation status 

```sh
curl --location --request GET '/simulation'
```

### POST /simulation/start

Starting simulation

```sh
curl --location --request POST '/simulation/start'
```

### DELETE /simulation/stop

Stopping simulation

```sh
curl --location --request DELETE '/simulation/stop'
```

### GET /simulation/current-round

Getting current week stats.

```sh
curl --location --request GET '/simulation/current-round'
```

## Contact

Muhammed İKİNCİ - muhammedikinci@outlook.com