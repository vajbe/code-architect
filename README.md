# Code Architect

This repository demonstrates various software architecture patterns, data pipeline implementations, and system design components using Go.

## Repository Structure

The codebase is organized into three main sections:

### 1. Flow Ingestion (`flow-ingestion/`)
A complete data ingestion and processing pipeline.
- **API**: REST API for data ingestion.
- **Flow Processor**: Backend processor for handling data flows.
- **Flow Generator**: Python utilities for generating test traffic and data flows.

### 2. Kafka Basics (`kafka-basics/`)
Foundational examples for Apache Kafka integration in Go.
- **Producer**: Example implementation of a Kafka producer.
- **Consumer**: Example implementation of a Kafka consumer.

### 3. Use Cases (`use-cases/`)
Standalone implementations of common system design patterns and components.

- **01.task-processor**
  A robust task processing service demonstrating:
  - Asynchronous task handling.
  - Observability integration with **Prometheus** and **Grafana**.
  - Docker Compose setup for easy deployment.

- **02.rate-limits**
  Implementation of rate limiting algorithms to control traffic flow and protect resources.

- **03.ttl-cache**
  A custom **Time-To-Live (TTL)** cache implementation for temporary data storage with automatic expiration.

- **04.lru-cache**
  A custom **Least Recently Used (LRU)** cache implementation for efficient memory management of frequently accessed data.

## Getting Started

Each subdirectory typically contains its own `go.mod` file and can be run independently. Refer to the specific sub-project directories for more detailed instructions on running and testing specific components.

### Prerequisites
- Go (1.20+)
- Docker & Docker Compose (for `flow-ingestion`, `kafka-basics`, and `task-processor`)
- Python 3.x (for `flow-generator`)
