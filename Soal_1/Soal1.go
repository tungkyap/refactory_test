package main

import (
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "bufio"
  // "strings"
  "strconv"
  "time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() {
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func callClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

var namaWarung, namaKasir string
var tanggal = time.Now()
var jumlahBeli, totalHarga, bayarTotal int
var jumlahProduk, hargaProduk, subTotal [10]int
var produk [10]string

func masukanDataWarung() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("Nama Warung : ")
    scanner.Scan()
    namaWarung = scanner.Text()
    dataWarung(namaWarung)
}

func dataWarung(warung string) string {
    return warung
}

func masukanDataKasir() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fmt.Printf("Nama Kasir : ")
    scanner.Scan()
    namaKasir = scanner.Text()
}

func masukanJumlahBeli() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("Masukan Jumlah Pembelian: ")
    scanner.Scan()
    jumlahBeli, _ = strconv.Atoi(scanner.Text())
    fmt.Println()

    for i := 0; i < jumlahBeli; i++ {
        fmt.Println()
        fmt.Println("Masukan Produk ke-"+strconv.Itoa(i+1))
        fmt.Printf("Nama Produk: ")
        scanner.Scan()
        produk[i] = scanner.Text()
        fmt.Printf("Jumlah Produk: ")
        scanner.Scan()
        jumlahProduk[i], _ = strconv.Atoi(scanner.Text())
        fmt.Printf("Harga Produk: ")
        scanner.Scan()
        hargaProduk[i], _ = strconv.Atoi(scanner.Text())

        subTotal[i] = jumlahProduk[i] * hargaProduk[i]
        totalHarga = totalHarga + subTotal[i]
    }
}

func cetakStruk() {
    var test string = namaWarung
    var tanggal = time.Now()
    fmt.Println("nilai =",dataWarung(test))
    fmt.Println("Tanggal :", tanggal.Format("2006/01/02 15:04:05"))
}

func masukanSemuaData() {
    var pilih string
    fmt.Println("Silakan, Masukan Data")
    masukanDataWarung()
    masukanDataKasir()
    masukanJumlahBeli()

    fmt.Println("nilai",namaWarung)

    fmt.Println()
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Apakah anda ingin cetak Data? (y/n)")
    scanner.Scan()
    pilih = scanner.Text()
    switch pilih {
        case "y" :
          cetakStruk()
        case "t" :
          tampilMenuProgram()
        default :
          fmt.Println("Pilihan invalid!")
          time.Sleep(1 * time.Second)
          os.Exit(0)

    }

}

func tampilMenuProgram() {
    callClear()
    fmt.Println("Program Cetak Struk")
    fmt.Println("-------------------")
    fmt.Println("1. Masukan Data ke Struk")
    fmt.Println("2. Keluar")
    fmt.Println("-------------------")
    fmt.Println()
    pilihMenuProgram()
}

func pilihMenuProgram() {
    var pilih int
    fmt.Printf("Pilih menu : ")
    fmt.Scanf("%d", &pilih)
    fmt.Println()
    switch pilih {
        case 1:
          masukanSemuaData()
        case 2:
          fmt.Println("Keluar dalam 2 detik...\n")
          time.Sleep(2 * time.Second)
          os.Exit(0)
        default:
          fmt.Printf("Pilihan menu tidak valid!")
          tampilMenuProgram()
    }
}

func main() {
  tampilMenuProgram()
}
