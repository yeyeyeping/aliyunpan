package localfile

import (
	"fmt"
	"github.com/tickstep/aliyunpan/internal/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestMyWalkFile(t *testing.T) {
	count := 0
	walkFunc := func(file SymlinkFile, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		count += 1
		fmt.Println("file: ", utils.ObjectToJsonStr(file, false))
		//fmt.Println("file: ", file)
		return nil
	}

	//curPath := "D:\\smb\\feny\\goprojects\\dev\\lks"
	curPath := "D:\\smb\\feny\\goprojects\\dev\\NS游戏备份"
	file := NewSymlinkFile(curPath)
	if err := WalkAllFile(file, walkFunc); err != nil {
		if err != filepath.SkipDir {
			fmt.Printf("警告: 遍历错误: %s\n", err)
		}
	}
	fmt.Println("count: ", count)
}
