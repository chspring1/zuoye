package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库连接
	dsn := "chenhh:123456@tcp(192.168.253.100:3306)/school_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	// 测试连接
	if err := db.Ping(); err != nil {
		log.Fatal("数据库连接测试失败:", err)
	}
	fmt.Println("数据库连接成功!")

	// 1. 插入新记录 - 学生姓名为"张三"，年龄为20，年级为"三年级"
	insertStudent(db)

	// 2. 查询所有年龄大于18岁的学生信息
	queryStudentsOver18(db)

	// 3. 更新张三的年级为"四年级"
	updateStudentGrade(db)

	// 4. 删除年龄小于15岁的学生记录
	deleteStudentsUnder15(db)
}

// 1. 插入新记录
func insertStudent(db *sql.DB) {
	// SQL语句：插入新学生记录
	insertSQL := `
        INSERT INTO students (name, age, grade) 
        VALUES (?, ?, ?)
    `

	result, err := db.Exec(insertSQL, "张三", 20, "三年级")
	if err != nil {
		log.Printf("插入学生记录失败: %v", err)
		return
	}

	// 获取插入记录的ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("获取插入ID失败: %v", err)
		return
	}

	// 获取影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("获取影响行数失败: %v", err)
		return
	}

	fmt.Printf("✅ 插入成功 - ID: %d, 影响行数: %d\n", lastInsertID, rowsAffected)
	fmt.Println("SQL语句:", insertSQL)
	fmt.Println("参数: 张三, 20, 三年级")
	fmt.Println("---")
}

// 2. 查询年龄大于18岁的学生
func queryStudentsOver18(db *sql.DB) {
	// SQL语句：查询年龄大于18岁的学生
	querySQL := `
        SELECT id, name, age, grade 
        FROM students 
        WHERE age > 18
        ORDER BY age DESC
    `

	rows, err := db.Query(querySQL)
	if err != nil {
		log.Printf("查询学生记录失败: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("✅ 年龄大于18岁的学生信息:")
	fmt.Println("SQL语句:", querySQL)
	fmt.Printf("%-5s %-10s %-5s %-10s\n", "ID", "姓名", "年龄", "年级")
	fmt.Println("----------------------------------------")

	count := 0
	for rows.Next() {
		var id int
		var name, grade string
		var age int

		err := rows.Scan(&id, &name, &age, &grade)
		if err != nil {
			log.Printf("扫描行数据失败: %v", err)
			continue
		}

		fmt.Printf("%-5d %-10s %-5d %-10s\n", id, name, age, grade)
		count++
	}

	if err = rows.Err(); err != nil {
		log.Printf("遍历结果集失败: %v", err)
		return
	}

	fmt.Printf("共查询到 %d 条记录\n", count)
	fmt.Println("---")
}

// 3. 更新张三的年级
func updateStudentGrade(db *sql.DB) {
	// SQL语句：更新张三的年级为四年级
	updateSQL := `
        UPDATE students 
        SET grade = ? 
        WHERE name = ?
    `

	result, err := db.Exec(updateSQL, "四年级", "张三")
	if err != nil {
		log.Printf("更新学生年级失败: %v", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("获取影响行数失败: %v", err)
		return
	}

	fmt.Printf("✅ 更新成功 - 影响行数: %d\n", rowsAffected)
	fmt.Println("SQL语句:", updateSQL)
	fmt.Println("参数: 四年级, 张三")

	if rowsAffected == 0 {
		fmt.Println("⚠️  警告: 没有找到姓名为'张三'的学生记录")
	}
	fmt.Println("---")
}

// 4. 删除年龄小于15岁的学生
func deleteStudentsUnder15(db *sql.DB) {
	// 先查询要删除的记录
	querySQL := `
        SELECT id, name, age, grade 
        FROM students 
        WHERE age < 15
    `

	rows, err := db.Query(querySQL)
	if err != nil {
		log.Printf("查询待删除记录失败: %v", err)
		return
	}

	fmt.Println("🗑️  将要删除的学生记录:")
	fmt.Printf("%-5s %-10s %-5s %-10s\n", "ID", "姓名", "年龄", "年级")
	fmt.Println("----------------------------------------")

	count := 0
	for rows.Next() {
		var id int
		var name, grade string
		var age int

		err := rows.Scan(&id, &name, &age, &grade)
		if err != nil {
			log.Printf("扫描行数据失败: %v", err)
			continue
		}

		fmt.Printf("%-5d %-10s %-5d %-10s\n", id, name, age, grade)
		count++
	}
	rows.Close()

	if count == 0 {
		fmt.Println("没有找到年龄小于15岁的学生记录")
		fmt.Println("---")
		return
	}

	// SQL语句：删除年龄小于15岁的学生
	deleteSQL := `
        DELETE FROM students 
        WHERE age < 15
    `

	result, err := db.Exec(deleteSQL)
	if err != nil {
		log.Printf("删除学生记录失败: %v", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("获取影响行数失败: %v", err)
		return
	}

	fmt.Printf("✅ 删除成功 - 影响行数: %d\n", rowsAffected)
	fmt.Println("SQL语句:", deleteSQL)
	fmt.Println("---")
}
