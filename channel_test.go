package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// kalo pake chanel pastikan ada yang mengirim dan ada yang menerima
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	//close chanel mau error mau sukses terserah pake defer
	//setelah selesai semua baru pake close kalo pake defer
	defer close(channel)
	//masukin data ke channel
	//channel <- "Eko"
	//mengambil data ke channel
	//data := <-channel
	//fmt.Println(data)
	//langsung dikirim ke parameter
	//fmt.Println(<-channel)

	//close chanel
	//close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Achmad Sidik Fauzi"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeRespown(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Achmad Sidik Fauzi"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespown(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Achmad Sidik Fauzi"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	//ada ,1 itu adalah buffer yang mengakibatkan channel dimasukkan ke buffer
	//nilai buffer berarti nilai data chanel max ada 3 data
	channel := make(chan string, 3)
	defer close(channel)

	//sebelumnya ini pasti error karena channel tidak ada yang mengeluarkan, cuma di masukkan aja
	//channel <- "Fauzi"
	//channel <- "Sidik"

	//jika ingin menggunakan go routine
	go func() {
		channel <- "Fauzi"
		channel <- "Sidik"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// range bisa menggunakan perulangan dan di ambil oleh range biar tidak memasukkan data 1 1
// untuk menerima data yang tidak tentu jumlahnya
func TestRangeChanel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		//harus di close biar pada perulangan terakhir tidak kenak deadlock
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data", data)
	}
	fmt.Println("Selesai")
}

// select chanel, untuk chanel > 1
func TestSelectChannel(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	defer close(chanel1)
	defer close(chanel2)

	go GiveMeRespown(chanel1)
	go GiveMeRespown(chanel2)

	counter := 0
	//jika tidak menggunakan for maka hanya akan mengambil 1 chanel yang tercepat datanya
	for {
		select {
		case data := <-chanel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-chanel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	defer close(chanel1)
	defer close(chanel2)

	go GiveMeRespown(chanel1)
	go GiveMeRespown(chanel2)

	counter := 0
	//jika tidak menggunakan for maka hanya akan mengambil 1 chanel yang tercepat datanya
	for {
		select {
		case data := <-chanel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-chanel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
