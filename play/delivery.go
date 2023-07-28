package play

import (
   "154.pages.dev/encoding/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

// downloadUrl
func (a App_File_Metadata) Download_URL() (int, []byte) {
   return a.m.Bytes(4)
}

// fileType
func (a App_File_Metadata) File_Type() (int, uint64) {
   return a.m.Uvarint(1)
}

// downloadUrl
func (s Split_Data) Download_URL() (int, []byte) {
   return s.m.Bytes(5)
}

// id
func (s Split_Data) ID() (int, []byte) {
   return s.m.Bytes(1)
}

func (d Delivery) Additional_File() []App_File_Metadata {
   var files []App_File_Metadata
   for {
      // additionalFile
      i, file := d.m.Message(4)
      if i == -1 {
         return files
      }
      files = append(files, App_File_Metadata{file})
      d.m = d.m[i+1:]
   }
}

func (d Delivery) Split_Data() []Split_Data {
   var splits []Split_Data
   for {
      // splitDeliveryData
      i, split := d.m.Message(15)
      if i == -1 {
         return splits
      }
      splits = append(splits, Split_Data{split})
      d.m = d.m[i+1:]
   }
}

// AndroidAppDeliveryData
type Delivery struct {
   m protobuf.Message
}

// SplitDeliveryData
type Split_Data struct {
   m protobuf.Message
}

// AppFileMetadata
type App_File_Metadata struct {
   m protobuf.Message
}

type File struct {
   Package_Name string
   Version_Code uint64
}

func (f File) OBB(file_type uint64) string {
   var b []byte
   if file_type >= 1 {
      b = append(b, "patch"...)
   } else {
      b = append(b, "main"...)
   }
   b = append(b, '.')
   b = strconv.AppendUint(b, f.Version_Code, 10)
   b = append(b, '.')
   b = append(b, f.Package_Name...)
   b = append(b, ".obb"...)
   return string(b)
}

func (f File) APK(id string) string {
   var b []byte
   b = append(b, f.Package_Name...)
   b = append(b, '-')
   if id != "" {
      b = append(b, id...)
      b = append(b, '-')
   }
   b = strconv.AppendUint(b, f.Version_Code, 10)
   b = append(b, ".apk"...)
   return string(b)
}
// downloadUrl
func (d Delivery) Download_URL() (int, []byte) {
   return d.m.Bytes(3)
}

func (h Header) Delivery(doc string, vc uint64) (*Delivery, error) {
   req, err := http.NewRequest(
      "GET", "https://play-fe.googleapis.com/fdfe/delivery", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "doc": {doc},
      "vc": {strconv.FormatUint(vc, 10)},
   }.Encode()
   h.Set_Agent(req.Header)
   h.Set_Auth(req.Header) // needed for single APK
   h.Set_Device(req.Header)
   res, err := client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   // ResponseWrapper
   mes, err := protobuf.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   // payload
   _, mes = mes.Message(1)
   // deliveryResponse
   _, mes = mes.Message(21)
   // status
   switch _, status := mes.Uvarint(1); status {
   case 2:
      return nil, errors.New("geo-blocking")
   case 3:
      return nil, errors.New("purchase required")
   case 5:
      return nil, errors.New("invalid version")
   }
   // appDeliveryData
   _, mes = mes.Message(2)
   return &Delivery{mes}, nil
}
