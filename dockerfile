# Gunakan image resmi Go versi terbaru berbasis Alpine
FROM golang:alpine

# Instal dependensi yang diperlukan untuk beberapa paket Go (jika ada)
RUN apk add --no-cache build-base git

# Set direktori kerja di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum terlebih dahulu (untuk memanfaatkan cache Docker)
COPY go.mod go.sum ./

# Unduh dependensi
RUN go mod tidy
RUN go mod download

# Salin seluruh file aplikasi setelah dependensi diunduh
COPY . .

# Ekspor port yang digunakan oleh aplikasi Go
EXPOSE 8080

# Jalankan aplikasi Go secara langsung
CMD ["go", "run", "cmd/main.go"]