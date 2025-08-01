package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

func selectmysqlAndWriteExcel() {
	// 构建连接字符串
	dsn := "chenhh:123456@tcp(192.168.253.100:3306)/chenhh"
	// 连接 MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer db.Close()
	// 测试连接
	err = db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	fmt.Println("连接数据库成功")

	// 查询表 employs 数据
	querySQL := "SELECT id, name, age, birthday, salary FROM employs"
	rows, err := db.Query(querySQL)
	if err != nil {
		fmt.Println("查询数据失败:", err)
		return
	}
	defer rows.Close()

	// 创建 Excel 文件
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName(f.GetSheetName(0), sheet)

	// 写表头
	headers := []string{"ID", "Name", "Age", "Birthday", "Salary"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// 写数据
	rowNum := 2
	for rows.Next() {
		var id int
		var name string
		var age int
		var birthday string
		var salary float64
		err := rows.Scan(&id, &name, &age, &birthday, &salary)
		if err != nil {
			fmt.Println("扫描数据失败:", err)
			return
		}
		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNum), id)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNum), name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNum), age)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", rowNum), birthday)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", rowNum), salary)
		rowNum++
	}
	if err := rows.Err(); err != nil {
		fmt.Println("遍历结果集失败:", err)
		return
	}

	// 保存 Excel 文件
	if err := f.SaveAs("employs.xlsx"); err != nil {
		fmt.Println("保存 Excel 文件失败:", err)
		return
	}
	fmt.Println("数据已成功写入 employs.xlsx")
}

func main() {
	selectmysqlAndWriteExcel() // 查询并写入 Excel
	fmt.Println("程序执行完毕")
}
