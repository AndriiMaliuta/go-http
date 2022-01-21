package _func

import (
	"net/http"
)

func CustHandle() {
	http.Handle("", nil)
}

//func cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
//
//	c := make(chan image.Image)
//	Returns
//	DB
//	struct instead of map
//		sp
//		:=
//		image.Point{
//		0, 0
//	}
//		channel
//		go func () {
//		Creates anonymous
//		newimage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
//		goroutine
//		for y := y1
//		y < y2
//		y = y + tileSize {
//		for x := x1
//		x < x2
//		x = x + tileSize {
//		r, g, b, _ := original.At(x, y).RGBA()
//		color := [3]float64{
//		float64(r), float64(g), float64(b)
//	}
//		nearest := db.nearest(color)
//		Calls nearest
//		file, err := os.Open(nearest)
//		method on DB to
//		if err == nil {
//		get best-fitting tile
//		img, _, err := image.Decode(file)
//		if err == nil {
//		t := resize(img, tileSize)
//		tile := t.SubImage(t.Bounds())
//		tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
//		draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
//	} else {
//		fmt.Println("error:", err)
//	}
//	} else {
//		fmt.Println("error:", nearest)
//	}
//		file.Close()
//	}
//	}
//		c <- newimage.SubImage(newimage.Rect)
//	}()
//		return c
//	}
