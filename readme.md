Ini adalah project cudos untuk mencari fraud transaksi, ada 3 keriteria terkait goal bnisnis ini:
1. Mencari setiap transaksi, dari per transaksi dipastikan diorder oleh user, dan kita harus melihat dari setiap transaksi user 1 jam kebelakang, jadi kuncinya ada di userId, dan transaction_date yang dimana nanti pencarian jumlah transaksi where userId = ?  and transaction_date between satu jam lalu dan saat ini. nah kita bisa dapet tuh jumlah data per user dari setiap transaksi berdasarkan kriteria waktu.
2. yang kedua saya jujur gak paham, sy masih mencari tahu arti dari z-score, tapi bukan berarti sy gak bisa developnya. saya gak paham konteknya.
3. Mencari nilai dari transaksi baru, dan mengkalkukasi dari rata2 amount transaksi user tersebut, apakah user tesebut tiba2 memasukan amount yang tidak lazim alias biasanya rata2 dia masukin amount diangka 1jt tiba2 dia masukan nilai diangka 5 jt, maka masuk dalam high fraud, yang perlu di cek kembali.

Dalam kali ini saya menggunakan Golang dengan handler Fiber. Design Pattern ada controller, repository dan model
- Model fungsinya sebagai agregator enitity antara gorm dan SQL
- Repository kumpulan perintah query
- Controller Logic base



Clone this repo
- `go mod download`
- `go mod tidy`
- `cp .env.example .env`
- Fill `.env` with your credential
- `go run main.go`


Sebelumnya Maaf jauh dari kata sempurna, mungkin butuh waktu lebih panjang lagi untuk lebih rapih dan cek dari segi performa latency