package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"unicode"
	"strconv"
	"math"
)

// The main function
func main() {
	input_matrix, rdm_cols, rdm_rows := get_input()   //function call to get inputs
	sum_mtrx_elements := sum_elements(input_matrix)   //function call to get the sum of squares of all elements
	sum_each_col := sum_col_sqr(input_matrix)         //function call to get the sum of squares of all elements of columns
	prob_each_col := cal_probab(sum_each_col, sum_mtrx_elements)   //function call to find probability of selection of each column
	sum_each_row := sum_row_sqr(input_matrix)      //function call to get the sum of squares of all elements of each row
	prob_each_row := cal_probab(sum_each_row, sum_mtrx_elements)    //function call to find probability of selection of each row
	scaled_mtrx_C := find_scaled_C(input_matrix, rdm_cols, prob_each_col)   //function call to find the scaled matrix C
	scaled_mtrx_R := find_scaled_R(input_matrix, rdm_rows, prob_each_row)   //function call to find the scaled matrix R
	mtrx_W := find_mtrx_W(scaled_mtrx_R, rdm_cols)     //function call to find matrix W
	find_mtrx_sigma(mtrx_W)     //function call to find matrix sigma
	mtrx_U := find_mtrx_U(mtrx_W)     //function call to find matrix U
	mtrx_CUR := find_mtrx_cur(scaled_mtrx_R, scaled_mtrx_C, mtrx_U)     //function call to find matrix CUR
	a_cur := sub_mtrx(input_matrix, mtrx_CUR)
	calc_frobenius_norm(a_cur)
}

//Function to get input from stdin
func get_input() ([][]string, []string, []string) {
	var input_matrix_temp [][]string
	var rdm_row_temp []string
	var rdm_col_temp []string
	var len_row_1 int
	flag_char := 0
	flag_char_count := 0
	flag_size := false
	flag_rdm_row := 0
	flag_rdm_col := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		str := strings.Fields(scanner.Text())
		if str[0] == "#" {
			continue
		} else if str[0] == "Random" && strings.Contains(str[1], "rows:") {
			rdm_row_temp = get_rdm_inputs(str)
		} else if str[0] == "Random" && strings.Contains(str[1], "columns:") {
			rdm_col_temp = get_rdm_inputs(str)
		} else {
			str, flag_char = check_space_char(str)
			flag_char_count += flag_char
			if len(input_matrix_temp) == 0 {
				len_row_1 = len(str)
				input_matrix_temp = append(input_matrix_temp, str)
			} else if len_row_1 == len(str) {
				input_matrix_temp = append(input_matrix_temp, str)
			} else {
				flag_size = true
				input_matrix_temp = append(input_matrix_temp, str)
			}
		}
	}
	if len(rdm_row_temp) == 0 || len(rdm_col_temp) == 0 || len(input_matrix_temp) == 0 {
		fmt.Println("Invalid Input!!!")
		os.Exit(0)
	}
	n1, _ := strconv.Atoi(rdm_row_temp[0])
	n2, _ := strconv.Atoi(rdm_row_temp[1])
	if n1 < 0 || n1 > len(input_matrix_temp) - 1 || n2 < 0 || n2 > len(input_matrix_temp) - 1 {
		flag_rdm_row += 1
	}
	n3, _ := strconv.Atoi(rdm_col_temp[0])
	n4, _ := strconv.Atoi(rdm_col_temp[1])
	if n3 < 0 || n3 > len_row_1 - 1 || n4 < 0 || n4 > len_row_1 - 1 {
		flag_rdm_row += 1
	}
	if flag_char_count == 0 && flag_size == false && len(rdm_col_temp) > 0 && len(rdm_row_temp) > 0 && len(rdm_row_temp) ==2 && len(rdm_col_temp) == 2 && flag_rdm_col == 0 && flag_rdm_row == 0 {
		// Printing the user inputs on the stdout
		fmt.Println("The input matrix is: ")
		for _, str := range input_matrix_temp {
			fmt.Printf("%s\n", str)
		}
		fmt.Println()
		fmt.Println("The input for random rows are: ", rdm_row_temp[0], rdm_row_temp[1])
		fmt.Println("The input for random columns are: ", rdm_col_temp[0], rdm_col_temp[1])
		fmt.Println()
	} else {
		fmt.Println("Invalid Input!!!")
		os.Exit(0)
	}
	return input_matrix_temp, rdm_col_temp, rdm_row_temp
}

//Function to handle spaces and validate if the input does not has character as an input
func check_space_char(temp []string) ([]string, int) {
	var r []string
	flag := 0
	for _, str := range temp {
		if str != " " {
			r_temp := []rune(str)
			if unicode.IsNumber(r_temp[0]) {
				r = append(r, str)
			} else {
				r = append(r, str)
				flag = 1
			}
		}
	}
	return r, flag
}

//Function to hanfle input for random rows and columns
func get_rdm_inputs(str []string) ([]string) {
	var temp string
	var arr []string
	for _, item := range str {
		temp = temp + item
	}
	for _, item := range temp {
		if unicode.IsNumber(item) {
			arr = append(arr, string(item))
		} else {
			continue
		}
	}
	return arr
}

//Function to calculate the sum of squares of each matrix element
func sum_elements(matrix [][] string) int {
	sum := 0
	for _, value := range matrix {
		for _, val := range value {
			num, _ := strconv.Atoi(val)
			sum += num * num
		}
	}
	return sum
}

//Function to calculate the sum of squares of all the elements in every coloumn
func sum_col_sqr(matrix [][]string) ([]int) {
	sum := make([]int, len(matrix[0]))
	for _, value := range matrix {
		for i := 0; i < len(value); i++ {
			num, _ := strconv.Atoi(value[i])
			sum[i] += num * num
		}
	}
    return sum
}

//Function to probability of selection of each column or row
func cal_probab(matrix []int, denominator int) ([]float64) {
	probability := make([]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		probability[i] = float64(matrix[i]) / float64(denominator)
	}
	return probability
}

//Function to find scaled_C matrix
func find_scaled_C(matrix [][]string, cols []string, probability []float64) ([][]float64){
	length_col := len(matrix[0])
	mtrx_C := make([][]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		mtrx_C[i] = make([]float64, len(cols))
	}
	col1, _ := strconv.Atoi(cols[0])
	col2, _ := strconv.Atoi(cols[1])
	for i := 0; i < len(matrix); i++ {
		num1, _ := strconv.Atoi(matrix[i][col1])
		num2, _ := strconv.Atoi(matrix[i][col2])
		mtrx_C[i][0] = float64(num1) / math.Sqrt(float64(length_col) * probability[col1])
		mtrx_C[i][1] = float64(num2) / math.Sqrt(float64(length_col) * probability[col2])
	}
	// Printing the user inputs on the stdout
	fmt.Println("The Scaled matrix C is: ")
	for _, item := range mtrx_C {
		fmt.Println(item)
	}
	fmt.Println()
	return mtrx_C
}

//Function to calculate the sum of squares of all the elements in every row
func sum_row_sqr(matrix [][]string) ([]int) {
	sum := make([]int, len(matrix))
	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[0]); i++ {
			num, _ := strconv.Atoi(matrix[j][i])
			sum[j] += num * num
		}
	}
    return sum
}

//Function to find scaled_C matrix
func find_scaled_R(matrix [][]string, rows []string, probability []float64) ([][]float64){
	length_row := len(matrix)
	mtrx_R := make([][]float64, len(rows))
	for i := 0; i < len(rows); i++ {
		mtrx_R[i] = make([]float64, len(matrix[0]))
	}
	row1, _ := strconv.Atoi(rows[0])
	row2, _ := strconv.Atoi(rows[1])
	for i := 0; i < len(matrix[0]); i++ {
		num1, _ := strconv.Atoi(matrix[row1][i])
		num2, _ := strconv.Atoi(matrix[row2][i])
		mtrx_R[0][i] = float64(num1) / math.Sqrt(float64(length_row) * probability[row1])
		mtrx_R[1][i] = float64(num2) / math.Sqrt(float64(length_row) * probability[row2])
	}
	// Printing the user inputs on the stdout
	fmt.Println("The Scaled matrix R is: ")
	for _, item := range mtrx_R {
		fmt.Println(item)
	}
	fmt.Println()
	return mtrx_R
}

//Function to find matrix W
func find_mtrx_W(matrix [][]float64, cols []string) ([][]float64) {
	mtrx_W := make([][]float64, 2)
	for i := 0; i < 2; i++ {
		mtrx_W[i] = make([]float64, 2)
	}
	col1, _ := strconv.Atoi(cols[0])
	col2, _ := strconv.Atoi(cols[1])
	mtrx_W[0][0] = matrix[0][col1]
	mtrx_W[0][1] = matrix[0][col2]
	mtrx_W[1][0] = matrix[1][col1]
	mtrx_W[1][1] = matrix[1][col2]
	fmt.Println("The matrix W is: ")
	for _, item := range mtrx_W {
		fmt.Println(item)
	}
	fmt.Println()
	return mtrx_W
}

//Function to multiply 2 float type matrices
func mult_mtrx(mtrx1 [][]float64, mtrx2 [][]float64) ([][]float64) {
	row1 := len(mtrx1)
	row2 := len(mtrx2)
	col1 := len(mtrx1[0])
	col2 := len(mtrx2[0])
	total := 0.0
	result_mtrx := make([][]float64, row1)
	for i := 0; i < row1; i++ {
		result_mtrx[i] = make([]float64, col2)
	}
	if row2 != col1 { 
		fmt.Println("Error: The matrix cannot be multiplied")
		os.Exit(0)
	} else {
		for i := 0; i < row1; i++ {
			for j := 0; j < col2; j++ {
				for k := 0; k < row2; k++ {
					total += float64(mtrx1[i][k]) * float64(mtrx2[k][j])
				}
				result_mtrx[i][j] = total
				total = 0
			}
		}
	}
	return result_mtrx
}

// Function to find transpose of a square matrix
func find_mtrx_sigma(mtrx1 [][]float64) {
	result_mtrx := make([][]float64, len(mtrx1))
	for i := 0; i < len(mtrx1); i++ {
		result_mtrx[i] = make([]float64, len(mtrx1[0]))
	}
	for i:= 0; i < len(mtrx1); i++ {
		for j := 0; j < len(mtrx1[0]); j++ {
			if i == 0 && j == 0 {
				result_mtrx[i][j] = mtrx1[1][1]
			} else if i == 1 && j == 1 {
				result_mtrx[i][j] = mtrx1[0][0]
			} else {
				result_mtrx[i][j] = mtrx1[i][j]
			}
		}
	}
	fmt.Println("The matrix sigma is: ")
	for _, item := range result_mtrx {
		fmt.Println(item)
	}
	fmt.Println()
}

//Function to find matrix U
func find_mtrx_U(mtrx1 [][]float64) ([][]float64) {
	result_mtrx := make([][]float64, len(mtrx1))
	for i := 0; i < len(mtrx1); i++ {
		result_mtrx[i] = make([]float64, len(mtrx1[0]))
	}
	for i:= 0; i < len(mtrx1); i++ {
		for j := 0; j < len(mtrx1[0]); j++ {
			if mtrx1[i][j] == 0{
				result_mtrx[i][j] = mtrx1[i][j]
			} else {
				result_mtrx[i][j] = 1 / mtrx1[i][j]
			} 	
		}
	}
	fmt.Println("The matrix U is: ")
	for _, item := range result_mtrx {
		fmt.Println(item)
	}
	fmt.Println()
	return result_mtrx
}

//Function to find matrix CUR
func find_mtrx_cur(mtrx_R [][]float64, mtrx_C [][]float64, mtrx_U [][]float64) ([][]float64) {
	temp := mult_mtrx(mtrx_C, mtrx_U)
	result_mtrx := mult_mtrx(temp, mtrx_R)
	fmt.Println("The matrix CUR is: ")
	for _, item := range result_mtrx {
		fmt.Println(item)
	}
	fmt.Println()
	return result_mtrx
}

func sub_mtrx(mtrx1 [][]string, mtrx2 [][]float64) ([][]float64) {
	result_mtrx := make([][]float64, len(mtrx1))
	for i := 0; i < len(mtrx1); i++ {
		result_mtrx[i] = make([]float64, len(mtrx1[0]))
	}
	if len(mtrx1) != len(mtrx2) || len(mtrx1[0]) != len(mtrx2[0]) {
		fmt.Println("Matrix Substraction not possible")
	} else {
		for i:= 0; i < len(mtrx1); i++ {
			for j := 0; j < len(mtrx1[0]); j++ {
				num1, _ := strconv.Atoi(mtrx1[i][j])
				result_mtrx[i][j] = float64(num1) - mtrx2[i][j]	
			}
		}
	}
	return result_mtrx
}

func calc_frobenius_norm(mtrx1 [][]float64) {
	sum := 0.0
	for i:= 0; i < len(mtrx1); i++ {
		for j := 0; j < len(mtrx1[0]); j++ {
			sum += mtrx1[i][j]	
		}
	}
	fmt.Println("Frobenius Norm(A – CUR) = ", math.Sqrt(sum))
}