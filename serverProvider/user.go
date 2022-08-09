package serverProvider

const (
	maxChunkSize = int64(5 << 20) // 5MB

	uploadDir = "./"
)

// type Chunk struct {
//	UploadID      string // unique id for the current upload.
//	ChunkNumber   int32
//	TotalChunks   int32
//	TotalFileSize int64 // in bytes
//	Filename      string
//	Data          io.Reader
//	UploadDir     string
//	ByteData      []byte
// }

//func (srv *Server) upload(w http.ResponseWriter, r *http.Request) {
//	err := srv.ProcessChunk(r)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	//err = srv.StorageProvider.Upload(r.Context(), chunk)
//	//if err != nil {
//	//	return
//	//}
//
//	w.Write([]byte("chunk processed check"))
//}

//func (srv *Server) uploadV2(w http.ResponseWriter, r *http.Request) {
//	var chunk models.Chunk
//	bucketName := "upload-image-c4c6e.appspot.com"
//
//	defer func(req *http.Request) {
//		if req.MultipartForm != nil { // prevent panic from nil pointer
//			if err := req.MultipartForm.RemoveAll(); err != nil {
//				logrus.Errorf("Unable to remove all multipart form. %+v", err)
//			}
//		}
//	}(r)
//
//	//r.Body = http.MaxBytesReader(w, r.Body, 51<<20)
//
//	if err := r.ParseMultipartForm(51 << 10); err != nil {
//		if err == io.EOF || err.Error() == "multipart: NextPart: unexpected EOF" {
//			logrus.Warn("EOF")
//		} else {
//			logrus.Errorf("[ParseMultipartForm] %s", err.Error())
//		}
//
//		return
//	}
//
//	// start readings parts
//	// 1. upload id
//	// 2. chunk number
//	// 3. total chunks
//	// 4. total file size
//	// 5. file name
//	// 6. chunk data
//
//	// 1
//	chunk.UploadID = r.FormValue("upload_id")
//
//	// dir to where we store our chunk
//	chunk.UploadDir = fmt.Sprintf("%s/%s", uploadDir, chunk.UploadID)
//
//	// 2
//	temp := r.FormValue("chunk_number")
//
//	parsedChunkNumber, err := strconv.ParseInt(temp, 10, 32)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user no. 2 %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	chunk.ChunkNumber = int32(parsedChunkNumber)
//
//	// 3
//	temp = r.FormValue("total_chunks")
//
//	parsedTotalChunksNumber, err := strconv.ParseInt(temp, 10, 32)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user parse int  %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	chunk.TotalChunks = int32(parsedTotalChunksNumber)
//
//	// 4
//	temp = r.FormValue("total_file_size")
//
//	parsedTotalFileSizeNumber, err := strconv.ParseInt(temp, 10, 64)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user no. 4  parse int %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	chunk.TotalFileSize = parsedTotalFileSizeNumber
//
//	// 5
//	fileName := r.FormValue("file_name")
//
//	//	6
//
//	file, _, err := r.FormFile("data")
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user no.  6 form file %v`, err)))
//	}
//
//	defer func() {
//		if err = file.Close(); err != nil {
//			logrus.Errorf("Unable to close file multipart form. %+v", err)
//		}
//	}()
//
//	//part, err := reader.NextPart()
//	//
//	//if err != nil {
//	//	w.Write([]byte(fmt.Sprintf(`get part user no.  6 %v`, err)))
//	//}
//	//chunk.Data = part
//
//	//filePath := fmt.Sprintf(`chunks/%v/%v-%s`, chunk.Filename, time.Now().Unix(), chunk.Filename)
//
//	filePath := "./chunk"
//
//	srv.Chunking(r.Context(), bucketName, file, filePath, fileName)
//
//	url, err := srv.StorageProvider.GetSharableURL(bucketName, filePath, time.Hour*24*365)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get shareable URL %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Write([]byte(fmt.Sprintf(`%v`, url)))
//}
//
////func (srv *Server) handleCompletedChunk(w http.ResponseWriter, r *http.Request) {
////	type request struct {
////		UploadID string `json:"uploadId"`
////		Filename string `json:"filename"`
////	}
////
////	var payload request
////	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
////		http.Error(w, err.Error(), http.StatusBadRequest)
////		return
////	}
////
////	// validate payload
////
////	if err := CompleteChunk(payload.UploadID, payload.Filename); err != nil {
////		http.Error(w, err.Error(), http.StatusInternalServerError)
////		return
////	}
////
////	w.Write([]byte("file processed"))
////
////}
//
//func getPartUser(expectedPart string, reader *multipart.Reader, buf *bytes.Buffer) error {
//	part, err := reader.NextPart()
//	if err != nil {
//
//		return fmt.Errorf("failed reading %s part %w", expectedPart, err)
//	}
//
//	if part.FormName() != expectedPart {
//		return fmt.Errorf("invalid form name for part. Expected %s got %s", expectedPart, part.FormName())
//	}
//
//	if _, err := io.Copy(buf, part); err != nil {
//		return fmt.Errorf("failed copying %s part %w", expectedPart, err)
//	}
//
//	return nil
//}
//
//func (srv *Server) Chunking(ctx context.Context, bucketName string, filem multipart.File, fileToBeChunkedPath, filename string) {
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
//	fileInfo, err := file.Stat()
//
//	var fileSize int64
//	fileSize = fileInfo.Size()
//
//	const fileChunk = 1024 // 1 MB, change this to your requirement
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
//		_, err := os.Create(fileName)
//
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
//
//		fmt.Println("Split to : ", fileName)
//	}
//}
//
//func (srv *Server) uploadV3(w http.ResponseWriter, r *http.Request) {
//	var chunk models.Chunk
//	bucketName := "upload-image-c4c6e.appspot.com"
//
//	defer func(req *http.Request) {
//		if req.MultipartForm != nil { // prevent panic from nil pointer
//			if err := req.MultipartForm.RemoveAll(); err != nil {
//				logrus.Errorf("Unable to remove all multipart form. %+v", err)
//			}
//		}
//	}(r)
//
//	//r.Body = http.MaxBytesReader(w, r.Body, 51<<20)
//
//	if err := r.ParseMultipartForm(51 << 10); err != nil {
//		if err == io.EOF || err.Error() == "multipart: NextPart: unexpected EOF" {
//			logrus.Warn("EOF")
//		} else {
//			logrus.Errorf("[ParseMultipartForm] %s", err.Error())
//		}
//
//		return
//	}
//
//	// start readings parts
//	// 1. upload id
//	// 2. chunk number
//	// 3. total chunks
//	// 4. total file size
//	// 5. file name
//	// 6. chunk data
//
//	// 1
//	chunk.UploadID = r.FormValue("upload_id")
//
//	// dir to where we store our chunk
//	chunk.UploadDir = fmt.Sprintf("%s/%s", uploadDir, chunk.UploadID)
//
//	// 2
//	temp := r.FormValue("chunk_number")
//
//	parsedChunkNumber, err := strconv.ParseInt(temp, 10, 32)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user no. 2 %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	chunk.ChunkNumber = int32(parsedChunkNumber)
//
//	// 3
//	temp = r.FormValue("total_chunks")
//
//	parsedTotalChunksNumber, err := strconv.ParseInt(temp, 10, 32)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user parse int  %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	chunk.TotalChunks = int32(parsedTotalChunksNumber)
//
//	// 4
//	temp = r.FormValue("total_file_size")
//
//	parsedTotalFileSizeNumber, err := strconv.ParseInt(temp, 10, 64)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user no. 4  parse int %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	chunk.TotalFileSize = parsedTotalFileSizeNumber
//
//	// 5
//	fileName := r.FormValue("file_name")
//
//	//	6
//
//	file, _, err := r.FormFile("data")
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get part user no.  6 form file %v`, err)))
//	}
//
//	defer func() {
//		if err = file.Close(); err != nil {
//			logrus.Errorf("Unable to close file multipart form. %+v", err)
//		}
//	}()
//
//	//part, err := reader.NextPart()
//	//
//	//if err != nil {
//	//	w.Write([]byte(fmt.Sprintf(`get part user no.  6 %v`, err)))
//	//}
//	//chunk.Data = part
//
//	//filePath := fmt.Sprintf(`chunks/%v/%v-%s`, chunk.Filename, time.Now().Unix(), chunk.Filename)
//
//	filePath := "./chunk"
//
//	srv.Chunking(r.Context(), bucketName, file, filePath, fileName)
//
//	url, err := srv.StorageProvider.GetSharableURL(bucketName, filePath, time.Hour*24*365)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`get shareable URL %v`, err)))
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Write([]byte(fmt.Sprintf(`%v`, url)))
//
//}
