package provider

//func Chunking(fileToBeChunkedPath, filename string) {
//
//	//fileToBeChunkedPath := "./somebigfile"
//
//	file, err := os.Open(fileToBeChunkedPath)
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	defer file.Close()
//
//	fileInfo, _ := file.Stat()
//
//	var fileSize int64
//	fileSize = fileInfo.Size()
//
//	const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement
//
//	// calculate total number of parts the file will be chunked into
//
//	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
//
//	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
//
//	for i := uint64(0); i < totalPartsNum; i++ {
//
//		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
//		partBuffer := make([]byte, partSize)
//
//		file.Read(partBuffer)
//
//		// write to disk
//		fileName := filename + strconv.FormatUint(i, 10)
//		NewChunkedFile, err := os.Create(fileName)
//
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		err = srv.StorageProvider.UploadV3(r.Context(), &chunk, bucketName, file, filePath, fileName)
//		if err != nil {
//			w.Write([]byte(fmt.Sprintf(`upload image on google cloud  %v`, err)))
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		//// write/save buffer to disk
//		//ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
//		//
//		//fmt.Println("Split to : ", fileName)
//	}
//}
