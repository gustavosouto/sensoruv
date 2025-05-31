# UVTrack Backend - Plataforma de Coleta e Visualização de Dados UV

Este repositório contém os componentes de backend e a configuração da plataforma de software para o sistema **UVTrack**, responsável por receber, processar, armazenar e visualizar os dados de exposição à radiação ultravioleta (UV) enviados pelo dispositivo vestível.

Este backend complementa o firmware e hardware localizados no repositório principal: [gustavosouto/UVTrack](https://github.com/gustavosouto/UVTrack).

## ☁️ Arquitetura da Plataforma

A plataforma é composta por um conjunto de serviços orquestrados, tipicamente via Docker Compose, para fornecer uma solução completa de ingestão e visualização de dados IoT:

1.  **Broker MQTT (Ex: Mosquitto):** Atua como o ponto central de comunicação, recebendo as mensagens em formato JSON publicadas pelo dispositivo ESP8266 via protocolo MQTT.
2.  **Serviço de Ingestão (Go):** Uma aplicação desenvolvida em Go (`main.go`) que subscreve ao tópico MQTT relevante no broker. É responsável por:
    *   Receber as mensagens JSON contendo os dados do sensor UV.
    *   Realizar o parsing e a validação dos dados.
    *   Escrever os dados validados no banco de dados InfluxDB.
3.  **Banco de Dados (InfluxDB):** Um banco de dados de séries temporais otimizado para armazenar eficientemente os dados de UV recebidos, indexados por tempo.
4.  **Plataforma de Visualização (Grafana):** Uma ferramenta web que se conecta ao InfluxDB para permitir a criação e visualização de dashboards interativos, exibindo o histórico e as tendências dos dados de exposição UV.

## 🛠️ Tecnologias Utilizadas

*   **Linguagem do Serviço:** Go (Golang)
*   **Protocolo de Comunicação:** MQTT
*   **Formato de Dados:** JSON
*   **Banco de Dados:** InfluxDB (Time Series Database)
*   **Visualização:** Grafana
*   **Orquestração:** Docker, Docker Compose
*   **Broker MQTT:** Mosquitto (ou similar)
*   **Controle de Versão:** Git, GitHub

## 🚀 Configuração e Execução (Usando Docker Compose)

O arquivo `docker-compose.yml` neste repositório facilita a configuração e execução de toda a pilha de backend.

**Pré-requisitos:**

*   Docker instalado ([https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/))
*   Docker Compose instalado ([https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/))

**Passos:**

1.  **Clonar o Repositório:**
    ```bash
    git clone https://github.com/gustavosouto/sensoruv.git
    cd sensoruv
    ```
2.  **Configuração (se necessário):**
    *   Verifique o arquivo `docker-compose.yml` para quaisquer configurações específicas de portas ou volumes.
    *   Verifique o arquivo de configuração do Mosquitto (`mosquitto-config/mosquitto.conf`) se precisar ajustar as configurações do broker (ex: autenticação).
    *   O serviço Go (`main.go`) pode precisar de variáveis de ambiente para configurar a conexão com o MQTT e InfluxDB (verificar o código ou o `docker-compose.yml`).
3.  **Iniciar os Serviços:**
    ```bash
    docker-compose up -d
    ```
    Este comando irá baixar as imagens necessárias (Mosquitto, InfluxDB, Grafana), construir a imagem para o serviço Go (se definido no compose) e iniciar todos os contêineres em segundo plano (`-d`).
4.  **Verificar os Serviços:**
    ```bash
    docker-compose ps
    ```
    Verifique se todos os contêineres (mosquitto, influxdb, grafana, go-service) estão em execução (`Up`).
5.  **Acessar o Grafana:**
    *   Abra seu navegador e acesse o Grafana na porta configurada (geralmente `http://localhost:3000`).
    *   Faça login (credenciais padrão costumam ser `admin`/`admin`, mas podem ser alteradas no `docker-compose.yml`).
    *   Configure o InfluxDB como uma fonte de dados (DataSource), apontando para o serviço InfluxDB (ex: `http://influxdb:8086`) e especificando o bucket/database correto.
    *   Crie ou importe dashboards para visualizar os dados UV que estão sendo enviados pelo dispositivo e armazenados no InfluxDB.
6.  **Parar os Serviços:**
    ```bash
    docker-compose down
    ```
    Este comando para e remove os contêineres definidos no `docker-compose.yml`.

## 🔌 Integração com o Dispositivo UVTrack

O firmware do dispositivo ESP8266 (no repositório [gustavosouto/UVTrack](https://github.com/gustavosouto/UVTrack)) deve ser configurado para:

*   Conectar-se à mesma rede Wi-Fi que hospeda o backend ou ter acesso à internet se o backend estiver na nuvem.
*   Apontar para o endereço IP ou hostname e porta do **Broker MQTT** configurado neste backend.
*   Utilizar as credenciais MQTT (usuário/senha), se configuradas no broker.
*   Publicar os dados no tópico MQTT esperado pelo serviço Go (ex: `uvtrack/data`).

## 📄 Código Fonte do Serviço Go

O código fonte principal do serviço de ingestão de dados está localizado em `main.go`. Ele utiliza bibliotecas Go para:

*   Conectar-se ao Broker MQTT (ex: `eclipse/paho.mqtt.golang`).
*   Subscrever a tópicos MQTT.
*   Receber e processar mensagens (parsing de JSON).
*   Conectar-se e escrever dados no InfluxDB (ex: `influxdata/influxdb-client-go`).

## 👥 Equipe

Consulte o README do repositório principal [gustavosouto/UVTrack](https://github.com/gustavosouto/UVTrack) para informações sobre a equipe de desenvolvimento.

## 🤝 Contribuições

Contribuições para melhorar o backend, a configuração ou a documentação são bem-vindas.

