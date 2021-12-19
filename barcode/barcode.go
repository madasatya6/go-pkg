package barcode

import (
	"image/png"
	"os"
	"fmt"
	"log"
	"time"
	"io/ioutil"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

/*** 
* @author madasatya6
* - image barcode jika expired akan dihapus otomatis 
*   sehingga tidak menumpuk sebagai sampah
* - jika folder belum ada maka secara otomatis create folder tsb 'barcode'
*   sehingga tidak menyebabkan error
* - hasil image bernama tanggal untuk memudahkan pengecekan tanggal kadaluarsa
* - image dianggap kadaluarsa jika sudah 24 jam yang lalu terakhir dibuat 
*/

var Directory = "app/static/barcode/"
var Extension = ".png"
var Zone = "Asia/Jakarta"

/* @return image name with error
*  Main Function : GenerateImage()
*/
func GenerateImage(key string, width, height int) (string, error) {

	var imageName string = GenerateRandomName("qrcode")

	//hapus tumpukan barcode yg expired
	go DeleteBarcodeIfExpired()

	// Create the barcode
	qrCode, err := qr.Encode(key, qr.M, qr.Auto)
	if err != nil {
		return imageName, err 
	}

	qrCode, err = barcode.Scale(qrCode, width, height)
	if err != nil {
		return imageName, err 
	}

	if err := CreateOrCheckDirectory(Directory); err != nil {
		return imageName, err 
	}

	file, err := os.Create(Directory + imageName + Extension)
	if err != nil {
		return imageName, err 
	}
	defer file.Close()

	err = png.Encode(file, qrCode)
	if err != nil {
		return imageName, err 
	}

	return imageName, nil 
}

func DeleteBarcodeIfExpired() error {
	 
	var images, err = ioutil.ReadDir(Directory)
	if err != nil {
		logErr(err)
		return err 
	}

	for _, image := range images {
		imageName := image.Name()
		explode := strings.Split(imageName, "-")
		dateTime := explode[0]
		if IsExpired(dateTime) {
			if err := os.Remove(Directory + imageName); err != nil {
				logErr(err)
			}
		}	
	}
	
	return nil 
}

func CheckDirectory(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil 
	}

	if os.IsNotExist(err) {
		return false, nil 
	}

	return false, err  
}

func CreateOrCheckDirectory(path string) error {
	isExist, err := CheckDirectory(path)
	if err != nil {
		return err 
	}

	if !isExist {
		//create directory
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err 
		}
	}

	return nil 
}

func GenerateRandomName(customName string) string {

	if customName != "" {
		customName = "-" + customName
	}

	timeZone, _ := time.LoadLocation(Zone)
	timeStamp := time.Now().In(timeZone).UnixNano() / int64(time.Millisecond) //millisecond

	return fmt.Sprintf("%v-%v%s", time.Now().Format("20060102150405"), timeStamp, customName)
}

func IsExpired(dateTime string) bool {

	//cek apakah parameter date time kurang dari 1 hari lebih 
	timeZone, err := time.LoadLocation(Zone)
	if err != nil {
		logErr(err)
		return false
	}

	now := time.Now().In(timeZone)
	cekTime, err := time.Parse("20060102150405", dateTime)
	if err != nil {
		logErr(err)
		return false
	} 

	status := now.After(cekTime)

	return status 
}

func logErr(err error) {
	log.Println("Barcode Error: ", err.Error())
}




