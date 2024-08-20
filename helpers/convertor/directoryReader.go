package convertor

import (
   "fmt"
   "os"
   "errors"
)

func ReadDirectory(path string) ([]string, error) {
   directory := "./" + path // The current directory

   fmt.Printf("opening dir: %s\n", directory)
   files, err := os.Open(directory) //open the directory to read files in the directory
   if err != nil {
      fmt.Println("error opening directory:", err) //print error if directory is not opened
      return []string{}, errors.New("error opening directory")
   }
   defer files.Close()    //close the directory opened
   fmt.Printf("closed dir: %s\n", directory)
   
   fileInfos, err := files.Readdir(-1)  //read the files from the directory
   fmt.Printf("read files from dir: %s\n", directory)
   if err != nil {
      fmt.Println("error reading directory:", err)  //if directory is not read properly print error message
      return []string{}, errors.New("error reading directory")
   }
   
   fileList := []string{}
   
   fmt.Printf("adding files to list:\n")
   for _, fileInfos := range fileInfos {
      fileName, err := TrimExtension(fileInfos.Name())
      if err != nil {
         return []string{}, errors.New("error trimming extension")
         }  
      fmt.Printf("File Before Trim: %s, File After Trim: %s\n", fileInfos.Name(), fileName)
      fileList = append(fileList, fileName)
   }

   return fileList, nil
}