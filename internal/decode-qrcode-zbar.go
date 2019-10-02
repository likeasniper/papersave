// -build ignore

package papersave

import (
	// "fmt"

	"image"
	_ "image/jpeg"
	_ "image/png"
	"errors"
	"os"
	"github.com/PeterCxy/gozbar"
)

func decodeQRCodesZbar(path string)  (data string, err error)  {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return
	}
	
	defer func() {
        if r := recover(); r != nil {
			err = errors.New("Method zbar failed")
        }
    }()

	i, _, _ := image.Decode(file)
	img := zbar.FromImage(i)

	s := zbar.NewScanner()
	s.SetConfig(0, zbar.CFG_ENABLE, 1)
	s.Scan(img)
	defer s.Destroy()

	// Order is inverted here
	var tmp []string
	img.First().Each(func(str string) {
		tmp = append(tmp, str)
	})
	for i := len(tmp)-1 ; i >= 0; i-- {
		data += tmp[i]
	}
	return
}


func init() {
	decodeFuncs[ZBAR] = decodeQRCodesZbar
}
