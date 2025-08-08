// 新建文件夹处理器
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 新建文件夹处理器
func mkdirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dir := r.FormValue("dir")
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "文件夹名不能为空", http.StatusBadRequest)
			return
		}
		path := name
		if dir != "" {
			path = dir + string(os.PathSeparator) + name
		}
		err := os.MkdirAll(path, 0755)
		if err != nil {
			http.Error(w, "创建文件夹失败", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "文件夹 %s 创建成功!", path)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 上传页面，支持输入目标子目录
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<html><body><form enctype="multipart/form-data" action="/upload" method="post">
目标子目录: <input type="text" name="dir" value="" placeholder="如 subdir，可留空"><br>
选择文件: <input type="file" name="file"><br>
<input type="submit" value="上传">
</form></body></html>`)
		return
	}
	if r.Method == "POST" {
		dir := r.FormValue("dir")
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "文件获取失败", http.StatusBadRequest)
			return
		}
		defer file.Close()

		savePath := header.Filename
		if dir != "" {
			// 创建目标子目录（如不存在）
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				http.Error(w, "创建目录失败", http.StatusInternalServerError)
				return
			}
			savePath = dir + string(os.PathSeparator) + header.Filename
		}

		out, err := os.Create(savePath)
		if err != nil {
			http.Error(w, "文件保存失败", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "文件写入失败", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "文件 %s 上传成功!", savePath)
	}
}

// 首页处理器，带上传表单，自动跟随iframe当前目录
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<html><body>
	<form id="uploadForm" enctype="multipart/form-data" action="/upload" method="post" target="uploadResult">
		<input type="hidden" id="dirInput" name="dir" value="">
		<input type="file" name="file" required>
		<input type="submit" value="上传到当前目录">
	</form>
	<form id="mkdirForm" action="/mkdir" method="post" target="mkdirResult" style="margin-top:10px;">
		<input type="hidden" id="mkdirDirInput" name="dir" value="">
		<input type="text" name="name" placeholder="新建文件夹名" required>
		<input type="submit" value="新建文件夹">
	</form>
	<iframe name="uploadResult" style="display:none;"></iframe>
	<iframe name="mkdirResult" style="display:none;"></iframe>
	<br>
	<iframe id="fileFrame" src="/files/" width="100%" height="600" frameborder="0"></iframe>

	<script>
	// 监听iframe跳转，自动同步目录到上传和新建文件夹表单
	function updateDir() {
		var frame = document.getElementById('fileFrame');
		var dir = decodeURIComponent(frame.contentWindow.location.pathname.replace(/^\\/files\\/?/, ''));
		document.getElementById('dirInput').value = dir;
		document.getElementById('mkdirDirInput').value = dir;
	}
	document.getElementById('fileFrame').addEventListener('load', updateDir);
	</script>
	</body></html>`)
}

func main() {
	// 上传接口
	http.HandleFunc("/upload", uploadHandler)
	// 新建文件夹接口
	http.HandleFunc("/mkdir", mkdirHandler)
	// 首页，带上传和新建文件夹表单、文件列表
	http.HandleFunc("/", indexHandler)
	// 文件列表用 /files/ 路径
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/files/", http.StripPrefix("/files/", fs))

	log.Println("文件服务器启动，访问 http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
