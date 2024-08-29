# OXTRO-API
OXTRO-API adalah sebuah proyek yang menyediakan API untuk mengelola pengguna dan organisasi. Proyek ini menggunakan mockery untuk menghasilkan mock dari interface yang ada.

## Dokumentasi

### Instalasi

Untuk menginstal proyek ini, ikuti langkah-langkah berikut:

1. Clone repositori ini.
   ```bash
   git clone <URL_REPOSITORI>
   ```
2. Masuk ke direktori proyek.
   ```bash
   cd OXTRO-API
   ```
3. Instal dependensi yang diperlukan.
   ```bash
   go mod tidy
   ```
4. Install mockery
   ```bash
   go install github.com/vektra/mockery/v2@v2.45.0
   ```
5. Install connectrpc
   ```bash
   go install github.com/bufbuild/buf/cmd/buf@latest
   go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
   ```
6. Install Database Migration

   ```bash
   go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

   Install Driver Postgres

   ```bash
   go install -tags 'postgres,mysql,sqlite' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

   Create Database Migration

   ```bash
   make migration-create name=create_users_table
   ```

   Up Migration

   ```bash
   make migration-up version=1
   ```

   Down Migration

   ```bash
   make migration-down version=1
   ```

   Pastikan database migration tidak dirty state, jika database migration dirty state maka akan terjadi error saat running database migration.

   Solusi jika database migration dirty state, cari table yang bernama `schema_migrations` lalu mundur 1 versi dan rubah dirty menjadi false

### How to Run

Untuk menjalankan proyek ini, gunakan perintah berikut:

- Windows
  ```bash
  make run-dev
  ```
- Linux
  ```bash
  make run
  ```

### Tech Stack

- Golang
- Gorm
- Postgres
- Docker
- Docker Compose
- Mockery
- Connectrpc
- Gin Gonic
- Logrus

### Example Handler
Buatlah file proto di directory `proto`. Lalu sesuaikan nama directory dengan fitur yang di butuhkan.

- buat directory `organization/v1` 
- buat file `organization.proto`
- buatkalah schema proto yang ingin di buat
  ```proto
  syntax = "proto3";

  package organization.v1;

  message RequestOrganization {
    string id = 1;
    string name = 2;
  }

  message ResponseOrganization{
   string id = 1;
   string name = 2;
  }

  Service Organization {
    rpc CreateOrganization(RequestOrganization) returns (ResponseOrganization);
  }
  ```
- generate file proto dengan perintah
  ```bash
  make gen-proto
  ```