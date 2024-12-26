package main

import (
	_ "encoding/base64"
	"fmt"
	"image"
	"image/color"

	// "image/jpeg"

	vidio "github.com/AlexEidt/Vidio"

	// "image/color"
	// "image/jpeg"
	"os"
	_ "strings"
)

func main(){
	fmt.Println("complete image data:")
	videoGrayScale()
	// test()
}

func videoGrayScale(){
	file,_ := os.Open("sunset.mp4")
	defer file.Close()
	video,_ := vidio.NewVideo("sunset.mp4")
	totalFrames := video.Frames()
	fmt.Printf("total frames in the video: %v\n",totalFrames)
	var frames  = make([][]*image.RGBA,totalFrames)
	
	for i := 0; i < totalFrames; i++ {
		frames[i],_ = video.ReadFrames(i)
	}
	
	// fmt.Printf("the frame is: %v\n",len(frames))
	grayFrames := make([]*image.RGBA,0)


	for i := 0; i < totalFrames; i++ {
		for _,frame := range frames[i] {
			// fmt.Print(index)
			bounds := frame.Bounds()
			grayFrame := image.NewRGBA(bounds)
	
			for y := bounds.Min.Y ; y<bounds.Max.Y ; y++ {
				for x := bounds.Min.X ; x<bounds.Max.X ; x++ {
					r,g,b,_ := frame.At(x,y).RGBA()
					grayfColor := uint8(float64(r>>8)*0.299+float64(g>>8)*0.587+float64(b>>8)*0.114)
					grayFrame.Set(x,y,color.Gray{Y: grayfColor})
				}
			}
			grayFrames = append(grayFrames, grayFrame)
		}
	}
	
	

	// fmt.Printf("total frames in the grayvid: %v\n",len(frames))


	newfile,_ := os.Create("graysunset.mp4")
	options := vidio.Options{
		FPS: video.FPS(),
		Bitrate: video.Bitrate(),
	}

	newvidWriter,_ := vidio.NewVideoWriter("graysunset.mp4",video.Width(),video.Height(),&options)
	// fmt.Printf("firstframe: %v\n",grayFrames[0].At(0,0))
	// allframes := make([]byte,0)
	// for _,frame := range grayFrames {
	// 	allframes = append(allframes, frame.Pix...)
	// }
	// fmt.Printf("%v",len(allframes))

	for i := 0; i < totalFrames; i++ {
		newvidWriter.Write(grayFrames[i].Pix)
	}


	defer video.Close()
	defer newfile.Close()
}

// func imageGrayScale(){
// 	file,_ := os.Open("hot-air-balloon.jpg")
	
// 	img,_,_ := image.Decode(file)
// 	bounds := img.Bounds()
// 	grayimg := image.NewRGBA(img.Bounds())

// 	for y := bounds.Min.Y ; y<bounds.Max.Y ; y++ {
// 		for x := bounds.Min.X ; x<bounds.Max.X ; x++ {
// 			r,g,b,_ := img.At(x,y).RGBA()
// 			gray := uint8(float64(r>>8)*0.299+float64(g>>8)*0.587+float64(b>>8)*0.114)
// 			grayimg.Set(x,y,color.Gray{Y: gray})
// 		}
// 	}



// 	newf,_ := os.Create("newImage.jpg")
// 	jpeg.Encode(newf,grayimg,nil)
	
// 	defer file.Close()
// 	defer newf.Close()
// }

// func test(){
// 	newf,_ := os.Create("video.mp4")
// 	defer newf.Close()
// 	video, _ := vidio.NewVideo("video.mp4")

// 	img := image.NewRGBA(image.Rect(0, 0, video.Width(), video.Height()))
// 	video.SetFrameBuffer(img.Pix)

// 	frame := 0
// 	for video.Read() {
// 		f, _ := os.Create(fmt.Sprintf("%d.jpg", frame))
// 		jpeg.Encode(f, img, nil)
// 		f.Close()
// 		frame++
// 	}
// }